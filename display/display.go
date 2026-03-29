package display

import (
	"fmt"
	"time"

	cal "gcal-week/calendar"
)

// Everforest Dark カラー
const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	fg     = "\033[38;2;211;198;170m" // #D3C6AA
	grey0  = "\033[38;2;122;132;120m" // #7A8478
	grey1  = "\033[38;2;133;146;137m" // #859289
	aqua   = "\033[38;2;131;192;146m" // #83C092
	yellow = "\033[38;2;219;188;127m" // #DBBC7F
	blue   = "\033[38;2;127;187;179m" // #7FBBB3
	orange = "\033[38;2;230;152;117m" // #E69875
	bgHL   = "\033[48;2;61;72;77m"    // #3D484D (今日の背景)
)

// 2行目以降のインデント: 曜日3 + スペース2 + 日付5 + スペース2 = 12文字
const indent = "            "

var dayNames = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func Render(days []cal.DayEvents) {
	today := time.Now()
	start := days[0].Date
	end := days[6].Date

	fmt.Printf("\n%s%sWeek: %s – %s%s\n",
		blue, bold,
		start.Format("2006/01/02 (Mon)"),
		end.Format("01/02 (Mon)"),
		reset,
	)
	fmt.Printf("%s%s%s\n", grey0, "────────────────────────────────────────", reset)

	for i, day := range days {
		isToday := day.Date.Format("2006-01-02") == today.Format("2006-01-02")

		dayColor := yellow
		if isToday {
			dayColor = orange + bold
		}

		prefix := ""
		if isToday {
			prefix = bgHL
		}

		dayLabel := fmt.Sprintf("%s%s%-3s%s", prefix, dayColor, dayNames[i], reset)
		dateLabel := fmt.Sprintf("%s%s%s", grey1, day.Date.Format("01/02"), reset)

		if len(day.Events) == 0 {
			fmt.Print(prefix)
			fmt.Printf("%s  %s  %s(予定なし)%s\n", dayLabel, dateLabel, grey0, reset)
			continue
		}

		for j, ev := range day.Events {
			if j == 0 {
				fmt.Printf("%s  %s  ", dayLabel, dateLabel)
			} else {
				fmt.Print(prefix + indent)
			}

			if ev.AllDay {
				fmt.Printf("%sTODO%s %s%s%s\n", aqua, reset, fg, ev.Title, reset)
			} else {
				fmt.Printf("%s%s%s %s%s%s\n", grey1, ev.Time, reset, fg, ev.Title, reset)
			}
		}
	}

	fmt.Printf("%s%s%s\n\n", grey0, "────────────────────────────────────────", reset)
}
