## 用法

来源：

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)

`databricks_auth` 是托管 Databricks Lakebase 提供的 PostgreSQL 身份验证扩展。官方目录为 PostgreSQL 16–18 列出的版本为 1.0。公开扩展列表没有说明可独立调用的身份验证 API，也没有提供可移植的源码包。

### 检查与启用

通过 Lakebase SQL Editor 或 PostgreSQL 客户端，使用具备权限的数据库角色，先检查供应商提供的版本及当前安装状态：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions WHERE name = 'databricks_auth';
```

供应商为目录中的扩展记录了标准安装方式：

```sql
CREATE EXTENSION IF NOT EXISTS databricks_auth;
```

此操作不会配置客户端凭据、授予应用访问权，也不能替代平台的身份验证设置。应遵循 Databricks 账号和数据库访问策略，不要移除已有的供应商管理安装。

### 版本边界

规划更新前，查询实际计算节点提供的版本和已安装版本。Lakebase 可能需要重启计算节点才能提供新版本，期间连接会中断。目录没有确认实现语言、许可证、SQL 函数、模式、依赖及预加载行为；本文不补造其他 API 或本地安装步骤。
