package xed

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

type Xed struct{}

//go:embed xed.dll
var dllBytes []byte

var (
	dll       *windows.LazyDLL
	procCache = make(map[string]*windows.LazyProc)
)

func init() {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	dir = filepath.Join(dir, "xed_dll_cache")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		panic(err)
	}
	if err := windows.SetDllDirectory(dir); err != nil {
		panic(err)
	}
	sha := sha256.Sum256(dllBytes)
	dllName := fmt.Sprintf("%s-%s.dll", "xed", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, dllBytes, 0o644); err != nil {
			panic(err)
		}
	}
	dll = windows.NewLazyDLL(dllName)
}

func getProc(name string) *windows.LazyProc {
	if p, ok := procCache[name]; ok {
		return p
	}
	p := dll.NewProc(name)
	procCache[name] = p
	return p
}
