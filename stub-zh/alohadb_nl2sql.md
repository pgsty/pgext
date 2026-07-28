## 用法

来源：

- [官方上游 README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [官方扩展控制文件 (alohadb_nl2sql.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_nl2sql/alohadb_nl2sql.control)
- [官方扩展 SQL (alohadb_nl2sql--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_nl2sql/alohadb_nl2sql--1.0.sql)

`alohadb_nl2sql` — 通过 LLM API 进行自然语言到 SQL 的翻译。用于相应的向量、模型或检索工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION alohadb_nl2sql;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `alohadb_explain_query(sql text)` 是一个扩展函数，返回 `text`。
- `alohadb_nl2sql(question text)` 是一个扩展函数，返回 `text`。
- `alohadb_nl2sql_execute(question text)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
