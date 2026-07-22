## 用法

来源：

- [腾讯云 PostgreSQL 官方扩展矩阵](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_pwdcheck` 是仅在腾讯云数据库 PostgreSQL 中提供的扩展。它只能在该托管服务边界内使用，密码策略语义应从腾讯云控制台或支持渠道获取：公开矩阵只确认可用性，并未公开 SQL 函数、策略参数、预加载行为或源代码。

### 核心流程

当前腾讯云矩阵为 PostgreSQL 10 至 17 列出版本 `1.0`，PostgreSQL 18 则没有条目。先确认实际实例引擎版本提供该扩展，再通过腾讯云支持的扩展流程启用：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'tencentdb_pwdcheck';

CREATE EXTENSION tencentdb_pwdcheck;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'tencentdb_pwdcheck';
```

如果第一个查询没有返回行，不要自行复制文件，也不要拿社区 `passwordcheck` 模块替代；应选择受支持的引擎版本或联系腾讯云。

### 提供商边界

引用的腾讯云页面只说明软件包可用性，并未说明检查哪些密码操作、如何配置规则、是否评估已有凭据，或是否需要重启。应把这些细节视为由提供商管理且随版本变化，并针对目标实例查阅控制台帮助或腾讯支持答复。

### 运维注意事项

在把策略应用到生产用户前，先用一次性角色测试启用流程与密码变更。数据库策略变更要与应用密钥轮换协调，并保留经过独立验证的管理员恢复路径；过严或理解有误的密码规则可能阻断用户创建、密码轮换或紧急访问。
