## 用法

来源：

- [官方上游 README](https://github.com/theqly/pg_crdt_jsonb/blob/34ae8b73a7c29c16f642d1abd126be4288d7b48a/README.md)
- [官方扩展控制文件 (pg_crdt_jsonb.control)](https://github.com/theqly/pg_crdt_jsonb/blob/34ae8b73a7c29c16f642d1abd126be4288d7b48a/pg_crdt_jsonb.control)
- [官方扩展 SQL (pg_crdt_jsonb--1.0.sql)](https://github.com/theqly/pg_crdt_jsonb/blob/34ae8b73a7c29c16f642d1abd126be4288d7b48a/pg_crdt_jsonb--1.0.sql)

`pg_crdt_jsonb` — 对于顶级数组中的每个元素，都会创建一个自己的时间戳。当应用程序数据需要这种类型、域或其操作符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_crdt_jsonb;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `crdt_jsonb_append(crdt_jsonb, jsonb)` 是一个扩展函数，返回 `crdt_jsonb`。
- `crdt_jsonb_in(cstring)` 是一个扩展函数，返回 `crdt_jsonb`。
- `crdt_jsonb_out(crdt_jsonb)` 是一个扩展函数，返回 `cstring`。
- `crdt_jsonb_recv(internal)` 是一个扩展函数，返回 `crdt_jsonb`。
- `crdt_jsonb_send(crdt_jsonb)` 是一个扩展函数，返回 `bytea`。
- `get_jsonb(crdt_jsonb)` 是一个扩展函数，返回 `jsonb`。
- `crdt_jsonb` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
