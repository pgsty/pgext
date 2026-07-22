## 用法

来源：

- [AWS RDS PostgreSQL 版本号指南](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.rds.version.html)
- [AWS PostgreSQL 密码加密指南](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL_Password_Encryption_configuration.html)
- [AWS autovacuum 诊断函数参考](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Functions.html)
- [AWS 扩展版本矩阵](https://docs.aws.amazon.com/AmazonRDS/latest/PostgreSQLReleaseNotes/postgresql-extensions.html)

`rds_tools` 是 AWS 维护、仅服务商提供的扩展，包含 Amazon RDS for PostgreSQL 管理函数。它不是可移植的开源扩展；可用版本与函数集合取决于准确的 RDS 引擎版本。当前 AWS 扩展矩阵在较新的引擎版本中列出 `1.9` 版。

### 创建与检查

使用适当权限的角色连接目标 RDS 数据库，确认可用性，并在每个需要其数据库局部对象的数据库中创建扩展。

```sql
SHOW rds.extensions;

SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'rds_tools';

CREATE EXTENSION rds_tools;
```

不要根据另一个 RDS 实例或大版本推断支持情况。AWS 说明引擎升级不会自动升级扩展；维护后应检查 `pg_extension.extversion` 与当前 AWS 矩阵。

### 已确认流程

在 RDS for PostgreSQL 15.2-R2 及之后适用版本中，可以返回 AWS RDS 补丁版本，而不仅是 PostgreSQL 社区版本。

```sql
SELECT rds_tools.rds_version();
```

在把实例改为强制 SCRAM 前，审计各角色密码的加密方式。AWS 要求管理员在范围内每个数据库/实例重复检查。

```sql
SELECT *
FROM rds_tools.role_password_encryption_type();
```

在 AWS 列明支持该特性的 RDS 小版本上，诊断 aggressive autovacuum 阻塞。为获得准确结果，应在事务 ID 最旧的数据库中运行查询。

```sql
SELECT blocker, database, blocker_identifier, wait_event,
       autovacuum_lagging_by, suggestion, suggested_action
FROM rds_tools.postgres_get_av_diag()
ORDER BY autovacuum_lagging_by DESC;
```

### 重要对象

- `rds_tools.rds_version()`：在支持版本上返回 RDS 引擎补丁标识。
- `rds_tools.role_password_encryption_type()`：为 SCRAM 迁移审计列出角色及其密码加密类型。
- `rds_tools.postgres_get_av_diag()`：在支持的小版本上报告可识别的 aggressive vacuum 阻塞及建议操作。

### 服务商与安全边界

函数可用性比扩展版本可用性更细。例如，AWS 只在 PostgreSQL 13–17 的特定小版本起记录 `postgres_get_av_diag()`；必须在目标实例检查函数目录与当前文档。

诊断输出可能包含 SQL 文本、角色信息、进程标识符，以及终止进程或破坏性操作建议。应限制访问并逐条审核 `suggested_action`，绝不能自动执行。`role_password_encryption_type()` 暴露的是密码*方法*而非明文密码，但结果仍属安全敏感信息。由于 AWS 没有把该扩展的实现 SQL 作为独立上游项目发布，不要虚构或依赖未记录的函数、签名或行为。
