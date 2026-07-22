## 用法

来源：

- [官方文档](https://docs.synchdb.com)
- [固定提交的上游 README](https://github.com/Hornetlabs/synchdb/blob/48d5d11e95d1c9711827ed9bade8255123434c87/README.md)
- [1.0 版本 SQL 定义](https://github.com/Hornetlabs/synchdb/blob/48d5d11e95d1c9711827ed9bade8255123434c87/synchdb--1.0.sql)
- [扩展控制文件](https://github.com/Hornetlabs/synchdb/blob/48d5d11e95d1c9711827ed9bade8255123434c87/synchdb.control)

`synchdb` 是一个混合 C 与 Java 的 PostgreSQL 扩展，用于从 MySQL、SQL Server、Oracle 和 OpenLog Replicator 源执行初始快照加变更数据捕获复制，并写入 PostgreSQL。它通过 JVM 嵌入 Debezium 连接器，在 PostgreSQL 内应用事件；它不是双向复制或 PostgreSQL 到 PostgreSQL 复制系统。

### 核心流程

安装扩展，注册命名连接器，启动其后台工作进程，并检查状态。上游对未指定的表和快照字段使用字面文本 `'null'`，因为注册函数是 strict 的。

```sql
CREATE EXTENSION synchdb CASCADE;

SELECT synchdb_add_conninfo(
    'mysqlconn',
    '127.0.0.1',
    3306,
    'cdc_user',
    'source-secret',
    'inventory',
    'postgres',
    'inventory.orders,inventory.customers',
    'null',
    'mysql'
);

SELECT synchdb_start_engine_bgw('mysqlconn');
SELECT * FROM synchdb_state_view;
```

更改连接器定义或执行维护前，先停止它：

```sql
SELECT synchdb_stop_engine_bgw('mysqlconn');
```

### 管理与可观测性

- `synchdb_add_conninfo(...)` 把连接器保存在 `synchdb_conninfo` 中；`synchdb_del_conninfo(name)` 将其删除。
- `synchdb_start_engine_bgw(name)` 启动快照加流处理，而其双参数重载选择快照模式。`synchdb_pause_engine(name)`、`synchdb_resume_engine(name)` 和 `synchdb_stop_engine_bgw(name)` 控制执行。
- `synchdb_state_view` 暴露阶段、状态、进程 ID、最近错误和源位点。
- `synchdb_genstats`、`synchdb_snapstats` 和 `synchdb_cdcstats` 暴露通用、快照和变更事件计数器；`synchdb_reset_stats(name)` 重置连接器计数器。
- 对象映射、类型转换、JMX、OpenLog Replicator 与位点函数支持高级连接器配置；应遵循与版本匹配的官方指南，不要猜测参数语义。

### 前置条件与运维

版本 `1.0` 需要 `pgcrypto`、JVM/JNI 运行时，以及与 C 库一同安装的 Java 组件。固定提交的 README 列出 Java 17 或更高版本、PostgreSQL 16 至 18 和类 Unix 操作系统；Oracle FDW 快照与 OpenLog Replicator 构建还会增加各自的原生依赖。

保存的连接器密码会加密，但初始密码仍会经过 SQL 语句，可能进入客户端历史、语句日志、审计日志或监控。应使用严格受控的管理会话与最小权限源账号。初始快照可能很大，而 DDL 与类型转换、位点、源保留、重启、重复交付、schema 漂移和应用错误决定正确性。用于生产前，应测试代表性数据源，监控 `synchdb_state_view` 和统计视图，验证行数与键，并演练暂停、恢复、重建和故障恢复。
