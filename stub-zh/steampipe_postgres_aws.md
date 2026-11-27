## 用法

来源：

- [CHANGELOG](https://github.com/turbot/steampipe-plugin-aws/blob/v1.37.0/CHANGELOG.md)

- [Upstream source](https://github.com/turbot/steampipe-plugin-aws/releases/tag/v1.37.0)
- [Upstream source](https://github.com/turbot/steampipe-plugin-aws/releases/download/v1.37.0/steampipe_postgres_aws.pg15.darwin_arm64.tar.gz)

- [Published 0.125.0 package page](https://database.dev/toliver38/steampipe_postgres_aws)
- [Official AWS plugin documentation](https://hub.steampipe.io/plugins/turbot/aws#postgres-fdw)
- [Official Postgres FDW configuration guide](https://steampipe.io/docs/steampipe_postgres/configure)

`steampipe_postgres_aws` 是 Turbot AWS Steampipe 插件的 PostgreSQL FDW 构建。它把对导入外部表的 SQL 扫描转换为实时 AWS API 调用，从现有 PostgreSQL 数据库中进行基础设施清点和配置分析。该软件包不是 Steampipe CLI；除非显式执行，否则不会把 AWS 资源复制到本地表。

### 核心流程

安装与操作系统、体系结构和 PostgreSQL 大版本匹配的二进制后，创建扩展，配置外部服务器，再把表导入某个模式：

```sql
CREATE EXTENSION steampipe_postgres_aws;

CREATE SERVER steampipe_aws
  FOREIGN DATA WRAPPER steampipe_postgres_aws
  OPTIONS (config 'profile = "default"
regions = ["us-east-1"]');

CREATE SCHEMA aws;
COMMENT ON SCHEMA aws IS 'Steampipe AWS FDW';

IMPORT FOREIGN SCHEMA aws
  FROM SERVER steampipe_aws
  INTO aws;

SELECT name, region
FROM aws.aws_s3_bucket
ORDER BY name;
```

服务器选项 `config` 是对换行敏感的 HCL 字符串，包含 AWS 插件连接参数。除了 profile，也可以使用 AWS 插件文档列出的凭据机制。任何环境变量或家目录下的文件都在 PostgreSQL 服务端操作系统用户的上下文中解析，而不是交互式客户端用户的上下文。

### 对象与查询模型

- 扩展/FDW：`steampipe_postgres_aws`。
- 外部服务器：由用户定义；每套不同的 AWS 账户/区域配置建立一个。
- 导入模式：由用户定义；`IMPORT FOREIGN SCHEMA` 会创建插件提供的外部表定义。
- 外部表：按服务组织的 `aws_*` 表，例如 `aws_s3_bucket`；可用列、过滤条件、权限和示例需查阅 Steampipe Hub 中每张表的文档。

查询会调用 AWS API，可能较慢、受到限流或分页影响、呈现最终一致性，或被 IAM 拒绝。当表文档指出必需或可选过滤条件时，应在 `WHERE` 子句中提供受支持的关键列。若需要可重复分析、连接或历史比较，应把快照物化到本地。

### 安全与版本边界

优先使用 AWS profile、工作负载角色或其他短期凭据来源，不要把静态密钥放入 `pg_foreign_server`。只授予查询目标表所需的 IAM 读取权限，并限制对服务器选项和已导入外部表的访问。外部扫描可能产生 AWS API 活动及相关服务费用。

已验证的官方 1.37.0 发布包含独立 PostgreSQL AWS 制品；检查的 PostgreSQL 15 macOS arm64 制品仍使用 SQL 版本 1.0。旧 database.dev 包仍是独立的 0.125.0 发布。必须按实际平台和 PostgreSQL 大版本选择制品，不能从插件标签推断新大版本支持；仅凭 SQL extversion 也无法确定插件修订。升级时应审查表定义变化，有计划地重新导入或重建外部表，并保留授权与依赖对象。安装说明包含破坏性的 DROP 示例，全新安装并不需要这些操作，不应直接照搬到现有部署。无需设置 shared_preload_libraries。

### 1.37.0 的变化

`aws_identitystore_user` 表增加 `user_status`，按 `id` 过滤时会返回完整字段。`aws_vpc_vpn_connection` 的选项增加了 BGP 日志设置。使用新字段前需核对已导入的表定义。

默认的 `aws_ec2_describe_attribute` 限流器可减少逐行 EC2 属性调用触发的节流，但不代表 AWS 配额不再受限。本版还修复了安全组规则标识不稳定，以及 SDK 调用失败被表现为数据缺失的问题。替换二进制后应重新核对已保存的资源清单和错误处理，并保留权限及依赖视图。
