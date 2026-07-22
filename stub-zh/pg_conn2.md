## 用法

来源：

- [官方 README](https://github.com/serpent7776/pg_conn2/blob/296c2d7086dd0021a636c57964a454961cecfead/README.md)
- [扩展 SQL API](https://github.com/serpent7776/pg_conn2/blob/296c2d7086dd0021a636c57964a454961cecfead/pg_conn2--0.1.sql)
- [libpq 连接实现](https://github.com/serpent7776/pg_conn2/blob/296c2d7086dd0021a636c57964a454961cecfead/pg_conn2.c)

`pg_conn2` 0.1 从 PostgreSQL 后端打开第二条 libpq 连接，连接同一本地服务器上的另一个数据库，并通过不透明的 `pg_conn2` 句柄执行 SQL。它只适合严格受控的管理实验：不会返回查询行，而且远端事务与调用者本地事务不具备原子性。

### 核心流程

实现把 Unix 套接字目录硬编码为 `/var/run/postgresql`，用户硬编码为 `postgres`；唯一可传入的连接字段是数据库名：

```sql
CREATE EXTENSION pg_conn2;

SELECT pg_conn2exec(
  pg_conn2make('analytics'),
  'CREATE TABLE IF NOT EXISTS import_log(ts timestamptz DEFAULT now())'
);
```

若多条语句需要在远端连接上共同成功或失败：

```sql
SELECT pg_conn2exec_many(
  pg_conn2make('analytics'),
  ARRAY[
    'INSERT INTO import_log DEFAULT VALUES',
    'ANALYZE import_log'
  ]
);
```

`pg_conn2exec_many` 会在远端依次发送 `BEGIN`、数组中的语句和 `COMMIT`。任何错误都会发送 `ROLLBACK` 并抛出异常；因此它的布尔结果表示成功时的 `true`，而不是通常可见的失败 `false`。`pg_conn2close(handle)` 可提前关闭；否则连接会在本地顶层事务提交或回滚时关闭。

### 边界与风险

`pg_conn2exec` 接受返回元组的语句，但会丢弃全部结果。每条语句通常在远端自动提交。即使 `pg_conn2exec_many` 也会在外围本地事务结束前提交远端事务，因此之后的本地回滚无法撤销远端操作。应避免两个会话相互等待的自连接锁模式。

认证遵循本地 `pg_hba.conf` 和 PostgreSQL 服务器进程的操作系统身份。数据库名未经转义就插入 libpq 连接文本；绝不能传入不可信输入，否则可注入额外连接参数。句柄仅在事务期间有效，不能有意义地序列化，也不应存入表。生产代码应选择维护中的 FDW、`dblink` 或应用侧连接，并显式处理凭据、结果和事务语义。
