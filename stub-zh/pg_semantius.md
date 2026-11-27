## 用法

来源：

- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/README.md)
- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/pg_semantius.control)
- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/pg_semantius--0.5.0.sql)
- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/META.json)
- [Official source](https://github.com/semantius/semantius/releases/tag/v0.5.0)

`pg_semantius` 0.5.0 在 PostgreSQL 18 上提供应用数据字典、RBAC、RLS 和消息队列。本文介绍带版本的 PostgreSQL 扩展分发方式；独立 CLI 可通过普通连接向托管数据库部署模式。

### 两步安装

使用 UTF8 数据库和具有 BYPASSRLS 属性的超级用户。逐条执行，外层不要开启事务，也不要添加 CASCADE：

```sql
CREATE EXTENSION pg_semantius;
CALL semantius.migrate();
SELECT semantius.version();
SELECT * FROM semantius.status();
SELECT * FROM semantius.pending();
```

第一条语句创建集群角色、安装器模式及函数。`semantius.migrate()` 将应用对象安装为普通数据库对象，并在每个迁移文件后提交事务。它将 `pgcrypto` 创建在 `public`；若该扩展已位于其他模式，安装会被拒绝。后续文件失败时，前面的文件仍已提交，下次调用会继续执行。

### 身份与应用访问

安装会创建 `authenticated`、`semantius_user`、`semantius_authenticator`、`semantius_owner` 四个无密码的 NOLOGIN 角色。所有者角色必须保持无 LOGIN、无成员。已有角色属性不兼容时会拒绝安装。另行配置应用登录，并让普通请求受到 RLS 约束。

会话模式信任应用层完成认证，并在请求事务内设置 `request.jwt.claims`。若调用者可直接使用能伪造这些设置的角色，就会突破该信任边界。需要校验受众时，应设置 `public._settings` 中的 `jwt_aud` 行；未设置时不检查受众。部署后检查 `semantius.status()` 和角色授权。

### 升级与恢复

安装新版扩展文件后，执行 `ALTER EXTENSION pg_semantius UPDATE`，再在事务块外调用迁移过程。扩展更新替换安装器，迁移调用才改变应用模式。`public._versions` 记录迁移校验值。更新或不兼容的一次性迁移历史可能在应用任何文件前触发拒绝，不应通过修改账本绕过保护。

备份应涵盖普通应用对象和角色，不能只依赖扩展成员关系。恢复到空的 UTF8 数据库，使用相同或更新的扩展文件及兼容的所有者角色；新集群需另行恢复登录凭据。保留数据库设置 `timezone = 'UTC'` 和 `intervalstyle = 'iso_8601'`，不包含创建数据库操作的普通转储不会自动携带它们。

`DROP EXTENSION pg_semantius` 移除安装器，但保留应用数据和集群角色。纯 SQL 安装器无需共享预加载或重启。使用会删除数据的完整卸载步骤前，应核对上游迁移与所有权边界。
