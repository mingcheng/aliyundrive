# aliyundrive sdk for go language

> 阿里云盘 Go SDK ，基于 https://github.com/chyroc/go-aliyundrive 。
> 因原作者对于细节的把控粒度的分歧、加上本人的精力有限，因此不再往原项目提交 PR，感谢理解。

## 和原版本有什么不同？

1. 重点关注接口以及功能的稳定性方面
2. 精简代码结构
3. 增加每个功能对应的测试用例
4. 去除不必要的第三方库（例如控制台二维码输出、自己造轮子的 HTTP 请求库等）
5. 增加实现配置保存的方式（内存、文件、以及 Redis）

## 更新记录

- 20220306 初始化重构版本

## 安装

```shell
go get github.com/mingcheng/aliyundrive
```

## 使用

### 初始化 SDK 实例

```go
client := aliyundrive.New()
```

### 登录

> [具体代码参考这里](./_examples/login-by-qrcode/main.go)


```go
user, err := ins.Auth.LoginByQrcode(context.TODO())
```

<img src="screenshots/login-by-qrcode.png" width="300px" >

### 读取文件

> [具体代码参考这里](./_examples/list-files/main.go)

```go
resp, err := ins.File.GetFileList(context.TODO(), &aliyundrive.GetFileListReq{
    DriveID:      driveID,
    ParentFileID: parentID,
    Marker:       next,
})
```

<img src="screenshots/list-files.png" width="300px" >

### 获取分享的内容

> [具体代码参考这里](./_examples/get-share/main.go)

```go
sharedInfo, err := ins.ShareLink.GetShareByAnonymous(ctx, &aliyundrive.GetShareByAnonymousReq{
    ShareID: shareID,
})
```

<img src="screenshots/get-share.png" width="300px" >

## 其他

- 阿里云盘命令行客户端: https://github.com/chyroc/aliyundrive-cli
