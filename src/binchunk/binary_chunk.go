package binchunk

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