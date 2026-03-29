package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func tokenPath() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "gcal-week", "token.json")
}

func GetClient(credentialsPath string) (*http.Client, error) {
	b, err := os.ReadFile(credentialsPath)
	if err != nil {
		return nil, fmt.Errorf("credentials.json が見つかりません: %w", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	token, err := loadToken(tokenPath())
	if err != nil {
		token, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		if err := saveToken(tokenPath(), token); err != nil {
			return nil, err
		}
	}

	return config.Client(context.Background(), token), nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Println("ブラウザで以下のURLを開いて認証してください:")
	fmt.Println(authURL)
	openBrowser(authURL)

	fmt.Print("認証コードを入力してください: ")
	var code string
	fmt.Scan(&code)

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("トークン取得失敗: %w", err)
	}
	return token, nil
}

func loadToken(path string) (*oauth2.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	return token, json.NewDecoder(f).Decode(token)
}

func saveToken(path string, token *oauth2.Token) error {
	// ディレクトリがなければ作成
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

func openBrowser(url string) {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		cmd, args = "open", []string{url}
	case "linux":
		cmd, args = "xdg-open", []string{url}
	default:
		return
	}
	exec.Command(cmd, args...).Start()
}
