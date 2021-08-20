# 1. 简介
1. SSO, single sign on, 单点登录。
    - sso 多用于多个应用之间的切换，例如百度论坛、百度知道、百度云、百度文库等，在其中一个系统中登录，（登录有效期内）切换到另一个系统的时候，不必再次输入用户名密码。
2. oauth2.0，开放授权，不兼容 oauth1.0.允许第三方应用代表用户获得访问权限。
    - 可以作为 web 应用、桌面应用和手机等设备提供专门的认证流程。
    - 例如，用 qq 账号登录豆瓣、美团、大众点评；用支付宝账号登录淘宝、天猫等。
3. sso 和 oauth2.0 在应用场景上的区别在于
    - SSO 是为了解决一个用户在鉴权服务器登陆过一次以后，可以在任何应用（通常是一个厂家的各个系统）中畅通无阻。
    - OAuth2.0 解决的是通过令牌（token）而不是密码获取某个系统的操作权限（不同厂家之间的账号共享）。
# 2. SSO身份认证流程
1. 文字说明
    1. 用户访问app1，浏览器向app1发出请求。【未登录状态】
    2. app1收到请求后，发现浏览器处于未登录状态，返回一个302，跳转访问cas，并附带returnurl【即app1自身地址】。
    3. 浏览器收到302后，访问cas。
    4. cas返回给浏览器一个登录页。
    5. 浏览器展示登录页，用户携带returnurl提交用户名密码到cas。
    6. cas进行身份认证，认证成功创建sso的session和认证ticket（ST），返回302，并将ticket一并返回给浏览器
    7. 浏览器记录cookie，带着ticket（ST）访问returnurl【app1地址】。（cookie是cas的）
    8. app1收到请求后，向cas确认ticket（ST）是否有效。
    9. cas验证ticket通过后，返回200给app1。
    10. app1记录登录状态并响应浏览器请求（步骤7），返回302（不带ST）。（添加app1的cookie）
    11. 浏览器收到302，访问app1地址，此时cookie和session齐全，身份认证完成，向浏览器返回请求的资源。之后（cookie、session有效期内）访问app1，不必再身份验证，直接返回请求资源。【app1登录完成】
    12. 用户访问app2，浏览器向app2发出请求。【app2未登录】
    13. app2收到请求，发现浏览器处于未登录状态，返回一个302，跳转访问cas，并附带returnurl【即aap2自身地址】。
    14. 浏览器收到302，访问cas。（由于步骤7已经记录的cas的cookie，浏览器会在请求时将cookie携带）
    15.cas验证身份，由于有cookie存在，直接返回票据ticket（ST2），并让浏览器重定向。
    16.浏览器带着ticket（ST2）访问returnurl。
    17.app2收到请求后，向cas确认ticket（ST2）是否有效。
    18.cas验证ticket通过后，返回200给app2。
    19.app2记录登录状态并响应浏览器请求（步骤16），返回302（不带ST）。（添加app2的cookie）
    20.浏览器收到302，访问app2地址，此时cookie和session齐全，身份认证完成，向浏览器返回请求的资源。之后（cookie、session有效期内）访问app2，不必再身份验证，直接返回请求资源。【app2登录完成】
2. 流程图
    - sso_process.jpg
# 3. 测试- 谷歌云平台 https://console.cloud.google.com/
1. project: altchain
2. OAuth 2.0 Client IDs
    1. clientID: 269392545684-79f8raerpes3r8q8b3pkgau3q221k2ar.apps.googleusercontent.com
    2. clientSecret: 5aaXtyztgmlCTkRkpa36HbFF