package construct

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func CreateBitLaunchAccount(email string, password string) error {
	creds := make(map[string]string)
	url := "https://app.bitlaunch.io"
	path := "/api/signup"
	creds["email"] = email
	creds["password"] = password
	credJson, _ := json.Marshal(creds)
	resp, err := http.Post(url+path, "application/json", bytes.NewReader(credJson))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	if resp.StatusCode == 200 {
		return nil
	}
	return errors.New("failed to create account")
}
