## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/logfmt/logfmt-1.0.0/README.md)
- [官方扩展控制文件 (logfmt.control)](https://api.pgxn.org/src/logfmt/logfmt-1.0.0/logfmt.control)
- [官方扩展 SQL (logfmt--1.0.0.sql)](https://api.pgxn.org/src/logfmt/logfmt-1.0.0/logfmt--1.0.0.sql)

`logfmt` — logfmt 已知与 PostgreSQL 16beta1 兼容。在收集或解释相应的 PostgreSQL 统计信息时，请使用它。在目标 PostgreSQL 构建上测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION logfmt;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `emit_test_logs()` 是一个扩展函数并返回 `void`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0.0`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
