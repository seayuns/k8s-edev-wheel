package main

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func main() {
	engine := gin.Default()
	gin.SetMode(gin.DebugMode)
	err := engine.Run()
	if err != nil {
		klog.Fatal("run server error, ", err)
		return
	}
}
