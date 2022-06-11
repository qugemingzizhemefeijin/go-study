### Golang 学习案例

Golang1.14.7 + redis5.0.9 for Windows

##### 1.相关软件下载

[Golang1.14.7下载](https://studygolang.com/dl/golang/go1.14.7.windows-amd64.zip "Golang1.14.7下载地址请点击") <br/>
[Redis5.0.9下载](https://github.com/tporadowski/redis/releases/download/v5.0.9/Redis-x64-5.0.9.msi "Redis5.0.9下载地址请点击") <br/>

#### 2.相关包地址

[go-redis地址](https://github.com/go-redis/redis) <br/>
[redigo地址](https://github.com/gomodule/redigo) <br/>

#### 3.vscode快捷键

| 名称 | 快捷键 |
| --- | --- |
| 删除当前行 | ctrl + shift + k |
| 向上/向下复制当前行 | shift + alt + ↓/↑ |
| 补全代码 | ctrl + . |
| 添加注释和取消注释 | ctrl + / |
| 快速修改 | ctrl + . |
| 快速格式化代码 | shift + alt + f |

#### 4.vscode 安装go环境无法安装gopls等插件，响应超时、失去连接等问题的简单解决方案

解决方案是修改代理，然后在`cmd`下面输入：
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=“sum.golang.org”
```

修改代理到国内的go，然后在 cmd 重新获取即可成功：
```
go get -v golang.org/x/tools/gopls
go get -u github.com/derekparker/delve/cmd/dlv
```

开启CGO
```
set CGO_ENABLED=1
```

#### 5.安装go protobuf

Linux下使用方法如下：
```
mkdir -p $GOPATH/src/google.golang.org
git clone https://e.coding.net/robinqiwei/googleprotobuf.git $GOPATH/src/google.golang.org/protobuf
```

Windows下使用方法：
```
// 找到你的GOPATH/src目录，新建google.golang.org文件夹
// 在google.golang.org目录下执行

git clone https://e.coding.net/robinqiwei/googleprotobuf.git protobuf
```

windows很难安装成功，可以使用如下步骤：
```
git clone git@github.com:golang/protobuf.git ./github.com/golang/protobuf
// git clone 完代码后，在这个文件夹下（protoc-gen-go）打开终端执行
cd github.com/golang/protobuf/protoc-gen-go
go install
// 会生成一个exe文件
// 复制里面的protoc-gen-go.exe(后面的grpc是我后面安装的，这个别管) 复制到goroot的bin目录下
// protoc --go_out=. hello.proto
```

##### 解决:protoc-gen-go: unable to determine Go import path for “*.proto”***

在使用`protoc`命令根据`*.proto`文件生成代码时报了错误，下面会提示相应的解决办法，这里我们使用的是第一种，稍微翻译一下就知道，可以通过`go_package`去指定生成的`go`文件的位置。

```
// 这里不能用 "./"，否则生成的go文件的 package 为 __，然后还要手动去改，为了避免就写成下面的样子
option go_package = "../所在包名"; // 指定生成go文件保存到当前包中
```

#### 6.安装 grpc

默认可以通过：
```
go get google.golang.org/grpc
```

如果安装不成功，则可以使用以下方式：
```
// 首先进入到GOPATH的src目录
cd $GOPATH/src
git clone git@github.com:grpc/grpc-go.git ./google.golang.org/grpc
git clone git@github.com:golang/net.git ./golang.org/x/net
git clone git@github.com:googleapis/go-genproto.git ./google.golang.org/genproto
git clone git@github.com:golang/text.git ./golang.org/x/text
go install google.golang.org/grpc
```

测试安装：
```
// 启动服务器
go run google.golang.org/grpc/examples/helloworld/greeter_server/main.go
// 启动客户端
go run google.golang.org/grpc/examples/helloworld/greeter_client/main.go
```

通过下面的命令可以通过`protobuf`生成`grpc`代码：
```
protoc --go_out=plugins=grpc:. hello.proto
```

### goland安装SDK报错

`Goland`配置`GOROOT`报错：`The selected directory is not a valid home for Go Sdk`

原因`goland`版本过低，当`go`版本大于`1.17`时就会报此异常。

编辑自己`go`安装目录下的：`go\src\runtime\internal\sys\zversion.go`，增加一行自己的版本:
```
const TheVersion = `go1.18`
```
