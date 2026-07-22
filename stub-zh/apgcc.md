## 用法

来源：

- [Aurora PostgreSQL 扩展版本](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`apgcc` 是早期 Aurora PostgreSQL 9.6 版本曾提供的 AWS 内部扩展，AWS 目前已明确标记其不再受支持。它没有公开的上游 SQL 契约，因此不要围绕它设计新工作负载，也不要假定它能在 Aurora 之外安装。

### 可用性检查

仅通过系统目录检查；除非具体 Aurora 版本和 AWS 支持指引明确允许，否则不要执行 `CREATE EXTENSION`。

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'apgcc';

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'apgcc';
```

### 运维说明

- AWS 只在较早的 Aurora PostgreSQL 9.6 版本中列出 1.0，并将这个内部扩展标记为不受支持。
- 已有安装属于迁移风险。升级到受支持的 Aurora 版本前，应盘点依赖、验证应用脱离该扩展后的行为，并遵循 AWS 升级指引。
- 引用的上游页面没有给出可移植 API、配置项或替代品；旧集群仍依赖它时应联系 AWS Support。
