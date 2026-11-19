## 用法

来源：

- [pgAudit 19.0 README](https://github.com/pgaudit/pgaudit/blob/19.0/README.md)
- [PostgreSQL 19 发行版](https://github.com/pgaudit/pgaudit/releases/tag/19.0)
- [发行版控制文件](https://github.com/pgaudit/pgaudit/blob/19.0/pgaudit.control)

`pgaudit` 通过 PostgreSQL 日志记录会话和对象事件。应选择匹配服务器的分支：PostgreSQL 19 对应 19.0，18 对应 18.1，17 对应 17.2，16 对应 16.2，15 对应 1.7.2，14 对应 1.6.4。各分支分别针对对应服务器构建，升级 PostgreSQL 不会使旧库自动具备 ABI 兼容性。

### 启用审计

将 `pgaudit` 加入 `shared_preload_libraries` 并重启服务器，再以超级用户创建扩展，之后才能设置 `pgaudit.log`：

```sql
CREATE EXTENSION pgaudit;
```

只有超级用户能修改审计设置。应规划并保护日志卷：SQL 文本和参数可能含敏感信息，广泛记录会占用大量磁盘。扩展提供审计记录，但本身并不等于满足合规要求。

### 配置参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pgaudit.log` | `none` | 要记录的语句类别：`READ`、`WRITE`、`FUNCTION`、`ROLE`、`DDL`、`MISC`、`MISC_SET`、`ALL` |
| `pgaudit.log_catalog` | `on` | 当所有关系都在 pg_catalog 中时记录语句 |
| `pgaudit.log_client` | `off` | 向客户端显示审计日志消息 |
| `pgaudit.log_level` | `log` | 审计条目的日志级别 |
| `pgaudit.log_parameter` | `off` | 在日志中包含语句参数 |
| `pgaudit.log_parameter_max_size` | `0` | 参数最大字节数（0=无限制） |
| `pgaudit.log_relation` | `off` | 在 SELECT/DML 中为每个关系生成单独的日志条目 |
| `pgaudit.log_rows` | `off` | 在日志中包含行数 |
| `pgaudit.log_statement` | `on` | 在日志中包含语句文本 |
| `pgaudit.log_statement_once` | `off` | 仅在第一个条目中记录语句文本 |
| `pgaudit.role` | (无) | 对象审计日志的主角色 |

### 会话审计日志

记录所有 DML 和 DDL，按关系详细记录：

```sql
SET pgaudit.log = 'write, ddl';
SET pgaudit.log_relation = on;
```

记录除杂项命令外的所有内容：

```sql
SET pgaudit.log = 'all, -misc';
```

输出示例：
```text
AUDIT: SESSION,1,1,DDL,CREATE TABLE,TABLE,public.account,create table account(...)
AUDIT: SESSION,2,1,READ,SELECT,,,select * from account
```

### 对象审计日志

向审计角色授予权限以控制记录哪些关系：

```sql
SET pgaudit.role = 'auditor';

GRANT SELECT, DELETE
   ON public.account
   TO auditor;
```

现在对 `account` 表的任何 `SELECT` 或 `DELETE` 操作都会被审计记录。

### 日志格式

条目为 CSV 格式，包含以下字段：`AUDIT_TYPE`、`STATEMENT_ID`、`SUBSTATEMENT_ID`、`CLASS`、`COMMAND`、`OBJECT_TYPE`、`OBJECT_NAME`、`STATEMENT`、`PARAMETER`。

### 日志边界

审计输出使用尽力而为的 PostgreSQL 日志机制，不具备事务持久性。回滚的语句可能留下记录，已提交事务也不保证具有持久化审计记录。应单独保护和收集日志，不能将其视为事务完整的账本。
