## 用法

来源：

- [官方上游 README](https://github.com/hydradb/grpc_fdw/blob/a594a3574b51b7beb79a10abbd522324ed35dc6e/README.md)
- [官方扩展控制文件 (grpc_fdw.control)](https://github.com/hydradb/grpc_fdw/blob/a594a3574b51b7beb79a10abbd522324ed35dc6e/grpc_fdw.control)
- [官方实现源代码](https://github.com/hydradb/grpc_fdw/blob/a594a3574b51b7beb79a10abbd522324ed35dc6e/src/lib.rs)

`grpc_fdw` — 外部数据包装器，用于将表操作委托给 gRPC 服务。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION grpc_fdw;

CREATE FOREIGN DATA WRAPPER grpc_fdw_handler HANDLER grpc_fdw_handler NO VALIDATOR;
CREATE SERVER user_srv FOREIGN DATA WRAPPER grpc_fdw_handler OPTIONS (server_uri 'http://[::1]:50051');
CREATE FOREIGN TABLE users (
    id integer,
    name text,
    email text
) SERVER user_srv OPTIONS (
    table_option '1',
    table_option2 '2'
);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `grpc_fdw_handler()` 是一个扩展函数，返回 `pg_sys::Datum`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
