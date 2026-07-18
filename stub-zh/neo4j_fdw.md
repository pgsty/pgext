## 用法

来源：

- [上游用法说明](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/Readme)
- [1.0 版本 SQL 对象](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/neo4j_fdw--1.0.sql)
- [回归示例](https://github.com/nuko-yokohama/neo4j_fdw/tree/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/test)

`neo4j_fdw` 是把 Neo4j Cypher 结果暴露为 PostgreSQL 外部表的 C 语言 FDW。server 提供 Neo4j URL，外部表列携带 Cypher 查询与预期结果列；查询的 `RETURN` 列表必须与 PostgreSQL 定义匹配。

```sql
CREATE EXTENSION neo4j_fdw;
CREATE SERVER neo4j_srv FOREIGN DATA WRAPPER neo4j_fdw
OPTIONS (url 'http://127.0.0.1:7474');
```

该扩展链接 `libcurl` 与 `json-c`，还提供临时 JSON 查询路径。它是较早的纯源码实现，与同名 Python/Multicorn 项目不同。不要让不可信用户控制凭据、Cypher 文本或选项；应在外部落实网络超时与 TLS，并在使用前验证类型、NULL、转义、下推和远端故障行为。
