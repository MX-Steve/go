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
```
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

# 10. another
## 1. mysql
1. 不使用强约束，如外键，使用软约束，通过代码限制