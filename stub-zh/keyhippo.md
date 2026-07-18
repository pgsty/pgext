## 用法

来源：

- [项目 README](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/README.md)
- [扩展专用 README](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/extension/README.md)
- [扩展 control 文件](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/extension/keyhippo.control)
- [1.2.5 版安装 SQL](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/extension/keyhippo--1.2.5.sql)

`keyhippo` 1.2.5 在保留行级安全策略的同时，为 Supabase/PostgREST 数据库增加 API Key 认证和基于角色的授权。它只保存密钥哈希，支持作用域、过期、撤销、轮换和声明，并提供用户组、角色与内置权限类型。

### 签发并使用密钥

```sql
CREATE EXTENSION keyhippo;

SELECT * FROM keyhippo.create_api_key('Primary API Key', 'default');
SELECT keyhippo.key_data();
SELECT keyhippo.authorize('manage_api_keys');
```

主要 API 包括 `keyhippo.create_api_key(text,text)`、`keyhippo.verify_api_key(text)`、`keyhippo.revoke_api_key(uuid)`、`keyhippo.rotate_api_key(uuid)` 和 `keyhippo.authorize(keyhippo.app_permission)`。API Key 通过 PostgREST 的 `x-api-key` 请求头传入；创建调用会返回明文密钥，调用方必须在签发时妥善保存。

### Supabase 依赖与安装副作用

这不是一个通用 PostgreSQL 认证扩展。其 SQL 依赖 `auth.users`、`auth.uid()`，以及 `anon`、`authenticated`、`authenticator`、`service_role` 角色。安装时会在缺失时创建 `pgcrypto`、`pg_net`、`pg_cron`，并创建四个固定 schema、RLS 策略、授权、一个 `auth.users` 触发器和默认 RBAC 数据。它还会修改 `authenticator` 的 `pgrst.db_pre_request` 设置，并通知 PostgREST 重新加载。

应用前必须审查安装 SQL。首次初始化会启用安装通知，并通过 `pg_net` 向 `https://app.keyhippo.com/api/ingest` 发送 HTTP 请求；审计 HTTP 日志默认关闭，但端点仍会写入配置。control 文件声称可重定位，而 SQL 实际创建固定 schema。所有 security-definer 函数、授权、模拟登录过程、外发 HTTP 行为和 PostgREST 角色变更都应按安全敏感迁移内容处理。
