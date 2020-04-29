package monitor

import (
	"errors"
	"github.com/741369/go_utils/log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// pingServer pings the http server to make sure the router is working.
func PingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infof(nil, "Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
