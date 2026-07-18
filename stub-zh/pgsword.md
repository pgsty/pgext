## 用法

来源：

- [扩展 control 文件](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/pgsword.control)
- [空扩展 SQL](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/pgsword--1.0.sql)
- [钩子实现](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/pgsword.c)
- [审核规则](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/rule.c)

`pgsword` 版本 `1.0` 是 PostgreSQL 10 时代的去哪儿 SQL 审核原型。安装 SQL 不创建任何数据库对象，动态库必须预加载才能注册解析和工具语句钩子。全局配置中应保持审核关闭：

```conf
shared_preload_libraries = 'pgsword'
pgsword.enabled = off
```

重启后注册扩展。超级用户只能在隔离的审核会话中启用它并提交候选 DDL：

```sql
CREATE EXTENSION pgsword;
SHOW pgsword.enabled;

SET pgsword.enabled = on;

-- Audit only: this statement is deliberately aborted after checks.
CREATE TABLE audit_candidate (
  id bigserial PRIMARY KEY,
  payload jsonb,
  created_at timestamptz
);
```

### 阻断行为与兼容性

启用 `pgsword.enabled` 后，大多数工具语句最终都会进入扩展的 `finishAudit()` 路径；即使全部规则通过，该路径也会抛出错误，候选 DDL 不会执行。规则会拒绝多种名称和类型，包括非小写标识符、名称不是 `id` 的主键、以 `timestamp` 代替 `timestamptz`，以及以 `json` 代替 `jsonb`。源码和 Makefile 面向 PostgreSQL 10.1 内部 API，且没有 README、许可证或回归测试。不能在普通会话中启用该动态库，现代服务器使用前必须移植并审查。
