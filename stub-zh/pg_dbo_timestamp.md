## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/README.md)
- [1.0.1 版本安装 SQL](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/scripts/pg_dbo_timestamp--1.0.1.sql)
- [DDL 事件函数源码](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/DATABASE/SCHEMA/public/FUNCTION/dbots_on_ddl_event.sql)
- [报告视图源码](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/DATABASE/SCHEMA/public/VIEW/dbots_object_timestamps.sql)
- [扩展 control 文件](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/pg_dbo_timestamp.control)

`pg_dbo_timestamp` 记录数据库对象最近一次观测到的 DDL 修改时间、当前用户、会话用户和客户端地址。它用于支持 pgCodeKeeper 等模式比较工具，并不是不可变或防篡改的安全审计日志。

### 核心流程

DDL 事件触发器会以禁用状态安装，创建扩展后必须手动启用：

```sql
CREATE SCHEMA schema_audit;
CREATE EXTENSION pg_dbo_timestamp SCHEMA schema_audit;
ALTER EVENT TRIGGER dbots_tg_on_ddl_event ENABLE;

GRANT USAGE ON SCHEMA schema_audit TO app_ddl;
GRANT SELECT, INSERT, UPDATE, DELETE
  ON schema_audit.dbots_event_data TO app_ddl;
GRANT SELECT
  ON schema_audit.dbots_object_timestamps TO schema_reader;

SELECT type, schema, name, last_modified, cur_user, ses_user, ip_address
FROM schema_audit.dbots_object_timestamps
ORDER BY last_modified DESC;
```

所有需要记录 DDL 的角色都必须能访问扩展模式，并拥有 `dbots_event_data` 的读写权限；否则事件函数会捕获错误、发出 warning，并允许 DDL 继续执行，从而留下陈旧元数据。

### 主要对象

`dbots_tg_on_ddl_event` 在创建与修改操作的 `ddl_command_end` 阶段运行 `dbots_on_ddl_event()`。`dbots_tg_on_drop_event` 为删除操作运行 `dbots_on_drop_event()`，安装后默认启用。`dbots_event_data` 保存目录类与对象标识符及最新元数据。`dbots_object_timestamps` 解析可读对象身份，并关联当前对象 ACL。

删除列会更新时间戳到父对象。由于 PostgreSQL 没有提供所需的事件触发器数据，GRANT 与 REVOKE 变化不会记录为修改事件；视图会在查询时读取当前 ACL。

### 信任与维护边界

获得 `dbots_event_data` DML 权限的角色可以修改或删除记录，而函数会有意把内部失败降级为 warning。表中每个对象只保留最新一次观测，并不是追加式事件流。应把这些数据当作工具缓存，监控 warning，并在权限变化、扩展升级、恢复或不支持的 DDL 可能造成陈旧状态时进行校准。

事件触发器作用于整个数据库，安装需要较高权限。应针对数据库使用的每条 DDL 部署路径、扩展管理操作和 PostgreSQL 大版本测试该扩展。
