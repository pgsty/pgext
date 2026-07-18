## 用法

来源：

- [华为云 RDS for PostgreSQL 支持的扩展](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0045.html)

`obs_fdw` `1.0.0` 版是华为云提供商扩展。当前华为 RDS for PostgreSQL 支持矩阵在 PostgreSQL 12–17 中列出它（9.5–11 未列出），但引用的公开页面没有说明外部服务器/表选项、凭据、格式、读写语义或权限。

### 提供商安全发现

尝试安装前先检查具体 RDS 实例，然后盘点已安装包装器，不要猜测选项：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions WHERE name = 'obs_fdw';
CREATE EXTENSION obs_fdw;
SELECT fdwname, fdwhandler::regproc, fdwvalidator::regproc
FROM pg_foreign_data_wrapper WHERE fdwname = 'obs_fdw';
```

应把可用性视为由 RDS 引擎版本、最新小版本、区域和提供商发布节奏共同控制。这不能证明扩展可移植到自管 PostgreSQL。创建服务器或映射前，应通过控制台/支持渠道取得实例对应的华为选项与凭据契约；使用权限受限的 OBS 身份，绝不要从无关 FDW 推断 secret 选项名。
