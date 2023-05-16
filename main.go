package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
		return
	}

	tokenUrl := "https://accounts.spotify.com/api/token"
	contentType := "application/x-www-form-urlencoded"
	clientId, _ := os.LookupEnv("CLIENT_ID")
	clientSecret, _ := os.LookupEnv("CLIENT_SECRET")
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)

	resp, err := http.Post(tokenUrl, contentType, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error getting token: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}

	fmt.Println(string(body))
}
