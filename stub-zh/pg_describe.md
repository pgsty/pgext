## 用法

来源：

- [pg_describe 1.0.0 README](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/README.md)
- [pg_describe 文档](https://sajonaro.github.io/pg_describe/)
- [pg_describe 1.0.0 控制文件](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/pg_describe.control)
- [pg_describe 1.0.0 SQL](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/sql/pg_describe--1.0.0.sql)

`pg_describe` 可以在不执行 SQL 语句的情况下报告其参数和结果列。它使用 PostgreSQL 的解析和分析能力，推断参数类型、线路协议可见的结果类型、源列来源，以及考虑外连接后的可空性。适用于代码生成、迁移检查和查询契约工具。

### 描述查询

```sql
CREATE EXTENSION pg_describe;

SELECT *
FROM pg_describe(
  'SELECT id, email FROM users WHERE id = $1'
);
```

`kind = 'param'` 的行描述 `$1`、`$2` 以及后续参数。`kind = 'column'` 的行描述结果列顺序、名称、类型 OID/名称、源表/列、基础 `NOT NULL` 状态，以及最终表达式是否确定为非空。

### 检查连接可空性

```sql
SELECT *
FROM pg_describe($query$
  SELECT o.id, c.email
  FROM orders AS o
  LEFT JOIN customers AS c ON c.id = o.customer_id
  WHERE o.placed_at >= $1
$query$);
```

即使 `customers.email` 声明为 `NOT NULL`，`result_not_null` 仍为 false，因为左连接可能以空值扩展该行。生成可空客户端类型时，这一区别很有用。

### 执行与安全边界

- 语句会被解析和分析，但不会执行。描述 `DELETE`、易变函数调用或高开销查询不会运行该语句。
- 正常的名称解析和权限检查仍然适用。调用者不能使用 `pg_describe` 检查其自身无权引用的对象。
- 参数类型必须能够从上下文推断；有歧义的 `$n` 参数仍会产生 PostgreSQL 分析错误。
- 结果描述的是 PostgreSQL 分析后的输出，而不是应用稍后组装的动态 SQL。

### 要求与注意事项

- 上游 1.0.0 要求 PostgreSQL 17；PostgreSQL 16 被描述为可能可用但未经测试。Pigsty 软件包面向 PostgreSQL 17 和 18。
- 扩展可重定位，不需要预加载或重启。
- 配套的 `pg-describe-gen` TypeScript 工具是独立的 npm 软件包。PostgreSQL 扩展无需它也能工作。
- 这是一个较新的 API。请在 CI 中固定扩展/工具版本，并在模式迁移时一并审查生成的变更。
