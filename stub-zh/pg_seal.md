## 用法

来源：

- [官方上游 README](https://github.com/allenvox/pg_seal/blob/3658d3608ba4ea3867cc0758149fed75730b88e3/README.md)
- [官方扩展控制文件 (pg_seal.control)](https://github.com/allenvox/pg_seal/blob/3658d3608ba4ea3867cc0758149fed75730b88e3/pg_seal.control)
- [官方实现源代码](https://github.com/allenvox/pg_seal/blob/3658d3608ba4ea3867cc0758149fed75730b88e3/src/lib.rs)

`pg_seal` — 一个用 Rust 编写的 PostgreSQL 扩展（pgrx），用于数据变更的加密审计日志。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_seal;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pg_seal_append` 是一个扩展函数。
- `pg_seal_hash` 是一个扩展函数。
- `pg_seal_verify()` 是一个扩展函数。
- `pg_seal_verify_detail()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `1.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
