## 用法

来源：

- [官方上游 README](https://github.com/samovers/ofarm2/blob/aa91d6097384838e8ba20efa0dbff41e540364da/deployment/postgresql/ofarm_ed25519/README.md)
- [官方扩展控制文件 (ofarm_ed25519.control)](https://github.com/samovers/ofarm2/blob/aa91d6097384838e8ba20efa0dbff41e540364da/deployment/postgresql/ofarm_ed25519/ofarm_ed25519.control)
- [官方扩展 SQL (ofarm_ed25519--1.0.sql)](https://github.com/samovers/ofarm2/blob/aa91d6097384838e8ba20efa0dbff41e540364da/deployment/postgresql/ofarm_ed25519/ofarm_ed25519--1.0.sql)

`ofarm_ed25519` — 该构建具有封闭的 linux/amd64/linux/arm64 输入集。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION ofarm_ed25519;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `ed25519_verify(public_key pg_catalog.bytea, signed_bytes pg_catalog.bytea, signature pg_catalog.bytea)` 是一个扩展函数，返回 `pg_catalog`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
