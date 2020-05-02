package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"time"

	"api-gin-web/controller"
	"api-gin-web/utils/errno"
	"github.com/741369/go_utils/log"

	//"api-gin-web/router/prome"
	"api-gin-web/utils"
	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func TraceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := utils.GetUuidString()
		ctx.Set("X-Request-Id", traceId)
	}
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("(/swagger)")
		if reg.MatchString(path) {
			return
		}

		if path == "/sd/health" || path == "/sd/ram" || path == "/upush/ping" || path == "/metric" {
			return
		}

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method
		ip := c.Request.Header.Get("ip")

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue
		c.Next()

		// Calculates the latency
		end := time.Now().UTC()
		latency := end.Sub(start)

		//prome.RequestDurations.WithLabelValues("api-gin-web", path, strconv.Itoa(c.Writer.Status())).Observe(latency.Seconds())
		//prome.ProcessingTime.WithLabelValues(path, "request").Observe(latency.Seconds())

		code, message := -1, ""

		// get code and message
		var response controller.Response

		if err := json.Unmarshal(blw.body.Bytes(), &response); err == nil && response.Message != "" {
			code = response.Code
			message = response.Message
		} else {
			log.Infof(c, "response body can not unmarshal to controller.Response struct, body: `%s`, err = %v \n", blw.body.Bytes(), err)
			code = errno.InternalServerError.Code
			message = ""
			//message = err.Error()
			//log.Errorf(err, "%s | %s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
		}

		if os.Getenv("ENV_GO") == "" || os.Getenv("ENV_GO") == "dev" || os.Getenv("ENV_GO") == "pre" {
			if pad.Right(method, 5, "") != "OPTIONS" {
				// 收集请求日志到es
				str, _ := json.Marshal(response.Data)
				returnData := map[string]interface{}{
					"request_date":      time.Now(),
					"request_path":      path,
					"request_cookie":    c.Request.Header.Get("cookie"),
					"request_ua":        c.Request.Header.Get("user-agent"),
					"authorization":     c.Request.Header.Get("authorization"),
					"request_url_param": c.Request.URL.RawQuery,
					"request_body":      string(bodyBytes),
					"latency_time":      latency.String(),
					"request_duration":  latency.Seconds(),
					"request_ip_addr":   ip,
					"request_method":    pad.Right(method, 5, ""),
					"response_status":   c.Writer.Status(),
					"return_code":       code,
					"return_message":    message,
					"return_data":       string(str),
				}
				log.Info(c, "request_response_data", returnData)
			}
		} else {
			log.Infof(c, "%s | %s | %s %s %d | {code: %d, message: %s, data: %v}", latency, ip, pad.Right(method, 5, ""), path, c.Writer.Status(), code, message, response.Data)
		}
	}
}
