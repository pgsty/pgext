## 用法

来源：

- [官方上游 README](https://github.com/jerrysievert/pllisp/blob/219df653a6b736626e45898e819852949612af93/README.md)
- [官方扩展控制文件 (pllisp.control)](https://github.com/jerrysievert/pllisp/blob/219df653a6b736626e45898e819852949612af93/pllisp.control)
- [官方扩展 SQL (pllisp--0.1.0.sql)](https://github.com/jerrysievert/pllisp/blob/219df653a6b736626e45898e819852949612af93/pllisp--0.1.0.sql)

`pllisp` — 该项目最初只是一个关于构建 Postgres 扩展的演示中开玩笑的产物，然后我决定继续开发它。当数据库代码必须在或与该过程语言进行交互时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pllisp;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pllisp_call_handler()` 是一个扩展函数，返回 `language_handler`。
- `pllisp_inline_handler(internal)` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
