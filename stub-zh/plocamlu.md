## 用法

来源：

- [官方扩展控制文件 (plocamlu.control)](https://github.com/higuoxing/plocaml/blob/6cbd404a94fa659785a3b6891a1cd1ea9c180594/plocamlu.control)
- [官方扩展 SQL (plocamlu--1.0.sql)](https://github.com/higuoxing/plocaml/blob/6cbd404a94fa659785a3b6891a1cd1ea9c180594/plocamlu--1.0.sql)

`plocamlu` — PL/OCaml 程序语言处理器用于 PostgreSQL。当数据库代码需要在该程序语言中运行或与其进行交互时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION plocamlu;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码片段，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `plocamlu_call_handler()` 是一个扩展函数，返回 `language_handler`。
- `plocamlu_inline_handler(internal)` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
