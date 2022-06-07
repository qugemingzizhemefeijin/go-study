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
```
