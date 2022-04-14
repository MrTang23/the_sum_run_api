package handle_run_request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sun_run_api/local_log"
)

func Try_run(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//请求方法错误
		w.Header().Set("Error:", "Please_Use_Post")
		w.WriteHeader(302)

	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Fatal("ParseForm err: ", err)
		}
		var req_map_data map[string][]map[string]string //解析post的body
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &req_map_data)

		u_name := fmt.Sprint(req_map_data["data"][0]["username"])
		u_code := fmt.Sprint(req_map_data["data"][0]["usercode"])
		u_distance := fmt.Sprint(req_map_data["data"][0]["userdistance"])
		u_mail := fmt.Sprint(req_map_data["data"][0]["usermail"])

		local_log.Output_Info("收到来自用户：" + u_name + ",跑" + u_distance + "米的请求.")
		fmt.Println("收到来自用户：" + u_name + ",跑" + u_distance + "米的请求.")
		justRun(u_code, u_distance, u_mail, u_name)
		w.WriteHeader(402)
	} else {
		//未知请求
		w.WriteHeader(404)
	}
}

