## 用法

来源：

- [当前 AWS Aurora PostgreSQL 扩展支持历史](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`apgcc` 曾是 AWS 内部扩展，并非可普遍安装的开源扩展。当前 AWS 历史表只在 Aurora PostgreSQL 9.6.3 和 9.6.6 中记录版本 `1.0`，后续 9.6 版本均标为 `NA`，并明确说明它已不再受支持。

目前没有可记录的功能 API。在 Aurora 集群上，可以用下面的诊断查询确认供应商是否仍向特定引擎版本暴露该扩展：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'apgcc';
```

当前受支持版本返回空结果属于预期行为。

### 注意事项

- 不要尝试提供同名第三方二进制。AWS 把 `apgcc` 定义为内部供应商扩展，且未发布源码或 SQL 接口。
- 目录版本 `1.0` 是历史元数据，不代表当前可用。实时可用性应以具体引擎版本的 AWS 表格和 `pg_available_extensions` 为准。
- 不要把这个已退役扩展与 `apg_plan_mgmt` 或其他当前支持的 Aurora 功能混淆。
