## 用法

来源：

- [固定提交的上游 README](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/README.md)
- [固定提交的 Rust 实现](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/src/lib.rs)
- [固定提交的 Cargo 元数据](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/Cargo.toml)
- [固定提交的扩展控制文件](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/s3_io.control)

s3_io 0.1.0 版以内嵌 Rust 函数提供基本的 S3 兼容对象操作：创建存储桶、完整写入与读取对象，以及通过 HEAD 请求检查对象是否存在。它不提供列举、删除、流式传输、分段传输或表导入/导出抽象。

### 环境与示例

PostgreSQL 启动前，其服务环境必须提供 S3_ENDPOINT_URL、AWS_ACCESS_KEY_ID、AWS_SECRET_ACCESS_KEY，以及可选的 AWS_SESSION_TOKEN。实现的默认区域为 us-east-1，并强制使用路径式寻址。

```sql
CREATE EXTENSION s3_io;

SELECT s3_create_bucket('analytics');

SELECT s3_put_object(
    'analytics',
    'hello.txt',
    convert_to('hello from PostgreSQL', 'UTF8'),
    content_type => 'text/plain'
);

SELECT convert_from(
    s3_get_object('analytics', 'hello.txt'),
    'UTF8'
);

SELECT s3_object_exists_lazy('analytics', 'hello.txt');
```

put 函数返回服务端 ETag。get 会把整个响应收集成一个 PostgreSQL bytea 值。

### 当前实现风险

每个后端都创建单线程 Tokio 运行时，并同步阻塞等待网络调用。大对象会占用后端内存，缓慢或不可用的端点会长期占住数据库会话。应设置严格的对象大小、语句超时、并发与网络出口控制；该 API 不适合批量传输。

凭据与客户端会缓存在后端进程内存中。SQL 形式的凭据还可能经活动视图、统计信息、日志或错误报告泄漏，因此应优先使用受保护的服务环境，并撤销所有函数的 PUBLIC 执行权限。只通过固定端点和存储桶的审计封装函数授予访问权限。

审阅的源码使用立即求值的默认参数，因此即使显式传入端点或密钥参数，仍会尝试读取相应环境变量。客户端缓存键还遗漏会话令牌，导致后续其他设置相同的调用可能复用先前客户端。应将它们视为 0.1.0 版缺陷。构建声明了 PostgreSQL 13 至 18 的 pgrx 特性，并要求超级用户安装，但每个目标版本和 S3 兼容服务都需要单独测试。
