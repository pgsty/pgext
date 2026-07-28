## 用法

来源：

- [官方上游 README](https://github.com/marvinirwin/terminal-doom/blob/053d773e13c9f8bc9c2b9fb1b09f7da1c549640c/build/README.md)
- [官方扩展控制文件 (doom.control)](https://github.com/marvinirwin/terminal-doom/blob/053d773e13c9f8bc9c2b9fb1b09f7da1c549640c/build/doom.control)
- [官方扩展 SQL (doom--0.0.1.sql)](https://github.com/marvinirwin/terminal-doom/blob/053d773e13c9f8bc9c2b9fb1b09f7da1c549640c/build/doom--0.0.1.sql)

`doom` — 我打算尝试将其作为 PostgreSQL 扩展运行，使用表作为输入和输出。当可用时，请运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 核心工作流

```sql
CREATE EXTENSION doom;
```

在目标数据库中安装扩展，运行上游提供的最小示例（如果可用），并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `doom()` 是一个扩展函数，并返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定来源。
