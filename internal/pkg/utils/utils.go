package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// BatchInsertQuery to create raw batch insert query.
func BatchInsertQuery(tableName string, rowCount int, colCount int, cols ...[]string) string {
	if rowCount <= 0 || colCount <= 0 {
		return ""
	}

	// Create column format. Ex. (?, ?, ?)
	colList := []string{}
	for i := 0; i < colCount; i++ {
		colList = append(colList, "?")
	}

	colFormat := fmt.Sprintf("(%s)", strings.Join(colList, ", "))

	rowList := []string{}
	for i := 0; i < rowCount; i++ {
		rowList = append(rowList, colFormat)
	}

	colNames := ""
	if len(cols) > 0 {
		if len(cols[0]) == colCount {
			colNames = fmt.Sprintf("(%s)", strings.Join(cols[0], ", "))
		}
	}

	return fmt.Sprintf("INSERT INTO %s %s VALUES %s",
		tableName,
		colNames,
		strings.Join(rowList, ", "),
	)
}

// GetMapKey to get key from map.
func GetMapKey(m map[int]string, value string) int {
	for i, v := range m {
		if strings.EqualFold(v, value) {
			return i
		}
	}
	return 0
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

// GetValueFromSplit to get value from splitted string.
// Will return empty string if out of index.
func GetValueFromSplit(str string, separator string, index int) string {
	splitStr := strings.Split(str, separator)
	if len(splitStr) <= index {
		return ""
	}
	return strings.TrimSpace(splitStr[index])
}

// StrToNum to convert string number to integer including comma removal (1,234 -> 1234).
func StrToNum(strNum string) int {
	strNum = strings.TrimSpace(strNum)
	strNum = strings.Replace(strNum, ",", "", -1)
	intNum, _ := strconv.Atoi(strNum)
	return intNum
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

// UniqueInt to remove duplicate in slice.
func UniqueInt(list []int) (unique []int) {
	keys := make(map[int]bool)
	for _, v := range list {
		if _, value := keys[v]; !value {
			keys[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

// Thousands to format int to thousands string format.
func Thousands(num int) string {
	str := strconv.Itoa(num)
	lStr := len(str)
	digits := lStr
	if num < 0 {
		digits--
	}
	commas := (digits+2)/3 - 1
	lBuf := lStr + commas
	var sbuf [32]byte // pre allocate buffer at stack rather than make([]byte,n)
	buf := sbuf[0:lBuf]
	// copy str from the end
	for si, bi, c3 := lStr-1, lBuf-1, 0; ; {
		buf[bi] = str[si]
		if si == 0 {
			return string(buf)
		}
		si--
		bi--
		// insert comma every 3 chars
		c3++
		if c3 == 3 && (si > 0 || num > 0) {
			buf[bi] = ','
			bi--
			c3 = 0
		}
	}
}
