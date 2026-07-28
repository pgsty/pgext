## Usage

Sources:

- [Official upstream README](https://github.com/evoludigit/pggit/blob/6d1feef9b9fa215802b9243640e8183586dcdf6a/README.md)
- [Official extension control file (pggit.control)](https://github.com/evoludigit/pggit/blob/6d1feef9b9fa215802b9243640e8183586dcdf6a/pggit.control)
- [Official extension SQL (pggit--0.1.3.sql)](https://github.com/evoludigit/pggit/blob/6d1feef9b9fa215802b9243640e8183586dcdf6a/pggit--0.1.3.sql)

`pggit` — **Git-like version control for PostgreSQL schemas.** Branch, merge, diff, and revert database schemas like you do with code. Use it when administering or automating the database behavior described above. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pggit;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pggit.abort_merge(p_merge_id uuid, p_reason text DEFAULT 'User aborted')` is an extension function and returns `void`.
- `pggit.add_dependency(p_dependent_name TEXT, p_depends_on_name TEXT, p_dependency_type TEXT DEFAULT 'generic')` is an extension function and returns `VOID`.
- `pggit.analyze_branch_for_pruning(p_branch_name TEXT)` is an extension function and returns `TABLE`.
- `pggit.analyze_cqrs_dependencies(command_schema text DEFAULT 'command', query_schema text DEFAULT 'query')` is an extension function and returns `TABLE`.
- `pggit.analyze_migration_intent(p_migration_content TEXT)` is an extension function and returns `TABLE`.
- `pggit.analyze_migration_size_impact(p_migration_content TEXT)` is an extension function and returns `TABLE`.
- `pggit.analyze_migration_with_ai(p_migration_id TEXT, p_migration_content TEXT, p_source_tool TEXT DEFAULT 'unknown')` is an extension function and returns `TABLE`.
- `pggit.analyze_migration_with_ai_enhanced(p_migration_id TEXT, p_migration_content TEXT, p_source_tool TEXT DEFAULT 'unknown')` is an extension function and returns `TABLE`.
- `pggit.analyze_query_performance()` is an extension function and returns `jsonb`.
- `pggit.analyze_schema_change_frequency(p_branch_name text, p_days integer DEFAULT 30)` is an extension function and returns `jsonb`.
- `pggit.apply_data_merge(p_merge_id UUID, p_source_branch TEXT, p_target_branch TEXT, p_resolution_strategy TEXT)` is an extension function and returns `VOID`.
- `pggit.apply_migration(p_version TEXT)` is an extension function and returns `VOID`.
- `pggit.apply_pruning_recommendation(p_recommendation_id INTEGER)` is an extension function and returns `TEXT`.
- `pggit.apply_retention_policy(p_policy JSONB DEFAULT '{"full_days": 30, "incremental_days": 7}')` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.3`.
- Install the confirmed extension dependencies first: `pgcrypto`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
