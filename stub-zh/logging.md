## 用法

来源：

- [扩展控制文件](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging.control)
- [0.1 版本扩展 SQL](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging--0.1.sql)
- [C 实现](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging.c)

`logging` 是一个历史 PostgreSQL 实验，用于在永久与不记录日志的目录状态之间切换堆表及其索引。它唯一的函数会直接修改 `pg_class.relpersistence`；该项目早于 PostgreSQL 受支持的 `ALTER TABLE SET LOGGED` 与 `ALTER TABLE SET UNLOGGED` 工作流，不应被当作现代数据转换工具。

### 核心流程

已发布的接口如下：

```sql
CREATE EXTENSION logging;

SELECT enable_logging('staging_events'::regclass::oid, false);
SELECT enable_logging('staging_events'::regclass::oid, true);
```

`false` 请求不记录日志状态，`true` 请求永久且写入 WAL 的状态。函数发生变更时返回 `true`，关系已经处于目标状态时返回 `false`。

### 运维边界

- `enable_logging(relid oid, mode boolean) -> boolean` 是唯一的 SQL 函数。
- 它要求超级用户权限，并在表上取得 `AccessExclusiveLock`。
- 它通过更新系统目录元组修改堆表及其当前的每个索引。
- 它没有公开预加载要求，版本 `0.1` 被标记为可重定位。

### 运维说明

不要用该扩展修改生产数据的持久性。直接改变 `relpersistence` 会绕过受支持 PostgreSQL DDL 所提供的存储重写与安全检查，而且 C 源码使用了现代 PostgreSQL 已移除的服务器 API。应改用 `ALTER TABLE staging_events SET UNLOGGED` 或 `ALTER TABLE staging_events SET LOGGED`，并在模式管理中显式记录持久性转换。

如果为了考古或旧服务器而保留该扩展，应在一次性数据上测试崩溃恢复、索引、备库、备份以及服务器版本兼容性。不记录日志的关系在崩溃后会被截断，且其复制到备库的方式不同于永久关系。
