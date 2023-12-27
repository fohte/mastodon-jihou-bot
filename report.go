package main

import (
	"fmt"
	"time"
)

type TimeReport struct {
	timeZone *time.Location
	time     time.Time
}

func NewTimeReport(timeZone *time.Location) *TimeReport {
	now := time.Now().In(timeZone)
	return &TimeReport{
		timeZone: timeZone,
		time:     time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, timeZone),
	}
}

func (tr *TimeReport) getClockEmoji() string {
	switch tr.time.Hour() {
	case 0, 12:
		return "🕛"
	case 1, 13:
		return "🕐"
	case 2, 14:
		return "🕑"
	case 3, 15:
		return "🕒"
	case 4, 16:
		return "🕓"
	case 5, 17:
		return "🕔"
	case 6, 18:
		return "🕕"
	case 7, 19:
		return "🕖"
	case 8, 20:
		return "🕗"
	case 9, 21:
		return "🕘"
	case 10, 22:
		return "🕙"
	case 11, 23:
		return "🕚"
	}
	return ""
}

// format:
// 📅 2023-12-28 (Tue)
// 🕛 00:00
// 📌 年末年始休暇: 残り n/n 日 (n %) [▓▓▓░░░░░░]
func (tr *TimeReport) CreateTimeReport() string {
	// 年末年始休暇の開始日と終了日 (仮設定)
	vacationStart := time.Date(2023, 12, 28, 0, 0, 0, 0, tr.timeZone)
	vacationEnd := time.Date(2024, 1, 4, 0, 0, 0, 0, tr.timeZone)

	// 現在日時から休暇までの残り日数とその割合
	daysUntilEnd := vacationEnd.Sub(tr.time).Hours() / 24
	totalVacationDays := vacationEnd.Sub(vacationStart).Hours() / 24
	remainingPercentage := int((daysUntilEnd / totalVacationDays) * 100)

	progressBar := ""
	for i := 0; i < 10; i++ {
		if i >= remainingPercentage/10 {
			progressBar += "░"
		} else {
			progressBar += "▓"
		}
	}

	s := ""

	if tr.time.Hour() == 0 {
		s += fmt.Sprintf("📅 %s (%s)\n", tr.time.Format("2006-01-02"), tr.time.Format("Mon"))
	}

	s += fmt.Sprintf("%s %s\n", tr.getClockEmoji(), tr.time.Format("15:04"))

	s += fmt.Sprintf("\n📌 年末年始休暇: 残り %.0f/%.0f 日 (%d %%) [%s]\n", daysUntilEnd, totalVacationDays, remainingPercentage, progressBar)

	return s
}
