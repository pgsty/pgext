## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/pg_ratelimit.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/pg_ratelimit--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/pg_ratelimit.c)

`pg_ratelimit` `1.0` provides shared-memory rate limits and session values. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_ratelimit;
SELECT ratelimit_check('demo', 100, 1, '1 second');
SELECT ratelimit_remaining('demo');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

Preload pg_ratelimit and restart before rate-limit operations. pg_ratelimit.max_entries bounds shared entries; token-bucket and sliding-window helpers track quotas, while session_set/get/touch manage expiring session records. Shared rate-limit state is lost on restart and is not transactional application data. Namespace keys and restrict reset/session operations; shared keys are not a tenant-security boundary.

This module corresponds to the former upstream name `alohadb_ratelimit`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
