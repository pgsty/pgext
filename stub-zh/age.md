


## 用法

> [age: AGE 图数据库扩展](https://github.com/apache/age)

Apache AGE 为 PostgreSQL 带来了图数据库能力，使用 openCypher 查询语言。它支持 SQL 与 Cypher 的混合查询、顶点和边上的属性索引，以及多图查询。

每个会话需要加载扩展：

```sql
CREATE EXTENSION age;
LOAD 'age';
SET search_path = ag_catalog, "$user", public;
```

### 图操作

创建图：

```sql
SELECT create_graph('my_graph');
```

创建顶点：

```sql
SELECT * FROM cypher('my_graph', $$
    CREATE (:Person {name: 'Alice', age: 30})
$$) AS (v agtype);

SELECT * FROM cypher('my_graph', $$
    CREATE (:Person {name: 'Bob', age: 25})
$$) AS (v agtype);
```

创建边：

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (a:Person), (b:Person)
    WHERE a.name = 'Alice' AND b.name = 'Bob'
    CREATE (a)-[e:KNOWS {since: 2020}]->(b)
    RETURN e
$$) AS (e agtype);
```

查询图：

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (v)-[r]-(v2)
    RETURN v, r, v2
$$) AS (v agtype, r agtype, v2 agtype);
```

### Cypher 查询功能

AGE 支持标准 Cypher 子句，包括 `MATCH`、`CREATE`、`SET`、`DELETE`、`RETURN`、`WITH`、`WHERE`、`ORDER BY`、`SKIP` 和 `LIMIT`。数据使用 `agtype` 数据类型存储，它扩展了 JSON，增加了顶点、边和路径等图特有类型。

可变长度路径模式匹配：

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (a:Person)-[:KNOWS*1..3]->(b:Person)
    RETURN a.name, b.name
$$) AS (source agtype, target agtype);
```

混合 SQL/Cypher 查询允许将图查询结果与关系表进行关联：

```sql
SELECT t.*, c.* FROM my_table t
JOIN cypher('my_graph', $$
    MATCH (n:Person) RETURN n.name, id(n)
$$) AS c(name agtype, id agtype)
ON t.graph_id = c.id;
```
