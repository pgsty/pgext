## 用法

来源：

- [bizniff@account 注册表页面](https://database.dev/bizniff/account)
- [database.dev 包安装指南](https://database.dev/docs/install-a-package)

`account` 是一个面向 Supabase 的纯 SQL Trusted Language Extension，为团队、权限、邀请和账单提供支持。注册表中的 `bizniff@account` `0.1.2` 会创建固定的 `account` schema，其中包含团队与成员表、配置、行级安全策略以及配套函数和触发器。

### 生成并审查迁移

先按 database.dev 要求安装 `pg_tle`，再用官方 `dbdev` 客户端生成迁移；应用前应审查生成内容。

```shell
dbdev add -o ./migrations -s extensions -v 0.1.2 package -n "bizniff@account"
```

应用生成的迁移后，已发布 SQL 会提供以下配置辅助函数：

```sql
SELECT account.get_config();
SELECT account.is_set('enable_account_teams');
```

该包假定运行环境与 Supabase 兼容，包括 `auth.users`、`auth.uid()`、`anon`、`authenticated`、`service_role` 角色以及 `extensions.uuid_generate_v4()`。它还会在 `auth.users` 上安装触发器，并对其对象授予权限。在普通 PostgreSQL 数据库中使用前，必须先处理这些依赖，并审查其 RLS 与权限模型。
