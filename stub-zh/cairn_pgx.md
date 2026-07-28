## 用法

来源：

- [官方上游 README](https://github.com/cairn-ehr/cairn-ehr/blob/bf5d2cf13c493edcba84685f2f8714851f64ea6e/README.md)
- [官方扩展控制文件 (cairn_pgx.control)](https://github.com/cairn-ehr/cairn-ehr/blob/bf5d2cf13c493edcba84685f2f8714851f64ea6e/extensions/cairn_pgx/cairn_pgx.control)
- [官方实现源代码](https://github.com/cairn-ehr/cairn-ehr/blob/bf5d2cf13c493edcba84685f2f8714851f64ea6e/extensions/cairn_pgx/src/lib.rs)

`cairn_pgx` — Cairn 在数据库中验证门 — COSE Sign1/Ed25519 事件验证（Spike 0002）。在实现相应的安全、审计或访问控制工作流时使用它。上游将其描述为一个概念证明。

### 核心工作流

```sql
CREATE EXTENSION cairn_pgx;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `cairn_actor_id` 是一个扩展函数。
- `cairn_attestation_ok` 是一个扩展函数。
- `cairn_blob_verify` 是一个扩展函数。
- `cairn_blob_verify_error` 是一个扩展函数。
- `cairn_body` 是一个扩展函数。
- `cairn_pgx_version()` 是一个扩展函数。
- `cairn_unseal_body` 是一个扩展函数。
- `cairn_verify` 是一个扩展函数。
- `cairn_verify_error` 是一个扩展函数。
- `cairn_wrap_dek` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.3.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 上游将该项目描述为一个概念证明。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
