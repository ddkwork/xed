package xed

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

const cmakeListName = "CMakeLists.txt"

type (
	examples struct {
		cFilePath  string
		name       string
		cmakeLists string
	}
)

func TestMakeExampleCmakePackages(t *testing.T) {
	projects := make([]examples, 0)
	filepath.Walk("kits", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}
		if strings.Contains(path, "cmake-build-debug") {
			return err
		}
		if filepath.Ext(path) != ".c" {
			return err
		}
		projects = append(projects, examples{
			cFilePath:  path,
			name:       stream.ToCamelUpper(stream.BaseName(path)),
			cmakeLists: "",
		})
		return err
	})
	subNames := map[string]string{}
	for _, project := range projects {
		switch project.name {
		case "xed", "Xed":
		default:
			project.name = strings.TrimPrefix(project.name, "xed")
			project.name = strings.TrimPrefix(project.name, "Xed")
			project.name = strings.TrimPrefix(project.name, "Examples")
		}
		project.name = strings.ReplaceAll(project.name, "Asmparse", "AsmParse")
		mylog.Success(project.name, project.cFilePath)
		subNames[project.name] = ""
		g := stream.NewGeneratedFile()
		g.P(`
cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
`)
		g.P("project(", project.name, " C)")
		g.P()

		g.P("include_directories(include)")
		g.P("link_directories(${CMAKE_SOURCE_DIR})")

		g.P("add_executable(", project.name)
		g.P(filepath.Base(project.cFilePath))
		g.P("xed-examples-util.c")
		g.P(")")

		g.P("target_link_libraries(", project.name, " xed)")
		projectRoot := "D:\\workspace\\workspace\\debuger\\xed\\kits\\xed-install-base-2024-11-27-win-x86-64\\examples"
		projectRoot = filepath.Join(projectRoot, "examples", project.name)
		stream.WriteTruncate(filepath.Join(projectRoot, cmakeListName), g.Bytes())
		stream.WriteTruncate(filepath.Join(projectRoot, filepath.Base(project.cFilePath)), stream.NewBuffer(project.cFilePath))
		stream.WriteTruncate(filepath.Join(projectRoot, "xed-examples-util.c"), stream.NewBuffer("kits/xed-install-base-2024-11-27-win-x86-64/examples/xed-examples-util.c"))
		stream.WriteTruncate(filepath.Join(projectRoot, "xed-examples-util.h"), stream.NewBuffer("kits/xed-install-base-2024-11-27-win-x86-64/examples/xed-examples-util.h"))
		stream.WriteTruncate(filepath.Join(projectRoot, "xed.lib"), stream.NewBuffer("kits/xed-install-base-2024-11-27-win-x86-64/lib/xed.lib"))
		stream.WriteTruncate(filepath.Join(projectRoot, "xed-ild.lib"), stream.NewBuffer("kits/xed-install-base-2024-11-27-win-x86-64/lib/xed-ild.lib"))
		s := "package " + project.name + `
import (
	"testing"

	"github.com/ddkwork/c2go"
)

func TestName(t *testing.T) {
	c2go.Run()
}
`
		stream.WriteTruncate(filepath.Join(projectRoot, project.name+"_test.go"), s)
		stream.CopyDir(filepath.Join(projectRoot, "include"), "kits/xed-install-base-2024-11-27-win-x86-64/include")
	}

	gXedUintTest := stream.NewGeneratedFile()
	gSub := stream.NewGeneratedFile()
	gXedUintTest.P("package xed")
	gXedUintTest.P("import \"testing\"")
	gSub.P("add_subdirectory(")
	for name := range subNames {
		gSub.P(name)
		gXedUintTest.P("func Test", name, "(t *testing.T) {}")
	}
	gSub.P(")")
	stream.WriteTruncate("kits/xed-install-base-2024-11-27-win-x86-64/examples/examples/CMakeLists.txt", gSub.Bytes())
	// stream.WriteGoFile("xed_test.go", gXedUintTest.Bytes())
}

const parseTest = `

asmparse -q vaddpd ymm1, ymm2, ymm3
asmparse -q -16 mov ax, dx
asmparse -q vaddpd ymm1{k3}, ymm2, ymm3
asmparse -q vaddpd ymm1{k2}{z}, ymm2, ymm3
asmparse -q vaddpd ymm1{k2}{z}, ymm2, ymmword [ebx]
asmparse -q vaddpd ymm1{k2}{z}, ymm2, ymmword ptr [ebx]
asmparse -q fdiv st(0), st(1)
asmparse -q cvtpi2pd xmm4, mm3
asmparse -q -64 mov cr0, rbx
asmparse -q mov ebx, dr1
asmparse -q call_near eax
asmparse -q call eax
asmparse -q jmp eax
asmparse -q rep cmpsb
asmparse -q repe cmpsb
asmparse -q repne cmpsb
asmparse -q lock  adc  dword [ebx], eax
asmparse -q lock cmpxchg dword ptr [ebx], esi
asmparse -q cmpxchg dword ptr [ebx], esi
asmparse -q lock mov dword ptr [ebx], esi
asmparse -q lock mov edi, dword ptr [ecx]
asmparse -q call far 0x1234:0x10dedead
asmparse -q call 0x1234:0x10dedead

asmparse -q mov dword ptr ds:[edi+5E8],eax
asmparse  mov dword ptr ds:[edi+5E8],eax

asmparse -q call 7FFC0A2B66B0
asmparse -q jmp 7FFC0A36B4C3
asmparse -q mov rax,qword ptr ss:[rsp+40]
asmparse -q or dword ptr ds:[rax+68],r14d
asmparse -q mov rax,qword ptr ss:[rsp+40]
asmparse -q mov rcx,qword ptr ds:[rax+30]
asmparse -q mov qword ptr ds:[7FFC0A4143A8],rcx
asmparse -q call 7FFC0A3093C8
asmparse -q mov ebx,eax
asmparse -q test eax,eax
asmparse -q jns 7FFC0A36B491
asmparse -q mov dword ptr ss:[rsp+28],eax
asmparse -q lea r8,qword ptr ds:[7FFC0A3DD270]
asmparse -q lea rax,qword ptr ds:[7FFC0A3DD148]
asmparse -q xor r9d,r9d
asmparse -q mov edx,DF2
asmparse -q mov qword ptr ss:[rsp+20],rax
asmparse -q lea rcx,qword ptr ds:[7FFC0A3C6C08]
asmparse -q call 7FFC0A2B66B0
asmparse -q jmp 7FFC0A36B4C3
pause`

func TestParseAssemble64(t *testing.T) {
	assemble64 := ParseAssemble[uint64](0, " mov rax,qword ptr ss:[rsp+40]", func(text string, value uint64) {
		mylog.Info(text, value)
	})
	mylog.Struct("assemble64", assemble64)

	assemble64 = ParseAssemble[uint64](0, " or dword ptr ds:[rax+68],r14d ", func(text string, value uint64) {
		mylog.Info(text, value)
	})
	mylog.Struct("assemble64", assemble64)

	assemble64 = ParseAssemble[uint64](0, "lea r8,qword ptr ds:[7FFC0A3DD270]", func(text string, value uint64) {
		mylog.Info(text, value)
	})
	mylog.Struct("assemble64", assemble64) // 失败了，感觉还是不够稳定

	// 00007FFC0A36B436 | E8 75B2F4FF              | call ntdll.7FFC0A2B66B0                 |
	// 00007FFC0A36B43B | E9 83000000              | jmp ntdll.7FFC0A36B4C3                  |

	// 00007FFC0A36B440 | 48:8B4424 40             | mov rax,qword ptr ss:[rsp+40]           | passed
	// 00007FFC0A36B445 | 44:0970 68               | or dword ptr ds:[rax+68],r14d           | passed
	// 00007FFC0A36B449 | 48:8B4424 40             | mov rax,qword ptr ss:[rsp+40]           | [rsp+40]:_fltused+1F78
	// 00007FFC0A36B44E | 48:8B48 30               | mov rcx,qword ptr ds:[rax+30]           | rcx:NtQueryInformationThread+14
	// 00007FFC0A36B452 | 48:890D 4F8F0A00         | mov qword ptr ds:[7FFC0A4143A8],rcx     | rcx:NtQueryInformationThread+14
	// 00007FFC0A36B459 | E8 6ADFF9FF              | call ntdll.7FFC0A3093C8                 |
	// 00007FFC0A36B45E | 8BD8                     | mov ebx,eax                             |
	// 00007FFC0A36B460 | 85C0                     | test eax,eax                            |
	// 00007FFC0A36B462 | 79 2D                    | jns ntdll.7FFC0A36B491                  |
	// 00007FFC0A36B464 | 894424 28                | mov dword ptr ss:[rsp+28],eax           |

	// 00007FFC0A36B468 | 4C:8D05 011E0700         | lea r8,qword ptr ds:[7FFC0A3DD270]      | r8:&"吚x\n€|$@", 00007FFC0A3DD270:"LdrpGetProcApphelpCheckModule"

	// 00007FFC0A36B46F | 48:8D05 D21C0700         | lea rax,qword ptr ds:[7FFC0A3DD148]     | 00007FFC0A3DD148:"Getting the shim engine exports failed with status 0x%08lx\n"
	// 00007FFC0A36B476 | 45:33C9                  | xor r9d,r9d                             |
	// 00007FFC0A36B479 | BA F20D0000              | mov edx,DF2                             |
	// 00007FFC0A36B47E | 48:894424 20             | mov qword ptr ss:[rsp+20],rax           |
	// 00007FFC0A36B483 | 48:8D0D 7EB70500         | lea rcx,qword ptr ds:[7FFC0A3C6C08]     | rcx:NtQueryInformationThread+14, 00007FFC0A3C6C08:"minkernel\\ntdll\\ldrinit.c"
	// 00007FFC0A36B48A | E8 21B2F4FF              | call ntdll.7FFC0A2B66B0                 |
	// 00007FFC0A36B48F | EB 32                    | jmp ntdll.7FFC0A36B4C3                  |
}

func TestParseAssemble32(t *testing.T) {
	assemble32 := ParseAssemble[uint32](0, "mov dword ptr ds:[edi+5E8],eax", func(text string, value uint32) {
		mylog.Info(text, value)
	})
	mylog.Struct("assemble32", assemble32)
}
