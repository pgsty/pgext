## 用法

来源：

- [README.md](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/README.md)
- [pg_deribit.control](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/pg_deribit.control)
- [Makefile](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/Makefile)
- [Dockerfile](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/Dockerfile)
- [sql/functions/environment.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/functions/environment.sql)
- [sql/functions/auth.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/functions/auth.sql)
- [sql/functions/internal_url_endpoint.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/functions/internal_url_endpoint.sql)
- [sql/endpoints/public_get_currencies.sql](https://github.com/rosssaunders/pg_deribit/blob/15a02c6ac725dadbac6dd26b75ac70d519a9e4c2/sql/endpoints/public_get_currencies.sql)

`pg_deribit` 1.0 通过 SQL 暴露 Deribit JSON-RPC 接口，依赖 Omnigres HTTP 扩展；上游容器使用 PostgreSQL 17。公共接口可读取市场元数据，需认证的接口可执行账户操作。

### 公共测试网工作流

```sql
CREATE EXTENSION pg_deribit CASCADE;
SELECT deribit.enable_test_net();
SELECT currency FROM deribit.public_get_currencies() ORDER BY currency;
```

### 依赖与配置

此扩展不可重定位，使用 `deribit` 模式，依赖 `pgcrypto`、`omni_http` 与 `omni_httpc`。安装遵循默认的超级用户限制。扩展仅定义 SQL 函数与类型，没有自身的共享库或预加载配置；依赖扩展必须已安装且可用。

默认连接生产端点。`deribit.enable_test_net()` 与 `deribit.disable_test_net()` 修改当前会话的 `deribit.set_test_net`，新连接需要重新启用测试网。调用需要出站网络，并依赖远端 API。

### 认证调用

`deribit.set_client_auth(client_id, client_secret)` 与 `deribit.set_access_token_auth(client_id, client_secret, access_token, refresh_token)` 设置会话凭据；`deribit.get_auth()` 读取这些凭据。应按需授权 SQL 访问，避免凭据进入查询日志或共享示例。私有封装包含会改变账户状态的操作：数据库事务回滚不能撤销已完成的远端动作。授权前应核对各端点签名。

凭据保存在会话配置中。上游设置函数使用 `%s` 插值而非 SQL 字面量转义，因此不能传入不可信的凭据字符串。函数保留默认的 `PUBLIC` 执行权限，应按部署需求收紧。私有请求与响应数据会持久化至 `deribit.internal_archive`，需要明确的保留策略。
