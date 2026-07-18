## 用法

来源：

- [已归档的上游仓库与 README](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/README.adoc)
- [源码 control 文件](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/neo4j-fdw.control)
- [最新正式 PGXN 版本](https://pgxn.org/dist/neo4j-fdw/)

`neo4j-fdw` 是一个 Python 集成，通过 Multicorn 把 Neo4j Cypher 查询结果暴露为 PostgreSQL 外部表。它可以把基本的比较与模式匹配条件下推到动态生成的 Cypher 查询中。

安装 `neo4jPg` Python 包、与其匹配的 Neo4j Python driver 以及 Multicorn 后，定义 server 和外部表：

```sql
CREATE EXTENSION multicorn;

CREATE SERVER multicorn_neo4j
FOREIGN DATA WRAPPER multicorn
OPTIONS (
    wrapper 'neo4jPg.neo4jfdw.Neo4jForeignDataWrapper',
    url 'neo4j://neo4j.example:7687',
    user 'neo4j',
    password 'replace-me'
);

CREATE FOREIGN TABLE neo4j_movie (movie text)
SERVER multicorn_neo4j
OPTIONS (
    cypher 'MATCH (n:Movie) RETURN n.title AS movie',
    database 'neo4j'
);

SELECT * FROM neo4j_movie;
```

外部表使用的 Cypher 表达式必须返回 collection，并为每个返回字段指定 alias。凭据存放在 foreign server 选项中，因此应相应限制系统目录可见性和对象权限。

### 归档状态与版本

该仓库于 2023 年归档。最新正式 PGXN 版本是 `4.3.1`，最终源码 control 文件则标记为 `4.3.2-dev`，与本目录采用的源码快照一致。上游兼容性文档只覆盖较旧的 Multicorn、Python、Neo4j 与 PostgreSQL 组合，其中 PostgreSQL 范围止于 12。应将其视为遗留软件，在仍受维护的 PostgreSQL 版本上使用前，必须验证完整依赖栈。
