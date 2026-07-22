## 用法

来源：

- [官方 database.dev 软件包页面](https://database.dev/garyaustin/custom_roles)
- [官方上游 README](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/README.md)
- [官方 custom_roles TLE README](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-roles/README.md)
- [官方 custom_roles SQL](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-roles/custom_roles--0.0.3.sql)

`custom_roles` 是面向 Supabase、以数据表为后端的角色系统，通过 database.dev 可信语言扩展分发。它把应用角色与 PostgreSQL 角色分开保存，并提供面向行级安全策略的辅助函数。它依赖 Supabase 的 `auth.users` 表和 `auth.uid()` 函数。

### 核心流程

为 registry 中的 `0.0.3` 版生成并审核迁移；当前 CLI 使用下方 registry 别名。生成的迁移会通过 `pg_tle` 安装 TLE，并创建带命名空间的扩展。

```bash
dbdev add -o ./migrations -s public -v 0.0.3 package -n "garyaustin@custom_roles"
```

先添加允许的角色名，再为已认证用户 UUID 分配一个或多个角色。

```sql
INSERT INTO public.custom_role_names (role_name)
VALUES ('Teacher'), ('Staff');

INSERT INTO public.custom_user_roles (user_id, role)
VALUES ('00000000-0000-0000-0000-000000000001', 'Teacher');
```

按照上游性能建议，在 RLS 策略中通过标量子查询调用角色辅助函数。

```sql
CREATE POLICY teacher_read
ON public.lesson
FOR SELECT
TO authenticated
USING ((SELECT public.user_has_role('Teacher')));
```

可选的 JWT 同步触发器安装后默认禁用。只有确实需要把角色复制到 `auth.users.raw_app_meta_data` 时才启用；JWT 刷新前，客户端看不到声明变化。

```sql
ALTER TABLE public.custom_user_roles
ENABLE TRIGGER on_custom_role_change;
```

### 重要对象

- `custom_role_names`：允许的角色名，预置 `RoleAdmin`。
- `custom_user_roles`：连接 `auth.users` 和角色名表的 `(user_id, role)` 分配关系。
- `user_has_role(text)`：检查当前 `auth.uid()` 是否具有单个角色。
- `user_role_in(text[])`：检查当前用户是否具有任一请求角色。
- `user_roles_match(text[])`：检查当前用户是否具有全部请求角色。
- `get_user_roles()`：返回当前用户的角色数组。
- `custom_roles_update_to_app_metadata()`：把 `user_roles` 数组写入 Supabase 应用元数据的触发器函数。

### 安全与限制

两个表都启用了 RLS。registry `0.0.3` 版允许已认证用户读取自己的分配记录，并允许具有 `RoleAdmin` 的用户管理分配；PostgreSQL 与 `service_role` 是否绕过策略或拥有更广权限，取决于 Supabase 部署。使用前应审计授权和每条策略。

发布的 `0.0.3` SQL 把函数 `search_path` 固定为 `public`，因此即使旧说明声称可以安装到其他模式，也不能安全地把它视为可迁移扩展。其禁用触发器名是 `on_custom_role_change`，并非 `custom_role_change`。registry SQL 的删除分支虽计算了受影响 UUID，后续却读取 `new.user_id`；启用触发器后删除角色可能报错或错误更新元数据。启用 JWT 同步前，应在受控迁移中复核并修补该函数。大型角色数组还会增大 JWT；上游也建议对拥有超过 1,000 个角色的用户实测 `get_user_roles()` 行为。
