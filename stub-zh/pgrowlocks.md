


## 用法

> [pgrowlocks: 显示行级锁信息](https://www.postgresql.org/docs/current/pgrowlocks.html)

pgrowlocks 显示表中哪些行当前被锁定、被哪些事务锁定以及锁模式。

### 函数

```sql
SELECT * FROM pgrowlocks('my_table');

 locked_row | locker | multi | xids  |     modes      |  pids
------------+--------+-------+-------+----------------+--------
 (0,1)      |    609 | f     | {609} | {"For Share"}  | {3161}
 (0,2)      |    609 | f     | {609} | {"For Share"}  | {3161}
 (0,3)      |    607 | f     | {607} | {"For Update"} | {3107}
```

### 返回列

| 列名 | 类型 | 描述 |
|--------|------|-------------|
| `locked_row` | tid | 被锁定行的元组 ID |
| `locker` | xid | 事务 ID（或多事务 ID） |
| `multi` | boolean | 如果 locker 是多事务则为 true |
| `xids` | xid[] | 所有锁持有者的事务 ID |
| `modes` | text[] | 锁模式：`For Key Share`、`For Share`、`For No Key Update`、`For Update` 等 |
| `pids` | integer[] | 锁持有后端的进程 ID |

### 查看被锁定行的内容

```sql
SELECT * FROM accounts AS a, pgrowlocks('accounts') AS p
WHERE p.locked_row = a.ctid;
```

### 访问权限

限制为超级用户、拥有 `pg_stat_scan_tables` 角色的成员以及对目标表有 `SELECT` 权限的用户。

### 注意事项

- 对目标表获取 `AccessShareLock`
- 不保证产生自洽的快照
- 对大表可能较慢
