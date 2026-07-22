## 用法

来源：

- [官方德文文档](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/README.de.md)
- [官方扩展 SQL](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/sql/plsh_mtb--1.0.sql)
- [官方备份控制脚本](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/src/plsh_mtb)

`plsh_mtb` 是基于 PL/SH 与 Korn shell 的控制器，允许数据库用户启动、暂停、继续或中止服务器端逻辑备份，但不向他们提供恢复接口。它以 PostgreSQL 操作系统账户运行可配置的转储命令，并在数据库表中记录控制状态。

### 前置条件与配置

安装 `plsh` 扩展、`mksh`、`plsh_mtb` 可执行文件以及数据库扩展。用户调用前，管理员必须定义自定义设置：

```conf
plsh_mtb.dump = 'pg_dump -Z 5 PGDATABASE -f BACKUPFILE.gz'
plsh_mtb.dir = '/srv/postgresql/tenant-backups'
plsh_mtb.log = 'syslog'
```

`PGDATABASE` 和 `BACKUPFILE` 是控制器替换的字符串占位符。目标目录与命令必须允许 PostgreSQL 服务账户写入和执行。

```sql
CREATE EXTENSION plsh;
CREATE EXTENSION plsh_mtb;
```

### 控制与状态

`customer_backup` 接受枚举值 `start`、`stop`、`continue` 和 `abort`：

```sql
SELECT customer_backup('start');
SELECT * FROM plsh_mtb_backups ORDER BY started DESC;
SELECT customer_backup('stop');
SELECT customer_backup('continue');
SELECT customer_backup('abort');
```

`plsh_mtb_backups` 记录 `filename`、开始与结束时间、状态和控制进程 PID。状态包括 `running`、`stopped`、`aborted`、`failed` 与 `done`；这些只是控制器元数据，不能证明备份完整或可恢复。

### 权限与恢复边界

- 安装 SQL 明确把 `plsh_mtb_backups` 的 `SELECT` 授予 `PUBLIC`，而 PostgreSQL 默认也把函数执行权授予 `PUBLIC`。应撤销二者，再把状态查看与控制权限分别授予预期租户角色。
- 配置的转储命令会由 shell 展开，并以 PostgreSQL 操作系统账户执行。只有可信管理员可以设置它；应正确引用路径、约束备份目录，并防止租户控制的数据库名或配置变成 shell 语法。
- `stop`、`continue` 与 `abort` 会发送操作系统信号。当前审阅脚本只检查 `/proc/<pid>` 是否存在，而更强的命令行验证被注释；陈旧元数据或 PID 重用可能让信号发给无关进程。
- 当前脚本的失败命令处理会调用未定义的 `mark_backup_failed` 辅助函数。因此转储失败后状态可能错误；依赖监控前必须修补并测试这条路径。
- 外部进程和文件系统与 PostgreSQL 不具有事务性。回滚 SQL 调用不会撤销已创建文件或信号，后端或会话故障也可能留下孤儿进程或陈旧记录。
- 扩展不提供恢复、保留、加密、异地复制、校验和、目录或恢复验证。必须单独补充这些控制，并定期使用真实制品执行恢复演练。
