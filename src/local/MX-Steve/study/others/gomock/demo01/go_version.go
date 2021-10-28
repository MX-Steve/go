package main

import "local/MX-Steve/study/others/gomock/demo01/spider"

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
