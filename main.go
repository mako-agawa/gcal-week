package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gcal-week/auth"
	"gcal-week/calendar"
	"gcal-week/display"
)

func main() {
	credPath := filepath.Join(os.Getenv("HOME"), ".config", "gcal-week", "credentials.json")

	client, err := auth.GetClient(credPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "認証エラー: %v\n", err)
		os.Exit(1)
	}

	days, err := calendar.FetchWeek(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "カレンダー取得エラー: %v\n", err)
		os.Exit(1)
	}

	display.Render(days)
}
