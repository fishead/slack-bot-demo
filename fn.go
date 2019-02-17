package demo

import (
	"net/http"

	echo "github.com/fishead/slack-bot-demo/echo"
)

// Echo response with echo command
func Echo(w http.ResponseWriter, r *http.Request) {
	echo.Echo(w, r)
}
