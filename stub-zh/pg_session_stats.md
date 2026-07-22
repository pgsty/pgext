## 用法

来源：

- [官方 README](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/README.md)
- [扩展控制文件](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/pg_session_stats.control)
- [空扩展 SQL](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/pg_session_stats--0.0.1.sql)
- [执行器钩子与 SQLite 写入器](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/main.c)

`pg_session_stats` 是一个面向 Linux 与 PostgreSQL 12 的实验，把后端 CPU 和 `/proc` 快照记录到外部 SQLite 数据库。它没有 SQL API：扩展脚本为空，唯一有效的启用路径是服务器预加载后重启。

### 核心流程

安装链接 SQLite 的二进制，选择由 PostgreSQL 拥有的文件位置，预加载动态库并重启。该文件包含进程状态与 I/O 信息，因此必须保护所在目录。

```conf
shared_preload_libraries = 'pg_session_stats'
pg_session_stats.path = '/var/lib/postgresql/pg_session_stats.sqlite'
```

SQLite 数据库包含一个 `log` 表，列为 `master_pid`、`my_pid`、`usage`、`procstatus` 与 `procio`。README 早于实现中的 `procio` 列。每次执行器完成都会追加一行；`usage` 是 `clock()` 返回的进程累计 CPU，而不是该查询的 CPU 增量。

### 可靠性与资源注意事项

尽管名称如此，钩子在每次 `ExecutorEnd` 都会写入，而不是只在会话结束时写入。SQLite 写入在查询后端同步执行，并使用五秒忙等待超时，因此锁竞争或文件系统延迟会拖慢应用查询。外部写入不属于 PostgreSQL 事务，回滚也不会撤销。

固定代码读取 Linux `/proc`，假设 PostgreSQL 12 的工作进程命令行行为，并使用脆弱的父 PID 推断。更严重的是，读取 `/proc` 后没有关闭文件，每次记录执行器完成都会泄漏文件描述符；长寿命后端最终可能耗尽描述符。错误路径也缺少可靠清理。不要在生产系统使用 `pg_session_stats`。若在隔离实验室评估，应限制会话寿命，监控文件描述符与 SQLite 增长，限制文件访问，并独立验证工作进程归属与 CPU 含义。
