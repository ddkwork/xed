package xed

import (
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/safemap"

	"github.com/ddkwork/app/bindgen/clang"
	"github.com/ddkwork/app/bindgen/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestName(t *testing.T) {
	const ignore = `// +build ignore

`
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, "cmake-build-debug") {
			return err
		}
		if filepath.Ext(path) == ".c" || filepath.Ext(path) == ".cpp" {
			b := stream.NewBuffer(path)
			if !strings.HasPrefix(b.String(), ignore) {
				b.ReplaceAll(strings.TrimSuffix(ignore, "\n"), "")
				b.InsertString(0, ignore).ReWriteSelf()
			}
		}
		return err
	})
}

func TestMergeHeader(t *testing.T) {
	names := new(safemap.M[string, int])
	debugIndex := 0
	includePath := "kits/xed-install-base-2024-11-27-win-x86-64/include/xed"
	var allNames []string
	filepath.Walk(includePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Base(path) == "xed-interface.h" {
			index := 0
			for _, s := range stream.ReadFileToLines(path) {
				if stream.IsIncludeLine(s) {
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
		// mylog.Info(debugIndex, filepath.Base(path))
		allNames = append(allNames, filepath.Base(path))
		return err
	})
	//for _, p := range names.List() {
	//	mylog.Success(k, p.Value)
	//}
	lastIndex := names.LastValue()

	for _, name := range allNames {
		if names.Has(name) {
			continue
		}
		lastIndex++
		names.Set(name, lastIndex)
	}
	b := stream.NewBuffer("")
	b.WriteStringLn("#define XED_DLL")
	b.WriteStringLn("#define XED_WINDOWS")
	sep := "------------------------------------------"
	for k := range names.Range() {
		// mylog.Success(k, p.Value)
		b.WriteStringLn("//" + sep + "start " + k + sep)
		// b.NewLine()
		for s := range stream.ReadFileToLines(filepath.Join(includePath, k)) {
			if stream.IsIncludeLine(s) {
				continue
			}
			b.WriteStringLn(s)
		}
		b.WriteStringLn("//" + sep + "end " + k + sep)
		b.NewLine()
	}
	stream.WriteTruncate("xed_merged.h", b.Bytes())
	clang.CheckHeadFile("xed_merged.h")
}

func TestBindXed(t *testing.T) {
	path := "xed_merged.h"
	path = "D:\\workspace\\workspace\\debuger\\xed\\kits\\xed-install-base-2024-11-27-win-x86-64\\include\\xed\\combined_header.h"
	// TestMergeHeader(t)
	clang.CheckHeadFile("D:\\workspace\\workspace\\debuger\\xed\\kits\\xed-install-base-2024-11-27-win-x86-64\\include\\xed\\combined_header.h")
	return
	pkg := gengo.NewPackage("xed")
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
