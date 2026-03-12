


## 用法

> [pg_freespacemap: 检查空闲空间映射](https://www.postgresql.org/docs/current/pgfreespacemap.html)

pg_freespacemap 提供函数来检查空闲空间映射（FSM），FSM 追踪关系中每个页面的可用空间。

### 函数

**单个页面的空闲空间：**

```sql
SELECT pg_freespace('my_table'::regclass, 0);  -- 块 0
```

**所有页面的空闲空间：**

```sql
SELECT * FROM pg_freespace('my_table'::regclass);

 blkno | avail
-------+-------
     0 |     0
     1 |     0
     2 |   224
     3 |  3456
     4 |  8160
```

### 示例：表膨胀分析

```sql
-- 有大量空闲空间的页面
SELECT blkno, avail
FROM pg_freespace('my_table'::regclass)
WHERE avail > 1000
ORDER BY avail DESC;

-- 关系中的总空闲空间
SELECT sum(avail) AS total_free_bytes,
       count(*) AS total_pages,
       count(*) FILTER (WHERE avail > 0) AS pages_with_free_space
FROM pg_freespace('my_table'::regclass);
```

### 注意事项

- FSM 值四舍五入到 `BLCKSZ` 的 1/256（通常为 32 字节）
- FSM 不会完全实时更新；值可能滞后于实际空闲空间
- 对于索引，仅追踪完全未使用的页面
- 访问限制为超级用户和 `pg_stat_scan_tables` 成员
