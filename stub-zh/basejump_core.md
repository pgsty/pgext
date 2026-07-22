## 用法

来源：

- [basejump_core 2.0.1 package](https://database.dev/basejump/basejump_core)
- [Official Basejump README](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/README.md)
- [Account schema and RPC implementation](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414161947_basejump-accounts.sql)
- [Invitation implementation](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414162100_basejump-invitations.sql)
- [Billing implementation](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414162131_basejump-billing.sql)

`basejump_core` 2.0.1 是 Basejump 的 Supabase 账户层数据库包。它添加个人与团队账户、成员和角色、邀请、计费状态、行级安全策略、触发器以及公共 RPC 函数。它面向 Supabase 数据库，并不是独立的认证或计费服务。

### 安装与启用

使用官方 database.dev 软件包名生成迁移，通过正常的 Supabase 迁移流程应用，再启用生成的扩展：

```bash
dbdev add -o ./migrations -s extensions -v 2.0.1 package -n "basejump@basejump_core"
```

```sql
CREATE EXTENSION IF NOT EXISTS basejump_core WITH SCHEMA extensions;
```

该软件包依赖 `auth.users`、`authenticated`、`service_role` 等 Supabase 身份与角色。SQL 还引用 `auth.uid()`、`extensions.uuid_generate_v4()`、`gen_random_bytes()` 等函数；应用迁移前，应确认 Supabase 项目已经提供这些前置对象。

### 账户工作流

公共函数应通过 Supabase RPC 调用，并使用调用者的已认证身份：

```sql
SELECT create_account('engineering', 'Engineering');
SELECT get_accounts();
SELECT get_account_by_slug('engineering');
SELECT get_account_members(get_account_id('engineering'));
```

所有者可通过 `update_account_user_role`、`remove_account_member`、`create_invitation`、`get_account_invitations`、`lookup_invitation`、`accept_invitation` 等函数管理成员和邀请。计费消费者使用 `get_account_billing_status`；可信计费 webhook 则通过 `service_role_upsert_customer_subscription` 同步状态。

### 数据库对象

- `basejump.accounts` 与 `basejump.account_user`：账户、成员和账户角色。
- `basejump.invitations`：一次性或 24 小时有效的邀请令牌。
- `basejump.billing_customers` 与 `basejump.billing_subscriptions`：计费服务商的客户和订阅状态。
- `basejump.config`：团队账户、计费和计费服务商的功能开关。
- 公共 RPC 函数：面向应用的支持接口；直接访问表会受授权和 RLS 限制。

### 安全与维护

该软件包会安装 RLS 策略、触发器、授权和多个 `SECURITY DEFINER` 函数。应用前应审阅生成的带版本迁移，保留文档定义的 `search_path`，并分别以 `anon`、`authenticated`、`service_role` 身份测试，而不能只用管理员验证。

账户、邀请和计费行都是持久业务状态。应同时备份模式与数据，把版本升级作为数据库迁移进行演练，并避免向不可信客户端暴露 webhook 凭据和 service-role 权限。即使普通读取受 RLS 限制，邀请令牌和计费载荷仍属于敏感数据。
