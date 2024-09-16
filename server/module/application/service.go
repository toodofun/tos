package application

import "github.com/MR5356/tos/persistence/database"

type Service struct {
	db *database.BaseMapper[*App]
}

func GetService() *Service {
	return &Service{
		db: database.NewMapper(database.GetDB(), &App{}),
	}
}

func (s *Service) ListApps() ([]*App, error) {
	return s.db.List(&App{})
}

func (s *Service) Initialize() error {
	if err := s.db.DB.AutoMigrate(&App{}); err != nil {
		return err
	}

	for _, app := range defaultApps {
		if err := s.db.DB.FirstOrCreate(app).Error; err != nil {
			return err
		}
	}

	return nil
}
