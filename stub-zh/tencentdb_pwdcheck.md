## 用法

来源：

- [腾讯云 PostgreSQL 扩展版本矩阵](https://cloud.tencent.com/document/product/409/75121)
- [腾讯云：在实例中创建扩展](https://cloud.tencent.com/document/product/409/121744)
- [腾讯云展示预加载模块的示例](https://cloud.tencent.com/document/product/409/111199)

`tencentdb_pwdcheck` 是腾讯云数据库 PostgreSQL 专有的强密码认证扩展。当前支持矩阵显示 PostgreSQL 10 至 17 提供 `1.0` 版本，PostgreSQL 18 不提供该包。腾讯云未公开其源码、control 文件、密码策略参数或 SQL API。

启用前先检查目标实例的实际目录：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'tencentdb_pwdcheck';
```

腾讯云文档展示的 `shared_preload_libraries` 中包含 `tencentdb_pwdcheck`。如果托管实例尚未预加载该模块，应在控制台参数设置中加入它，保存参数变更并等待实例重启；随后在需要扩展对象的各数据库中创建扩展：

```sql
CREATE EXTENSION tencentdb_pwdcheck;
```

修改 `shared_preload_libraries` 会重启托管实例，因此应安排维护时段，并确保客户端能够自动重连。云厂商没有公开各项策略参数，不应假定社区 `passwordcheck` 的配置同样适用；应以所选引擎版本的控制台与腾讯云支持信息为准。
