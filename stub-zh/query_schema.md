## 用法

来源：

- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/README.md)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/query_schema.control)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/sql/query_schema--1.0.sql)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/src/query-schema.c)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/Dockerfile)

`query_schema` `1.0` 使用 PostgreSQL 的解析器与分析器，将查询结果描述为 Effect 模式源文本。上游将其限定为本地开发工具；现有发行制品面向 Alpine 上的 PostgreSQL 18，提供 amd64 和 arm64 架构。

### 工作流程

以超级用户身份安装到开发数据库。模块不要求预加载或重启服务：

```sql
CREATE EXTENSION query_schema;
CREATE TEMP TABLE example_accounts (id integer NOT NULL, label text);
SELECT query_schema('SELECT id, label FROM example_accounts');
SELECT * FROM query_shape('SELECT id, label FROM example_accounts WHERE id = $1');
```

`query_schema(text)` 返回基于 `Schema.Struct` 的文本表达式。`query_shape(text)` 返回 `path`、`pg_type`、`nullable`、`schema`、`note` 列，包含推断出的参数类型及嵌套 JSON、数组路径。参数行的 `nullable` 值为空。

### 分析边界

每次传入一条语句。实现接受 `SELECT`、`INSERT`、`UPDATE`、`DELETE`、`MERGE`；数据修改语句描述其 `RETURNING` 字段。实用命令和多条语句会被拒绝。函数分析语句，并对查询执行重写，不执行结果计划。

分析使用当前数据库对象和名称解析规则。外连接、条件、聚合、CTE、枚举及 JSON 构造函数都会影响输出。无法确定的可空性保守表示为可空。顶层类型映射遵循 `@effect/sql-pg`，嵌套 JSON 则遵循 PostgreSQL 的 JSON 转换，两种表示可能不同。

这是开发辅助工具，不能证明任意输入安全，也不能保证推断模式完整验证运行时结果。应结合真实客户端解码器复核生成模式和注释。SQL 函数采用 PostgreSQL 的常规执行授权；共享开发数据库前应限制访问。已检查的源码未声明许可证，复用授权尚未核实。
