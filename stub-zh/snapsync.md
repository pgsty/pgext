## 用法

来源：

- [扩展及客户端文档](https://github.com/capraconsulting/skjera/blob/67da21683b0c4d38449c2230178afa64b1f875f0/packages/snapsync/README.md)
- [扩展控制文件](https://github.com/capraconsulting/skjera/blob/67da21683b0c4d38449c2230178afa64b1f875f0/packages/snapsync/extension/snapsync.control)
- [版本化 SQL](https://github.com/capraconsulting/skjera/blob/67da21683b0c4d38449c2230178afa64b1f875f0/packages/snapsync/extension/snapsync--0.1.0.sql)

`snapsync` 是 Skjera 仓库中的纯 SQL 扩展，通过事务快照轮询表变更。控制文件版本为 `0.1.0`；上游要求 PostgreSQL 14 或更高版本。它使用固定的 `snapsync` 模式，不可迁移到其他模式，无需共享库、预加载或重启。

### 跟踪表

安装扩展文件后，启用扩展并由表所有者开启跟踪：

```sql
CREATE EXTENSION snapsync;

CREATE TABLE attendees (id uuid PRIMARY KEY, display_name text);
SELECT snapsync.track('attendees');
```

跟踪操作会添加 `xid` 列、索引及触发器。更新时记录写入事务 ID，删除时生成墓碑记录。只支持单列键，键值以文本表示，不支持复合键。安装配置声明 `superuser = false`，但安装角色仍需具备 SQL 脚本使用的数据库及对象权限。

### 一致地拉取变更

每页变更及对应行的读取必须在同一个 `REPEATABLE READ` 事务内完成：

```sql
BEGIN ISOLATION LEVEL REPEATABLE READ;
SELECT changed, deleted, next_token, complete, more
FROM snapsync.changes('attendees', NULL, 1000);
COMMIT;
```

首次拉取传入空令牌。在该事务内，按读者通常的访问规则读取 `changed` 中键值对应的行，并从客户端副本删除 `deleted` 返回的键。应用整页后才能保存 `next_token`；`more` 为真时继续读取下一页。

`complete` 为真表示旧令牌不可用，客户端正在重新获取完整副本。应取完全部分页后再替换旧副本。并发提交可能导致重复返回某些行，消费端必须允许重复交付。

### 权限与维护

跟踪和维护函数已撤销 `PUBLIC` 执行权限。墓碑记录通过行级安全策略检查读者对原表的表级 `SELECT` 权限；读取变更行时仍须执行应用授权和行过滤。如果可见性取决于另一张表，该表变更时还应触碰被跟踪行。

`snapsync.prune(interval)` 删除旧墓碑。仍需使用已清理历史的客户端必须重新获取全量数据；保留时间应长于长事务的持续时间。逻辑恢复后，在提供拉取服务前调用 `snapsync.ensure_current()`，避免恢复的事务 ID 静默复用旧同步历史。截断表也会强制重新获取副本。`snapsync.untrack(regclass)` 停止跟踪，但保留事务 ID 列。
