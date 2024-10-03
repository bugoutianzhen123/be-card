// 爬虫
package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asynccnu/be-card/domain"
	"github.com/asynccnu/be-card/repository"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

const (
	url1 = "http://vcard.ccnu.edu.cn/index.php/index/index/cardreport.html"
	url2 = "http://vcard.ccnu.edu.cn/index.php/index/index/virtualreport.html"
)

//"88Gar7ySd4TeT9Q7O510h9Q6e0pacL2lZ3kmMRwxL7zIU0Y27yvD4IDz1EUYRFYp4K129yqy97HJo4XWLMggipTBH8Hxwp%3DOS4jCW9gLV4yfr"

func GetRecordOfConsumptionByVCard(msg domain.ServiceMsg) ([]domain.ResponseRecordsOfConsumption, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 添加表单字段
	_ = writer.WriteField("key", msg.Key)
	_ = writer.WriteField("start_time", msg.StartTime)
	writer.Close()
	var url string
	switch msg.Type {
	case repository.Card:
		url = url1
		break
	case repository.Virtual:
		url = url2
		break
	default:
		return nil, errors.New("卡片类型错误")
	}
	// 创建请求
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 获取响应的 Content-Type
	contentType := resp.Header.Get("Content-Type")

	// 检查 Content-Type 是否包含 "application/json"
	if !strings.Contains(contentType, "application/json") {
		return nil, fmt.Errorf("预期 JSON 响应，但收到 Content-Type: %s, 内容: %s", contentType, string(body))
	}

	var response Response
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return nil, err
	}
	return response.Data.List, nil
	//records, err := RecordsOfConsumptionToResponse(response.Data.List)
	//if err != nil {
	//	return nil, err
	//}
	//return records, nil
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []domain.ResponseRecordsOfConsumption `json:"list"`
		Sql  string                                `json:"sql"`
	} `json:"data"`
}
