## sample web project based on blockchain(hyperledger fabric)


### 1.Start Service
`./runApp.sh`

### 2.Send Request 
- Register

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

- Enroll

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

- Create Channel

```
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTQ4NjU1OTEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE0OTQ4NjE5OTF9.yWaJhFDuTvMQRaZIqg20Is5t-JJ_1BP58yrNLOKxtNI" 
  -H "content-type: application/json" 
  -d '{
	"channelName":"mychannel",
	"channelConfigPath":"../fixtures/channel/mychannel.tx"
}'
```

**OUTPUT:**
```
{
    "success":true,
    "message":"Channel 'mychannel' created Successfully"
}
```


