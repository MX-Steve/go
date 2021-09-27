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
