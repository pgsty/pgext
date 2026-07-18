## Usage

Sources:

- [Pinned current upstream README](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/README.md)
- [Pinned configuration implementation](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/src/guc.rs)
- [Pinned executor-hook implementation](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/src/hooks.rs)
- [Pinned Cargo metadata](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/Cargo.toml)
- [Pinned extension control file](https://github.com/doctolib/pg_no_seqscan/blob/1f7e7d5e50fdd5cb33e1b3a1ebf698d9f6ada804/pg_no_seqscan.control)

pg_no_seqscan version 0.1.1 is a development and CI guardrail. Its executor hook walks query plans, finds sequential-scan nodes in selected databases, schemas, or tables, and either raises an error or emits a notice. Upstream explicitly says it is not for production.

### Preload and exercise the guardrail

```conf
shared_preload_libraries = 'pg_no_seqscan'
enable_seqscan = off
jit_above_cost = 800000000000
pg_no_seqscan.check_schemas = 'public'
pg_no_seqscan.level = 'Error'
```

```sql
CREATE EXTENSION pg_no_seqscan;

CREATE TABLE no_seqscan_demo AS
SELECT generate_series(1, 10000) AS id;

SELECT * FROM no_seqscan_demo WHERE id = 123;
```

After adding a suitable index, the example can use an index plan and pass. Version 0.1.1 fixes parallel EXPLAIN behavior and declares pgrx builds for PostgreSQL 14 through 18.

### Guardrail, not enforcement

A sequential scan can be the correct plan for small tables or broad queries, while an index scan can still be slow. The extension also forces enable_seqscan off and recommends a very high JIT threshold, so test plans differ deliberately from normal production planning. Errors can block migrations, maintenance, analytics, catalog tools, and legitimate full-table reads.

Any user can set pg_no_seqscan.level to Off because that setting is USERSET, and a query containing the documented skip comment bypasses detection. Therefore this cannot enforce a security or performance policy against an untrusted user. Scope settings are superuser-only, but table matching is name-list based and must be tested with schemas, partitions, CTEs, subqueries, views, and utility commands.

Error and notice messages include the query text and rendered plan, which can put literals into client output or server logs. Use synthetic CI data, restrict log access, preload only on dedicated development instances, measure hook overhead on complex plans, and keep ordinary EXPLAIN review and production observability as the authoritative process.
