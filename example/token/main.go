package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aoisensi/go-yasuna/pkg/yasuna"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	oauth2 := &yasuna.OAuth2{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		RedirectURI:  os.Getenv("TWITTER_REDIRECT_URL"),
	}
	url := oauth2.AuthCodeURL(yasuna.ScopeAll)
	fmt.Println("Please open the URL above in your browser and enter the code:")
	fmt.Println(url)
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		state := r.FormValue("state")
		twitter, err := oauth2.Exchange(code, state)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("AccessToken:", twitter.Token.AccessToken)
		fmt.Println("RefreshToken:", twitter.Token.RefreshToken)
		if !twitter.Token.Expire.IsZero() {
			fmt.Println("Expires at about:", twitter.Token.Expire)
		} else {
			fmt.Println("The token does not expire.")
		}
		w.Write([]byte("You can close this window now.\nBack to the terminal."))
		cancel()
	})
	srv := &http.Server{Addr: ":" + os.Getenv("PORT")}
	go srv.ListenAndServe()

	<-ctx.Done()
	srv.Shutdown(ctx)
}
