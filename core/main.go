package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

var cliFlag string

func Init() {
	flag.StringVar(&cliFlag, "path", "", "Just for demo")
	flag.Parse()
}

func main() {
	Init()
	fmt.Println(cliFlag)
	r := gin.Default()

	root := cliFlag
	r.GET("/*filepath", func(c *gin.Context) {
		path := c.Param("filepath")
		fmt.Println(path)
		if strings.HasPrefix(path, "/assets") {
			imgPath := root + path
			fmt.Println(imgPath)
			if Exists(imgPath) {
				c.File(imgPath)
			} else {
				c.File(root + "/assets/blockchains/kcc/img.png")
			}
		}
	})

	r.Run(":9099") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
