## 用法

来源：

- [官方 README](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/README.md)
- [1.0 版本扩展 SQL](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/pg_follower--1.0.sql)
- [应用工作进程实现](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/pg_follower_apply.c)

`pg_follower` 是一个教育用途的逻辑复制实验，用于在下游 PostgreSQL 服务器上重建有限种类的 SQL 变更。上游明确说明不要用于生产；它不能替代内置逻辑复制。

### 核心流程

在上游设置 `wal_level = logical` 并重启服务器，然后在两个数据库中创建扩展。在下游使用指向上游的 libpq 连接字符串启动工作进程：

```sql
-- Run in both upstream and downstream databases.
CREATE EXTENSION pg_follower;

-- Run only on the downstream.
SELECT start_follow('host=upstream.example user=replicator dbname=app');
```

工作进程会在上游创建临时逻辑复制槽，接收扩展输出插件产生的类 SQL 消息，并在下游通过 SPI 应用。

### 对象索引

- `start_follow(connection_string text) RETURNS void` 启动下游后台工作进程。
- `detect_ddl() RETURNS event_trigger` 捕获受支持的表 DDL。
- 事件触发器 `test_trigger` 在 `CREATE TABLE` 和 `DROP TABLE` 后触发。

### 运维说明

版本 `1.0` 可重定位，且未声明预加载要求。扩展必须安装在两端，上游角色需要逻辑复制连接能力。应防止连接字符串中的密码暴露在 SQL 文本、日志和监控中。

文档只支持 `INSERT`、`CREATE TABLE`、`DROP TABLE` 和 `TRUNCATE`。不支持 `UPDATE` 和 `DELETE`；表约束与选项会被忽略。`UNLOGGED`、`TEMPORARY`、分区表或继承表以及类型化表子句会被拒绝。临时复制槽、手工构造的 SQL 流、后台进程和不完整的 DDL 模型存在严重的正确性、故障转移、安全和恢复缺口。只应使用隔离的测试数据。
