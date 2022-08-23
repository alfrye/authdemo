package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	http.HandleFunc("/bfauth/demoapp/status/", authenticate)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
		os.Exit(1)

	}
	fmt.Printf("Serving running on port 3333")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("authenticate handler fired")

	http.SetCookie(w, &http.Cookie{
		Name:     "AUTHSESSION",
		Value:    uuid.New().String(),
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(time.Hour),
	})
	// w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("x-myheader", "simplauth")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login succeed"))
}
