
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pg_command_fw;
> ALTER SYSTEM SET pg_command_fw.block_truncate = on;
> ALTER SYSTEM SET pg_command_fw.production_schemas = 'public,payments';
> SELECT pg_reload_conf();
> ```
>
> 来源：[README](https://github.com/rustwizard/pg_command_fw)

`pg_command_fw` 是 PostgreSQL 命令防火墙。它通过 `ProcessUtility` 钩子拦截 DDL 和 utility 命令，并通过 post-parse analyze 钩子拦截部分危险的内置文件读取函数。每个命令类别都由独立的 GUC 控制。

## 安装

扩展必须预加载：

```ini
shared_preload_libraries = 'pg_command_fw'
```

然后在数据库中启用：

```sql
CREATE EXTENSION pg_command_fw;
```

## 命令类别

上游 README 记录了以下防火墙类别：

- `TRUNCATE`
- `DROP TABLE`
- `ALTER SYSTEM`
- `LOAD`
- `COPY ... PROGRAM`
- 普通 `COPY`
- `pg_read_file()`、`pg_read_binary_file()` 和 `pg_stat_file()`

其中部分类别仅阻止非超级用户，另一些则连超级用户也会阻止。只有当超级用户未被列入 `pg_command_fw.blocked_roles` 时，才会免于非超级用户类检查。

## 重要 GUC

- `pg_command_fw.enabled` 用于整体启用或禁用所有检查
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

## 审计日志

扩展会将拦截到的命令写入 `command_fw.audit_log`。README 中描述的字段包括：

- 时间戳
- 会话用户和当前用户
- 原始查询文本
- 命令类型
- 目标模式或对象
- 客户端地址
- 是否被阻止
- 内部阻止原因

## 示例

在生产模式下阻止 `TRUNCATE` 和 `DROP TABLE`：

```sql
ALTER SYSTEM SET pg_command_fw.block_truncate = on;
ALTER SYSTEM SET pg_command_fw.block_drop_table = on;
ALTER SYSTEM SET pg_command_fw.production_schemas = 'public,payments';
ALTER SYSTEM SET pg_command_fw.hint = 'Contact your DBA to request access';
SELECT pg_reload_conf();
```

阻止特定角色执行任何受防火墙管控的命令：

```sql
ALTER SYSTEM SET pg_command_fw.blocked_roles = 'app_deploy';
SELECT pg_reload_conf();
```
