package auth

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pkg/browser"
	"github.com/rs/cors"
)

const webAuthURL = "https://poesitory.edgarallanohms.com/webauth"

var webAuthToken string
var webAuthSrv http.Server

func WebAuthentication() string {
	mux := http.NewServeMux()
	mux.HandleFunc("/webauth", webauthHttp)
	fmt.Println("Performing web auth")
	go openBrowserAfterDelay()
	webAuthSrv = http.Server{Addr: "127.0.0.1:25622", Handler: cors.Default().Handler(mux)}
	if err := webAuthSrv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}
	return fmt.Sprintf("User %s", webAuthToken)
}

func webauthHttp(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	webAuthToken = string(b)
	w.WriteHeader(http.StatusOK)
	go (func() {
		webAuthSrv.Shutdown(context.Background())
	})()
}

func openBrowserAfterDelay() {
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	if len(webAuthToken) == 0 {
		fmt.Printf("Open %s in your browser, and accept web authorization\n", webAuthURL)
		fmt.Println()
		fmt.Println()
		fmt.Println("To use a different authorization method, restart the command with the -p or -u flags.")
		browser.OpenURL(webAuthURL)
	}
}
