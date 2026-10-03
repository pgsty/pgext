## 用法

来源：

- [Version 4.0.9 SQL](https://github.com/pgroonga/pgroonga/blob/4.0.9/data/pgroonga_database.sql)
- [Version 4.0.9 control](https://github.com/pgroonga/pgroonga/blob/4.0.9/pgroonga_database.control)
- [Version 4.0.9 implementation](https://github.com/pgroonga/pgroonga/blob/4.0.9/src/pgroonga-database.c)
- [Official recovery procedure](https://pgroonga.github.io/reference/functions/pgroonga-database-remove.html)

`pgroonga_database` 4.0.9 是用于恢复损坏的 PGroonga 内部数据库的辅助扩展。它只提供一个 SQL 函数，用来删除数据库目录及适用表空间目录中的 PGroonga 文件，不提供搜索访问方法。

### 恢复流程

普通索引损坏可能只需 REINDEX 即可修复。只有内部 Groonga 数据库本身损坏、确定需要重建时，才使用此模块。安排恢复窗口，先断开所有使用 PGroonga 的会话；文件被删除时，残留会话可能崩溃。

在尚未打开任何 PGroonga 索引的新管理连接中执行：

```sql
CREATE EXTENSION pgroonga_database;
SELECT pgroonga_database_remove();
```

执行后立即断开该连接，再建立新连接，对**每一个** PGroonga 索引执行 REINDEX，从 PostgreSQL 表数据重新创建内部数据库。完成所有受影响索引的重建和检查后，再恢复应用流量。

### 返回值与边界

`pgroonga_database_remove()` 结束清理循环时返回 true。表空间所有权检查未通过时，循环会提前停止，但函数仍可能返回 true，因此不能仅凭返回值认定所有位置都已清理。其他失败可能报错。它直接删除内部文件，既不导出文件，也不重建索引。不要在清理连接中使用其他 PGroonga 功能。此操作不是例行清理、卸载命令，也不能仅靠包裹在 SQL 事务中就保证安全。

控制文件未将扩展标记为受信任或可迁移模式。C 实现在遍历位置时检查表空间所有权，应使用拥有所需位置的管理员，并确认清理和全量索引重建完成。模块不需要预加载，仅应为恢复任务启用。
