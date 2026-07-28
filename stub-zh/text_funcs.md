## 用法

来源：

- [官方上游 README](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/README.md)
- [官方扩展控制文件 (text_funcs.control)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/3_text_funcs/text_funcs.control)
- [官方扩展 SQL (text_funcs--1.0.sql)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/3_text_funcs/sql/text_funcs--1.0.sql)

`text_funcs` — 文本函数教程。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION text_funcs;

SELECT hello();
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `reverse_text(text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
