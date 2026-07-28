## 用法

来源：

- [官方扩展控制文件 (hello_pg.control)](https://github.com/stevelauc/debug_macos_build_failure_pgrx/blob/98ffdcf0dc09d2f657c43f36f5d48778c4ae2665/hello_pg.control)
- [官方实现源代码](https://github.com/stevelauc/debug_macos_build_failure_pgrx/blob/98ffdcf0dc09d2f657c43f36f5d48778c4ae2665/src/lib.rs)
- [官方 Rust 包清单](https://github.com/stevelauc/debug_macos_build_failure_pgrx/blob/98ffdcf0dc09d2f657c43f36f5d48778c4ae2665/Cargo.toml)

`hello_pg` — 最小的 pgrx hello-world 扩展，用于重现 macOS 构建问题。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION hello_pg;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hello_hello_pg()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
