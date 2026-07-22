## 用法

来源：

- [AWS Aurora PostgreSQL 扩展支持表](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)
- [Amazon Aurora 产品页](https://aws.amazon.com/rds/aurora/)

`apgunit` 曾是 AWS 内部的 Aurora PostgreSQL 扩展，并不是可移植的社区扩展。AWS 目前明确说明该扩展已不再受支持。此条目只适合用于识别提到它的旧 Aurora 集群或转储；不要围绕它设计新的数据库功能。

### 历史边界

AWS 支持表仅在旧版 `Aurora PostgreSQL 9.6.3` 与 `Aurora PostgreSQL 9.6.6` 中列出 `apgunit` 版本 `1.0`，后续版本则显示为不可用。AWS 没有公开该组件的 control 文件、SQL API、配置参考或迁移接口。

因此不存在可记录的当前受支持启用流程。尤其不要假定现代 Aurora 集群会接受 `CREATE EXTENSION`、自托管 PostgreSQL 存在同名软件包，或可通过 `shared_preload_libraries` 恢复该组件。

### 运维建议

当接手的 Schema 或迁移引用 `apgunit` 时：

1. 确认它实际使用时的 Aurora 引擎版本。
2. 向 AWS Support 询问历史行为与受支持替代方案。
3. 在升级到不再提供该扩展的引擎版本前，移除或替换依赖对象。
4. 在独立目标集群上验证修改后的 Schema 与应用行为。

应把目录中的版本视为历史元数据，而不是 `apgunit` 今天仍可安装或使用的证据。
