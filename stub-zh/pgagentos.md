## 用法

来源：

- [pgagentos.control](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/pgagentos.control)
- [README.md](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/README.md)
- [Makefile](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/Makefile)
- [sql/schemas/00_extensions.sql](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/sql/schemas/00_extensions.sql)
- [sql/schemas/02_aos_auth.sql](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/sql/schemas/02_aos_auth.sql)
- [sql/rls/rls_policies.sql](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/sql/rls/rls_policies.sql)

`pgagentos` 安装用于模型／事件记录、租户、角色设定、技能、会话、记忆和检索的关系型基础设施。它是纯 SQL 扩展，不是自动执行模型的运行时。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pgagentos;
INSERT INTO aos_auth.tenant(name) VALUES ('example_team') RETURNING tenant_id;
SELECT * FROM aos_core.job LIMIT 10;
```

### 运行边界

需要 PostgreSQL 14 及以上、`vector`、`pgcrypto` 和 PL/pgSQL。由超级用户安装，无需预加载。SQL 函数用于记录工作，模型执行及凭据由外部应用负责。行策略依赖 `aos_auth.current_tenant()` 和会话租户设置；能修改该设置的调用方不构成独立的身份认证边界。向租户开放数据库会话前，应审查授权、对象所有者和 RLS 绕过行为。仅创建模式不能证明多租户隔离安全。
