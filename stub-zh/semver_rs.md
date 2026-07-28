## 用法

来源：

- [官方上游 README](https://github.com/fabmation-gmbh/pg_semver-rs/blob/0bc1aa00db74b824852027cbc6a369a5ccbd3f10/README.md)
- [官方扩展控制文件 (semver_rs.control)](https://github.com/fabmation-gmbh/pg_semver-rs/blob/0bc1aa00db74b824852027cbc6a369a5ccbd3f10/semver_rs.control)
- [官方实现源码](https://github.com/fabmation-gmbh/pg_semver-rs/blob/0bc1aa00db74b824852027cbc6a369a5ccbd3f10/src/lib.rs)

`semver_rs` — The awesome [pg-semver][] 扩展但实现于 Rust。它支持 [pg-semver][] 所有操作。当应用程序数据需要此类型、领域或其运算符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION semver_rs;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 表记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
