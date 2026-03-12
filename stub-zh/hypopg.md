


## 用法

> [hypopg: PostgreSQL 假想索引](https://github.com/HypoPG/hypopg)

HypoPG 允许创建仅存在于当前会话中的假想（虚拟）索引，这些索引可被 `EXPLAIN`（不带 `ANALYZE`）在查询规划时考虑。这使得无需实际创建索引即可测试索引对查询的影响。

### 函数

| 函数 | 描述 |
|------|------|
| `hypopg_create_index(query text)` | 使用 CREATE INDEX 语法创建假想索引 |
| `hypopg_list_indexes()` | 列出当前会话中的所有假想索引 |
| `hypopg_drop_index(oid)` | 通过 OID 删除特定假想索引 |
| `hypopg_reset()` | 删除所有假想索引 |
| `hypopg()` | 以类似 pg_index 的格式返回假想索引 |

### 工作流程

创建测试表并查看基线执行计划：

```sql
CREATE TABLE hypo AS SELECT id, 'line ' || id AS val FROM generate_series(1, 10000) id;
ANALYZE hypo;
EXPLAIN SELECT * FROM hypo WHERE id = 1;
-- Seq Scan on hypo (cost=0.00..170.00 rows=1 width=15)
```

创建假想索引：

```sql
SELECT * FROM hypopg_create_index('CREATE INDEX ON hypo (id)');
--  indexrelid |      indexname
-- ------------+----------------------
--       13543 | <13543>btree_hypo_id
```

查看使用假想索引后的执行计划：

```sql
EXPLAIN SELECT * FROM hypo WHERE id = 1;
-- Index Scan using <13543>btree_hypo_id on hypo (cost=0.04..8.06 rows=1 width=15)
```

列出和管理假想索引：

```sql
SELECT * FROM hypopg_list_indexes();
SELECT * FROM hypopg_drop_index(13543);
SELECT * FROM hypopg_reset();
```

### 限制

- 只有不带 `ANALYZE` 的 `EXPLAIN` 才会考虑假想索引
- 假想索引仅存在于当前后端会话中
- 其他并发连接不受影响
- 索引名称和部分 CREATE INDEX 选项会被忽略
