## Usage

Sources:

- [Project README](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/README.md)
- [Extension control file](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/yeti.control)
- [Version 1.14.1 upgrade SQL](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/yeti--1.14.0--1.14.1.sql)
- [Shared-memory rate limiter](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/src/f_tbf_rate_check.c)
- [LNP resolver functions](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/src/f_lnp_resolve.c)

`yeti` 1.14.1 provides specialized native helper functions for the YETI softswitch. Its surface includes LNP endpoint and resolution calls, tag operations, randomized regular-expression replacements, template processing, DNS SRV ranking, and a shared-memory token-bucket rate check. It is not a general-purpose telephony abstraction.

### Install and enable the rate limiter

Only `tbf_rate_check()` requires server preloading. Configure the library and restart PostgreSQL before relying on it:

```conf
shared_preload_libraries = 'yeti_pg_ext'
```

```sql
CREATE SCHEMA yeti_ext;
CREATE EXTENSION yeti WITH SCHEMA yeti_ext;

SELECT yeti_ext.tbf_rate_check(1, 1001, 5.0);
```

If the library was not preloaded, `tbf_rate_check()` emits a warning and returns `true`. This is fail-open behavior: missing configuration disables enforcement rather than rejecting traffic. Include a startup probe that verifies preload state before serving requests.

### Volatile state and network boundary

The token buckets live only in one PostgreSQL instance's shared memory, use a fixed 5,000-entry hash, and are lost on restart or failover. They are not coordinated across primaries, replicas, or multiple clusters. Hook chaining, contention, clock behavior, rate semantics, and capacity exhaustion must be load-tested; do not use this mechanism as the sole billing, abuse, or security limit.

LNP functions send phone-number-related requests to configured external endpoints and can return null or an error field on transport failure. Set finite timeouts, restrict endpoint configuration, protect phone-number data in SQL and logs, and define retry and circuit-breaker behavior outside the extension. Pin the YETI application and extension versions together and follow the complete SQL upgrade chain; helper signatures changed in 1.14.1.
