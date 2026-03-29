package calendar

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type Event struct {
	Title  string
	Time   string // "HH:MM" or "" for all-day
	AllDay bool
}

type DayEvents struct {
	Date   time.Time
	Events []Event
}

func FetchWeek(client *http.Client) ([]DayEvents, error) {
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	// 今週の月曜〜日曜を計算
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	monday := now.AddDate(0, 0, -(weekday - 1))
	monday = time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, monday.Location())
	sunday := monday.AddDate(0, 0, 7)

	events, err := srv.Events.List("primary").
		TimeMin(monday.Format(time.RFC3339)).
		TimeMax(sunday.Format(time.RFC3339)).
		SingleEvents(true).
		OrderBy("startTime").
		Do()
	if err != nil {
		return nil, err
	}

	// 曜日ごとにグルーピング
	days := make([]DayEvents, 7)
	for i := range days {
		days[i].Date = monday.AddDate(0, 0, i)
		days[i].Events = []Event{}
	}

	for _, item := range events.Items {
		var t time.Time
		allDay := false

		if item.Start.DateTime != "" {
			t, _ = time.Parse(time.RFC3339, item.Start.DateTime)
		} else {
			t, _ = time.Parse("2006-01-02", item.Start.Date)
			allDay = true
		}

		idx := int(t.Sub(monday).Hours() / 24)
		if idx < 0 || idx >= 7 {
			continue
		}

		ev := Event{Title: item.Summary, AllDay: allDay}
		if !allDay {
			ev.Time = t.Format("15:04")
		}
		days[idx].Events = append(days[idx].Events, ev)
	}

	return days, nil
}
