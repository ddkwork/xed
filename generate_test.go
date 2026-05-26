package xed

import (
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:  "include",
		OutputDir:   ".",
		PackageName: "xed",
		HeaderOrder: []string{"xed-interface.h"},
		BindDll:     true,
		DllName:     "xed.dll",
		Predefined: `
#define XED_ENCODER
#define XED_ENCODE_ORDER_MAX_OPERANDS 5
#define XED_ENCODE_ORDER_MAX_ENTRIES 35
#define stderr ((void*)0)
int fprintf(void*, const char*, ...);
int fflush(void*);
void abort(void);
`,
		ExtraIncludeDirs: []string{
			"include",
			"include/public/xed",
		},
		DllFuncFilter: func(name string) bool {
			return strings.HasPrefix(name, "xed_")
		},
	}})
}
