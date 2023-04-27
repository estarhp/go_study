package main

import "time"

type wangyi_user struct {
	id        int
	lv        int
	city      int
	nickName  string
	avatarImg string
	create    time.Time
	loveList  []string
	profile   []struct{}
}
