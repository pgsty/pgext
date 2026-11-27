## 用法

来源：

- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/README.md)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/pg_repack_safe.control)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/pg_repack_safe--1.0.sql)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/doc/installation.md)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/doc/limits.md)
- [Official source](https://api.pgxn.org/src/pg_repack_safe/pg_repack_safe-1.0.0/doc/removal.md)

`pg_repack_safe` 1.0.0 防止 PostgreSQL 19 内置并发 REPACK 后，使用重写前快照的读取静默看到空表。受影响事务会收到 SQLSTATE 40001，应用须重试事务。它不保护独立的 pg_repack 扩展，也不覆盖所有表重写命令。

### 启用与检查

在主库及全部备库安装匹配动态库，将其加入现有预加载列表，先重启备库，再重启主库：

```ini
shared_preload_libraries = 'pg_repack_safe'
```

由超级用户在每个需要保护的数据库启用：

```sql
CREATE EXTENSION pg_repack_safe;
SELECT repack_safe.status();
```

结果必须为 true。只创建 SQL 扩展而不预加载，不会提供保护。分发和动态库版本为 1.0.0，SQL 扩展有意使用版本 1.0。它依赖 `plpgsql`，不受信任、不可重定位，使用固定的 `repack_safe` 模式。

### 运行边界

扩展使用资源管理器 149 写入自定义 WAL。每台重放服务器都需要兼容动态库，否则重放可能停止。主备须使用相同主版本的动态库，并将库文件和配置纳入恢复计划。应用须重试整个失败事务。

部分不经过执行器的读取不受保护，包括 amcheck 和 pgrowlocks。外键检查可能返回 SQLSTATE 23503，而非重试错误码；上游还记录了备库重放空隙。这些边界意味着它不能普遍修复所有 MVCC 问题。

### 移除

仍存在受保护的重写记录时，不应直接移除预加载动态库。先停止并发重整，按移除指南由超级用户执行记录裁剪和 VACUUM TRUNCATE 流程，确认相关关系存储归零，再删除扩展并退役动态库。此流程须协调备库重放和备份。
