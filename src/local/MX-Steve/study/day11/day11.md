# 1. review
## 1. gin 
1. restful
    - delete
    - get
    - post
    - put
2. render
    - html
    - xml
    - yaml
3. router and group
```go
router.GET("/", indexHandler)
router.POST("/", indexHandler)
router.Any("/", indexHandler)
```

# 2. today
## 1. params
1. querystring
    - http://www.baidu.com?search=key
2. path params
    - http://127.0.0.1:8080/posts/2021/09/26
3. form params

## 2. logrus
```go
var log = logrus.New()
func initLogrus() (err error) {
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Error("fail to log to file")
		return
	}
	log.Out = file
	gin.SetMode(gin.ReleaseMode) // only use online
	// gin.DefaultWriter = log.Out
	gin.DefaultWriter = io.MultiWriter(log.Out, os.Stdout)
	log.Level = logrus.InfoLevel
	return
}
```
## 3. Cookie
1. 最大支持 4k
2. 可能被拦截和窃取
## 4. Session
1. 保存在服务端的数据，记录用户状态
2. 安全性更高
3. 存储方式
    - 内存
    - redis
    - mysql

# 10. another
## 1. mysql
1. 不使用强约束，如外键，使用软约束，通过代码限制


洞  hole
半路    halfway
无意义的    meaningless
互联网  internet
理解    understanding
开始n   beginning
日历    calender
删除    delete
马来西亚    Maraysia
高兴    glad
舒服地  comfortably
添加    add
代理    agent
特征    charactor
表演者  performer
事实上  actually/truly
决心    resolution  **
每周    weekly
怪物    monster
谋杀    murder
优势    advantage
收集，共同的    collect
评论    comment
乘船旅游    cruise
信封    envelope
小说    fiction
仓鼠    hamster
诚实的  sincerely
特别地  particularly
预言    
等级，收视率
令人失望的
赤道
幸运地
规矩
拥护者
收藏家
发型
表述
术语
导师
成年人
娱乐
愤怒的
再也不
亚洲人
服务员
评论
电梯
地球
毕业生
中午
海报
学期
真诚地
战略
套装
不舒服的
结果，影响
宿舍
银行员
瓶子
年长的
急坏的
刺穿
迷人的
葬礼
许可
突发的
圆圈
可信任的
传统
餐巾纸
同伴
化学
电子的
伤心
英国
庆祝
程序
煤炭
创造性
祖先
生态系统
环保的
网站
通常的
戏弄
实验室
品牌
绅士
灵感
表演
历史学家
感谢的
墨西哥
温暖
海岸