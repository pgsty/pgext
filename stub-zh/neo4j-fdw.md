## 用法

来源：

- [已复核 commit 的归档官方 README](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/README.adoc)
- [已复核 commit 的随附 SQL](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/sql/neo4j-fdw.sql)
- [Multicorn 包装器实现](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/neo4jPg/neo4jfdw.py)
- [扩展 control 文件](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/neo4j-fdw.control)
- [官方 PGXN 发行页](https://pgxn.org/dist/neo4j-fdw/)

`neo4j-fdw` 是一个已归档的 Python 集成，通过 Multicorn 把 Neo4j Cypher 结果暴露为 PostgreSQL 外部表。它不是原生 FDW 库：PostgreSQL 服务器进程会导入单独安装的 `neo4jPg` Python 包、Neo4j Python 驱动和 Multicorn。

### 核心流程

安装彼此兼容的 Python、Multicorn 与 Neo4j 驱动栈后，创建 Multicorn 服务器，并把一个带别名的 Cypher 结果映射为一张外部表：

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

SELECT * FROM neo4j_movie WHERE movie ILIKE '%matrix%';
```

Cypher 表达式必须返回集合，并把每个返回字段取别名为外部表列。简单比较与模式条件可被重写为 Cypher；可选且格式严格的 `/*WHERE{...}*/` 占位符可告诉包装器在何处更高效地插入条件。

### 打包与兼容性边界

规范项目名是 `neo4j-fdw`，但已复核的源码并不是可靠的现代扩展包。control 文件报告 `4.3.2-dev`，仓库安装的却是未版本化 SQL 文件，而不是匹配的版本化扩展脚本。该 SQL 还请求旧式 `plpythonu` 扩展，却把函数声明为 `plpython3u`。应验证并修补打包，而不能假定 `CREATE EXTENSION "neo4j-fdw"` 可以成功。

上述 FDW 表路径需要 Multicorn 和 Python 包；可选的直接 `cypher()` SQL 辅助函数还需要不可信 PL/Python。上游声明的 PostgreSQL 兼容范围止于 12，仓库已于 2023 年归档。凭据保存在外部服务器选项中。应限制目录与服务器权限，使用专用 Neo4j 账户和加密传输，并针对确切的 Neo4j 驱动与服务器组合测试查询生成。
