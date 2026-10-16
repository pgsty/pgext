## 用法

来源：

- [1.0.1 upgrade](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/src/upgrades/pgmorbac--1.0.0--1.0.1.sql)
- [1.0.1 changes](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/CHANGELOG.md)
- [pgmorbac.control](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/pgmorbac.control)
- [README.md](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/README.md)
- [SECURITY.md](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/SECURITY.md)
- [Makefile](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/Makefile)
- [src/authorization.sql](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/src/authorization.sql)
- [src/rls.sql](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/src/rls.sql)
- [src/system_rls.sql](https://github.com/crudylabs/pgmorbac/blob/11ef1715abdf10ea610dd86368955d7ecd17efdd/src/system_rls.sql)

`pgmorbac` 在固定的 `morbac` 模式中实现按组织授权。角色、活动、视图、许可、禁止和委派构成策略模型，可由应用的 RLS 策略调用。

### 核心用法

```sql
CREATE EXTENSION pgmorbac;
SELECT morbac.refresh_hierarchy_cache();
```

在配置策略和可信会话身份之后：

```sql
SELECT morbac.is_allowed(
  '00000000-0000-0000-0000-000000000001'::uuid,
  NULL::uuid, 'read', 'documents');
```

### 运行边界

由管理员在 PostgreSQL 13 及以上安装，此 SQL 扩展不需要预加载或重启。创建组织和角色、分配应用用户 UUID、编译策略，再将 `morbac.rls_check` 接入对应业务表。仅安装扩展不会自动为业务表开启 RLS。

使用 `morbac.is_allowed` 判断对象访问权限，使用 `morbac.is_allowed_nocache` 绕过授权结果缓存，使用 `morbac.refresh_hierarchy_cache` 刷新层级数据。`morbac.has_permission` 用于判断界面功能是否可用，不能替代对象级授权检查。行的组织为 NULL 表示尚未归属组织的对象，而不是所有组织。

通常，较高优先级的规则生效，相同优先级时禁止优先。但列入 `morbac.system_principals` 的用户会跳过禁止规则检查，仍需存在适用的许可。应同时保护该表、策略表、上下文函数与委派管理。应用必须提供经过认证的 `morbac.user_id` 和组织上下文，用户可控的会话参数不能作为身份证明。应显式授予必要权限，并以实际应用角色验证访问。本文对应所引源码修订的 1.0.1 版本。

### 1.0.1 的授权计算开销

未命中缓存的授权检查会集中解析相关组织与角色组合，避免为其他租户的规则重复计算。权限判定保持 1.0.0 的语义。安装匹配的 SQL 文件后执行 `ALTER EXTENSION pgmorbac UPDATE TO '1.0.1'`，并用实际应用角色验证 RLS 行为。
