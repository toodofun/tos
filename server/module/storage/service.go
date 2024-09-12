package storage

import (
	"github.com/MR5356/tos/config"
	"github.com/MR5356/tos/util/storagemanager"
	"io"
	"path/filepath"
)

type Service struct {
}

func GetService() *Service {
	return &Service{}
}

func (s *Service) ListDirectory(directoryPath string) ([]*storagemanager.FileInfo, error) {
	sm, err := storagemanager.GetStorage(storagemanager.LocationTypeLocal, config.Current().Storage.Root)
	if err != nil {
		return nil, err
	}
	return sm.ListDirectory(directoryPath)
}

func (s *Service) Upload(fileName string, fileContent io.Reader, targetPath string, mode string) error {
	sm, err := storagemanager.GetStorage(storagemanager.LocationTypeLocal, config.Current().Storage.Root)
	if err != nil {
		return err
	}

	// 如果文件已存在，忽略
	if sm.Exists(filepath.Join(targetPath, fileName)) && mode == "ignore" {
		return nil
	}
	return sm.Upload(fileName, fileContent, targetPath)
}

func (s *Service) Exists(filePath string) bool {
	sm, err := storagemanager.GetStorage(storagemanager.LocationTypeLocal, config.Current().Storage.Root)
	if err != nil {
		return false
	}

	return sm.Exists(filePath)
}

func (s *Service) GetSpecialPath() []*storagemanager.FileInfo {
	sm, err := storagemanager.GetStorage(storagemanager.LocationTypeLocal, config.Current().Storage.Root)
	if err != nil {
		return nil
	}
	return sm.GetSpecialPath()
}

func (s *Service) Initialize() error {
	return nil
}
