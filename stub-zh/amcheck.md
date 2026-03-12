

## 用法

> [amcheck: 验证关系完整性的函数](https://www.postgresql.org/docs/current/amcheck.html)

`amcheck` 扩展提供验证 B-tree 索引、GIN 索引和堆（表）数据逻辑一致性的函数，在不修改数据的情况下检测损坏。

### B-Tree 索引验证

```sql
-- 轻量级检查（AccessShareLock，适合生产环境）
SELECT bt_index_check('my_index');

-- 带堆全索引验证
SELECT bt_index_check('my_index', heapallindexed => true);

-- 全面检查，包含父/子不变量（ShareLock，阻止写入）
SELECT bt_index_parent_check('my_index');

-- 最全面：rootdescend 从根节点重新查找每个元组
SELECT bt_index_parent_check('my_index',
    heapallindexed => true,
    rootdescend => true,
    checkunique => true);
```

### 检查所有系统目录索引

```sql
SELECT bt_index_check(c.oid), c.relname, c.relpages
FROM pg_index i
JOIN pg_opclass op ON i.indclass[0] = op.oid
JOIN pg_am am ON op.opcmethod = am.oid
JOIN pg_class c ON i.indexrelid = c.oid
JOIN pg_namespace n ON c.relnamespace = n.oid
WHERE am.amname = 'btree' AND n.nspname = 'pg_catalog'
  AND c.relpersistence != 't' AND c.relkind = 'i'
  AND i.indisready AND i.indisvalid
ORDER BY c.relpages DESC LIMIT 10;
```

### GIN 索引验证

```sql
SELECT gin_index_check('my_gin_index');
```

### 堆（表）验证

```sql
-- 基本堆检查
SELECT * FROM verify_heapam('my_table');

-- 带 TOAST 验证（更慢）
SELECT * FROM verify_heapam('my_table', check_toast => true);

-- 检查特定块范围
SELECT * FROM verify_heapam('my_table', startblock => 0, endblock => 1000);

-- 在第一个损坏块处停止
SELECT * FROM verify_heapam('my_table', on_error_stop => true);
```

返回每个检测到的问题的行：

| 列 | 类型 | 描述 |
|--------|------|-------------|
| `blkno` | bigint | 存在损坏的块号 |
| `offnum` | integer | 损坏元组的偏移量 |
| `attnum` | integer | 属性号（如果是特定列） |
| `msg` | text | 问题描述 |

### 函数概览

| 函数 | 锁 | 用途 |
|----------|------|----------|
| `bt_index_check(index, heapallindexed, checkunique)` | AccessShareLock | 常规生产检查 |
| `bt_index_parent_check(index, heapallindexed, rootdescend, checkunique)` | ShareLock | 全面验证 |
| `gin_index_check(index)` | AccessShareLock | GIN 索引验证 |
| `verify_heapam(relation, on_error_stop, check_toast, skip, startblock, endblock)` | AccessShareLock | 表/堆损坏检测 |

所有 `amcheck` 错误都是真阳性。诊断后使用 `REINDEX` 或时间点恢复进行修复，可配合 `pageinspect` 使用。
