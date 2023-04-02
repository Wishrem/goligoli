# goligoli

## User模块
第一次用rpc的尝试
设计混乱，为了实现功能而写的代码

## Video模块
在User模块上进行了大量的改进

将业务逻辑进行更细致的拆分，文件目录说明如下
```
api-gateway # 网关，用于统一接收所有http请求，并调用handler方法，向其他对应的业务服务器发送请求
|-- cmd # 程序入口
|-- handler # http请求参数处理，检验所需参数和调用service方法，向对应业务服务器发送请求
|-- router # 路由
|-- service # grpc的客户端，用于向对应业务服务器发送请求

video # video模块，接收video有关的业务请求（comment等模块于video一致，不包括user）
|-- cmd # 程序入口
|-- model # 持久层，用于与数据库进行相关交互
|-- proto # proto文件和grpc生成的go代码
    |-- pb # grpc生成的go代码
|-- service # 业务逻辑层，用于响应api-gateway中service发出的请求，在这里组装业务逻辑，并装载数据调用model的方法，对数据库进行操作
```

设计了以下模块：
1. erp模块
用于统一返回前端错误码，屏蔽程序内部的错误
2. logger模块（依托答辩）
用于记录程序内部错误

## Comment等模块于Video模块雷同

## 流程图
![Process]()

## 数据关系
![ER]()