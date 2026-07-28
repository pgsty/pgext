## 用法

来源：

- [官方上游 README](https://github.com/serxka/tag_parser/blob/18280fc48cc212866a864b7f1741d7b8c75d12d8/Readme.md)
- [官方扩展控制文件 (tag_parser.control)](https://github.com/serxka/tag_parser/blob/18280fc48cc212866a864b7f1741d7b8c75d12d8/tag_parser.control)
- [官方扩展 SQL (tag_parser--0.1.0.sql)](https://github.com/serxka/tag_parser/blob/18280fc48cc212866a864b7f1741d7b8c75d12d8/tag_parser--0.1.0.sql)

`tag_parser` — 一个用 C 编写的简单全文搜索解析器，适用于 PostgreSQL。它比默认的 tsearch 解析器要简单得多，只会根据逗号边界拆分词素。可以通过更改 BREAK_CHAR 的定义来将其更改为任何其他字符。使用它来进行相应的文本搜索、解析或语言学工作流。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION tag_parser;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `tagpsr_end(internal)` 是一个扩展函数，返回 `void`。
- `tagpsr_gettoken(internal, internal, internal)` 是一个扩展函数，返回 `internal`。
- `tagpsr_lextype(internal)` 是一个扩展函数，返回 `internal`。
- `tagpsr_start(internal, int4)` 是一个扩展函数，返回 `internal`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
