


## 用法

- 来源：[README](https://github.com/rustwizard/pg_command_fw/blob/master/README.md)

`pg_command_fw` 是 PostgreSQL command firewall。它通过 `ProcessUtility` hook 拦截 DDL 和 utility commands，并通过 post-parse analyze hook 阻断部分内置文件读取函数。每个命令类别都有自己的 GUC 控制。

### 启用扩展

该扩展必须 preload：

```ini
shared_preload_libraries = 'pg_command_fw'
```

然后在数据库中启用：

```sql
CREATE EXTENSION pg_command_fw;
```

Pigsty 包元数据记录版本 `0.1.0`，覆盖 PostgreSQL 15-18，并说明需要 preload 才能在所有会话中激活 hooks。上游 README 也记录 PostgreSQL 15-18 支持。

### 命令类别

上游 README 记录了这些 firewall 类别：

- `TRUNCATE`：`pg_command_fw.block_truncate`，默认 `on`，阻断非超级用户。
- `DROP TABLE`：`pg_command_fw.block_drop_table`，默认 `off`，启用后阻断非超级用户。
- `ALTER SYSTEM`：`pg_command_fw.block_alter_system`，默认 `on`，阻断所有人。
- `LOAD`：`pg_command_fw.block_load`，默认 `on`，阻断所有人。
- `COPY ... PROGRAM`：`pg_command_fw.block_copy_program`，默认 `on`，阻断所有人。
- 普通 `COPY`：`pg_command_fw.block_copy`，默认 `off`，启用后阻断非超级用户。
- `pg_read_file()`、`pg_read_binary_file()` 和 `pg_stat_file()`：`pg_command_fw.block_read_file`，默认 `on`，阻断所有人。

部分类别只阻断非超级用户，另一些类别阻断包括超级用户在内的所有人。除非显式列入 `pg_command_fw.blocked_roles`，超级用户只会豁免非超级用户类别。

### 重要 GUC

- `pg_command_fw.enabled`：启用或禁用所有检查
- `pg_command_fw.block_truncate`
- `pg_command_fw.block_drop_table`
- `pg_command_fw.production_schemas`
- `pg_command_fw.block_alter_system`
- `pg_command_fw.block_load`
- `pg_command_fw.block_copy_program`
- `pg_command_fw.block_copy`
- `pg_command_fw.block_read_file`
- `pg_command_fw.blocked_roles`
- `pg_command_fw.hint`
- `pg_command_fw.audit_log_enabled`

设置 `production_schemas` 时，`DROP TABLE` 检查仅限这些 schema 中显式带 schema 的表名；README 说明未限定名称不会通过 `search_path` 解析。

### 审计日志

扩展会在 `command_fw.audit_log` 中记录被拦截的命令。README 记录的列包括：

- timestamp
- session 和 current user 名称
- 原始 query text
- command type
- target schema 或 object
- client address
- 命令是否被阻断
- 内部阻断原因

被阻断的审计插入是 best-effort，因为该行会随被阻断事务回滚；请使用 PostgreSQL server log 作为被阻断事件的权威记录。

### 示例

在生产 schema 中阻断 `TRUNCATE` 和 `DROP TABLE`：

```sql
ALTER SYSTEM SET pg_command_fw.block_truncate = on;
ALTER SYSTEM SET pg_command_fw.block_drop_table = on;
ALTER SYSTEM SET pg_command_fw.production_schemas = 'public,payments';
ALTER SYSTEM SET pg_command_fw.hint = 'Contact your DBA to request access';
SELECT pg_reload_conf();
```

阻断某个角色执行任何受管命令：

```sql
ALTER SYSTEM SET pg_command_fw.blocked_roles = 'app_deploy';
SELECT pg_reload_conf();
```

在维护会话中临时禁用 firewall：

```sql
SET pg_command_fw.enabled = off;
TRUNCATE big_table;
SET pg_command_fw.enabled = on;
```
