# 设置项目
为您的项目创建一个目录，并将其初始化为 Go Module：

```shell
$ mkdir gqlgen-todos
$ cd gqlgen-todos
$ go mod init github.com/[username]/gqlgen-todos
$ go get github.com/99designs/gqlgen
```

# 构建服务器
## 创建项目骨架
$ go run github.com/99designs/gqlgen init
这将创建默认的包布局。如果需要，您可以在 gqlgen.yml 中修改这些路径。

├── go.mod
├── go.sum
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
├── graph
│   ├── generated            - A package that only contains the generated runtime
│   │   └── generated.go
│   ├── model                - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go
│   ├── resolver.go          - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go  - the resolver implementation for schema.graphql
└── server.go                - The entry point to your app. Customize it however you see fit
## 定义您的架构
在编写代码之前，使用 GraphQL模式定义语言来描述你的 API 。在schema目录下定义 .graphqls文件定义 model、input、Query、Mutation

```graphql
type Mutation {
  createUser(input: NewUser!): User!
}
```
然后执行如下命令生成代码到model、resolver等目录
```shell
go run -mod=mod github.com/99designs/gqlgen generate
```

- mutation是 增删改的行为，里面包含对数据改动的方法
- input是mutation方法中的输入参数，也就是前端给后端的，后面的User是后端返回的所有数据，前端可以在查询语句中自定义返回想要的字段
- Query 是读取的行为，里面有查询数据的方法
##  实现解析器
上面提到的query和mutation的方法 都会自动生成函数名在resolve.go 中，只需要我们改动函数的逻辑即可。
```go
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
    //panic(fmt.Errorf("not implemented: CreateUser - createUser"))
    coll := r.DB.Database(db).Collection("user")
    u := &model.User{
    ID:        uuid.New().String(),
    Name:      input.Name,
    CreatedOn: int(time.Now().Unix()),
    UpdatedOn: int(time.Now().Unix()),
    Deleted:   false,
    IsActive:  input.IsActive,
    }
    _, err := coll.InsertOne(ctx, u)
    if err != nil {
    return nil, err
    }
    return u, nil
}

```

# 查询方式

客户端可以打开http://localhost:9090/ 界面进行查询 或者postman选graphql 进行查询

```graphql endpoint
query{
    getUser(id:"xx"){
    id
    name
    }
}

mutation{
    createUser(input:{
    id:"xx",
    name: "my_user01",
    isActive:true
    }) {
    id
    name
    isActive
    updatedOn

}
}

```

