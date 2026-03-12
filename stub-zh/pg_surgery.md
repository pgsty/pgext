

## 用法

> [pg_surgery: 对损坏的关系执行手术修复的扩展](https://www.postgresql.org/docs/current/pgsurgery.html)

`pg_surgery` 扩展提供对损坏关系执行底层手术修复的函数。这些是最后手段的工具，误用可能导致数据损坏。

### 函数

#### heap_force_kill

将行指针标记为"死亡"而不检查元组，强制移除不可访问的元组。

```sql
heap_force_kill(regclass, tid[]) RETURNS void
```

```sql
-- 元组在访问时导致错误
SELECT * FROM t1 WHERE ctid = '(0, 1)';
-- ERROR: could not access status of transaction 4007513275

-- 强制杀死有问题的元组
SELECT heap_force_kill('t1'::regclass, ARRAY['(0, 1)']::tid[]);

-- 元组现在已被移除
SELECT * FROM t1 WHERE ctid = '(0, 1)';
-- (0 rows)
```

#### heap_force_freeze

将元组标记为已冻结而不检查元组数据，使可见性信息损坏的元组变为可访问。

```sql
heap_force_freeze(regclass, tid[]) RETURNS void
```

```sql
-- VACUUM 因可见性信息损坏而失败
VACUUM t1;
-- ERROR: found xmin 507 from before relfrozenxid 515

-- 找到有问题的元组
SELECT ctid FROM t1 WHERE xmin = 507;
--  ctid
-- -------
--  (0,3)

-- 强制冻结元组
SELECT heap_force_freeze('t1'::regclass, ARRAY['(0, 3)']::tid[]);

-- 元组现在已冻结（xmin 变为 2 = FrozenTransactionId）
SELECT ctid FROM t1 WHERE xmin = 2;
--  ctid
-- -------
--  (0,3)
```

### 何时使用

- `heap_force_kill`：当元组因事务状态损坏导致访问错误，且数据可以丢弃时
- `heap_force_freeze`：当 VACUUM 因元组的 xmin 早于 relfrozenxid 而失败时，或当元组因可见性信息损坏而不可见时

这些函数在设计上是不安全的，只应在正常恢复方法失败后作为最后手段使用。
