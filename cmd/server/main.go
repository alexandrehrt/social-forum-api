package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"social-api/config"
	"social-api/internal/initializers"

	"github.com/spf13/viper"
)

func init() {
	_, err := config.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	config.ConnectToDB()
}

func main() {
	r := chi.NewRouter()

	initializers.Routes(r)

	port := viper.Get("WEB_SERVER_PORT")

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		panic(err)
	}
}
