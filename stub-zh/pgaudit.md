

## 用法

> [pgaudit: 开源 PostgreSQL 审计日志](https://github.com/pgaudit/pgaudit)

pgAudit 通过标准 PostgreSQL 日志功能提供详细的会话和/或对象审计日志，生成政府、金融或 ISO 认证所需的审计追踪。

```sql
CREATE EXTENSION pgaudit;
```

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
```
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
