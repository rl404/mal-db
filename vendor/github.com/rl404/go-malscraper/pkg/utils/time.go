package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// StrToTime to convert MyAnimeList date string to time.Time type.
func StrToTime(str string) (t time.Time, isValid bool) {
	str = strings.TrimSpace(str)
	if str == "" {
		return time.Time{}, false
	}

	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)

	// Now
	if strings.ToLower(str) == "now" {
		return now, true
	} else

	// 5 seconds ago || 1 second ago
	if match, _ := regexp.MatchString(`([0-9]+) seconds? ago`, str); match {
		secondStr := GetValueFromSplit(str, " ", 0)
		second := time.Duration(StrToNum(secondStr))
		return now.Add(-(second * time.Second)), true
	} else

	// 27 minutes ago || 1 minute ago
	if match, _ := regexp.MatchString(`([0-9]+) minutes? ago`, str); match {
		minuteStr := GetValueFromSplit(str, " ", 0)
		minute := time.Duration(StrToNum(minuteStr))
		return now.Add(-(minute * time.Minute)), true
	} else

	// 8 hours ago || 1 hour ago
	if match, _ := regexp.MatchString(`([0-9]+) hours? ago`, str); match {
		hourStr := GetValueFromSplit(str, " ", 0)
		hour := time.Duration(StrToNum(hourStr))
		return now.Add(-(hour * time.Hour)), true
	} else

	// Today, 8:48 AM
	if match, _ := regexp.MatchString(`Today, ([0-9]{1,2}):([0-9]{2}) (AM|PM)`, str); match {
		today := now.Format("Jan 2, 2006")
		str = strings.Replace(str, "Today,", today, -1)
		return StrToTime(str)
	} else

	// Yesterday, 2:45 AM
	if match, _ := regexp.MatchString(`Yesterday, ([0-9]{1,2}):([0-9]{2}) (AM|PM)`, str); match {
		today := now.AddDate(0, 0, -1).Format("Jan 2, 2006")
		str = strings.Replace(str, "Yesterday,", today, -1)
		return StrToTime(str)
	} else

	// Apr 17, 8:57 PM
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}), ([0-9]{1,2}):([0-9]{2}) (AM|PM)`, str); match {
		str = strings.Replace(str, ",", now.Format(", 2006"), -1)
		return StrToTime(str)
	} else

	// Apr 9, 2017 10:57 PM
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}), ([0-9]{4}) ([0-9]{1,2}):([0-9]{2}) (AM|PM)`, str); match {
		return ParseTime("Jan 2, 2006 3:04 PM", str), true
	} else

	// Apr 7, 2013, 01:58 (JST)
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}), ([0-9]{4}), ([0-9]{1,2}):([0-9]{2}) \(([A-Z]+)\)`, str); match {
		return ParseTime("Jan 2, 2006, 15:04 (MST)", str), true
	} else

	// Apr 3, 1998
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}), ([0-9]{4})$`, str); match {
		return ParseTime("Jan 2, 2006", str), true
	} else

	// Apr ??, 1998
	if match, _ := regexp.MatchString(`([a-z]+) ([?]{2}), ([0-9]{4})$`, str); match {
		str = strings.Replace(str, "??", "1", -1)
		return StrToTime(str)
	} else

	// ??? ??, 1998
	if match, _ := regexp.MatchString(`([?]+) ([?]{2}), ([0-9]{4})$`, str); match {
		str = GetValueFromSplit(str, " ", 2)
		return StrToTime(str)
	} else

	// Jan, 2021
	if match, _ := regexp.MatchString(`([a-z]+), ([0-9]{4})$`, str); match {
		str = strings.Replace(str, ",", " 1,", -1)
		return StrToTime(str)
	} else

	// Jan 2021
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{4})$`, str); match {
		str = strings.Replace(str, " ", ", ", -1)
		return StrToTime(str)
	} else

	// Jan 3,
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}),$`, str); match {
		return ParseTime("Jan 2,", str), true
	} else

	// Jan 3
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2})$`, str); match {
		return ParseTime("Jan 2", str), true
	} else

	// 2021
	if match, _ := regexp.MatchString(`^([0-9]{4})$`, str); match {
		return ParseTime("Jan 2, 2006", "Jan 1, "+str), true
	} else

	// 10-25-02
	if match, _ := regexp.MatchString(`^([0-9]{2})-([0-9]{2})-([0-9]{2})$`, str); match {
		return ParseTime("01-02-06", str), true
	} else

	// 10-??-02
	if match, _ := regexp.MatchString(`^([0-9]{2})-([?]{2})-([0-9]{2})$`, str); match {
		str = strings.Replace(str, "??", "01", -1)
		return ParseTime("01-02-06", str), true
	} else

	// ??-??-02
	if match, _ := regexp.MatchString(`^([?]{2})-([?]{2})-([0-9]{2})$`, str); match {
		str = strings.Replace(str, "??", "01", -1)
		return ParseTime("01-02-06", str), true
	}

	return time.Time{}, false
}

var months = map[string]int{
	"Jan": 1,
	"Feb": 2,
	"Mar": 3,
	"Apr": 4,
	"May": 5,
	"Jun": 6,
	"Jul": 7,
	"Aug": 8,
	"Sep": 9,
	"Oct": 10,
	"Nov": 11,
	"Dec": 12,
}

// StrToDate to convert MyAnimeList date string to splitted date (year, month, day).
func StrToDate(str string) (y, m, d int) {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0, 0, 0
	}

	str = strings.Replace(str, ",", "", -1)
	spltStr := strings.Split(str, " ")

	// Apr 3 1998
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}) ([0-9]{4})$`, str); match {
		return StrToNum(spltStr[2]), months[spltStr[0]], StrToNum(spltStr[1])
	} else

	// Apr ?? 1998
	if match, _ := regexp.MatchString(`([a-z]+) ([?]{2}) ([0-9]{4})$`, str); match {
		return StrToNum(spltStr[2]), months[spltStr[0]], 0
	} else

	// ??? ?? 1998
	if match, _ := regexp.MatchString(`([?]+) ([?]{2}) ([0-9]{4})$`, str); match {
		return StrToNum(spltStr[2]), 0, 0
	} else

	// Jan 2021
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{4})$`, str); match {
		return StrToNum(spltStr[1]), months[spltStr[0]], 0
	} else

	// Jan 3
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2})$`, str); match {
		return 0, months[spltStr[0]], StrToNum(spltStr[1])
	} else

	// 2021
	if match, _ := regexp.MatchString(`^([0-9]{4})$`, str); match {
		return StrToNum(spltStr[0]), 0, 0
	} else

	// Apr 9 2017 10:57 PM
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}) ([0-9]{4}) ([0-9]{1,2}):([0-9]{2}) (AM|PM)`, str); match {
		return StrToNum(spltStr[2]), months[spltStr[0]], StrToNum(spltStr[1])
	} else

	// Apr 7, 2013, 01:58 (JST)
	if match, _ := regexp.MatchString(`([a-z]+) ([0-9]{1,2}) ([0-9]{4}) ([0-9]{1,2}):([0-9]{2}) \(([A-Z]+)\)`, str); match {
		return StrToNum(spltStr[2]), months[spltStr[0]], StrToNum(spltStr[1])
	} else

	// 10-25-02
	if match, _ := regexp.MatchString(`^([0-9]{2})-([0-9]{2})-([0-9]{2})$`, str); match {
		spltStr = strings.Split(str, "-")
		y := StrToNum(spltStr[2])
		if y >= 69 {
			y += 1900
		} else {
			y += 2000
		}
		return y, StrToNum(spltStr[0]), StrToNum(spltStr[1])
	} else

	// 10-??-02
	if match, _ := regexp.MatchString(`^([0-9]{2})-([?]{2})-([0-9]{2})$`, str); match {
		spltStr = strings.Split(str, "-")
		y := StrToNum(spltStr[2])
		if y >= 69 {
			y += 1900
		} else {
			y += 2000
		}
		return y, StrToNum(spltStr[0]), 0
	} else

	// ??-??-02
	if match, _ := regexp.MatchString(`^([?]{2})-([?]{2})-([0-9]{2})$`, str); match {
		spltStr = strings.Split(str, "-")
		y := StrToNum(spltStr[2])
		if y >= 69 {
			y += 1900
		} else {
			y += 2000
		}
		return y, 0, 0
	}

	return 0, 0, 0
}

// ParseTime is time parser wrapper of built-in library `time.ParseInLocation`.
func ParseTime(layout, str string) time.Time {
	t, _ := time.ParseInLocation(layout, str, time.Local)
	return t
}

// GetDuration to get duration in seconds.
// Example: "3 hr. 30 min. 20 sec." => 12620.
func GetDuration(duration string) (seconds int) {
	duration = strings.TrimSpace(duration)
	if duration == "" {
		return 0
	}

	// Count hour.
	r := regexp.MustCompile(`\d+ hr.`)
	match := r.FindStringSubmatch(duration)
	if len(match) > 0 {
		h := GetValueFromSplit(match[0], " ", 0)
		seconds += StrToNum(h) * 3600
	}

	// Count minute.
	r = regexp.MustCompile(`\d+ min.`)
	match = r.FindStringSubmatch(duration)
	if len(match) > 0 {
		m := GetValueFromSplit(match[0], " ", 0)
		seconds += StrToNum(m) * 60
	}

	// Count minute.
	r = regexp.MustCompile(`\d+ sec.`)
	match = r.FindStringSubmatch(duration)
	if len(match) > 0 {
		s := GetValueFromSplit(match[0], " ", 0)
		seconds += StrToNum(s)
	}

	return seconds
}

// SecondToString to convert int seconds to string duration.
// Example: 12620 => 03:30:20.
func SecondToString(seconds int) string {
	d, _ := time.ParseDuration(fmt.Sprintf("%vs", seconds))
	duration := d.String()

	h, m, s := "00", "00", "00"

	// Count hour.
	r := regexp.MustCompile(`\d+h`)
	match := r.FindStringSubmatch(duration)
	if len(match) > 0 {
		h = strings.Replace(match[0], "h", "", -1)
		if len(h) == 1 {
			h = "0" + h
		}
	}

	// Count minute.
	r = regexp.MustCompile(`\d+m`)
	match = r.FindStringSubmatch(duration)
	if len(match) > 0 {
		m = strings.Replace(match[0], "m", "", -1)
		if len(m) == 1 {
			m = "0" + m
		}
	}

	// Count minute.
	r = regexp.MustCompile(`\d+s`)
	match = r.FindStringSubmatch(duration)
	if len(match) > 0 {
		s = strings.Replace(match[0], "s", "", -1)
		if len(s) == 1 {
			s = "0" + s
		}
	}

	return fmt.Sprintf("%v:%v:%v", h, m, s)
}
