# 1. 创建并使用 swagger-editor
1. 创建 swagger-editor
]# docker run -d  -p 9104:8080 --name swagger-editor  swaggerapi/swagger-editor
2. 打开浏览器，编写 swagger
3. 导出 swagger.yaml
# 2. 使用新导出的 swagger 创建  swagger-ui 展示
1. 创建 swagger-ui
]# docker run -d  -p 9105:8080 -e SWAGGER_JSON=/swagger/swagger-minerops.json -v /swagger:/swagger  --name swagger-minerops swaggerapi/swagger-ui
2. 访问浏览器即可