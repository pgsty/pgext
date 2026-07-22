## 用法

来源：

- [官方 README](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/README.md)
- [0.1 版 SQL 实现](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/auth0--0.1.sql)
- [扩展控制文件](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/auth0.control)

`auth0` 是一个面向部分 Auth0 Authentication 与 Management API 操作的 SQL 客户端。它通过必需的 `http` 扩展从 PostgreSQL 发出 HTTPS 请求，使查询能够创建、读取或更新 Auth0 用户，并触发密码变更消息。

### 核心流程

先创建依赖与扩展，再为当前数据库角色保存租户配置。SQL 实现使用 `set_config`；README 中名称不同的 setter 并不存在于 `0.1` 版中。

```sql
CREATE EXTENSION http;
CREATE EXTENSION auth0;

SELECT auth0.set_config('domain', 'example.auth0.com');
SELECT auth0.set_config('client_id', 'client-id');
SELECT auth0.set_config('client_secret', 'client-secret');
SELECT auth0.set_config('connection', 'Username-Password-Authentication');

SELECT auth0.get_user_by_email('alice@example.com', 'user_id,email,name');
```

当缓存令牌为空或已过期时，`get_api_token()` 会执行客户端凭证交换。`get_user()`、`get_user_by_email()`、`create_user()` 与 `update_user()` 调用 Management API；`change_password_prompt()` 调用数据库连接端点。

### 运维边界

配置与缓存的访问令牌由 `ALTER ROLE` 写入 PostgreSQL 的每角色设置，并非保存在外部密钥存储中。应限制对该角色与系统目录的访问，并制定明确的密钥轮换方案。这些函数会在数据库后端内同步发出网络请求；因此 API 延迟、限流、TLS/DNS 故障及 Auth0 权限都会影响 SQL 执行。

实现会用调用方文本拼接部分请求路径，且并未覆盖通用 Auth0 API。只向可信角色授予执行权限，在应用侧校验标识符与 JSON 载荷，并在事务中采用前测试 HTTP 状态与错误行为。
