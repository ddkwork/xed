package xed

import (
	"bytes"
	"unsafe"

	"github.com/ddkwork/golibrary/mylog"
	"golang.org/x/sys/windows"
)

func StringByCBuffer(data []byte) string {
	b := new(bytes.Buffer)
	for _, elem := range data {
		if elem == 0 {
			break
		}
		b.WriteByte(elem)
	}
	return b.String()
}

type Status int

const (
	ParseError Status = iota // 解析错误
	ParseOk                  // 解析成功
)

func (x Status) String() string {
	switch x {
	case ParseError:
		return "Error"
	case ParseOk:
		return "Ok"
	}
	return "Unknown"
}

type Parse struct {
	Address     uint64 // 指令指针（用于相对寻址）
	Bytes       []byte // 目标缓冲区
	Instruction string // 指令文本
	Error       string // 错误文本（发生错误时）
	Status      Status // 解析状态
	// todo 格式化汇编信息成一行
}

func ParseAssemble[T uint32 | uint64](address T, instruction string, callback func(text string, value T)) Parse {
	isX64 := false
	dllPath := ""
	switch any(address).(type) {
	case uint64:
		isX64 = true
		dllPath = "XEDParse64.dll"
	case uint32:
		dllPath = "XEDParse32.dll"
	}
	const (
		XedParseMaxBufSize = 256 // 这种都没必要导出，因为它只是c内存对齐需要的，不符合go的习惯
		XedParseMaxAsmSize = 16
	)
	type cXedParse struct {
		X64       bool                     // 使用 64 位指令
		Cip       uint64                   // 指令指针（用于相对寻址）
		DestSize  uint                     // 目标大小（由 XedParse 返回） todo 这里应该是 uint64 才对
		CbUnknown uintptr                  // 未知操作数回调，使用 uintptr 存储回调函数指针
		Dest      [XedParseMaxAsmSize]byte // 目标缓冲区
		Instr     [XedParseMaxBufSize]byte // 指令文本
		Error     [XedParseMaxBufSize]byte // 错误文本（发生错误时）
	}
	handle := windows.MustLoadDLL(dllPath)
	defer func() { mylog.Check(handle.Release()) }()
	type parseUnknownCallBackType func(textPtr uintptr, valuePtr uintptr) uintptr
	cbUnknown := windows.NewCallback(parseUnknownCallBackType(func(textPtr uintptr, valuePtr uintptr) uintptr {
		text := (*[XedParseMaxBufSize]byte)(unsafe.Pointer(textPtr))
		value := (*[8]byte)(unsafe.Pointer(valuePtr)) // todo 这里不用区分位数？
		mylog.Warning("unknown operand: ", StringByCBuffer(text[:]), " ---> ", T(value[0]))
		callback(StringByCBuffer(text[:]), T(value[0])) // todo 这些东西需要一个x86asm的xed文本汇编库就方便调试bug了
		return 0
	}))
	xedParse := &cXedParse{
		X64:       isX64,
		Cip:       0,
		DestSize:  0,
		CbUnknown: cbUnknown,
		Dest:      [16]byte{},
		Instr:     [256]byte{},
		Error:     [256]byte{},
	}
	copy(xedParse.Instr[:], instruction)
	proc := mylog.Check2(handle.FindProc("XEDParseAssemble"))
	ret, _, _ := proc.Call(uintptr(unsafe.Pointer(xedParse)))
	return Parse{
		Address:     uint64(address),
		Bytes:       xedParse.Dest[:xedParse.DestSize],
		Instruction: instruction,
		Error:       StringByCBuffer(xedParse.Instr[:]),
		Status:      Status(ret),
	}
}
