## 用法

来源：

- [官方上游 README](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/README.md)
- [官方扩展控制文件 (toast.control)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/7_toast/toast.control)
- [官方扩展 SQL (toast--1.0.sql)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/7_toast/sql/toast--1.0.sql)

`toast` — 本教程介绍 TOAST 扩展。在管理或自动化上述数据库行为时使用它。请使用上方链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION toast;

SELECT hello();
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `compare_storage(text, text)` 是一个扩展函数，返回 `void`。
- `force_detoast(text)` 是一个扩展函数，返回 `text`。
- `text_size_info(text)` 是一个扩展函数，返回 `int`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
