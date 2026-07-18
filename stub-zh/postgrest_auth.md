## 用法

来源：

- [已复核提交的 postgrest_auth README](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/README.md)
- [已复核提交的 postgrest_auth.control](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth.control)
- [版本 0.1 的安装 SQL](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth--0.1.sql)
- [上游 Makefile](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/Makefile)

`postgrest_auth` 是一个尚未完成的扩展骨架，并非可用的 PostgREST 认证实现。已复核版本 `0.1` 的安装 SQL 为空：它不会创建任何函数、表、角色、策略、钩子或认证 API。

因此不存在功能性用法示例。如果本地有修复后或手工安装的副本，下面的查询只能确认扩展登记状态，不能确认认证行为：

```sql
SELECT extname, extversion
FROM pg_extension
WHERE extname = 'postgrest_auth';
```

### 注意事项

- README 明确写着“仍在开发中”，且没有记录 API 或设置流程。
- Makefile 引用了带连字符的 `postgrest-auth.control` 和 `postgrest-auth--0.1.sql` 文件名，但仓库中是下划线文件。按已复核状态，常规 PGXS `DATA` 和控制文件查找与已跟踪文件并不匹配。
- 不要根据项目名称推断它支持 PostgREST JWT、角色切换、用户管理或行级安全集成。已复核源码中没有这些功能。
