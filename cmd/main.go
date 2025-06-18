package main

import (
	"fmt"
	"os"

	"github.com/rlaope/ainit/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("인수를 하나 넣어주세요.")
		return
	}

	// TODO 지금은 content 자연어 하나만 받지만 project name 받고 그걸로 mkdir 하고 go project settings 하도록
	// TODO 1. ai로 readme 만들기
	// TODO 2. go project .gitignore, .gitattributes, Makefile, cmd/main.go 에 hello world 출력하는거 자동으로 만들어지도록

	arg := os.Args[1]
	pkg.GenerateProejct(arg)
}