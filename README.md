# OAuth2-Google

## 使用OAuth进行认证和授权的过程如下所示

1. 用户打开客户端以后，客户端要求用户给予授权。
2. 用户同意给予客户端授权。
3. 客户端使用上一步获得的授权，向认证服务器申请令牌。
4. 认证服务器对客户端进行认证以后，确认无误，同意发放令牌。
5. 客户端使用令牌，向资源服务器申请获取资源。
6. 资源服务器确认令牌无误，同意向客户端开放资源

[Google Dev Console](https://console.developers.google.com)

## Content Example

    Content: {
        "id": "112417997111972571448",
        "email": "luxuze1994@gmail.com",
        "verified_email": true,
        "name": "xuze lu",
        "given_name": "xuze",
        "family_name": "lu",
        "picture": "https://lh6.googleusercontent.com/-pg0aK22eny0/AAAAAAAAAAI/AAAAAAAAAAA/ACHi3resmgnqhJgJrdyev5Nsme2wQMKc0w/photo.jpg",
        "locale": "zh-CN"
    }

## [理解OAuth 2.0 - 阮一峰](http://www.ruanyifeng.com/blog/2014/05/oauth_2_0.html)
