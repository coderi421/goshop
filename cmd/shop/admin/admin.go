package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/coderi421/goshop/app/shop/admin"
)

func main() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand.New(randSrc)

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	admin.NewApp("admin-server").Run()
}
