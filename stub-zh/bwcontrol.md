## 用法

来源：

- [上游 README](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/README.md)
- [扩展 control 文件](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/bwcontrol.control)
- [SQL 安装脚本](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/bwcontrol--1.0.sql)
- [C 实现](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/bwcontrol.c)

`bwcontrol` 是定制版 `Bottledwater` 变更数据捕获进程及其 `Kafka Connect` 集成的 PostgreSQL 侧控制器。它在扩展自有元数据中维护表白名单、启动或停止外部进程、检查进程状态，并通过 HTTP 更新连接器配置。control 文件与目录记录的版本均为 `1.0`。

### 最小状态检查

先配置外部路径和端点，再创建扩展并调用只读的状态入口：

```sql
CREATE EXTENSION bwcontrol;
SELECT pg_get_status_ingest();
```

实现会读取 `bw.bwpath`、`bw.kafka_broker`、`bw.schema_registry`、`bw.consumer` 和 `bw.consumer_sub`。上游工作流还要求复制连接权限、定制版 Bottledwater、Kafka、schema registry 与 Kafka Connect。

白名单通过 `pg_add_ingest_table()`、`pg_del_ingest_table()`、`pg_add_ingest_column()` 和 `pg_del_ingest_column()` 管理。进程及连接器操作使用 `pg_resume_ingest()`、`pg_suspend_ingest()`、`pg_make_kafka_connect()` 和 `pg_delete_kafka_connect()`。

这是较老且高度依赖具体环境的 C 代码，没有当前 PostgreSQL 兼容矩阵。控制函数会启动 shell 命令、管理复制槽、发起 HTTP 请求并构造动态 SQL，但安装脚本把所有函数都标为 `IMMUTABLE`，也没有撤销默认执行权限。应把所有入口视为特权管理 API，限制 `EXECUTE`，校验每个标识符和连接器名称，并在用于生产前于隔离环境验证恢复流程。
