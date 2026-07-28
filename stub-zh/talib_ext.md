## 用法

来源：

- [官方扩展控制文件（talib_ext.control）](https://github.com/sachaarbonel/pg_talib-rs/blob/fcc0ce3cb1475d8f12be7f2421f19b2e4e89384b/talib_ext/talib_ext.control)
- [官方实现源代码](https://github.com/sachaarbonel/pg_talib-rs/blob/fcc0ce3cb1475d8f12be7f2421f19b2e4e89384b/talib_ext/src/lib.rs)
- [官方 Rust 包清单](https://github.com/sachaarbonel/pg_talib-rs/blob/fcc0ce3cb1475d8f12be7f2421f19b2e4e89384b/talib_ext/Cargo.toml)

`talib_ext` — 用于在 PostgreSQL 中进行技术分析的扩展。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION talib_ext;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hello_talib_ext()` 是一个扩展函数。
- `macd` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
