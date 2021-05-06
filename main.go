package main

import (
	"Filestore-SERVER/handler"
	"fmt"
	"net/http"
)

func main() {
	// 静态资源处理
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	//动态路由设置
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download",handler.DownloadHandler)
	http.HandleFunc("/file/update",handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)
	// 秒传接口
	http.HandleFunc("/file/fastupload", handler.HTTPInterceptor(
		handler.TryFastUploadHandler))

	//用户相关接口
	http.HandleFunc("/user/signup",handler.SignupHandler)
	http.HandleFunc("/user/signin",handler.SignInHandler)
	http.HandleFunc("/user/info",handler.HTTPInterceptor(handler.UserInfoHandler))


	// 监听端口
	fmt.Println("上传服务正在启动, 监听端口:8080...")
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Printf("Fail to start server,err:%s", err.Error())
	}
}

