


## Usage

> [age: AGE graph database extension](https://github.com/apache/age)

Apache AGE brings graph database capabilities to PostgreSQL using the openCypher query language. It enables hybrid querying that combines SQL and Cypher, property indexes on vertices and edges, and the ability to query multiple graphs.

Each session requires loading the extension:

```sql
CREATE EXTENSION age;
LOAD 'age';
SET search_path = ag_catalog, "$user", public;
```

### Graph Operations

Create a graph:

```sql
SELECT create_graph('my_graph');
```

Create vertices:

```sql
SELECT * FROM cypher('my_graph', $$
    CREATE (:Person {name: 'Alice', age: 30})
$$) AS (v agtype);

SELECT * FROM cypher('my_graph', $$
    CREATE (:Person {name: 'Bob', age: 25})
$$) AS (v agtype);
```

Create edges:

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (a:Person), (b:Person)
    WHERE a.name = 'Alice' AND b.name = 'Bob'
    CREATE (a)-[e:KNOWS {since: 2020}]->(b)
    RETURN e
$$) AS (e agtype);
```

Query the graph:

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (v)-[r]-(v2)
    RETURN v, r, v2
$$) AS (v agtype, r agtype, v2 agtype);
```

### Cypher Query Features

AGE supports standard Cypher clauses including `MATCH`, `CREATE`, `SET`, `DELETE`, `RETURN`, `WITH`, `WHERE`, `ORDER BY`, `SKIP`, and `LIMIT`. Data is stored using the `agtype` data type, which extends JSON with graph-specific types for vertices, edges, and paths.

Pattern matching with variable-length paths:

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (a:Person)-[:KNOWS*1..3]->(b:Person)
    RETURN a.name, b.name
$$) AS (source agtype, target agtype);
```

Hybrid SQL/Cypher queries allow joining graph results with relational tables:

```sql
SELECT t.*, c.* FROM my_table t
JOIN cypher('my_graph', $$
    MATCH (n:Person) RETURN n.name, id(n)
$$) AS c(name agtype, id agtype)
ON t.graph_id = c.id;
```
