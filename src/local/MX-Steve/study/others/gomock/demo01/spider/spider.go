//go:generate mockgen -destination mock_spider.go -package spider local/MX-Steve/study/others/gomock/demo01/spider Spider
package spider

type Spider interface {
	GetBody() string
}
