package application

import (
	"github.com/MR5356/tos/persistence/database"
	"github.com/google/uuid"
)

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

type Theme string

type App struct {
	Title        string `json:"title"`
	Icon         string `json:"icon"`
	Page         string `json:"page"`
	Width        uint   `json:"width"`
	Height       uint   `json:"height"`
	X            uint   `json:"x"`
	Y            uint   `json:"y"`
	Theme        Theme  `json:"theme"`
	Background   string `json:"background"`
	Singleton    bool   `json:"singleton"`
	BanUninstall bool   `json:"banUninstall"`
	FixOnDock    bool   `json:"fixOnDock"`
	FixOnDesk    bool   `json:"fixOnDesk"`

	database.BaseModel
}

func (a *App) TableName() string {
	return "applications"
}

var defaultApps = []*App{
	{
		BaseModel: database.BaseModel{
			ID: uuid.MustParse("ab95a563-f9c8-1bd2-5cf6-9e93c80d2c3b"),
		},
		Title:        "文件管理器",
		Icon:         "internal://icon-oss",
		Page:         "internal://finder",
		FixOnDock:    true,
		FixOnDesk:    true,
		BanUninstall: true,
	},
	{
		BaseModel: database.BaseModel{
			ID: uuid.MustParse("d79d66f2-2932-0753-81a1-3b66ae6da94a"),
		},
		Title:        "启动台",
		Icon:         "internal://icon-app",
		Page:         "system://launchpad",
		Background:   "linear-gradient(to right, #4e54c8, #8f94fb)",
		FixOnDock:    true,
		BanUninstall: true,
	},
	{
		BaseModel: database.BaseModel{
			ID: uuid.MustParse("40e3595a-7c10-cc09-08fb-54683ff39d74"),
		},
		Title:        "终端",
		Icon:         "internal://icon-terminal",
		Page:         "internal://terminal",
		Theme:        ThemeDark,
		FixOnDock:    true,
		FixOnDesk:    true,
		BanUninstall: true,
	},
	{
		BaseModel: database.BaseModel{
			ID: uuid.MustParse("ef708add-8c00-f5de-4279-40115b52321d"),
		},
		Title:        "应用商店",
		Icon:         "internal://icon-app-store",
		Page:         "internal://app-store",
		Background:   "linear-gradient(to right, #06beb6, #48b1bf)",
		FixOnDesk:    true,
		FixOnDock:    true,
		Singleton:    true,
		BanUninstall: true,
	},
	{
		BaseModel: database.BaseModel{
			ID: uuid.MustParse("039b3754-2eed-5c77-67e1-00549cfc44ce"),
		},
		Title:        "设置",
		Icon:         "internal://icon-setting",
		Page:         "internal://setting",
		Background:   "linear-gradient(to right, #536976, #292e49)",
		FixOnDock:    true,
		Singleton:    true,
		BanUninstall: true,
	},
}
