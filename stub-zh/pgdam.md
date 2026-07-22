## 用法

来源：

- [AWS Aurora PostgreSQL 扩展支持矩阵](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)
- [AWS Aurora PostgreSQL 扩展指南](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.html)
- [AWS 扩展版本概览](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Extensions.html)

`pgdam` 是 AWS 为 Amazon Aurora PostgreSQL 提供的扩展。AWS 当前在受支持的 Aurora 引擎版本中列出版本 `1.7`，但没有发布开源仓库或公开的 SQL 对象/函数参考。应把它视为提供方管理组件，并在目标集群上检查实际可用接口。

### 核心流程

先确认确切 Aurora 引擎版本提供该扩展，并且集群的 `rds.allowed_extensions` 策略允许安装：

```sql
SELECT name, default_version, installed_version, comment
FROM pg_available_extensions
WHERE name = 'pgdam';

CREATE EXTENSION pgdam;

SELECT extname, extversion, extnamespace::regnamespace
FROM pg_extension
WHERE extname = 'pgdam';
```

使用 `\dx+ pgdam` 和该集群的 PostgreSQL 系统目录检查提供方安装的对象与权限。不要从同名第三方项目推断函数。

### 平台边界

可用性与版本绑定到 Aurora PostgreSQL 引擎版本。引擎大版本升级时扩展不会自动升级；AWS 建议单独检查并升级扩展。根据集群配置，安装可能需要 `rds_superuser` 或委派的扩展管理权限。

### 注意事项

这里引用的 AWS 公开材料只能证明 `pgdam` `1.7` 的可用性，不能证明其应用 API 或行为。缺少提供方对象参考时，没有足够权威证据发布函数调用、配置参数或运行保证。

应在非生产 Aurora 集群测试已安装扩展，记录实际对象定义与授权，并在依赖它之前向 AWS 确认支持范围。升级前重新核对引擎专属矩阵，因为支持版本和允许扩展策略都可能独立变化。
