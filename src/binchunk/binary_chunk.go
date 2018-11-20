package binchunk

type binaryChunk struct {
	header						// 头部
	upvaluesSize 	byte		// 主函数upvalue数量
	mainFunc 		*Prototype	// 主函数原型
}

type header struct {

}

type Prototype struct {

}