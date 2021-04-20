package main

/*
默认从$GOPATH/src下导入
➜  cd $GOPATH
➜  pwd
/Users/luzihang/go
➜  cd src
*/
import (
	"源代码/5.init/lib1"
	"源代码/5.init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
