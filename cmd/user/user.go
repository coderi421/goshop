package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/coderi421/goshop/app/user/srv"
)

func main() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand.New(randSrc)

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// 通过 NewApp 创建一个 app.App 对象
	srv.NewApp("user-service").Run()
	// 通过 Run 方法启动服务
}
