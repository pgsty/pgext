## Usage

Sources:

- [docker/raodb/README](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/README)
- [docker/raodb/ext/rdb_catalog/c.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/c.sql)
- [docker/raodb/ext/rdb_catalog/Makefile](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/Makefile)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.0.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.0.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.0--1.1.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.0--1.1.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.1--1.2.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.1--1.2.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.2--1.3.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.2--1.3.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.3--1.4.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.3--1.4.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.4--1.5.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.4--1.5.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.5--1.6.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.5--1.6.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.6--1.7.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.6--1.7.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog.control](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog.control)

`rdb_catalog` 1.7 is a SQL extension from the RAODB development tree. It installs catalog wrappers and diagnostic views in `public`; it does not provide a monitoring collector or automatic tuning.

### Core Workflow

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_qualstats;
CREATE EXTENSION rdb_catalog;
SELECT * FROM public.v$databases;
SELECT * FROM public.ses$blocked;
SELECT * FROM public.vac$needed;
```

### Operational Boundaries

Install as an administrator. Although the control file has an empty dependency list, the 1.4→1.6 SQL scripts directly query `pg_qualstats` and `pg_stat_statements`. Prepare those extensions in the same database and configure their collection/preload requirements before loading the complete 1.7 chain. This extension has no shared library of its own. PostgreSQL-major compatibility and a redistribution license are not stated in the inspected sources.

The `v$` views cover objects, privileges, partitions and statistics; `ses$` views expose sessions, blocking and locks; `stx$` views identify statistics conditions; `ind$` views summarize index and predicate observations; `vac$` views cover vacuum activity and freeze risk. Some catalog sources are restricted to privileged users. Grant access selectively: view results can expose database structure, session queries and user-mapping options.

Treat diagnostic names as heuristics. For example, `ind$safe_to_drop` uses current usage counters and excludes primary-key names, but does not prove that an index is safe to remove. Counter resets, unobserved workloads and constraints must be reviewed separately. The scripts perform no automatic index removal.

The default version is assembled from the 1.0 install file plus all seven consecutive upgrade scripts; install the complete SQL set. Inspect it against the target server and dependency versions before use. Preserve any changes to the bundled wait-event lookup tables in backups; this is source-only catalog intake, not a tested Pigsty package.
