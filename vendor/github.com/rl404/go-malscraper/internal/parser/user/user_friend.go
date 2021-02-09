package user

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type friend struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.UserFriend
}

// GetFriends to get user friend list.
func (p *parser) GetFriends(a *goquery.Selection) []model.UserFriend {
	v := friend{area: a, cleanImg: p.cleanImg}
	v.setDetails()
	return v.data
}

func (f *friend) setDetails() {
	friendList := []model.UserFriend{}
	f.area.Find(".majorPad").Find(".friendHolder").Each(func(i int, friend *goquery.Selection) {
		friendArea := friend.Find(".friendBlock")
		friendList = append(friendList, model.UserFriend{
			Username:    f.getName(friendArea),
			Image:       f.getImage(friendArea),
			LastOnline:  f.getLastOnline(friendArea),
			FriendSince: f.getFriendSince(friendArea),
		})
	})
	f.data = friendList
}

func (f *friend) getName(friendArea *goquery.Selection) string {
	name, _ := friendArea.Find("a").Attr("href")
	return utils.GetValueFromSplit(name, "/", 4)
}

func (f *friend) getImage(friendArea *goquery.Selection) string {
	image, _ := friendArea.Find("a img").Attr("src")
	return utils.URLCleaner(image, "image", f.cleanImg)
}

func (f *friend) getLastOnline(friendArea *goquery.Selection) time.Time {
	t, _ := utils.StrToTime(friendArea.Find(".friendBlock div:nth-of-type(3)").Text())
	return t
}

func (f *friend) getFriendSince(friendArea *goquery.Selection) time.Time {
	friendSince := friendArea.Find(".friendBlock div:nth-of-type(4)").Text()
	friendSince = strings.Replace(friendSince, "Friends since", "", -1)
	t, _ := utils.StrToTime(friendSince)
	return t
}
