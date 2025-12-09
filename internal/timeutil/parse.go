package timeutil

import (
	"errors"
	"strconv"
	"strings"
	"time"
)



func ParseDateString(date string, loc *time.Location) (time.Time, error){

	if loc == nil {
		loc = time.Local
	}

	trimmed := strings.TrimSpace(date)
	lower := strings.ToLower(trimmed)

	// 1) keywords: today, tomorrow, +Nd
	nowLocal := time.Now().In(loc)
	midnight := time.Date(nowLocal.Year(), nowLocal.Month(), nowLocal.Day(), 0, 0, 0, 0, loc)

	if lower == "today" {
		return EndOfDay(midnight), nil
	}
	if lower == "tomorrow" {
		tomorrow := midnight.AddDate(0, 0, 1)
		return EndOfDay(tomorrow), nil
	}

	if strings.HasPrefix(lower, "+") && strings.HasSuffix(lower, "d") {
		nstr := strings.TrimSuffix(strings.TrimPrefix(lower, "+"), "d")
		n, err := strconv.Atoi(nstr)
		if err != nil {
			return time.Time{}, errors.New("invalid relative days")
		}
		future := midnight.AddDate(0, 0, n)
		return EndOfDay(future), nil
	}
		
// 2) structured parsing — try layouts in order.
	// Use the original `trimmed` (case preserved) for parsing.
	layouts := []string{
		"2006-01-02",        // date-only -> apply EndOfDay
		time.RFC3339,        // includes time -> keep exact time
		"02-01-2006",        // date-only -> apply EndOfDay
		"January 2 2006",    // date-only -> apply EndOfDay
		"Jan 2 2006",        // date-only -> apply EndOfDay
	}

	for _, layout := range layouts {
		t, err := time.ParseInLocation(layout, trimmed, loc)
		if err != nil {
			continue
		}

		// If layout is RFC3339 (time included), return exact parsed time.
		// For date-only layouts, normalize to end of day.
		if layout == time.RFC3339 {
			return t, nil
		}
		// date-only case
		return EndOfDay(t), nil
	}

	return time.Time{}, errors.New("invalid date format")

}



func EndOfDay(t time.Time) time.Time{
	if t.IsZero() {
		return time.Time{}
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func ToUTC(t time.Time) time.Time{
	return t.In(time.UTC)
}