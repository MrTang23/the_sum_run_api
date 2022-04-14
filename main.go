package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sun_run_api/handle_run_request"
	"sun_run_api/local_log"
	"time"
)

func main() {
	local_log.Output_Info("服务器正在启动中...")
	log.Println("服务器启动中...")
	//开机特效
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
		fmt.Printf("\r%.0f%%[%s]", float64(i)/49*100, h)
	}
	//启动服务
	http.HandleFunc("/run", handle_run_request.Try_run)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		local_log.Output_Error(fmt.Sprintln("服务器启动失败,Listen&serve error: ", err))
	}
}
