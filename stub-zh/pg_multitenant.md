## 用法

来源：

- [官方 README](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/README.md)
- [扩展实现](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/src/pgmt.rs)
- [上游示例](https://github.com/gurkanindibay/pg_multitenant/blob/b24ebb692a61f8ae30da2eb7c455edd3cc9c85ac/scripts/test.sql)

`pg_multitenant` 是一个实验性的共享数据库、共享 Schema 多租户辅助工具。它创建 PostgreSQL 行级安全策略，把指定租户列与当前角色名进行比较，并在 `pgmt.table_tenant_column` 中记录已标记列。

### 核心流程

当前可用的上游示例采用 `user` 策略：每个租户以名称等于租户键文本形式的角色连接或切换身份。

```sql
CREATE EXTENSION pg_multitenant;

CREATE TABLE public.tenant_item (
    tenant_id uuid NOT NULL,
    item_id bigint GENERATED ALWAYS AS IDENTITY,
    payload jsonb,
    PRIMARY KEY (tenant_id, item_id)
);

SET pgmt.tenant_strategy = 'user';
SELECT pgmt.mark_tenant_column(
    'public',
    'tenant_item',
    'tenant_id'
);

CREATE ROLE "11111111-1111-1111-1111-111111111111";
GRANT SELECT, INSERT, UPDATE, DELETE
ON public.tenant_item
TO "11111111-1111-1111-1111-111111111111";

SET ROLE "11111111-1111-1111-1111-111111111111";
SELECT * FROM public.tenant_item;
RESET ROLE;
```

移除策略与元数据条目：

```sql
SELECT pgmt.unmark_tenant_column('public', 'tenant_item');
```

`pgmt.mark_tenant_column(schema_name, table_name, tenant_column)` 启用 RLS、创建名为 `tenant_isolation_policy` 的策略，并记录表 OID 与列。`pgmt.unmark_tenant_column(schema_name, table_name)` 禁用 RLS、删除该策略并删除元数据行。

### 安全与实现限制

已复核版本不能视为完整的租户隔离框架。源码把两个 GUC 都注册为 `pgmt.tenant_strategy`；因此宣传的 `value` 策略及 `pgmt.tenant_value` 路径无法按当前实现安全使用。辅助函数还会把传入标识符未经标识符引用直接插入动态 SQL，使用固定策略名，并且重复标记不是幂等操作。

如果周边设计没有额外约束，PostgreSQL 表所有者、超级用户及具有 `BYPASSRLS` 的角色可以绕过普通 RLS。生成的策略用 `tenant_column::TEXT` 比较 `current_user`，所以角色命名本身就是安全契约的一部分。只应使用可信标识符，测试所有者与特权角色行为，在 `pg_policy` 中检查生成策略，并在生产使用前完成独立安全审查。control 版本为 `1.0`，冻结的 Cargo 清单则报告 `0.1.0`，目标为 PostgreSQL `16`。
