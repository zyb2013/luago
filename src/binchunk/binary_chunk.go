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

type binaryChunk struct {
	header // 头部
	upvaluesSize byte // 主函数upvalue数量
	mainFunc *Prototype	// 主函数原型
}

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

type Prototype struct {

}