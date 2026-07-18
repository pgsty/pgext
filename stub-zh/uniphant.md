## 用法

来源：

- [固定提交的上游 README](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/README.md)
- [1.5 版应用 SQL](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/uniphant--1.5.sql)
- [固定提交的扩展 control 文件](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/uniphant.control)
- [固定提交的 WebAuthn 验证函数](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/FUNCTIONS/api/verify_assertion.sql)

uniphant 1.5 版是一个结合 PostgreSQL、PostgREST、WebAuthn、nginx 和浏览器客户端的全栈演示项目。它在固定的 public/api 模式中实现用户、凭据、角色、资源、权限、访问令牌、WebAuthn 注册/登录，以及生成 OpenAPI 元数据的功能。

### 仅用于演示的安装

预期的数据库角色必须先存在，webauthn 依赖通过 CASCADE 安装：

```sql
CREATE ROLE web_anon NOLOGIN;
CREATE ROLE postgrest NOLOGIN;

CREATE EXTENSION uniphant
WITH SCHEMA public
CASCADE;

SELECT api.openapi_swagger();
```

完整演示还需要 PostgREST、可信反向代理、浏览器 WebAuthn 支持、非本地使用时的 TLS，以及相关 PostgreSQL 加密依赖。README 的部署方案面向 Ubuntu 20.04、PostgreSQL 13、PostgREST 7.0.1 和旧版请求头约定。

### 安全敏感副作用

安装会向 web_anon 授予 api/webauthn 的 USAGE 权限，把 web_anon 角色授予 postgrest，并向 web_anon 授予 api 与 public 中全部现有表的 SELECT 权限。它还会创建大量 SECURITY DEFINER 函数。第一个成功注册的非匿名用户会自动获得管理员角色。这些是有意设置的演示捷径，在现有数据库中会产生重大影响。

不要把 uniphant 安装到通用或生产环境的 public 模式。现代部署必须审查每项授权和定义者权限函数、移除首用户引导逻辑、验证固定搜索路径、适配 PostgREST 请求设置、强制使用可信代理头与 TLS、保护访问令牌存储，并重新测试当前 WebAuthn 语义。评估应使用空的一次性数据库，并对比安装前后生成的对象与权限清单。
