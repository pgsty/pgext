## 用法

来源：

- [database.dev 上的 basejump_core 2.0.1](https://database.dev/basejump/basejump_core)
- [Basejump 产品与 RPC 概览](https://usebasejump.com/)
- [已复核 commit 的账户 API 迁移](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414161947_basejump-accounts.sql)

`basejump_core` 是 Basejump 的 Supabase 账户层数据库包。2.0.1 版创建个人与团队账户、账户成员和角色、邀请、计费记录、RLS 策略、触发器及公共 RPC 函数。先用 dbdev 把软件包生成到 Supabase 迁移目录，应用迁移后再启用扩展：

```bash
dbdev add -o ./migrations -s extensions -v 2.0.1 package -n "basejump@basejump_core"
```

```sql
CREATE EXTENSION IF NOT EXISTS basejump_core WITH SCHEMA extensions;

SELECT create_account('team-slug', 'Your Team Name');
SELECT get_accounts();
```

公共 RPC 函数使用调用者的 Supabase 身份。例如，`create_account` 为当前已认证用户创建账户，`get_accounts` 返回该用户可见的账户。

### 注意事项

- 这是面向 Supabase 的模式软件包，不是独立认证服务器；它依赖 `auth.users`、`authenticated` 和 `service_role` 等 Supabase 对象与角色。
- 扩展会安装 security-definer 函数、触发器、授权和 RLS 策略。将生成的迁移应用到既有项目之前，应复核配置与权限。
- 账户和计费数据保留在数据库中；应把升级与备份按模式迁移对待，而不是按无状态应用部署对待。
