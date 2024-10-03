//go:build ignore
// +build ignore

// 测试接受信息
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

const (
	url1 = "http://vcard.ccnu.edu.cn/index.php/index/index/virtualreport.html"
	url2 = "http://vcard.ccnu.edu.cn/index.php/index/index/cardreport.html"
)

func main() {
	//StudentId := "2023214557"
	//grade := StudentId[:4]
	//fmt.Println("StudentId:", grade)
	test()
}

func test() {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 添加表单字段
	_ = writer.WriteField("key", "88Gar7ySd4TeT9Q7O510h9Q6e0pacL2lZ3kmMRwxL7zIU0Y27yvD4IDz1EUYRFYp4K129yqy97HJo4XWLMggipTBH8Hxwp%3DOS4jCW9gLV4yfr")
	_ = writer.WriteField("start_time", "2024 09")

	writer.Close()

	// 创建请求
	req, err := http.NewRequest("POST", url2, &buf)
	if err != nil {
		panic(err)
	}

	// 设置请求头
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, _ := ioutil.ReadAll(resp.Body)
	var response Response
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(response.Data.List)
	//fmt.Println(string(body))

}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []Transaction `json:"list"`
		Sql  string        `json:"sql"`
	} `json:"data"`
}

type Transaction struct {
	SMT_TIMES string `json:"SMT_TIMES"`
	//SMT_INMONEY      string `json:"SMT_INMONEY"`
	//SMT_OUTMONEY     string `json:"SMT_OUTMONEY"`
	//MONEY            string `json:"MONEY"`
	//SMT_DEALDATETIME string `json:"SMT_DEALDATETIME"`  //创建时间
	SMT_ORG_NAME string `json:"SMT_ORG_NAME"` //窗口
	SMT_DEALNAME string `json:"SMT_DEALNAME"` //消费方式
	//BeforeMoney  string `json:"before_money"`
	AfterMoney string `json:"after_money"`
	Money      string `json:"money"`
}

//{
//"SMT_TIMES": "902",
//"SMT_INMONEY": "2350",
//"SMT_OUTMONEY": "950",
//"MONEY": "-1400",
//"SMT_DEALDATETIME": "2024-09-26 12:04:23",
//"SMT_ORG_NAME": "锅巴饭",
//"SMT_DEALNAME": "消费",
//"before_money": "23.50",
//"after_money": "9.50",
//"money": "-14.00"
//},
//
//{
//"SMT_TIMES": "0",
//"SMT_INMONEY": "1300",
//"SMT_OUTMONEY": "0",
//"MONEY": "-1300",
//"SMT_DEALDATETIME": "2024-09-26 18:21:16",
//"SMT_ORG_NAME": "铁板饭",
//"SMT_DEALNAME": "第三方支付消费",
//"before_money": "13.00",
//"after_money": "0.00",
//"money": "-13.00"
//},
