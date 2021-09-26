# 1. go module
# 2. Gin 框架
## 1. 渲染
1. JSON
```
c.JSON(状态码,能够被序列化的数据gin.H{})
```
2. HTML
```
c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "log website",
})
```
3. template
```
router.LoadHTMLGlob("templates/*")
```
4. statics
```
router.Static("/st", "./statics")
```
5. XML
```
c.XML(http.StatusOK, "login.html", gin.H{
		"msg": "log website",
})
```
# 3. ex: book manager system

