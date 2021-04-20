package main

/*
默认从$GOPATH/src下导入,拷贝包代码到src文件夹下
➜  cd $GOPATH
➜  pwd
/Users/luzihang/go
➜  cd src
*/
import (
	"GolangStudy/5.init/lib1"
	"GolangStudy/5.init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
