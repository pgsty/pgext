## 用法

来源：

- [官方扩展控制文件](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/condaversion.control)
- [官方上游文档](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/README.md)

`condaversion` — 使用 C++ 实现的 PostgreSQL Conda 包版本号数据类型扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C++`。

```sql
CREATE EXTENSION "condaversion";
SELECT extversion
FROM pg_extension
WHERE extname = 'condaversion';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
