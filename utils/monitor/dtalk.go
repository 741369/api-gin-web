/**********************************************
** @Des:
** @Author: lg1024
** @Last Modified time: 2019/12/31 上午11:38
***********************************************/

package monitor

import (
	"api-gin-web/utils/errno"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/741369/go_utils/log"
	"github.com/741369/go_utils/xlcurl"
	"github.com/spf13/viper"
	"os"
	"time"
)

type DTalk struct {
	webhook string
}

var httpClient *xlcurl.Client

func init() {
	httpClient = xlcurl.DefaultClient()
}

func New(webhook string) *DTalk {
	if webhook == "" {
		log.Infof(nil, "dtalk new webhook error")
		return nil
	}
	return &DTalk{webhook: webhook}
}

func (d *DTalk) Send(ctx context.Context, title, content string) {
	body := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": title,
			"text":  content,
		},
	}

	enc, err := json.Marshal(body)
	if err != nil {
		log.Infof(ctx, "send monitor dtalk json error, err = %v", err)
		return
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, _, err = httpClient.SimplePost(ctx, d.webhook, enc, 5*time.Second, headers)
	if err != nil {
		log.Infof(ctx, "send monitor dtalk request error, err = %v", err)
	}
}

func PushDingTalk(serviceTitle, method, debugName, description, httpStatus string) error {
	envGo := os.Getenv("ENV_GO")
	if envGo == "" {
		envGo = "prod"
	}

	phone := viper.GetString("monitor.phone")
	urlStr := viper.GetString("monitor.webhook")
	apiName := viper.GetString("appname")
	hostName, _ := os.Hostname()
	hostName = fmt.Sprintf("<font color=#0000FF> %s </font>", hostName)

	postData := `
		{
    "msgtype":"markdown",
    "markdown":{
        "title": "%s",
		"text":"### %s
		 - 服务主机：%s
		 - 服务方法：%s
		 - 告警级别：%s
		 - 告警名称：%s
		 - 告警环境：%s
		 - 错误信息：%s
		 > 服务请求状态码 %s
		@%s"
		},
		"at":{
			"atMobiles":[
				"%s"
			],
			"isAtAll":false
		}
	}
	`

	if debugName == "error" {
		description = fmt.Sprintf("<font color=#FF0000> %s </font>", description)
	} else {
		description = fmt.Sprintf("<font color=#32CD32> %s </font>", description)
	}

	postData = fmt.Sprintf(postData, apiName, apiName, hostName, method, debugName, serviceTitle, envGo, description, httpStatus, phone, phone)
	//req, _ := http.NewRequest("POST", urlStr, strings.NewReader(postData))
	//req.Header.Add("content-type", "application/json")

	headers := make(map[string]string)
	ctx := context.Background()
	cancel, resp, err := httpClient.Post(ctx, urlStr, bytes.NewReader([]byte(postData)), 5*time.Second, headers)
	if err != nil {
		log.Infof(nil, "PushDingTalk, %s, %v", urlStr, err)
		return errno.InternalServerError
	}
	defer cancel()
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//log.Printf("PushDingTalk2 %s", string(body))
	return nil
}
