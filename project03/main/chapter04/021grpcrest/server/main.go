package main

import (
	hs "../service"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

// gRPC 服务一般用于集群内部通信，如果需要对外暴露服务一般会提供等价的 REST 接口。
// 通过 REST 接口比较方便前端 JavaScript 和后端交互。开源社区中的 grpc-gateway 项目就实现了将 gRPC 服务转为 REST 服务的能力。

// 我们首先为 gRPC 定义了 Get 和 Post 方法，然后通过元扩展语法在对应的方法后添加路由信息。
// 其中 “/get/{value}” 路径对应的是 Get 方法，{value} 部分对应参数中的 value 成员，结果通过 json 格式返回。
// Post 方法对应 “/post” 路径，body 中包含 json 格式的请求信息。

// 然后通过以下命令安装 protoc-gen-grpc-gateway 插件：
// go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
// go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

// 当使用命令编译的时候提示找不到包：
// google/api/annotations.proto: File not found.
// test.proto:5:1: Import "google/api/annotations.proto" was not found or had errors.

// 去github上将对应的包下载下来放在$GOPATH/src下，例如这里缺失google/api。
// 去gooogleapis将项目下载下来，并将googles目录整个放到$GOPATH/src

// 首先通过 runtime.NewServeMux() 函数创建路由处理器，然后通过 RegisterRestServiceHandlerFromEndpoint 函数
// 将 RestService 服务相关的 REST 接口中转到后面的 gRPC 服务。
// grpc-gateway 提供的 runtime.ServeMux 类也实现了 http.Handler 接口，因此可以和标准库中的相关函数配合使用。

// 当 gRPC 和 REST 服务全部启动之后，就可以用 curl 请求 REST 服务了：
// curl localhost:8080/get/gopher
// curl localhost:8080/post -X POST --data '{"value":"grpc"}'

// 在对外公布 REST 接口时，我们一般还会提供一个 Swagger 格式的文件用于描述这个接口规范。
// go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
// protoc -I. \
//  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
//  --swagger_out=. \
//  hello.proto

// 然后会生成一个 hello.swagger.json 文件。这样的话就可以通过 swagger-ui 这个项目，在网页中提供 REST 接口的文档和测试等功能。

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := hs.RegisterRestServiceHandlerFromEndpoint(
		ctx, mux, "localhost:5000",
		[]grpc.DialOption{grpc.WithInsecure()},
		)

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", mux)
}