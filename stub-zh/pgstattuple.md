


## 用法

> [pgstattuple: 元组级别统计](https://www.postgresql.org/docs/current/pgstattuple.html)

pgstattuple 提供函数来获取表和索引的元组级别统计信息，包括死元组计数和空闲空间指标。

### 表统计

```sql
SELECT * FROM pgstattuple('my_table');

 table_len          | 81920
 tuple_count        | 1000
 tuple_len          | 72000
 tuple_percent      | 87.89
 dead_tuple_count   | 50
 dead_tuple_len     | 3600
 dead_tuple_percent | 4.39
 free_space         | 6320
 free_percent       | 7.71
```

**近似（更快）版本**，使用可见性映射和 FSM：

```sql
SELECT * FROM pgstattuple_approx('my_table');
-- 返回：table_len、scanned_percent、approx_tuple_count、approx_tuple_len、
--       dead_tuple_count、dead_tuple_len、approx_free_space 等
```

### B-Tree 索引统计

```sql
SELECT * FROM pgstatindex('my_index');

 version            | 4
 tree_level         | 2
 index_size         | 1327104
 root_block_no      | 3
 internal_pages     | 1
 leaf_pages         | 160
 empty_pages        | 0
 deleted_pages      | 0
 avg_leaf_density   | 89.27
 leaf_fragmentation | 0
```

### GIN 索引统计

```sql
SELECT * FROM pgstatginindex('my_gin_index');
-- 返回：version、pending_pages、pending_tuples
```

### Hash 索引统计

```sql
SELECT * FROM pgstathashindex('my_hash_index');
-- 返回：version、bucket_pages、overflow_pages、bitmap_pages、
--       unused_pages、live_items、dead_tuples、free_percent
```

### 页面计数

```sql
SELECT pg_relpages('my_table');
```

### 访问权限

限制为超级用户和拥有 `pg_stat_scan_tables` 权限的角色。
