

## 用法

> [pgspider_ext: 用于远程 PGSpider 服务器的外部数据包装器](https://github.com/pgspider/pgspider_ext)

PGSpider 扩展将 PostgreSQL 变为分布式查询引擎，通过创建虚拟分区表，透明地并行查询多个远程节点的数据。

### 设置子服务器

首先，使用各自的 FDW 为每个数据源创建服务器：

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

### 创建子外部表

```sql
CREATE FOREIGN TABLE t1_pg1_child (i int, t text)
  SERVER pgsrv1 OPTIONS (table_name 't1');
CREATE FOREIGN TABLE t1_pg2_child (i int, t text)
  SERVER pgsrv2 OPTIONS (table_name 't1');
```

### 创建 PGSpider 分区表

创建 PGSpider 服务器和带有额外分区键列的分区父表：

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

默认情况下，PGSpider 使用 `[父表名]_child` 模式查找子表。使用 `child_name` 指定不同的名称。

### 跨所有节点查询

```sql
SELECT * FROM t1;
 i  | t | node
----+---+-------
 10 | a | node1
 11 | b | node1
 20 | c | node2
 21 | d | node2
```

查询会自动并行分发到所有子节点。WHERE 子句和聚合函数会下推到子节点：

```sql
SET enable_partitionwise_aggregate TO on;
SELECT count(*), node FROM t1 GROUP BY node;
```

**注意：** 仅支持 SELECT 操作；不支持 INSERT、UPDATE 和 DELETE。
