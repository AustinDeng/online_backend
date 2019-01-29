package main

import (
	"online_backend/package/setting"
	"online_backend/model"
	"online_backend/routes"
)

func main() {
	setting.Setup()
	model.Setup()

	routersInit := routes.InitRouter()
	routersInit.Run(":8080")
	// readTimeout := setting.ServerSetting.ReadTimeout
	// writeTimeout := setting.ServerSetting.WriteTimeout
	// endPoint := fmt.Sprintf("%d:", setting.ServerSetting.HTTPPort)
	// maxHeaderBytes := 1 << 20

	// server := &http.Server{
	// 	Addr:         endPoint,
	// 	Handler:      routersInit,
	// 	ReadTimeout:  readTimeout,
	// 	WriteTimeout: writeTimeout,
	// 	MaxHeaderBytes: maxHeaderBytes,
	// }

	// log.Printf("[info] start http server listening %s", endPoint)

	// server.ListenAndServe()
}
