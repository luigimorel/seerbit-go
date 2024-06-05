package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type KeyRequest struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func GenerateAuthToken(privateKey, publicKey string) (string, error) {
	requestData := KeyRequest{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	resp, err := http.Post("https://seerbitapi.com/api/v2/encrypt/key", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(body)
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", body)
		return "", err
	}

	return tokenResponse.Token, nil
}
