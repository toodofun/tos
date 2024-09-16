package storagemanager

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var fileManager StorageManager = &LocalStorage{}
var ErrRootPath = errors.New("root path error")
var ErrCreatePath = errors.New("create path error")

var specialPath = []*FileInfo{
	{
		Name: "Root",
		Path: "/",
	},
}

type LocalStorage struct {
	root string
}

func NewLocalStorage(args ...interface{}) (*LocalStorage, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("%w: root path is required", ErrRootPath)
	}

	root, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("%w: root path type error, required string but got %T", ErrRootPath, args[0])
	}

	root, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(root)
	if os.IsNotExist(err) {
		err := os.MkdirAll(root, os.ModePerm)
		if err != nil {
			logrus.Errorf("create root path error: %v", err)
			return nil, ErrCreatePath
		}
	}

	ls := &LocalStorage{
		root: root,
	}

	// 创建特殊目录
	for _, v := range specialPath {
		p, _ := ls.getRealPath(v.Path)
		_ = os.MkdirAll(p, os.ModePerm)
	}

	return ls, nil
}

func (lfm *LocalStorage) GetSpecialPath() []*FileInfo {
	return specialPath
}

func (lfm *LocalStorage) Close() error {
	return nil
}

func (lfm *LocalStorage) Upload(fileName string, fileContent io.Reader, targetPath string) error {
	targetPath, err := lfm.getRealPath(targetPath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(targetPath, os.ModePerm)
	if err != nil {
		logrus.Errorf("create target path error: %v", err)
		return ErrCreatePath
	}

	file, err := os.Create(filepath.Join(targetPath, fileName))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = io.Copy(file, fileContent)
	return err
}

func (lfm *LocalStorage) Download(filePath string) (string, io.ReadCloser, error) {
	filePath, err := lfm.getRealPath(filePath)
	if err != nil {
		return "", nil, err
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}

	// 确保文件在 root 目录下存在
	stat, err := file.Stat()
	if err != nil {
		_ = file.Close()
		return "", nil, err
	}

	// 确保不是目录
	if stat.IsDir() {
		_ = file.Close()
		return "", nil, errors.New("cannot download a directory")
	}

	// 返回文件名和文件读取流
	return filepath.Base(filePath), file, nil
}

func (lfm *LocalStorage) GetFileInfo(filePath string) (*FileInfo, error) {
	filePath, err := lfm.getRealPath(filePath)
	if err != nil {
		return nil, err
	}

	stat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	return &FileInfo{
		Name:    stat.Name(),
		Path:    filePath,
		Size:    stat.Size(),
		IsDir:   stat.IsDir(),
		ModTime: stat.ModTime(),
	}, nil
}

func (lfm *LocalStorage) ListDirectory(directoryPath string) ([]*FileInfo, error) {
	// 获取真实路径
	directoryPath, err := lfm.getRealPath(directoryPath)
	if err != nil {
		return nil, err
	}

	fileInfos := make([]*FileInfo, 0)
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		logrus.Warnf("read directory error: %s %v", directoryPath, err)
		return nil, ErrFileNotFound
	}

	// 创建两个切片分别存储文件夹和文件
	var dirs []*FileInfo
	var files []*FileInfo

	// 遍历目录内容
	for _, entry := range entries {
		file, err := entry.Info()
		if err != nil {
			logrus.Warnf("get file info error: %s %v", entry.Name(), err)
			continue
		}

		relativePath, err := filepath.Rel(lfm.root, filepath.Join(directoryPath, file.Name()))
		if err != nil {
			logrus.Warnf("get relative path error: %s %v", file.Name(), err)
			continue
		}

		info := &FileInfo{
			Name:    file.Name(),
			Path:    filepath.ToSlash("/" + relativePath),
			Size:    file.Size(),
			IsDir:   file.IsDir(),
			ModTime: file.ModTime(),
		}

		// 分别存储文件夹和文件
		if file.IsDir() {
			dirs = append(dirs, info)
		} else {
			files = append(files, info)
		}
	}

	// 对文件夹和文件分别按名称排序
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name < dirs[j].Name
	})
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	// 合并文件夹和文件
	fileInfos = append(fileInfos, dirs...)
	fileInfos = append(fileInfos, files...)

	return fileInfos, nil
}

func (lfm *LocalStorage) Copy(srcPath, dstPath string) error {
	srcPath, err := lfm.getRealPath(srcPath)
	if err != nil {
		return err
	}
	dstPath, err = lfm.getRealPath(dstPath)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = srcFile.Close()
	}()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = dstFile.Close()
	}()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return dstFile.Sync()
}

func (lfm *LocalStorage) Move(srcPath, dstPath string) error {
	srcPath, err := lfm.getRealPath(srcPath)
	if err != nil {
		return err
	}
	dstPath, err = lfm.getRealPath(dstPath)
	if err != nil {
		return err
	}

	return os.Rename(srcPath, dstPath)
}

func (lfm *LocalStorage) Remove(filePath string) error {
	filePath, err := lfm.getRealPath(filePath)
	if err != nil {
		return err
	}
	return os.RemoveAll(filePath)
}

func (lfm *LocalStorage) Exists(filePath string) bool {
	filePath, err := lfm.getRealPath(filePath)
	if err != nil {
		return false
	}

	_, err = os.Stat(filePath)
	return !os.IsNotExist(err)
}

func (lfm *LocalStorage) Share(filePath string) (string, error) {
	filePath, err := lfm.getRealPath(filePath)
	if err != nil {
		return "", err
	}

	if !lfm.Exists(filePath) {
		return "", ErrFileNotFound
	}
	return "not support", nil
}

// getRealPath 计算出相对于 root 的安全路径，确保不能访问 root 目录之外的内容
func (lfm *LocalStorage) getRealPath(path string) (string, error) {
	// 禁止使用环境变量符号和特殊符号
	if strings.Contains(path, "~") || strings.Contains(path, "$") {
		return "", errors.New("invalid path: contains environment variables or special characters")
	}

	// 清理传入的路径，移除冗余的符号
	cleanPath := filepath.Clean(filepath.FromSlash(path))

	//// 获取 root 目录的绝对路径
	absPath := filepath.Join(lfm.root, cleanPath)

	// 确保传入路径不会超出 root 目录
	if !strings.HasPrefix(absPath, lfm.root) {
		return "", errors.New("invalid path: path attempts to escape root directory")
	}

	return absPath, nil
}
