package main

import "local/MX-Steve/study/day09/03chat/server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}
