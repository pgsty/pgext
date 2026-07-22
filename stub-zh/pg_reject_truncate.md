## 用法

来源：

- [ProcessUtility 钩子实现](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.c)
- [扩展 SQL](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate--1.0.sql)
- [官方预加载配置](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.conf)
- [扩展控制文件](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.control)

`pg_reject_truncate` 安装一个实用命令钩子，用 `TRUNCATE IS NOT ALLOWED` 拒绝每条 `TRUNCATE` 语句。它是粗粒度的全服务器保护，没有按表、按角色或管理员绕过机制；它不是完整的数据丢失防护策略。

### 配置与行为

在所有服务器进程中预加载该库，并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_reject_truncate'
```

版本 `1.0` 的扩展 SQL 不包含可调用 SQL 对象。`CREATE EXTENSION pg_reject_truncate` 可以把扩展记录到系统目录，但 SQL 脚本从未引用共享库，因此不会自行加载钩子。预加载才是真正的激活步骤。在隔离的诊断会话中，`LOAD 'pg_reject_truncate'` 只会为该后端激活钩子。

加载后，任何直接的 `TRUNCATE` 都会被拒绝，包括超级用户发出的命令。实现没有用于临时关闭或限定规则范围的 GUC 或函数。

### 边界与钩子兼容性

钩子只检查节点类型为 `T_TruncateStmt` 的实用语句。它不会阻止 `DELETE`、`DROP TABLE`、分区操作、表重写、文件系统丢失或其他机制执行的破坏性操作。权限、备份、审计和经过测试的恢复流程才应是主要控制措施。

实现加载时会保存之前的 `ProcessUtility_hook`，但对每条允许的命令都直接调用 `standard_ProcessUtility`，而不是调用保存的钩子。因此，在 `pg_reject_truncate` 之前加载的实用钩子可能被非 `TRUNCATE` 命令绕过。与审计、安全、DDL 复制或其他实用钩子扩展组合时，除非进行源码级修复与集成测试，否则加载顺序并不安全。

核验的源码最后更新于 2020 年，并使用依赖 PostgreSQL 版本的钩子签名。应针对实际服务器主版本构建和测试，在重启后核验钩子顺序，并执行允许与拒绝两类命令测试后再依赖该保护。
