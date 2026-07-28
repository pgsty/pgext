## 用法

来源：

- [官方上游 README](https://github.com/rom8726/fuzzysearch/blob/b61c645cd0916ca4b7d077e4c5e4e73abcadaa5f/README.md)
- [官方扩展控制文件 (fuzzysearch.control)](https://github.com/rom8726/fuzzysearch/blob/b61c645cd0916ca4b7d077e4c5e4e73abcadaa5f/fuzzysearch.control)
- [官方扩展 SQL (fuzzysearch--1.0.sql)](https://github.com/rom8726/fuzzysearch/blob/b61c645cd0916ca4b7d077e4c5e4e73abcadaa5f/fuzzysearch--1.0.sql)

`fuzzysearch` — PostgreSQL 扩展，用于字符串模糊匹配。使用它来进行相应的文本搜索、解析或语言学工作流。请使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION fuzzysearch;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `jaro_winkler_match(text, text)` 是一个扩展函数，返回 `float`。
- `levenshtein_match(text, text)` 是一个扩展函数，返回 `integer`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
