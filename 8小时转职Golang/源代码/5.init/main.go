package main

/*
默认从$GOPATH/src下导入,拷贝包代码到src文件夹下
➜  cd $GOPATH
➜  pwd
/Users/luzihang/go
➜  cd src
*/
import (
	//"GolangStudy/5.init/lib1"
	//"GolangStudy/5.init/lib2"

	// 匿名导包,如果main函数中不使用，不会报错；无法使用当前包的方法，但是会执行当前包内部的init()方法
	_ "GolangStudy/5.init/lib1"

	// 起一个别名导包
	//mylib2 "GolangStudy/5.init/lib2"

	// (不建议使用：如果各个包有相同方法名称，会出现歧义)
	//相当于，lib2包里的全部方法，导入本包的作用中，lib2中的全部方法可以直接使用API来调用，不需要lib2.API来调用
	. "GolangStudy/5.init/lib2"
)

func main() {
	//lib1.Lib1Test()
	//lib2.Lib2Test()

	// 起一个别名导包
	//mylib2.Lib2Test()

	// 相当于，lib2包里的方法名称， 属于当前包里的方法名
	Lib2Test()
}
