package main

import (
	"local/MX-Steve/study/others/gomock/demo01/spider"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSpider := spider.NewMockSpider(mockCtl)
	mockSpider.EXPECT().GetBody().Return("go1.8")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.17" {
		t.Error("Get wrong version ", goVer)
	}
}
