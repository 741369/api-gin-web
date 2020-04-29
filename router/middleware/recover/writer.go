package recovery

import (
	"api-gin-web/utils/monitor"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"sync"
)

const (
	ModeDev  = "debug"
	ModeProd = "release"
)

var (
	runmode string
	appname string
	dtalker *monitor.DTalk
	mu      sync.Mutex
)

func lazyInit() {
	runmode = viper.GetString("runmode")
	appname = viper.GetString("appname")
}

func NewWriter() io.Writer {
	lazyInit()
	switch runmode {
	case ModeDev:
		return new(notify)
	case ModeProd:
		return new(notify)
	default:
		return new(notify)
	}
}

type notify struct{}

func (l notify) Write(p []byte) (n int, err error) {
	text := fmt.Sprintf("## 项目%s告警, 环境%s：panic\n---------%s", appname, runmode, string(p))
	fmt.Println(runmode, appname, "=====", text)
	getDtalker().Send(context.Background(), "异常 - panic", text)
	return
}

func getDtalker() *monitor.DTalk {
	if dtalker != nil {
		return dtalker
	}
	mu.Lock()
	defer mu.Unlock()
	if dtalker != nil {
		return dtalker
	}
	webhook := viper.GetString("monitor.webhook")
	if webhook == "" {
		return nil
	}
	dtalker = monitor.New(webhook)
	return dtalker
}
