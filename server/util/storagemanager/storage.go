package storagemanager

import (
	"errors"
	"io"
	"time"
)

const (
	LocationTypeLocal = "local"
)

var (
	ErrPathFormat   = errors.New("path format error")
	ErrFileNotFound = errors.New("file not found")
	ErrStorageType  = errors.New("storage type error")
)

func GetStorage(locationType string, args ...interface{}) (fm StorageManager, err error) {
	switch locationType {
	case LocationTypeLocal:
		return NewLocalStorage(args...)
	default:
		return nil, ErrStorageType
	}
}

type StorageManager interface {
	// GetSpecialPath 获取特殊路径
	GetSpecialPath() []*FileInfo
	// Upload 上传文件
	Upload(fileName string, fileContent io.Reader, targetPath string) error
	// Download 下载文件，返回文件名和文件内容
	Download(filePath string) (string, io.ReadCloser, error)
	// GetFileInfo 获取文件信息
	GetFileInfo(filePath string) (*FileInfo, error)
	// ListDirectory 列出目录下的文件
	ListDirectory(directoryPath string) ([]*FileInfo, error)
	// Copy 复制文件
	Copy(srcPath, dstPath string) error
	// Move 移动文件，重命名文件
	Move(srcPath, dstPath string) error
	// Remove 删除文件
	Remove(filePath string) error
	// Exists 判断文件是否存在
	Exists(filePath string) bool
	// Share 创建文件分享链接
	Share(filePath string) (string, error)
	// Close 关闭
	Close() error
}

type FileInfo struct {
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Size    int64     `json:"size"`
	IsDir   bool      `json:"isDir"`
	ModTime time.Time `json:"modTime"`
}
