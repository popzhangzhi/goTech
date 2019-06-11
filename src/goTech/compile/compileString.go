package compile

//var Name = "zhangzhi"
// go tool compile -S compileString.go
// 得到
// go.string."zhangzhi" SRODATA dupok size=8
//        0x0000 7a 68 61 6e 67 7a 68 69                          zhangzhi
//"".Name SDATA size=16
//        0x0000 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
//        rel 0+8 t=1 go.string."zhangzhi"+0
// 字符串的汇编币整形来的复杂。 go.string."zhangzhi" 字符串标识。 SRODATA 只读内存标识 dupok多个字符串标识（go.string."zhangzhi")
// 只使用一个
// "".Name 只读标识 SDATA 。是16字节。对应的结构 reflect.StringHeader结构体类型
//type reflect.StringHeader struct {
//    Data uintptr
//    Len  int
//}
// 前8字节对应着字符串的指针地址，后8字节对应着字符串长度。
//汇编最后一行得有空行,汇编如下
// GLOBL ·NameData(SB),$8
//DATA ·NameData(SB)/8,$"zhangzhi"
//
//GLOBL ·Name(SB),$16
//DATA ·Name+0(SB)/8,$·NameData(SB)
//DATA ·Name+8(SB)/1,$0x08

//因为汇编底层是引用NameData.如果NameData改变了。Name也变化了，这和字符串的只读定义是冲突的
//所以修改成当前汇编，见.s文件
//在新的结构中，Name符号对应的内存从16字节变为24字节，多出的8个字节存放底层的“zhangzhi”字符串。·Name符号前16个字节依然对应reflect.StringHeader结构体
// ：Data部分对应$·Name+16(SB)，表示数据的地址为Name符号往后偏移16个字节的位置；Len部分依然对应6个字节的长度。
// 这是C语言程序员经常使用档技巧。

//var NameData [8]byte
var Name string
