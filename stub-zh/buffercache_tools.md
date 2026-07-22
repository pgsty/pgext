## 用法

来源：

- [官方 README](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/README.md)
- [官方扩展 SQL](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/buffercache_tools--1.0.sql)
- [官方 C 实现](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/buffercache_tools.c)

`buffercache_tools` 提供底层 PostgreSQL 共享缓冲区检查与修改工具。它面向受控调试和故障注入；重写标签、标脏、刷盘或失效缓冲区的函数可能破坏一致性，不能作为日常管理工具使用。

### 核心流程

以超级用户安装扩展，然后先进行只读检查：

```sql
CREATE EXTENSION buffercache_tools;

SELECT * FROM pg_show_buffer(42);
SELECT * FROM pg_show_relation_buffers('public.orders');
```

`pg_show_buffer(integer)` 描述一个共享缓冲区。`pg_show_relation_buffers(text)` 列出某个关系的缓冲区，也可以报告当前会话的本地临时缓冲区。`pg_read_page_into_buffer(text, text, bigint)` 将指定关系分支与块读入缓冲区缓存。

### 修改接口与安全性

修改函数族包括 `pg_change_buffer`、`pg_change_relation_fork_buffers`、`pg_change_relation_buffers`、`pg_change_database_buffers`、`pg_change_tablespace_buffers`、`pg_change_all_valid_buffers` 和 `pg_change_buffer_by_page`。支持的模式包括 `mark_dirty`、`flush`、`invalidate`、`change_spcoid`、`change_dboid`、`change_relnumber`、`change_forknum` 与 `change_blocknum`。

实现要求超级用户执行检查和修改函数。它拒绝处理其他会话的临时关系缓冲区，且临时缓冲区不可修改。即使在测试集群中，也应先保留可恢复副本、隔离并发流量，并预期标签修改或失效操作可能导致写入丢失、错误页面读取、崩溃或重启恢复。测试修改接口时优先使用可丢弃集群。
