package main

import (
	"seal-x-polarhare-backend/internal/provider"
	"seal-x-polarhare-backend/internal/router"
	"seal-x-polarhare-backend/pkg/log"
)

func Init() {
	provider.Init()
	log.Info("所有模块初始化完成...")
}

func main() {
	Init()
	r := router.SetupRoutes()

	log.Info("服务器即将启动于 :8080")
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
