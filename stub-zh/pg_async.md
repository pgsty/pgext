## 用法

来源：

- [官方 README](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/README.md)
- [官方扩展控制文件](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async.control)
- [官方扩展 SQL](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async--1.0.sql)
- [官方 hook 实现](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async.c)

`pg_async` 1.0 是一个预加载模块，它重新实现 PostgreSQL 异步通知处理，使 `LISTEN`、`NOTIFY` 和 `UNLISTEN` 可以在物理热备恢复期间运行；恢复之外则委托给 PostgreSQL 的普通实现。通知只存在于单个服务器实例，不会在主库与备库之间复制。

### 安装与流程

安装扩展文件，并在主库创建扩展，使目录对象传到备库。把库加入每个需要运行它的服务器，然后重启。

```conf
shared_preload_libraries = 'pg_async'
```

```sql
LISTEN replica_events;

NOTIFY replica_events, 'refresh complete';

SELECT pg_listening_channels();
SELECT pg_notification_queue_usage();

UNLISTEN replica_events;
```

在恢复期间，process-utility hook 会拦截 SQL 命令并使用模块的私有队列实现。Listen/unlisten 变化及通知何时可见仍由事务提交控制。

### 已安装函数

- `pg_listen(text)` 与 `pg_unlisten(text)` 提供注册频道的函数形式。
- `pg_unlisten_all()` 删除当前后端的全部注册。
- `pg_notify(text, text)` 把频道和载荷放入队列。
- `pg_listening_channels()` 列出当前后端监听的频道。
- `pg_notification_queue_usage()` 报告本地队列的占用比例。

这些名称与 PostgreSQL 内置函数重叠。Schema 位置和 `search_path` 决定未限定调用解析到哪个函数；普通场景使用 utility command 更清晰。

### 恢复与版本边界

固定源码嵌入了修改版 PostgreSQL 13 和 14 `async.c` 内部实现，并拒绝通过 `shared_preload_libraries` 之外的方式加载。它没有 PostgreSQL 15 或更高版本的实现分支。这本质上是打包成扩展的服务器内部补丁，必须严格匹配大版本源码。

队列属于本节点和数据库。主库发送的通知不会到达备库，备库提升/重启也会丢弃临时监听状态。应用必须容忍重复、漏事件、断连和提升；需要可靠投递时应使用持久表或消息系统。必须在隔离且版本匹配的 PostgreSQL 构建上测试 hook/信号交互、队列耗尽、长事务和故障切换。
