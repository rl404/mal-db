package utils

import (
	"strconv"
	"strings"
)

// StrToNum to convert string number to integer including comma removal (1,234 -> 1234).
func StrToNum(strNum string) int {
	strNum = strings.TrimSpace(strNum)
	strNum = strings.Replace(strNum, ",", "", -1)
	intNum, _ := strconv.Atoi(strNum)
	return intNum
}

// StrToFloat to convert string number to float64 including comma removal (1,234.56 -> 1234.56).
func StrToFloat(strNum string) float64 {
	strNum = strings.TrimSpace(strNum)
	strNum = strings.Replace(strNum, ",", "", -1)
	floatNum, _ := strconv.ParseFloat(strNum, 64)
	return floatNum
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

// Thousand to convert int to string with ',' thousand.
func Thousand(num int) string {
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

// Ellipsis to substring long string.
func Ellipsis(str string, l int) string {
	if len(str) < l {
		return str
	}
	return str[:l] + "..."
}
