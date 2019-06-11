package compile

//var Id int = 9527
// 用go tool compile， -S代表输出汇编格式。但并非是能完整运行的汇编。编译器的汇编，在某些特殊标识是手工的汇编在不能使用的。防止编译器的规则被打破
// go tool compile -S compile.go
// 编写汇编compile.s
// 修改代码声明汇编定义的Id，注释var Id = 9527
// int 在mac下 是int64 = 8字节
// compile.go 和compile.s 相当于c中的.h和.c文件
// 在main进行引入使用汇编的声明，引入方式和平常包引用方式一致import 包路径
var Id int
