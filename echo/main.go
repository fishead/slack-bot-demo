package echo

import (
	"net/http"
	"os"

	"encoding/json"
	"flag"

	"github.com/nlopes/slack"
)

// Echo echo message
func Echo(w http.ResponseWriter, r *http.Request) {
	var (
		verificationToken string
	)

	flag.StringVar(&verificationToken, "token", "YOUR_VERIFICATION_TOKEN_HERE", os.Getenv("SLACK_VERIFICATION_TOKEN"))
	flag.Parse()

	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(verificationToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := &slack.Msg{Text: s.Text}
	b, err := json.Marshal(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
