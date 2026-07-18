## 用法

来源：

- [固定提交的当前上游 README](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/README.md)
- [0.1.0 版安装 SQL](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/pg_recyclebin--0.1.0.sql)
- [固定提交的 Cargo 元数据](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/Cargo.toml)
- [固定提交的备份恢复测试矩阵](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/docs/BACKUP_RESTORE_TEST_MATRIX.md)

pg_recyclebin 0.1.0 版拦截选定的 DROP 与 TRUNCATE 操作。被删除的表会移动并改名到 flashback_recycle；被截断的数据则会复制到那里。flashback 下的元数据记录原所有者、模式、操作、依赖和保留期限，使条目可以列出、恢复或永久清除。

### 预加载与恢复示例

```conf
shared_preload_libraries = 'pg_recyclebin'
```

```sql
CREATE EXTENSION pg_recyclebin;

CREATE TABLE recycle_demo (id integer);
INSERT INTO recycle_demo VALUES (1), (2);

DROP TABLE recycle_demo;

SELECT * FROM flashback_list_recycled_tables();
SELECT flashback_restore('recycle_demo');

SELECT * FROM recycle_demo ORDER BY id;
```

当前项目记录支持 PostgreSQL 14 至 18，使用 pgrx 0.16.1。后台工作进程会清理过期条目，默认连接数据库是 postgres；若安装在其他数据库，应设置 flashback.database_name。运行时限制控制保留天数、表数量、总大小、工作进程间隔和排除模式。

### 恢复功能不是备份

该钩子会有意改变普通 DROP/TRUNCATE 语义，包括多表命令、CASCADE 和 DROP SCHEMA CASCADE。它可能移动分区、重建依赖视图和传入外键、保留策略/触发器，或复制整张被截断的表。这些操作会增加 DDL 延迟、存储/WAL 压力、锁交互和新故障模式。必须针对确切版本测试事务回滚、依赖、分区、标识序列、所有权、RLS 和带引号名称。

回收站无法防护 DROP DATABASE、集群或磁盘丢失、临时表或排除表、特权清除、保留期淘汰、扩展故障或损坏。仍需保留独立且经过验证的备份与时间点恢复。应限制恢复/清除函数、监控回收站存储，并在依赖它处理误操作前演练普通恢复和批量恢复。
