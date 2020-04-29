package monitor

import (
	"github.com/741369/go_utils/log"
	"net/http"

	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Shows OK as the ping-pong result
// @Description Shows OK as the ping-pong result
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /sd/health [get]
func HealthCheck(c *gin.Context) {
	message := c.ClientIP() + "==" + c.Request.Header.Get("X-Real-IP") + "==" +
		c.Request.Header.Get("http_x_Forwarded-For") + "==" +
		c.Request.Header.Get("http_X-Real-Ip") + "==" +
		c.Request.Header.Get(" X-Forwarded-Host") + "==" +
		c.Request.Header.Get("RemoteIp")

	ip := GetIp(c.Request)
	log.Infof(nil, "%s===%s", message, ip)
	c.String(http.StatusOK, message)
}

func GetIp(req *http.Request) (clientIP string) {
	header := req.Header
	clientIP = header.Get("Http_x_real_ip")
	if strings.TrimSpace(clientIP) != "" {
		return
	}
	clientIP = header.Get("X-Original-Forwarded-For")
	if strings.TrimSpace(clientIP) != "" {
		return
	}
	clientIP = header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP != "" {
		return clientIP
	}
	clientIP = header.Get("X-Real-Ip")
	if strings.TrimSpace(clientIP) != "" {
		return
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
