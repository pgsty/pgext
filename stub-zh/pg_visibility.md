


## 用法

> [pg_visibility: 检查可见性映射](https://www.postgresql.org/docs/current/pgvisibility.html)

pg_visibility 提供函数来检查和验证可见性映射（VM），VM 追踪哪些页面仅包含对所有事务可见的元组。

### 函数

**单页可见性：**

```sql
-- 特定块的 VM 位
SELECT * FROM pg_visibility_map('my_table', 0);
-- 返回：all_visible、all_frozen

-- VM 位加上页面的 PD_ALL_VISIBLE 标志
SELECT * FROM pg_visibility('my_table', 0);
-- 返回：all_visible、all_frozen、pd_all_visible
```

**所有页面的可见性：**

```sql
-- 每个页面的 VM 位
SELECT * FROM pg_visibility_map('my_table');
-- 返回：blkno、all_visible、all_frozen

-- 每个页面的 VM 位加上 PD_ALL_VISIBLE
SELECT * FROM pg_visibility('my_table');
-- 返回：blkno、all_visible、all_frozen、pd_all_visible
```

**汇总：**

```sql
SELECT * FROM pg_visibility_map_summary('my_table');
-- 返回：all_visible（计数）、all_frozen（计数）
```

### 损坏检测

```sql
-- 查找 all-frozen 页面上实际未冻结的元组
SELECT * FROM pg_check_frozen('my_table');

-- 查找 all-visible 页面上实际非全可见的元组
SELECT * FROM pg_check_visible('my_table');
```

如果任一函数返回行，则可见性映射已损坏。

### 修复

```sql
-- 截断可见性映射（强制完整 VACUUM 重建）
SELECT pg_truncate_visibility_map('my_table');
-- 然后执行：VACUUM my_table;
```

### 访问权限

函数需要超级用户或 `pg_stat_scan_tables` 角色。`pg_truncate_visibility_map` 需要超级用户。
