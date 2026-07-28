## 用法

来源：

- [官方上游 README](https://github.com/bkircher/pg_base62/blob/19f8807844d59f44a776f43752b1c23635ccc5e2/README.md)
- [官方扩展控制文件 (pg_base62.control)](https://github.com/bkircher/pg_base62/blob/19f8807844d59f44a776f43752b1c23635ccc5e2/pg_base62.control)
- [官方实现源代码](https://github.com/bkircher/pg_base62/blob/19f8807844d59f44a776f43752b1c23635ccc5e2/src/lib.rs)

`pg_base62` — 一个用于将 UUID 编码为 Base62 字母表并将其解码回 UUID 的 PostgreSQL 扩展。它使用 pgrx 用 Rust 编写。在相应的 SQL 或数据库实用工具工作流中使用它。请使用上述链接的上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_base62;

SELECT base62_encode('f81d4fae-7dec-11d0-a765-00a0c91e6bf6'::uuid);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `base62_decode` 是一个扩展函数。
- `base62_encode` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
