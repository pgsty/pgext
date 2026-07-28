## 用法

来源：

- [官方扩展控制文件 (rsam.control)](https://github.com/nekit2-002/rsam/blob/5052ad151b6442f5b066c6b47f9e0448fe9506e0/postgres-rust-table-am/rsam.control)
- [官方实现源代码](https://github.com/nekit2-002/rsam/blob/5052ad151b6442f5b066c6b47f9e0448fe9506e0/postgres-rust-table-am/src/lib.rs)
- [官方 Rust 包清单](https://github.com/nekit2-002/rsam/blob/5052ad151b6442f5b066c6b47f9e0448fe9506e0/postgres-rust-table-am/Cargo.toml)

`rsam` — 一个用于 Postgresql 的新访问方法仓库，用 Rust 编写。当应用程序需要此特定数据库功能时，请使用它。在目标 PostgreSQL 构建上测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION rsam;
```

在目标数据库中安装扩展，在可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 元数据记录版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
