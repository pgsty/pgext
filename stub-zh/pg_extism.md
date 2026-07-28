## 用法

来源：

- [官方上游 README](https://github.com/mhmd-azeez/pg_extism/blob/1f955f47b9853fe594b26dbb0056854639034e1c/README.md)
- [官方扩展控制文件 (pg_extism.control)](https://github.com/mhmd-azeez/pg_extism/blob/1f955f47b9853fe594b26dbb0056854639034e1c/pg_extism.control)
- [官方实现源代码](https://github.com/mhmd-azeez/pg_extism/blob/1f955f47b9853fe594b26dbb0056854639034e1c/src/lib.rs)

`pg_extism` — 一个 Extism 示例，展示了如何使用 pgrx 在 PostgreSQL 中运行 Extism 插件。当数据库代码需要在该过程语言中运行或与其进行交互时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_extism;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小版本，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `extism_call` 是一个扩展函数。
- `extism_define` 是一个扩展函数。
- `to_lowercase` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源代码进行比对。
