package utils

import (
	"regexp"
	"strings"
	"time"
)

// ImageURLCleaner to clean dirty image url. For examples:
//   https://cdn.myanimelist.net/r/80x120/images/manga/3/214566.jpg?s=48212bcd0396d503a01166149a29c67e => https://cdn.myanimelist.net/images/manga/3/214566.jpg
//   https://cdn.myanimelist.net/r/76x120/images/userimages/6098374.jpg?s=4b8e4f091fbb3ecda6b9833efab5bd9b => https://cdn.myanimelist.net/images/userimages/6098374.jpg
//   https://cdn.myanimelist.net/r/76x120/images/questionmark_50.gif?s=8e0400788aa6af2a2f569649493e2b0f => empty string
func ImageURLCleaner(str string) string {
	str = strings.TrimSpace(str)
	match, _ := regexp.MatchString("(questionmark)|(qm_50)|(na.gif)", str)

	if match {
		return ""
	}

	str = strings.Replace(str, "v.jpg", ".jpg", -1)
	str = strings.Replace(str, "t.jpg", ".jpg", -1)
	str = strings.Replace(str, "_thumb.jpg", ".jpg", -1)
	str = strings.Replace(str, "userimages/thumbs", "userimages", -1)

	r := regexp.MustCompile(`r\/\d{1,3}x\d{1,3}\/`)
	str = r.ReplaceAllString(str, "")
	r = regexp.MustCompile(`\?.+`)
	return r.ReplaceAllString(str, "")

}

// VideoURLCleaner to clean dirty video url. For examples:
//   https://www.youtube.com/embed/qig4KOK2R2g?enablejsapi=1&wmode=opaque&autoplay=1 => https://www.youtube.com/watch?v=qig4KOK2R2g
//   https://www.youtube.com/embed/j2hiC9BmJlQ?enablejsapi=1&wmode=opaque&autoplay=1 => https://www.youtube.com/watch?v=j2hiC9BmJlQ
func VideoURLCleaner(str string) string {
	str = strings.TrimSpace(str)
	r := regexp.MustCompile(`\?.+`)
	str = r.ReplaceAllString(str, "")
	return strings.Replace(str, "embed/", "watch?v=", -1)
}

// URLCleaner is wrapper for image and video url cleaner for easier call.
func URLCleaner(str string, URLType string, isNeeded ...bool) string {
	if len(isNeeded) > 0 {
		if !isNeeded[0] {
			return str
		}
	}

	switch strings.ToLower(URLType) {
	case "image", "images", "img", "i":
		return ImageURLCleaner(str)
	case "video", "videos", "vid", "v":
		return VideoURLCleaner(str)
	default:
		return str
	}
}

// GetCurrentSeason to get current season ("spring", "summer", "fall", "winter").
func GetCurrentSeason() string {
	return GetSeasonName(int(time.Now().Month()))
}

// GetSeasonName to get season name ("spring", "summer", "fall", "winter").
func GetSeasonName(m int) string {
	switch {
	case m >= 1 && m < 4:
		return "winter"
	case m >= 4 && m < 7:
		return "spring"
	case m >= 7 && m < 10:
		return "summer"
	case m >= 10 && m <= 12:
		return "fall"
	default:
		return ""
	}
}
