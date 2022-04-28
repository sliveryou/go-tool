package filex

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

// Ext returns the lowercase extension of the file name, excluding the dot.
func Ext(fileName string) string {
	ext := filepath.Ext(fileName)
	if ext != "" {
		ext = strings.ToLower(ext[1:])
	}
	return ext
}

// Size returns the file size in bytes.
// If the file does not exist or cannot be accessed, it returns -1.
func Size(fileName string) int64 {
	f, err := os.Stat(fileName)
	if nil != err {
		return -1
	}
	return f.Size()
}

// DirSize returns the number of files contained in the dictionary and size in bytes.
// If the dictionary is empty or does not exist or cannot be accessed, it returns 0, 0.
func DirSize(rootPath string) (fileNum, dirSize int64) {
	var wg sync.WaitGroup
	fileSizes := make(chan int64)

	wg.Add(1)
	go walkDirSize(rootPath, &wg, fileSizes)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	for fileSize := range fileSizes {
		fileNum++
		dirSize += fileSize
	}
	return fileNum, dirSize
}

// MD5 returns the MD5 hash string of the file content by file name.
func MD5(fileName string) (string, error) {
	return Hash(fileName, md5.New())
}

// SHA1 returns the SHA1 hash string of the file content by file name.
func SHA1(fileName string) (string, error) {
	return Hash(fileName, sha1.New())
}

// SHA256 returns the SHA256 hash string of the file content by file name.
func SHA256(fileName string) (string, error) {
	return Hash(fileName, sha256.New())
}

// SHA512 returns the SHA512 hash string of the file content by file name.
func SHA512(fileName string) (string, error) {
	return Hash(fileName, sha512.New())
}

// Hash returns the hash string of the file content by file name and hash algorithm.
func Hash(fileName string, h hash.Hash) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}

	bytes := h.Sum(nil)
	return hex.EncodeToString(bytes), nil
}

// Read reads the file by file name and returns the contents.
func Read(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

// ReadToSlice returns the file contents slice separated by row,
// and excluding the end-of-line bytes.
func ReadToSlice(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	slice := make([]string, 0)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				return slice, nil
			}
			return nil, err
		}
		slice = append(slice, string(line))
	}
}

// Write writes data to the file by file name with perm, default perm is 0666.
// If the file does not exist, Write creates it with permissions perm,
// otherwise Write truncates it before writing.
func Write(fileName string, data []byte, perm ...os.FileMode) error {
	if dir := path.Dir(fileName); dir != "" {
		if err := os.MkdirAll(dir, 0o777); err != nil {
			return err
		}
	}
	pe := getPerm(0o666, perm...)
	return ioutil.WriteFile(fileName, data, pe)
}

// Append appends data to the end of the file by file name with perm, default perm is 0666.
// If the file does not exist, Append creates it with permissions perm,
// otherwise Append appends data to the end of the file.
func Append(fileName string, data []byte, perm ...os.FileMode) error {
	if dir := path.Dir(fileName); dir != "" {
		if err := os.MkdirAll(dir, 0o777); err != nil {
			return err
		}
	}
	pe := getPerm(0o666, perm...)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, pe)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

// Copy copies src file contents to dest file contents with perm, default perm is 0666.
// If dest file already exists, Copy truncates it before copying.
func Copy(srcName, destName string, perm ...os.FileMode) error {
	sf, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer sf.Close()

	if dir := path.Dir(destName); dir != "" {
		if err = os.MkdirAll(dir, 0o777); err != nil {
			return err
		}
	}

	pe := getPerm(0o666, perm...)
	df, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, pe)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	return err
}

// Rename renames (moves) old name to new name.
// If new name already exists and is not a directory, Rename replaces it.
func Rename(oldName, newName string) error {
	if !IsExist(oldName) {
		return nil
	}
	return os.Rename(oldName, newName)
}

// Remove removes the file or empty directory by file name.
func Remove(fileName string) error {
	if !IsExist(fileName) {
		return nil
	}
	return os.Remove(fileName)
}

// Mkdir creates a directory by file path with perm,
// along with any necessary parents, default perm is 0777.
// If the file path is already a directory, Mkdir does nothing and returns nil.
func Mkdir(filePath string, perm ...os.FileMode) error {
	pe := getPerm(0o777, perm...)
	return os.MkdirAll(filePath, pe)
}

// Deldir deletes file path and any children it contains.
// It deletes everything it can but returns the first error it encounters.
// If the file path does not exist, Deldir does nothing and returns nil.
func Deldir(filePath string) error {
	return os.RemoveAll(filePath)
}

// AbsPath returns an absolute representation of file path.
func AbsPath(filePath string) (string, error) {
	return filepath.Abs(filePath)
}

// IsAbsPath reports whether the file path is absolute.
func IsAbsPath(filePath string) bool {
	return path.IsAbs(filePath)
}

// IsExist reports whether the file path already exists.
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// IsFile reports whether the file path describes a file.
func IsFile(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// IsDir reports whether the file path describes a dictionary.
func IsDir(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

var semaphore = make(chan struct{}, 20)

func entries(dirPath string) []os.FileInfo {
	semaphore <- struct{}{}
	defer func() { <-semaphore }()
	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil
	}
	return entries
}

func walkDirSize(dirPath string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range entries(dirPath) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dirPath, entry.Name())
			go walkDirSize(subDir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func getPerm(defaultPerm os.FileMode, perm ...os.FileMode) os.FileMode {
	pe := defaultPerm
	if len(perm) > 0 {
		pe = perm[0]
	}
	return pe
}
