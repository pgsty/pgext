## 用法

来源：

- [官方上游 README](https://github.com/changtonghf/pgyaml/blob/6472c3a2b0fd7d2c6affc735f358fb6ef970ee0a/README.md)
- [官方扩展控制文件 (pgyaml.control)](https://github.com/changtonghf/pgyaml/blob/6472c3a2b0fd7d2c6affc735f358fb6ef970ee0a/pgyaml.control)
- [官方扩展 SQL (pgyaml--1.0.sql)](https://github.com/changtonghf/pgyaml/blob/6472c3a2b0fd7d2c6affc735f358fb6ef970ee0a/pgyaml--1.0.sql)

`pgyaml` — 在 YAML 和 jsonb 之间进行转换。使用它来执行相应的 SQL 或数据库实用程序工作流。在目标 PostgreSQL 构建上测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pgyaml;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小 SQL 代码，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `yaml_to_jsonb(text)` 是一个扩展函数，返回 `jsonb`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
