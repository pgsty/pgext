## 用法

来源：

- [官方 README](https://github.com/usebasejump/supabase-test-helpers/blob/d90a51f197d49f0a2af155462669cf1dc0218534/README.md)
- [0.0.6 版扩展 SQL](https://github.com/usebasejump/supabase-test-helpers/blob/d90a51f197d49f0a2af155462669cf1dc0218534/supabase_test_helpers--0.0.6.sql)
- [database.dev 包页面](https://database.dev/basejump/supabase_test_helpers)

`basejump-supabase_test_helpers` 是 Basejump Supabase 测试辅助扩展在 database.dev 中的包名。它创建用于测试用户、JWT/角色模拟、行级安全断言和冻结时间的函数。上游项目明确建议只在事务化 pgTAP 测试中启用，不要在生产中长期安装。

### 测试事务流程

通过 dbdev 安装包，然后在最终回滚的事务内创建服务商命名的扩展：

```sql
SELECT dbdev.install('basejump-supabase_test_helpers');

BEGIN;
CREATE EXTENSION "basejump-supabase_test_helpers";

SELECT plan(2);

CREATE TABLE public.test_accounts (id bigint PRIMARY KEY);
ALTER TABLE public.test_accounts ENABLE ROW LEVEL SECURITY;

SELECT tests.rls_enabled('public', 'test_accounts');
SELECT tests.create_supabase_user('test_owner', 'owner@example.test');

SELECT * FROM finish();
ROLLBACK;
```

仓库的原生控制文件与 SQL 使用 `supabase_test_helpers`，database.dev 则暴露带命名空间的 `basejump-supabase_test_helpers` 包。应遵循安装机制报告的扩展名称，不要手工重命名控制文件。

### 重要辅助函数

- `tests.create_supabase_user` 向 `auth.users` 插入合成用户并返回 UUID；`tests.get_supabase_user` 与 `tests.get_supabase_uid` 按易记的测试标识检索该用户。
- `tests.authenticate_as`、`tests.authenticate_as_service_role` 与 `tests.clear_authentication` 修改事务本地角色和 `request.jwt.claims` 设置，以测试 Supabase 授权路径。
- `tests.rls_enabled` 针对单表或模式内所有表返回 pgTAP 断言。
- `tests.freeze_time` 与 `tests.unfreeze_time` 配合 `test_overrides.now` 让依赖时间的测试可重复。

安全定义者函数若固定了自身搜索路径，只有测试临时将 `test_overrides` 放在该函数路径中的 `pg_catalog` 之前时，冻结时间才会生效。只能在测试事务内这样做，并让回滚恢复目录设置。

### 前置条件与安全性

扩展依赖 `pgtap`，并假设存在包含 `auth.users`、`anon`、`authenticated` 与 `service_role` 角色以及其 SQL 所用 UUID 辅助函数的 Supabase 布局。多个辅助函数是安全定义者，并且会刻意模拟高权限应用角色。测试应在隔离数据库或一次性分支中运行，绝不能针对生产身份执行。

仓库包含 0.0.6 安装脚本，但签入的控制文件仍把 0.0.5 设为默认版本。应通过包管理器固定所需版本，并在 `pg_extension` 中确认安装结果。回滚会移除事务内创建的对象，但被测应用代码产生的外部副作用未必具有事务性。
