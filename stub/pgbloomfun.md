## Usage

Sources:

- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/README)
- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/Makefile)
- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/ext/pgbloomfun.control)
- [Official source](https://github.com/saaros/pgbloomfun/blob/89ff830fed3f37fadd1e30ae9b64fc50f0abf60e/ext/pgbloomfun.sql)

`pgbloomfun` 0.1 stores scalable Bloom filters in byte arrays. It supplies functions rather than an index access method. This is legacy source from 2013; compatibility with current PostgreSQL releases has not been established.

### Core Workflow

```sql
CREATE EXTENSION pgbloomfun;
SELECT bloom_check(bloom_add(bloom_init(500, 3, 0.01), 'hello'), 'hello');
```

`bloom_init(initial_capacity, growth_factor, error_rate)` creates a filter. `bloom_add` returns the updated byte array; save that result when maintaining a stored filter. `bloom_check` tests possible membership and `bloom_stats` reports filter statistics.

### Limits

Positive matches can be false positives and need exact verification. Filters do not track table changes automatically. Installation uses C functions and requires a superuser; the extension is relocatable and no preload is specified. Build metadata substitutes control version 0.1 and selects JSON output where supported. Verify restoration and extension compatibility before using old stored filters on a new server.
