## 用法

来源：

- [官方上游文档](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/Readme)
- [1.0 版本 SQL 对象](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/neo4j_fdw--1.0.sql)
- [1.0 版本外部表示例](https://github.com/nuko-yokohama/neo4j_fdw/tree/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/test)

`neo4j_fdw` 是一个 C 语言外部数据封装器，把 Neo4j Cypher 请求返回的行暴露为 PostgreSQL 外部表。它面向 Neo4j 历史 HTTP Cypher 端点，不应与同名的 Python 或 Multicorn 项目混淆。

### 核心流程

先使用端点 URL 创建 server，再定义外部表，其 `query` 选项是 JSON 请求。Cypher `RETURN` 列表必须按位置与 PostgreSQL 列匹配，并使用兼容类型。

```sql
CREATE EXTENSION neo4j_fdw;

CREATE SERVER neo4j_srv
FOREIGN DATA WRAPPER neo4j_fdw
OPTIONS (url 'http://127.0.0.1:7474/db/data/cypher');

CREATE FOREIGN TABLE neo4j_people (
    name text,
    gender text,
    location text
)
SERVER neo4j_srv
OPTIONS (
    query '{"query":"START n=node(*) RETURN n.name AS name, n.gender? AS gender, n.location? AS location"}'
);

SELECT name, location FROM neo4j_people;
```

### 对象与边界

- `neo4j_fdw` 是外部数据封装器；server 选项是 `url`，外部表选项是 `query`。
- `exec_cypher(text, text)` 直接发送端点 URL 与 JSON 请求，并返回 Neo4j JSON 响应。
- 扩展链接 `libcurl` 和 `json-c`。版本 `1.0` 文档没有定义身份验证、TLS 校验、超时、写入支持或规划器下推保证。

示例使用 Neo4j 已过时的 `START` 语法和旧版 `/db/data/cypher` 端点，因此可能与现代 Neo4j 版本不兼容。应把查询文本和 URL 视为只由管理员控制的输入，限制出站网络访问，并在使用前针对确切 Neo4j 版本验证身份验证、转义、NULL 与类型转换、本地与远程过滤、取消和故障行为。
