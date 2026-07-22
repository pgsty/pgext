## 用法

来源：

- [DagDB 引擎 README](https://github.com/norayr-m/dagdb-engine/blob/7cd3c50ccceca47f61b21724a6262518ceadb1e7/README.md)
- [`pg_dagdb` Rust 实现](https://github.com/norayr-m/dagdb-engine/blob/7cd3c50ccceca47f61b21724a6262518ceadb1e7/dagdb/pg_dagdb/src/lib.rs)
- [`pg_dagdb` 控制文件](https://github.com/norayr-m/dagdb-engine/blob/7cd3c50ccceca47f61b21724a6262518ceadb1e7/dagdb/pg_dagdb/pg_dagdb.control)

`pg_dagdb` 是 PostgreSQL 到实验性 DagDB 分层有向无环图引擎的客户端桥接。SQL 扩展本身不存储或求值图；它通过固定 Unix 套接字向独立的本地守护进程发送 DagDB DSL 命令，并从固定共享内存文件读取行结果。

### 核心流程

查询扩展前，应在 PostgreSQL 服务器主机上安装并启动匹配的 `dagdb_daemon`：

```sql
CREATE EXTENSION pg_dagdb;

SELECT dagdb_status();
SELECT dagdb_tick(100);

SELECT *
FROM dagdb_exec('NODES AT RANK 2 WHERE truth=1');
```

对于返回节点行的命令，`dagdb_exec` 暴露 `node_id`、`rank`、`truth` 和 `node_type`；状态类命令则把守护进程响应放在 `message` 中。

### 对象索引

- `dagdb_status() RETURNS text` 发送 `STATUS`，返回守护进程响应或 `OFFLINE` 消息。
- `dagdb_tick(n integer DEFAULT 1) RETURNS text` 请求 GPU 引擎运行若干 tick。
- `dagdb_exec(command text)` 返回包含可空 `node_id`、`rank`、`truth`、`node_type` 和 `message` 列的表。

### 运维说明

目录版本 `0.0.0` 表明这是预发布接口。控制文件将扩展标记为不可重定位、仅超级用户可安装且不受信；未声明预加载要求。实现硬编码 `/tmp/dagdb.sock` 和 `/tmp/dagdb_shm_file`，因此守护进程必须在数据库服务器上运行，并具有兼容的文件权限和二进制布局。这些路径也形成主机本地信任边界：应限制扩展调用，并防止其他主机用户访问套接字和共享内存文件。

错误以文本行返回，而不是 SQL 异常；缺失或格式错误的共享内存结果可能只返回守护进程消息或空结果。上游项目明确列出了尚未验证的限制，并采用 Apple Silicon/Metal 设计。应将守护进程和扩展固定到同一提交，把它视为实验性组件，且不要从这个桥接推断事务一致性或 PostgreSQL 持久性。
