# LoginSystem

> 技術使用:Gin-framework, Gorm, MySQL, Redis  
> MVC架構，RESTful API  
> 實作一個使用cookie&session的會員登入系統web server，會員登入後可以對product進行查詢、更新、增加、刪除  

## Quick start
先讓redis跟mysql啟動  
在MySQL創建一個schema並取名為interview  
schema名稱可以在.env更改  

```
go run main.go
```

## API 介紹
> 除了GET，DELETE方法之外其他都需要使用JSON檔案進行新增或更新的動作  
> JSON檔案格式可以在controllers裡的 input struct裡看到  
***User 基本路由***  

```
註冊用戶: POST   /user/register
用戶登入: POST   /user/login
用戶登出: GET    /user/logout
```
***Authorization***  
```
目前用戶: GET      /user/auth/current_user
取得產品: GET      /user/auth/product/:pid
創建產品: POST     /user/auth/product
更新產品: PUT      /user/auth/product/:pid
刪除產品: DELETE   /user/auth/product/:pid
```
