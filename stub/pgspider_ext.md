


## Usage

> [pgspider_ext: Foreign data wrapper for remote PGSpider servers](https://github.com/pgspider/pgspider_ext)

PGSpider Extension turns PostgreSQL into a distributed query engine by creating virtual partitioned tables that transparently query data across multiple remote nodes in parallel.

### Setup Child Servers

First, create servers for each data source using their respective FDWs:

```sql
CREATE EXTENSION pgspider_ext;
CREATE EXTENSION postgres_fdw;

CREATE SERVER pgsrv1 FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '127.0.0.1', port '5433', dbname 'postgres');
CREATE SERVER pgsrv2 FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '127.0.0.1', port '5434', dbname 'postgres');

CREATE USER MAPPING FOR CURRENT_USER SERVER pgsrv1
  OPTIONS (user 'user', password 'pass');
CREATE USER MAPPING FOR CURRENT_USER SERVER pgsrv2
  OPTIONS (user 'user', password 'pass');
```

### Create Child Foreign Tables

```sql
CREATE FOREIGN TABLE t1_pg1_child (i int, t text)
  SERVER pgsrv1 OPTIONS (table_name 't1');
CREATE FOREIGN TABLE t1_pg2_child (i int, t text)
  SERVER pgsrv2 OPTIONS (table_name 't1');
```

### Create PGSpider Partitioned Table

Create a PGSpider server and a partitioned parent table with an extra partition key column:

```sql
CREATE SERVER spdsrv FOREIGN DATA WRAPPER pgspider_ext;
CREATE USER MAPPING FOR CURRENT_USER SERVER spdsrv;

CREATE TABLE t1 (i int, t text, node text)
  PARTITION BY LIST (node);

CREATE FOREIGN TABLE t1_pg1 PARTITION OF t1
  FOR VALUES IN ('node1') SERVER spdsrv;
CREATE FOREIGN TABLE t1_pg2 PARTITION OF t1
  FOR VALUES IN ('node2') SERVER spdsrv
  OPTIONS (child_name 't1_pg2_child');
```

By default, PGSpider finds child tables using the pattern `[parent_table_name]_child`. Use `child_name` to specify a different name.

### Query Across All Nodes

```sql
SELECT * FROM t1;
 i  | t | node
----+---+-------
 10 | a | node1
 11 | b | node1
 20 | c | node2
 21 | d | node2
```

Queries automatically fan out to all child nodes in parallel. WHERE clauses and aggregate functions are pushed down to child nodes:

```sql
SET enable_partitionwise_aggregate TO on;
SELECT count(*), node FROM t1 GROUP BY node;
```

**Note:** Only SELECT operations are supported; INSERT, UPDATE, and DELETE are not supported.
