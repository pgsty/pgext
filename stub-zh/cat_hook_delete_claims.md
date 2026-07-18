## 用法

来源：

- [database.dev 软件包页面](https://database.dev/joel/cat_hook_delete_claims)
- [Supabase 官方自定义访问令牌 Hook 文档](https://supabase.com/docs/guides/auth/auth-hooks/custom-access-token-hook)
- [dbdev 官方软件包安装文档](https://github.com/supabase/dbdev/blob/master/website/content/docs/install-a-package.mdx)

cat_hook_delete_claims 0.0.1 是 database.dev 上用于 Supabase 自定义访问令牌 Hook 的可信语言扩展软件包。它安装一个 PL/pgSQL 函数，返回删除可选 user_metadata 成员后的传入 claims。它只修改新签发 JWT 里的声明；不会删除 Auth 用户存储的元数据，也不会撤销已签发令牌。

### 安装与测试

dbdev 依赖 pg_tle，并为选定的软件包版本生成迁移。该软件包的 SQL 会显式向 public 模式中的函数授权，因此应针对 public 模式生成迁移，并在应用前审查：

```bash
dbdev add -o ./migrations -s public -v 0.0.1 package -n joel@cat_hook_delete_claims
```

应用生成的迁移后，直接测试 Hook 契约：

```sql
SELECT public.custom_access_token_delete_user_metadata(
  jsonb_build_object(
    'claims',
    jsonb_build_object(
      'sub', '00000000-0000-0000-0000-000000000001',
      'role', 'authenticated',
      'user_metadata', jsonb_build_object('display_name', 'Ada')
    )
  )
);
```

只有在结果仍保留真实 Auth 流程所用的每一个必需声明后，才应将这个带模式限定的函数配置为 Supabase 自定义访问令牌 Hook。

### 声明与权限边界

Supabase 会在签发或刷新令牌之前运行该 Hook，并拒绝缺少必需声明的输出。软件包只删除 user_metadata，app_metadata 和其他所有成员仍会保留。密码、OAuth、匿名、刷新、MFA 和服务集成可能产生不同的合法声明集，都应测试。格式异常的输入或缺失的 claims 对象超出该软件包的最小实现范围，应在应用自有包装函数中以失败关闭方式处理。

软件包向 supabase_auth_admin 授予执行权，并从 authenticated 和 anon 撤销执行权，但没有撤销 PostgreSQL 默认授给 PUBLIC 的 EXECUTE 权限。因此，仅撤销这两个具名角色并不会限制调用。安装后应撤销 PUBLIC，再仅授权给 Auth Hook 角色。由于授权语句硬编码了 public 模式，应将函数保留在 public；若要改变位置，应在自有迁移中复制并加固这个小函数。

database.dev 软件包依赖 pg_tle。dbdev 文档警告，对包含 TLE 的数据库做逻辑恢复可能失败，并建议使用物理备份。应锁定 0.0.1 版本、保留生成的迁移，并在生产使用前演练备份、恢复、禁用 Hook 和回滚。
