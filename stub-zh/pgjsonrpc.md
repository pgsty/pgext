## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pgjsonrpc/pgjsonrpc-0.0.1/README.md)
- [官方扩展控制文件 (pgjsonrpc.control)](https://api.pgxn.org/src/pgjsonrpc/pgjsonrpc-0.0.1/pgjsonrpc.control)
- [官方扩展 SQL (pgjsonrpc.sql)](https://api.pgxn.org/src/pgjsonrpc/pgjsonrpc-0.0.1/sql/pgjsonrpc.sql)

`pgjsonrpc` — JSON-RPC 2.0 实现作为 PostgreSQL 扩展。用于相应的 SQL 或数据库实用程序工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgjsonrpc;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `jsonrpc.echo(p_request JSON)` 是一个扩展函数，返回 `JSON`。
- `jsonrpc.error_response(p_request JSON, p_code INTEGER, p_message TEXT, p_data JSON)` 是一个扩展函数，返回 `JSON`。
- `jsonrpc.execute(p_request TEXT)` 是一个扩展函数，返回 `JSON`。
- `jsonrpc.get_response(p_request JSON, p_jsonrpc TEXT, p_id TEXT, p_result JSON, p_code INTEGER, p_message TEXT)` 是一个扩展函数，返回 `JSON`。
- `jsonrpc.success_response(p_request JSON, p_result ANYELEMENT)` 是一个扩展函数，返回 `JSON`。
- `jsonrpc.methods` 是一个由扩展安装或管理的表。
- `jsonrpc` 是一个由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
