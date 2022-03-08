# aliyundrive sdk for go language

> 阿里云盘 Go SDK ，基于 https://github.com/chyroc/go-aliyundrive 。
> 因原作者对于细节的把控粒度的分歧、加上本人的精力有限，因此不再往原项目提交 PR，感谢理解。

<!-- TOC depthfrom:2 orderedlist:false -->

- [和原版本有什么不同？](#%E5%92%8C%E5%8E%9F%E7%89%88%E6%9C%AC%E6%9C%89%E4%BB%80%E4%B9%88%E4%B8%8D%E5%90%8C)
- [更新记录](#%E6%9B%B4%E6%96%B0%E8%AE%B0%E5%BD%95)
- [安装](#%E5%AE%89%E8%A3%85)
- [使用](#%E4%BD%BF%E7%94%A8)
  - [初始化](#%E5%88%9D%E5%A7%8B%E5%8C%96)
    - [持久化配置](#%E6%8C%81%E4%B9%85%E5%8C%96%E9%85%8D%E7%BD%AE)
  - [Token](#token)
  - [信息](#%E4%BF%A1%E6%81%AF)
    - [用户](#%E7%94%A8%E6%88%B7)
  - [文件和目录](#%E6%96%87%E4%BB%B6%E5%92%8C%E7%9B%AE%E5%BD%95)
    - [列出文件和目录](#%E5%88%97%E5%87%BA%E6%96%87%E4%BB%B6%E5%92%8C%E7%9B%AE%E5%BD%95)
    - [新建目录](#%E6%96%B0%E5%BB%BA%E7%9B%AE%E5%BD%95)
    - [文件和目录信息](#%E6%96%87%E4%BB%B6%E5%92%8C%E7%9B%AE%E5%BD%95%E4%BF%A1%E6%81%AF)
    - [上传文件](#%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6)
    - [删除文件和目录](#%E5%88%A0%E9%99%A4%E6%96%87%E4%BB%B6%E5%92%8C%E7%9B%AE%E5%BD%95)
    - [移动文件和目录](#%E7%A7%BB%E5%8A%A8%E6%96%87%E4%BB%B6%E5%92%8C%E7%9B%AE%E5%BD%95)
    - [重命名文件和目录](#%E9%87%8D%E5%91%BD%E5%90%8D%E6%96%87%E4%BB%B6%E5%92%8C%E7%9B%AE%E5%BD%95)
  - [分享](#%E5%88%86%E4%BA%AB)
    - [列出分享](#%E5%88%97%E5%87%BA%E5%88%86%E4%BA%AB)
    - [新建分享](#%E6%96%B0%E5%BB%BA%E5%88%86%E4%BA%AB)
    - [修改分享](#%E4%BF%AE%E6%94%B9%E5%88%86%E4%BA%AB)
    - [取消分享](#%E5%8F%96%E6%B6%88%E5%88%86%E4%BA%AB)
- [常见问题](#%E5%B8%B8%E8%A7%81%E9%97%AE%E9%A2%98)
  - [如何获取 Refresh Token？](#%E5%A6%82%E4%BD%95%E8%8E%B7%E5%8F%96-refresh-token)

<!-- /TOC -->

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

### 初始化

```go
client := aliyundrive.New(...OptionFunc)
```

同时，选用持久化方案可以保存登录信息，建议自行配置对应的持久化以及日志方案。

#### 持久化配置

本 SDK 已经实现了部分常见的持久化方案，例如内存（没有持久化）、Redis 以及文件。如果您需要自行持久化方案，实现 Store 的 interface 即可，参见 store 目录中的具体实现。

例如，使用 Redis 作为持久化方案，可以参考以下代码示例：

```golang
redisStore, err := NewRedisStore(&redis.Options{
Addr: redisAddr,
})

if err != nil {
return err
}

cli := aliyundrive.New(WithStore(redisStore))
```

### Token

阿里云盘主要使用两种 token 信息，用于辨别用户以及获取用户对应的权限：

1. `access token` 用于获得用户对应的权限信息
2. `refresh token` 用于获得更新用户的 `access token` 信息

注意，这两个 token 都是具有有效期的（`access token` 的有效期按照小时来计算），有关具体如何获得 `refresh token`，可以参考以下常见问题的章节。

当获得 `refresh token` 以后，即可或许对应的 `access token` 以及用户信息，可以使用

```golang
token, err := cli.RefreshToken(context.TODO(), &RefreshTokenReq{
RefreshToken: refreshToken,
})
```

同时，SDK 会根据配置的持久化方案，来保存登录凭证。

### 信息

当通过 refresh token 获得用户信息以及 access token 以后，既可以通过使用

```golang
islogin := cli.IsLogin(context.TODO())
```

来判断用户是否已经登录成功，当登录成功后即可进行下一步的操作。

#### 用户

获取已登录用户的本身信息，可以使用

```golang
self, err := cli.MySelf(context.TODO())
```

详细可以参见 `self_test.go` 这个测试用例文件的内容。

同时，如果需要查看对应的用户权限已经网盘的资源等内容，可以使用

```golang
info, err := cli.PersonalInfo(context.TODO())
```

这个函数调用。详细可以参见 `personal_info_test.go` 这个测试用例文件的内容。

### 文件和目录

文件调用是 SDK 中的核心功能。在了解文件和目录之前，首先需要了解预设的几个模式以及常量。目前，SDK 封装了几个比较常用的常量，分别对应了以下的内容（以下内容来自 `type.go` 文件）。

类型

- `TypeFolder` // 类型为目录
- `TypeFile` // 类型为文件

建立资源的模式

- `ModeRefuse` // 当发现目标已经有存在的对象，则拒绝建立对应的目标对象
- `ModeAutoRename` // 当有目标对象存在时，自动重命名需要上传的文件或者目录资源

同时，每个文件和目录对象都会有对应的父节点，网盘对应的根节点统一为 `RootFileID` 这个常量。

#### 列出文件和目录

列出文件和目录有比较复杂的参数以及请求和相应结构，需要了解部分 SQL 相关的信息，可以方便更好的调用这个函数本身的功能。

简单的调用列出根目录的文件内容，可以调用

```golang
files, err := cli.Lists(context.TODO(), &FileListReq{
DriveID: self.DefaultDriveID,
})
```

即可，具体的参数相见 `FileListReq` 这个结构体其中的内容。

其中，`self.DefaultDriveID` 为获取用户信息中能够得到，它普遍是目前操作网盘对应的默认 DriveID ，因此很重要。

#### 新建目录

建立目录非常简单，例如在根目录下建立 just-for-example 目录以下的调用即可：

```golang
dirInfo, err := cli.CreateFolder(context.TODO(), &CreateFolderReq{
DriveID:      self.DefaultDriveID,
ParentFileID: RootFileID,
Name:         "just-for-example",
})
```

这里需要说明的是子目录的建立方式，可以使用以下的调用。

```golang
dirInfo, err := cli.CreateFolder(context.TODO(), &CreateFolderReq{
  DriveID:       self.DefaultDriveID,
  ParentFileID:  RootFileID,
  Name:          "just/for/example",
  CheckNameMode: ModeRefuse,
})
```

需要注意的是

- 子目录的父节点目录需要指定，否则会建立不成功
- 返回的目录信息是最子节点的信息，例如以上的建立是 example 这个子目录的信息
- 如果 CheckNameMode 的类型不是 ModeRefuse，那么如果有同名名录的情况下将在父目录建立重新命名的子目录。例如，以上的目录如果有对应的目录，那么将会建立 example(1) 这个目录。
- 如果 CheckNameMode 类型为 ModeRefuse，那么将会返回已经存在的目录信息

#### 文件和目录信息

用于获取文件和目录的信息，具体的调用方式可以参考 `get_test.go` 这个文件。

#### 上传文件

具体的调用方式可以参考 `get_test.go` 这个文件。

#### 删除文件和目录

具体的调用方式可以参考 `trash_test.go` 这个文件，注意云盘默认不会直接删除文件而是将文件移动到回收站中。

#### 移动文件和目录

具体的调用方式可以参考 `move_test.go` 这个文件。

#### 重命名文件和目录

具体的调用方式可以参考 `rename_test.go` 这个文件，注意这只是文件或者目录名称的修改，如果是需要移动文件，则参考移动文件和目录这个章节。

### 分享

分享部分的功能可以直接参见 `share_test.go` 这个测试用例。

## 常见问题

### 如何获取 Refresh Token？

可以在登录阿里网盘客户端后，打开 Web 控制台粘贴输入 JavaScript 代码

```javascript
JSON.parse(localStorage.token).refresh_token;
```

返回的字符串内容就是 Refresh Token 。目前 Token 的有效期未知，目前的测试情况看来有几天的时间。
