## 用法

来源：

- [官方上游 README](https://github.com/godwhoa/pbloom/blob/fc189b436337ae0812800c06903162d0b6ee1fe9/README.md)
- [官方扩展控制文件 (pbloompg.control)](https://github.com/godwhoa/pbloom/blob/fc189b436337ae0812800c06903162d0b6ee1fe9/pg/pbloompg.control)
- [官方实现源代码](https://github.com/godwhoa/pbloom/blob/fc189b436337ae0812800c06903162d0b6ee1fe9/pg/src/lib.rs)

`pbloompg` — 可移植的布隆过滤器，用于在 Go、Rust 和 Postgres 之间创建、序列化和查询。当 SQL 需要这些特殊功能或聚合时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pbloompg;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pbloom_add` 是一个扩展函数。
- `pbloom_contains` 是一个扩展函数。
- `pbloom_create` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
