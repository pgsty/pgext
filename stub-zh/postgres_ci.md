## 用法

来源：

- [锁定提交的扩展控制文件](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/postgres_ci.control)
- [完整 0.3 版安装 SQL](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/postgres_ci--0.3.sql)
- [锁定提交的 PUBLIC 权限授予](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/packages.sql)
- [锁定提交的用户创建函数](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/functions/users/add.sql)
- [锁定提交的密码重置函数](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/functions/password/reset.sql)
- [锁定提交的 GitHub 密钥读取函数](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/functions/project/get_github_secret.sql)

postgres_ci 0.3 是一个已废弃 Postgres-CI 服务的数据库后端。它存储应用用户和会话、项目和 webhook 密钥、Git 提交、构建、构建部件和测试、工作队列与通知。它依赖 PL/pgSQL、pgcrypto 和 pg_trgm，并创建多个硬编码模式。

### 隔离安装审计

只应安装到可丢弃的隔离数据库，并在调用任何服务 API 前检查定义者接口面：

```sql
CREATE EXTENSION postgres_ci CASCADE;

SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_ci';

SELECT n.nspname, p.proname, p.prosecdef
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
WHERE n.nspname IN (
  'auth', 'hook', 'users', 'build', 'project',
  'password', 'notification', 'postgres_ci'
)
ORDER BY n.nspname, p.proname;

REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA
  auth, hook, users, build, project,
  password, notification, postgres_ci
FROM PUBLIC;
```

应用函数实现登录、webhook 摄取、项目变更以及工作进程获取/接受流程。它们是一套不可分割的历史服务 API，不是通用 CI 原语。

### 严重身份验证与权限缺陷

几乎每个服务函数都是 SECURITY DEFINER，且没有固定的安全 search_path。安装过程显式向 PUBLIC 授予所有服务模式的使用权和所有函数的执行权。如果不撤销这些授权，任何数据库角色都能访问特权例程。用户创建函数接受应用级 is_superuser 标志，却不验证调用者；password.reset 可替换其他用户的密码；project.get 和 project.get_github_secret 会返回已存储 webhook 密钥。

密码使用快速加盐 SHA-1 摘要，而不是当代密码 KDF。GitHub webhook 密钥以明文存储。会话保存在不记日志表中，崩溃后可能消失，并使用一小时滑动过期。错误消息会区分未知用户和密码无效，从而可用于枚举账户。

不得向不可信角色或面向网络的应用暴露该扩展。撤销 PUBLIC 只是第一步：每个定义者函数都需要固定的 pg_catalog 优先 search_path、显式调用者授权、最小权限所有权、密钥加密与轮换、当代密码迁移、限流、审计日志和对抗性测试。安全替代方案比改造这份代码更可靠。

扩展不可重定位，并会创建名称宽泛的 auth、hook、users、build、project、password 和 notification 模式。它的最后代码提交在 2016 年，README 没有可用运维文档，目录生命周期为已废弃。只应为读取或导出历史状态而保留它，不得用它启动新 CI 服务。
