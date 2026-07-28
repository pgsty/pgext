## 用法

来源：

- [官方上游 README](https://github.com/eugwne/pllj/blob/d5d63a265b60200a7606b5cb5c92824837fd95cf/README.md)
- [官方扩展控制文件 (pllju.control)](https://github.com/eugwne/pllj/blob/d5d63a265b60200a7606b5cb5c92824837fd95cf/pllju.control)
- [官方扩展 SQL (pllju--0.1.sql)](https://github.com/eugwne/pllj/blob/d5d63a265b60200a7606b5cb5c92824837fd95cf/pllju--0.1.sql)

`pllju` — LuaJIT FFI PostgreSQL 语言扩展。当数据库代码必须在该过程语言中运行或与其进行交互时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pllju;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pllj_call_handler_u()` 是一个扩展函数，返回 `language_handler`。
- `pllj_inline_handler_u(internal)` 是一个扩展函数，返回 `VOID`。
- `pllj_validator_u(oid)` 是一个扩展函数，返回 `VOID`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
