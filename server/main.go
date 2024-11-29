package main

import (
	"embed"
	"stellar/cmd"
	"stellar/common"
)

//go:embed stellar.yaml
var fs embed.FS // Go 1.16 版本之后提供的将静态资源打包的方法，写法固定，可以将目录也打包

func main() {
	common.FS = fs
	cmd.Execute()
}
