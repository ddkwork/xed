package xed

import (
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/stream"

	"github.com/ddkwork/app/bindgen/clang"
	"github.com/ddkwork/app/bindgen/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestMergeHeader(t *testing.T) {
	names := stream.NewOrderedMap("", 0)
	debugIndex := 0

	var allNames []string
	filepath.Walk("kits/xed-install-base-2024-11-27-win-x86-64/include/xed", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Base(path) == "xed-interface.h" {
			index := 0
			for _, s := range stream.ToLines(path) {
				if strings.HasPrefix(s, "#include") || strings.HasPrefix(s, "# include") {
					index++
					s = strings.TrimPrefix(s, "#include")
					s = strings.TrimPrefix(s, "# include")
					split := strings.Split(s, " ")
					// mylog.Struct("", split)
					for _, s := range split {
						s = strings.TrimSpace(s)
						if strings.HasSuffix(s, `.h"`) {
							s = mylog.Check2(strconv.Unquote(s))
							names.Set(s, index)
						}
					}
				}
			}
		}
		debugIndex++
		mylog.Info(debugIndex, filepath.Base(path))
		allNames = append(allNames, filepath.Base(path))
		return err
	})
	//for _, p := range names.List() {
	//	mylog.Success(p.Key, p.Value)
	//}

	last := names.Last()
	lastIndex := last.Value

	for _, name := range allNames {
		if names.Has(name) {
			continue
		}
		lastIndex++
		names.Set(name, lastIndex)
	}
	for _, p := range names.List() {
		mylog.Success(p.Key, p.Value)
	}
}

func TestBindXed(t *testing.T) {
	TestMergeHeader(t)
	pkg := gengo.NewPackage("xed")
	path := "xed.h"
	mylog.Check(pkg.Transform("xed", &clang.Options{
		Sources: []string{path},
		// AdditionalParams: []string{},
	}),
	)
	mylog.Check(pkg.WriteToDir("tmp"))
}

func TestDisasMacho(t *testing.T)   {}
func TestEx1(t *testing.T)          {}
func TestExCpuid(t *testing.T)      {}
func TestXed(t *testing.T)          {}
func TestExAgen(t *testing.T)       {}
func TestMin(t *testing.T)          {}
func TestTester(t *testing.T)       {}
func TestAvltree(t *testing.T)      {}
func TestDisasFilter(t *testing.T)  {}
func TestEnc22(t *testing.T)        {}
func TestEx3(t *testing.T)          {}
func TestEx8(t *testing.T)          {}
func TestDecPrint(t *testing.T)     {}
func TestDisasElf(t *testing.T)     {}
func TestEx5Enc(t *testing.T)       {}
func TestNmSymtab(t *testing.T)     {}
func TestReps(t *testing.T)         {}
func TestDllDiscovery(t *testing.T) {}
func TestDot(t *testing.T)          {}
func TestExIld2(t *testing.T)       {}
func TestFindSpecial(t *testing.T)  {}
func TestEx6(t *testing.T)          {}
func TestEx9Patch(t *testing.T)     {}
func TestExIld(t *testing.T)        {}
func TestSize(t *testing.T)         {}
func TestUtil(t *testing.T)         {}
func TestDotPrep(t *testing.T)      {}
func TestEnc21(t *testing.T)        {}
func TestEnc23(t *testing.T)        {}
func TestSymbolTable(t *testing.T)  {}
func TestAsmParseMain(t *testing.T) {}
func TestDisasHex(t *testing.T)     {}
func TestEx4(t *testing.T)          {}
func TestTables(t *testing.T)       {}
func TestAsmParse(t *testing.T)     {}
func TestDisasRaw(t *testing.T)     {}
func TestEncLang(t *testing.T)      {}
func TestEx7(t *testing.T)          {}
