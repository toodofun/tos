package storage

import (
	"errors"
	"github.com/MR5356/tos/config"
	"github.com/MR5356/tos/persistence/database"
	"github.com/MR5356/tos/util/cacheutil"
	"github.com/MR5356/tos/util/storagemanager"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"time"
)

var (
	defaultLocationId = uuid.MustParse("ae35c8f1-aad8-b5e2-b5ca-0187ec7cafa1")
)

type Service struct {
	cache *cacheutil.CountdownCache[storagemanager.StorageManager]
	db    *database.BaseMapper[*Storage]
}

func GetService() *Service {
	return &Service{
		cache: cacheutil.NewCountdownCache[storagemanager.StorageManager](time.Minute * 30),
		db:    database.NewMapper(database.GetDB(), &Storage{}),
	}
}

func (s *Service) GetStorageManager(id uuid.UUID) (storagemanager.StorageManager, error) {
	location, err := s.db.Detail(NewStorageWithID(id))
	if err != nil {
		logrus.Errorf("get storage manager error: %v", err)
		return nil, errors.New("获取存储管理器失败")
	}
	locationType := location.LocationType

	if sm, ok := s.cache.Get(locationType); ok {
		return sm, nil
	}
	switch locationType {
	case storagemanager.LocationTypeLocal:
		if sm, err := storagemanager.GetStorage(storagemanager.LocationTypeLocal, config.Current().Storage.Root); err != nil {
			return nil, err
		} else {
			s.cache.Set(locationType, sm, func() {
				_ = sm.Close()
			})
			return sm, nil
		}
	default:
		return nil, storagemanager.ErrStorageType
	}
}

func (s *Service) ListStorages() ([]*Storage, error) {
	res := make([]*Storage, 0)
	if err := s.db.DB.Model(&Storage{}).Select("id, title, location_type").Find(&res).Error; err != nil {
		logrus.Errorf("list storages error: %v", err)
		return nil, errors.New("获取存储列表失败")
	}
	return res, nil
}

func (s *Service) ListDirectory(id uuid.UUID, directoryPath string) ([]*storagemanager.FileInfo, error) {
	sm, err := s.GetStorageManager(id)
	if err != nil {
		return nil, err
	}
	return sm.ListDirectory(directoryPath)
}

func (s *Service) Upload(id uuid.UUID, fileName string, fileContent io.Reader, targetPath string, mode string) error {
	sm, err := s.GetStorageManager(id)
	if err != nil {
		return err
	}

	// 如果文件已存在，忽略
	if sm.Exists(filepath.Join(targetPath, fileName)) && mode == "ignore" {
		return nil
	}
	return sm.Upload(fileName, fileContent, targetPath)
}

func (s *Service) Exists(id uuid.UUID, filePath string) bool {
	sm, err := s.GetStorageManager(id)
	if err != nil {
		return false
	}

	return sm.Exists(filePath)
}

func (s *Service) GetSpecialPath(id uuid.UUID) []*storagemanager.FileInfo {
	sm, err := s.GetStorageManager(id)
	if err != nil {
		return nil
	}
	return sm.GetSpecialPath()
}

func (s *Service) Initialize() error {
	if err := s.db.DB.AutoMigrate(&Storage{}); err != nil {
		return err
	}

	defaultLocalStorage := NewStorageWithID(defaultLocationId)
	defaultLocalStorage.LocationType = storagemanager.LocationTypeLocal
	defaultLocalStorage.Args = config.Current().Storage.Root
	defaultLocalStorage.Title = "默认存储"

	if err := s.db.DB.FirstOrCreate(defaultLocalStorage).Error; err != nil {
		return err
	}
	return nil
}
