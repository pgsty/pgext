## 用法

来源：

- [官方 database.dev 软件包页面](https://database.dev/mansueli/rls_helpers)

`rls_helpers` 提供一组小型 PL/pgSQL 过程，用于在测试行级安全策略时模拟 Supabase 风格的已认证或匿名身份。它适合开发与测试会话；切换会话身份会影响同一连接上随后执行的所有受 RLS 保护的查询。

### 核心流程

官方软件包流程使用 `dbdev` CLI 生成带版本的迁移。将迁移应用到目标数据库前应先审阅其内容：

```bash
dbdev add -o ./migrations -s extensions -v 1.0.0 package -n "mansueli@rls_helpers"
```

在显式事务中运行 RLS 测试：选择模拟身份、执行应用查询，并在结束时恢复会话：

```sql
BEGIN;

CALL auth.login_as_user('rodrigo@contoso.com');
SELECT * FROM private.user_documents;

CALL auth.logout();
ROLLBACK;
```

要测试未认证请求对应的策略，可使用匿名辅助过程：

```sql
BEGIN;
CALL auth.login_as_anon();
SELECT * FROM public.visible_documents;
CALL auth.logout();
ROLLBACK;
```

### 用户接口过程

- `auth.login_as_user(text)` 按电子邮件地址选择身份，供后续策略检查使用。
- `auth.login_as_anon()` 将会话切换到匿名测试上下文。
- `auth.logout()` 清除模拟登录上下文。

### 运维说明

上游页面将这些辅助过程定位为 RLS 测试或模拟工具，而不是应用认证系统。应在隔离连接上测试，用事务包住整个过程，并在连接返回连接池前调用 `auth.logout()` 或回滚。软件包页面记录的版本为 `1.0.0`；需要检查生成的迁移，并针对待测 Supabase/PostgreSQL 环境核对其中的角色、模式与权限假设。
