## 用法

来源：

- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/README.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/quickstarts/postgres.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/reference/rls.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/reference/jwt-claims.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docker/postgresql/cloudsync.control.in)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/src/cloudsync.h)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/src/postgresql/sql_postgresql.c)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/src/postgresql/migrations/README.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/LICENSE)
- [Official source](https://github.com/sqliteai/sqlite-sync/releases/tag/1.2.4)
- [Official source](https://github.com/sqliteai/sqlite-sync/releases/download/1.2.4/cloudsync-postgresql18-macos-arm64-1.2.4.tar.gz)

`cloudsync` 是 SQLite AI CRDT 同步栈的 PostgreSQL 端。发行版 1.2.4 安装 SQL 扩展版本 1.2。PostgreSQL 集成仍处于 beta 阶段，网络同步还需要独立的 CloudSync 服务和 SQLite 客户端集成。

### 初始化表

安装与服务器平台及 PostgreSQL 大版本匹配的官方文件。已发布的 PostgreSQL 制品面向 15、17、18，Linux 二进制分别要求至少 15.14、17.6、18.0。无需共享预加载。由管理员及表所有者完成初始化：

```sql
CREATE EXTENSION cloudsync;
SELECT cloudsync_version();
CREATE TABLE sync_tasks (
  id text PRIMARY KEY,
  title text NOT NULL DEFAULT ''
);
SELECT cloudsync_init('sync_tasks');
INSERT INTO sync_tasks VALUES ('task-1', 'Review offline changes');
SELECT cloudsync_is_enabled('sync_tasks');
```

控制文件声明扩展可信且可重定位，但安装会创建 C 语言函数，初始化还可能创建数据库级 DDL 事件触发器。需核对安装角色的实际权限，并由管理员完成设置，不应向应用角色授予不受信任语言的创建权。触发器使用 `plpgsql`，通过 `public.app_schema_version` 跟踪模式变化。

在 PostgreSQL 中显式创建同步表。允许多个客户端独立写入时，使用全局唯一的文本或 UUID 标识，并保持列类型兼容。初始化会添加触发器和 CRDT 影子元数据，不会启动复制工作进程。SQLite 端的网络函数、独立服务的认证配置属于其他组件。

### 同步对象与访问

`cloudsync_version()` 返回完整库版本，`cloudsync_siteid()` 和 `cloudsync_db_version()` 标识同步状态。`cloudsync_is_enabled`、`cloudsync_enable`、`cloudsync_disable` 控制表跟踪。`cloudsync_payload_encode`、`cloudsync_payload_decode`、`cloudsync_payload_apply` 交换二进制变更载荷，其计数不能替代最终行数据核验。

服务验证客户端 JWT 后，设置 `request.jwt.claims` 与数据库角色，再应用 RLS。能够直接设置这些声明的数据库客户端不在该认证边界内。仅授予所需的表、序列、影子表和函数权限，将初始化与清理函数留给管理员。RLS 拒绝写入时抛出 SQLSTATE `42501`，并回滚该次应用语句。删除标记含有主键，会传递给同步该表的所有设备，因此敏感主键需要另行设计处理。

### 模式变更与恢复

按官方流程使用 `cloudsync_begin_alter`、`cloudsync_commit_alter` 包围模式变更，并协调各客户端模式兼容性。恢复规划需同时覆盖应用数据、影子表、站点标识、同步版本和服务状态。关闭跟踪不会撤销已经送达的变更。清理与终止操作会删除同步状态，应作为明确的管理操作处理。

PostgreSQL 文本不能包含 NUL 字符，文档说明 SQLite 文本会在第一个 NUL 处截断并记录警告。二进制值应存入 BYTEA，并验证跨引擎类型转换。CRDT 收敛不能保证应用业务约束，也不能代替权限控制。

### 升级

保持客户端、服务和扩展制品兼容。补丁发行版替换库，SQL 主次版本保持不变。按上游的重启和重连流程确保会话使用新二进制。SQL 主次版本变化时，安装配套升级脚本，并在各数据库执行 `ALTER EXTENSION cloudsync UPDATE`。1.2.4 制品提供从 1.0 经 1.1 到 1.2 的路径。升级后用 `cloudsync_version()` 确认二进制版本，并核验同步行与权限。
