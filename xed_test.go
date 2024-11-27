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
		return err
	})
	for _, p := range names.List() {
		mylog.Success(p.Key, p.Value)
	}
}

// xed-address-width-enum.h
// xed-agen.h
// xed-attribute-enum.h
// xed-attributes.h
// xed-build-defines.h
// xed-category-enum.h
// xed-chip-enum.h
// xed-chip-features.h
// xed-common-defs.h
// xed-common-hdrs.h
// xed-convert-table-init.h
// xed-cpuid-group-enum.h
// xed-cpuid-rec-enum.h
// xed-cpuid-rec.h
// xed-decode.h
// xed-decoded-inst-api.h
// xed-decoded-inst.h
// xed-disas.h
// xed-encode-check.h
// xed-encode-direct.h
// xed-encode.h
// xed-encoder-gen-defs.h
// xed-encoder-hl.h
// xed-encoder-iforms.h
// xed-error-enum.h
// xed-exception-enum.h
// xed-extension-enum.h
// xed-flag-action-enum.h
// xed-flag-enum.h
// xed-flags.h
// xed-format-options.h
// xed-gen-table-defs.h
// xed-get-time.h
// xed-iclass-enum.h
// xed-iform-enum.h
// xed-iform-map.h
// xed-iformfl-enum.h
// xed-ild-enum.h
// xed-ild.h
// xed-immdis.h
// xed-immed.h
// xed-init-pointer-names.h
// xed-init.h
// xed-inst.h
// xed-interface.h
// xed-isa-set-enum.h
// xed-isa-set.h
// xed-machine-mode-enum.h
// xed-mapu-enum.h
// xed-nonterminal-enum.h
// xed-operand-accessors.h
// xed-operand-action-enum.h
// xed-operand-action.h
// xed-operand-convert-enum.h
// xed-operand-ctype-enum.h
// xed-operand-ctype-map.h
// xed-operand-element-type-enum.h
// xed-operand-element-xtype-enum.h
// xed-operand-enum.h
// xed-operand-storage.h
// xed-operand-type-enum.h
// xed-operand-values-interface.h
// xed-operand-visibility-enum.h
// xed-operand-width-enum.h
// xed-patch.h
// xed-portability.h
// xed-print-info.h
// xed-reg-class-enum.h
// xed-reg-class.h
// xed-reg-enum.h
// xed-reg-role-enum.h
// xed-rep-prefix.h
// xed-state.h
// xed-syntax-enum.h
// xed-types.h
// xed-util.h
// xed-version.h
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
