


## 用法

来源：

- [README](https://github.com/HypoPG/hypopg/blob/4e564ebbd46e1ab365ad0d05ca4a2f19cf316ba7/README.md)
- [Control](https://github.com/HypoPG/hypopg/blob/4e564ebbd46e1ab365ad0d05ca4a2f19cf316ba7/hypopg.control)
- [Release 1.4.4](https://github.com/HypoPG/hypopg/releases/tag/1.4.4)

HypoPG 允许创建仅存在于当前会话中的假想（虚拟）索引，这些索引可被 `EXPLAIN`（不带 `ANALYZE`）在查询规划时考虑。这使得无需实际创建索引即可测试索引对查询的影响。

### 函数

| 函数 | 描述 |
|------|------|
| `hypopg_create_index(query text)` | 使用 CREATE INDEX 语法创建假想索引 |
| `hypopg_list_indexes` | 列出当前会话中的所有假想索引 |
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


### 1.4.4 与升级边界

上游 1.4.4 修复大 OID 下的假想索引名称、依赖对象失效、INCLUDE 列数量限制和 EXPLAIN 失败后的状态，并补充 PostgreSQL 19 兼容性修复。假想索引仍仅存在于单个后端，只影响不带 ANALYZE 的 EXPLAIN。项目不提供扩展升级脚本，安装匹配文件后需逐库删除并重新创建扩展。重连或重置后端可清除假想状态；上游发布不代表对应软件包已经提供。
