## 用法

来源：

- [官方 README](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/README.md)
- [扩展 SQL 实现](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/generic_plan--1.0.sql)
- [PostgreSQL 类型转换规则](https://www.postgresql.org/docs/current/typeconv.html)

`generic_plan` 1.0 无需参数值即可为参数化 SQL 语句生成 `EXPLAIN` 输出。它适合分析从 `pg_stat_statements` 或日志复制的语句，但只返回通用估算计划，绝不会产生执行测量值。

### 核心流程

```sql
CREATE EXTENSION generic_plan;

SELECT plan
FROM generic_plan(
  'SELECT * FROM orders WHERE customer_id = $1',
  verbose  => true,
  costs    => true,
  settings => true,
  format   => 'TEXT'
);
```

`generic_plan(query, verbose, costs, settings, format)` 每个输出行返回一条文本记录。支持 `TEXT`、`XML`、`JSON` 和 `YAML` 格式。函数用 `unknown` 参数类型预备语句，设置 `plan_cache_mode=force_generic_plan`，执行不带 `ANALYZE` 的 `EXPLAIN`，然后释放临时语句。

### 解释与限制

因为所有占位符都从 `unknown` 开始，PostgreSQL 可能选择与真实客户端语句不同的重载或操作符，也可能因歧义失败。生成计划忽略实际参数选择率、运行时间、真实行数、缓冲区和 I/O。应把它当作诊断线索，并尽量用真实类型语句复现。

实现统计的是占位符出现次数，而不是不同参数编号。重复占位符、只用 `$2` 而没有 `$1` 的编号空洞、美元引用字符串以及 SQL 注释中的 `$n` 文本都可能生成错误参数列表或报错。它只识别普通单引号字符串，拒绝中间分号，并使用会话预备语句名 `_stmt_`，可能与已有同名语句冲突。函数是 `SECURITY INVOKER`；调用者仍需拥有解析和规划所传 SQL 所需的权限。
