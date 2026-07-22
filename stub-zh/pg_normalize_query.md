## 用法

来源：

- [官方 README](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/README.md)
- [官方扩展 SQL](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/pg_normalize_query--1.0.sql)
- [官方实现](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/pg_normalize_query.c)

`pg_normalize_query` 将有效 SQL 语句中的常量替换为编号参数，生成类似 `pg_stat_statements` 归一化结果的指纹文本。它适合在日志或应用遥测中按语句形态聚合查询。

### 核心流程

```sql
CREATE EXTENSION pg_normalize_query;

SELECT pg_normalize_query(
  $$SELECT * FROM pg_proc WHERE proname = 'pg_normalize_query'$$
);
-- SELECT * FROM pg_proc WHERE proname = $1

SELECT pg_normalize_query(
  $$SELECT oid FROM pg_class WHERE relkind IN ('r', 'i') LIMIT 10$$
);
-- SELECT oid FROM pg_class WHERE relkind IN ($1, $2) LIMIT $3
```

### API

扩展只公开一个 strict、immutable 函数：`pg_normalize_query(query text) RETURNS text`。实现会调用 PostgreSQL 原始解析器，在语法树中遍历常量位置，再把对应的词法片段重写为 `$1`、`$2` 等参数。编号替换时会计入输入中已有的外部参数。

### 解释与兼容性

输入必须对构建该扩展时使用的 PostgreSQL 解析器语法有效；解析错误会正常抛出。结果保留大部分原始拼写和空白，因此它只是常量归一化，而不是规范化 SQL 格式器。实现还明确指出，等价查询仍可能产生不同的归一化表示。

本次复核的上游快照面向 PostgreSQL `9.4` 到预发布 `13`。由于它复制了历史 PostgreSQL/libpg_query 中与解析器紧密相关的逻辑，在更新的服务端大版本上依赖跨升级稳定指纹之前，应同时验证编译结果和有代表性的语法。
