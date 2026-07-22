## 用法

来源：

- [1.0 版 SQL API](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc--1.0.sql)
- [gRPC C-core 实现](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc.c)
- [扩展控制文件](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc.control)

`pg_grpc` 是围绕少数 gRPC C-core 调用构建的实验性、极低层包装器。它在后端进程全局变量中维护一个通道、调用与完成队列，只提供返回布尔值的分步函数；它不提供 protobuf 编码、服务 stub、响应数据、状态详情、截止时间、元数据或完整 RPC 抽象。

### 已暴露流程

可调用顺序是创建通道、为方法创建一次调用、提交原始文本载荷，再推进完成队列。

```sql
CREATE EXTENSION pg_grpc;

SELECT grpc_insecure_channel_create('127.0.0.1:50051');
SELECT grpc_channel_create_call('/example.Service/Method');
SELECT grpc_call_start_batch('raw request bytes');
```

SQL API 还声明了 `grpc_secure_channel_create(target)` 与 `grpc_completion_queue_next()`。每个函数都只返回 `boolean`；没有任何 SQL 函数会返回收到的消息。乱序调用某一步会触发扩展错误。

### 原型与安全边界

所审阅实现把 NULL 凭证指针传给安全通道构造函数，使用无限调用截止时间，在每个后端保存可变状态，并让 `grpc_completion_queue_next()` 循环直到队列关闭。这些路径可能失败、无限期阻塞数据库会话，或留下不适合后续调用的状态。批处理代码既不向 SQL 暴露 protobuf 序列化，也不返回接收缓冲区，因此成功的 `true` 不能证明已经得到应用层 RPC 响应。

所有函数都能从 PostgreSQL 服务端进程发起出站网络活动，并且管理员若不主动撤销，默认会授予 `PUBLIC` 执行权。应撤销执行权限、仅允许可信目的地，并且不要把该原型用于生产 RPC 或不可信输入。应优先采用维护良好的外部服务，或能强制 TLS 凭证、截止时间、载荷限制、取消及响应/状态处理的完整扩展。
