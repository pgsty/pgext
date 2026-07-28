## 用法

来源：

- [官方上游 README](https://github.com/kaznak/pgx_basex_codec/blob/8a272c5ac73ad3965e0a37f15e415cc56df33e89/README.md)
- [官方扩展控制文件 (pgx_basex_codec.control)](https://github.com/kaznak/pgx_basex_codec/blob/8a272c5ac73ad3965e0a37f15e415cc56df33e89/pgx_basex_codec.control)
- [官方实现源代码](https://github.com/kaznak/pgx_basex_codec/blob/8a272c5ac73ad3965e0a37f15e415cc56df33e89/src/lib.rs)

`pgx_basex_codec` — 一个使用 pgrx 实现 BaseX 编码和解码的 PostgreSQL 扩展。用于相应的 SQL 或数据库实用工具工作流。请使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgx_basex_codec;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件记录了扩展的版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
