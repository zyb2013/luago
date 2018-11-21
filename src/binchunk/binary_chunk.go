package binchunk

const (
	LUA_SIGNATURE = "\x1bLua"
	LUAC_VERSION = 0x53
	LUAC_FORMAT = 0
	LUAC_DATA = "\x19\x93\r\n\x1a\n"
	CINT_SIZE = 4
	CSZIET_SIZE = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE = 8
	LUAC_INT = 0x5678
	LUAC_NUM = 370.5
)

const (
	TAG_NIL = 0x00
	TAG_BOOLEAN = 0x01
	TAG_NUMBER = 0x03
	TAG_INTEGER = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR = 0x14
}

// 二进制chunk
type binaryChunk struct {
	header // 头部
	upvaluesSize byte // 主函数upvalue数量
	mainFunc *Prototype	// 主函数原型
}

// 头部
type header struct {
	signature [4]byte // 魔数 0x1B4C7561 魔数主要起快速识别文件格式的作用
	version byte // 版本号
	format byte // 格式号
	luacData [6]byte
	cintSize byte
	sizetSize byte
	instructionSize byte
	luaIntegerSize byte
	luaNumberSize byte
	luacInt int64
	luacNum float64
}

// 函数原型
type Prototype struct {
	Source string
	LineDefined uint32
	LastLineDefined uint32
	NumParams byte
	IsVararg byte
	MaxStackSize byte // 寄存器数量
	Code []uint32 // 指令表
	Constants []interface{} // 常量表
	Upvalues []Upvalue
	Protos []*Prototype
	LineInfo []uint32
	LocVars []LocVar
	UpvalueNames []string
}