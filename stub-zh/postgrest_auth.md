## 用法

来源：

- [已复核提交的 postgrest_auth README](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/README.md)
- [已复核提交的 postgrest_auth.control](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth.control)
- [版本 0.1 的安装 SQL](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth--0.1.sql)
- [上游 Makefile](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/Makefile)

`postgrest_auth` 是一个尚未完成的扩展骨架，并非可用的 PostgREST 认证组件。已复核版本 `0.1` 的安装 SQL 为空，不会创建任何函数、表、角色、策略、钩子或认证 API。

### 核心流程

这里没有可记录的功能工作流。如果本地有修复后或手工安装并登记的副本，下面的查询只能报告目录登记状态，不能测试认证：

```sql
SELECT extname, extversion
FROM pg_extension
WHERE extname = 'postgrest_auth';
```

### 对象索引

已复核的安装脚本没有定义任何 SQL 对象。具体来说，它不提供登录函数、JWT 校验器、角色切换辅助函数、凭据表或行级安全策略。

### 运维边界

- README 写明“仍在开发中”，没有记录 API、配置或启用顺序。
- Makefile 使用 `postgrest-auth.control` 与 `postgrest-auth--0.1.sql`，但仓库跟踪的文件是 `postgrest_auth.control` 与 `postgrest_auth--0.1.sql`。因此按已复核状态，常规 PGXS control 与数据文件名和仓库内容并不匹配。
- 不要根据项目名称推断 PostgREST 安全模型。应用仍需独立审查认证、角色、权限与行级安全设计。
