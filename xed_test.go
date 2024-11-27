package xed

import (
	"github.com/ddkwork/app/bindgen/clang"
	"github.com/ddkwork/app/bindgen/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"io/fs"
	"path/filepath"
	"testing"
)

func TestMergeHeader(t *testing.T) {
	filepath.Walk("kits/xed-install-base-2024-11-27-win-x86-64/include/xed", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		println(path)
		return err
	})
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
