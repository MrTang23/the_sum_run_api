package handle_run_request

import (
	"encoding/json"
	. "fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sun_run_api/local_log"
	"sun_run_api/utils"
	"time"
)

//输入应跑的距离
func justRun(imeiCode string, distance string, reciever string, u_name string) bool {
	ios := "http://client4.aipao.me/api"
	iosb := false
	println("个人Imei：" + imeiCode)
	println("跑步距离：" + distance)

	randomGenerateTable()
	//第一次GET请求 获取token请求
	//安卓机型
	req, _ := http.NewRequest("GET", apiRoot+"/%7Btoken%7D/QM_Users"+
		"/Login_AndroidSchool?IMEICode="+imeiCode, nil)
	req.Header.Add("version", appVersion)
	resInfo, _ := (&http.Client{}).Do(req)
	dataInfo, _ := ioutil.ReadAll(resInfo.Body)
	// log.Println(string(dataInfo))
	resInfo.Body.Close()

	returnData := &returnInfo{}
	_ = json.Unmarshal(dataInfo, returnData)
	//苹果机型
	if !returnData.Success {
		req, _ := http.NewRequest("GET", ios+"/%7Btoken%7D/QM_Users"+
			"/LoginSchool?IMEICode="+imeiCode, nil)
		req.Header.Add("Version", appVersionForIOS)
		resInfo, _ := (&http.Client{}).Do(req)
		dataInfo, _ := ioutil.ReadAll(resInfo.Body)
		//若此处失败 则imei失效
		log.Println(string(dataInfo))
		resInfo.Body.Close()
		_ = json.Unmarshal(dataInfo, returnData)
		iosb = true
	}
	url := apiRoot
	if iosb {
		url = ios
	}

	UserId := string(Sprintf("%d", returnData.Data.UserId))

	timespan := Sprintf("%d", time.Now().UnixNano()/1e6)
	nonce := Sprintf("%d", 100000+rand.Intn(9900000)) //Nonce是或Number once的缩写，在密码学中Nonce是一个只被使用一次的任意或非重复的随机数值。 在加密技术中的初始向量和加密散列函数都发挥着重要作用，在各类验证协议的通信应用中确保验证信息不被重复使用以对抗重放攻击(Replay Attack)。
	sign := strings.ToUpper(MD5(returnData.Data.Token + nonce + timespan + UserId))

	time.Sleep(1 * time.Second)

	runTime, runStep := randomGenerateInfo() //2400

	var x, _ = strconv.Atoi(distance)
	var y = x + rand.Intn(5) //生成跑步步数
	runDistance := strconv.Itoa(y)

	//第二次GET请求 提交跑步开始信息
	client := &http.Client{}

	requestRun, _ := http.NewRequest("GET", url+"/"+
		returnData.Data.Token+"/QM_Runs/SRS?S1="+longtitude+"&S2="+latitute+"&S3="+distance, nil)

	requestRun.Header.Add("nonce", nonce)       //第一次请求获得
	requestRun.Header.Add("timespan", timespan) //第一次请求获得
	requestRun.Header.Add("sign", sign)         //第一次请求获得
	requestRun.Header.Add("version", appVersion)
	requestRun.Header.Add("Accept", "text/html")
	requestRun.Header.Add("User-Agent", UserAgent)
	requestRun.Header.Add("Accept-Encoding", "gzip")
	requestRun.Header.Add("Connection", "Keep-Alive")
	//发起请求
	resRun, _ := client.Do(requestRun)

	infoData, _ := ioutil.ReadAll(resRun.Body)
	// log.Println(string(infoData))
	resRun.Body.Close()

	returndata := &returnRun{}

	//json转结构体·
	_ = json.Unmarshal(infoData, returndata)

	//第三期GET请求 获取跑步结束信息
	resEnd, _ := http.Get(url + "/" + returnData.Data.Token + "/QM_Runs/ES?S1=" +
		returndata.Data.RunId + "&S4=" + encrypt(runTime) + "&S5=" + encrypt(runDistance) +
		"&S6=" + returndata.Data.Routes + "&S7=1&S8=" + Sprintf("%s", table) + "&S9=" + encrypt(runStep))
	dataEnd, _ := ioutil.ReadAll(resEnd.Body)

	resEnd.Body.Close()
	returnEnd := &returnEnd{}
	_ = json.Unmarshal(dataEnd, returnEnd)

	//处理跑步结果
	if returnEnd.Success {
		//调用邮件函数
		local_log.Output_Info("用户：" + u_name + " 跑步成功.")
		utils.Mail(reciever, true)
		return true
	} else {
		local_log.Output_Info("用户：" + u_name + " 跑步失败.")
		utils.Mail(reciever, false)
		return false
	}
}
