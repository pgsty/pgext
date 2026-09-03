## 用法

来源：

- [EDB SPL Check 官方文档](https://www.enterprisedb.com/docs/pg_extensions/spl_check/)
- [官方安装指南](https://www.enterprisedb.com/docs/pg_extensions/spl_check/installing/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/spl_check/configuring/)

`spl_check` 对 EPAS SPL 与 PL/pgSQL routine 进行静态检查，发现 type error、无效 object reference、dead code、缺失 return、可疑 cast 与部分 SQL-injection pattern。

### 启用

安装匹配的 EPAS 软件包并创建扩展：

```sql
CREATE EXTENSION spl_check;
```

Active/manual check 不要求预加载。该 provider-only 扩展支持 EDB Postgres Advanced Server。

### 检查 Routine

```sql
SELECT *
FROM spl_check_function_tb('public.f1()');

SELECT *
FROM spl_check_function(
  'public.f1()',
  fatal_errors := false
);
```

`spl_check_function` 支持 text、JSON 与 XML output。当 routine name 存在重载时，应传入完整 signature 或 `regprocedure`。

### Passive Mode 与审查

Passive mode 会在 routine 执行前检查并可提示 compatibility warning，但静态分析无法证明 runtime correctness，也无法完全证明 dynamic SQL 安全。应把发现项作为审查输入，而不是自动改写。诊断可能暴露 SQL text 与 object name，应限制可检查 routine definition 的角色，并在应用 workload 全面启用 passive check 前测试 warning-level 变化。

