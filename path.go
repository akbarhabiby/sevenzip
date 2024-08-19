package sevenzip

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type muString struct {
	val string
	sync.Mutex
}

func (ss *muString) Get() string {
	ss.Lock()
	defer ss.Unlock()
	return ss.val
}

func (ss *muString) Set(s string) {
	ss.Lock()
	ss.val = s
	ss.Unlock()
}

var binPath muString
var lookPath = exec.LookPath

func SetPath(path string) {
	binPath.Set(path)
}

func GetPath() string {
	return binPath.Get()
}

func findPath() (path string, err error) {
	const exe = "7z"
	path = GetPath()
	if path != "" {
		// 7z has already been found, return
		return
	}
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}
	path, err = lookPath(filepath.Join(exeDir, exe))
	if err == nil && path != "" {
		SetPath(path)
		return
	}
	path, err = lookPath(exe)
	if err == nil && path != "" {
		SetPath(path)
		return
	}
	dir := os.Getenv("7Z_PATH")
	if dir == "" {
		return "", fmt.Errorf("%s not found", exe)
	}
	path, err = lookPath(filepath.Join(dir, exe))
	if err == nil && path != "" {
		SetPath(path)
		return
	}
	return "", fmt.Errorf("%s not found", exe)
}
