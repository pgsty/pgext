## 用法

来源：

- [README](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/README.md)
- [Control](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgauth/pgauth.control)
- [SQL](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgauth/pgauth--0.1.0.sql)

`pgauth` 提供口令哈希、签名 JWT 会话以及行级安全策略所用的 SQL 辅助函数。此 0.1.0 PGEverything 源码快照面向其 PostgreSQL 16 部署。身份声明存放在可修改的会话设置中，因此应通过可信应用网关使用，不能用来隔离可执行任意 SQL 的客户端。

### 核心用法

安装 `pgcrypto`、`pgjwt` 和 `pgauth`，并确保 PL/pgSQL 可用。由扩展所有者初始化私有签名密钥；PGEverything 容器从 `PGE_JWT_SECRET` 配置它。应生成新的私有值，不能沿用容器开发默认值。这些辅助功能不需要 pgauth 预加载库。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pgjwt;
CREATE EXTENSION pgauth;
INSERT INTO auth.secret (value)
VALUES (encode(gen_random_bytes(32), 'hex'));

SELECT auth.register('alice@example.com', :'user_password');
BEGIN;
SELECT auth.authenticate(auth.login('alice@example.com', :'user_password'));
SELECT auth.uid(), auth.role();
COMMIT;
```

应安全提供 psql 变量 `user_password`。认证声明仅在当前事务内有效，认证与受保护操作必须放在同一事务中。本例只用于初始化新数据库，不应随意更换已有签名密钥。

### SQL 接口

- `auth.register(text,text,text)` 对口令求哈希并返回用户 UUID；可选第三参数保存角色声明，不会创建 PostgreSQL 角色。
- `auth.login(text,text)` 返回一小时有效的 JWT；凭据错误时返回 NULL。
- `auth.authenticate(text)` 验证签名和过期时间，设置 `auth.claims`，并返回布尔结果；`auth.uid()` 与 `auth.role()` 读取这些声明。
- `auth.users` 与 `auth.secret` 是扩展配置表，逻辑备份包含其数据，应按照认证数据保护备份。

### RLS 与安全边界

通过 `auth.uid()` 定义应用表策略，并使用非超级用户的应用角色；表所有者及其他可绕过 RLS 的角色不会受到通常的限制。使用角色声明授权前，应限制公开注册及由调用者指定的角色参数。扩展公开了 API 的执行权限，但限制了对密钥表和密钥读取函数的直接访问。

不可信 SQL 可以直接修改 `auth.claims`，因此仅凭这两个读取函数不能认证此类客户端。安全定义者函数的路径也包含 `public`，不应允许不可信用户写入该模式。扩展没有刷新令牌 API，创建固定的 `auth` 模式，且上游没有明确的许可证声明。
