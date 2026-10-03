## Usage

Sources:

- [pgagentos.control](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/pgagentos.control)
- [README.md](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/README.md)
- [Makefile](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/Makefile)
- [sql/schemas/00_extensions.sql](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/sql/schemas/00_extensions.sql)
- [sql/schemas/02_aos_auth.sql](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/sql/schemas/02_aos_auth.sql)
- [sql/rls/rls_policies.sql](https://github.com/pgAgentOS/pgAgnetOS/blob/82d6ae60c1efdaf9c02d158fa25ed7198e5a6b2b/sql/rls/rls_policies.sql)

`pgagentos` installs relational agent infrastructure: model/event records, tenants, personas, skills, conversations, memory and retrieval schemas. It is a SQL-only extension, not an autonomous model runtime.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pgagentos;
INSERT INTO aos_auth.tenant(name) VALUES ('example_team') RETURNING tenant_id;
SELECT * FROM aos_core.job LIMIT 10;
```

### Operational Boundaries

Requires PostgreSQL 14 or later, `vector`, `pgcrypto` and PL/pgSQL. Install as a superuser; no preload is required. Use the SQL functions to record work, while an external application handles model execution and credentials. Row policies depend on `aos_auth.current_tenant()` and a session tenant setting; callers able to change that setting are not an independent authentication boundary. Review grants, ownership and RLS bypass before exposing database sessions to tenants. Schema creation alone does not establish secure multi-tenant operation.
