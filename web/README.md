## 区块链web


### 1.启动服务
`./runApp.sh`

### 2.通过CURL进行交互数据
- 注册

* Register user in Organization - **Org1**:

`curl -s -X POST http://localhost:4000/register -H "content-type: application/x-www-form-urlencoded" -d 'username=Jim&password=123456&orgName=org1'`

**OUTPUT:**

```
{
  "success": true,
  "secret": "RaxhMgevgJcm",
  "message": "Jim enrolled Successfully",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTQ4NjU1OTEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE0OTQ4NjE5OTF9.yWaJhFDuTvMQRaZIqg20Is5t-JJ_1BP58yrNLOKxtNI"
}
```

- 登录

* Enroll user in Organization - **Org1**:

`curl -s -X POST http://localhost:4000/enroll -H "content-type: application/x-www-form-urlencoded" -d 'username=Jim&password=123456orgName=org1'`

**OUTPUT:**

```
{
  "success": true,
  "message": "Jim enrolled Successfully",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTQ4NjU1OTEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE0OTQ4NjE5OTF9.yWaJhFDuTvMQRaZIqg20Is5t-JJ_1BP58yrNLOKxtNI"
}
```
