## Usage

Sources:

- [Archived upstream repository](https://github.com/RhodiumToad/uuid-freebsd)
- [Upstream README](https://github.com/RhodiumToad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/README.md)
- [Version 2.0 SQL definitions](https://github.com/RhodiumToad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/uuid-freebsd--2.0.sql)
- [Upstream build metadata](https://github.com/RhodiumToad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/Makefile)
- [Current PostgreSQL UUID functions](https://www.postgresql.org/docs/current/functions-uuid.html)

`uuid-freebsd` is a FreeBSD-specific UUID generator backed by the operating system's UUID implementation. Because the extension name contains a hyphen, quote it in SQL:

```sql
CREATE EXTENSION "uuid-freebsd";

SELECT uuid_generate_v4();
SELECT uuid_generate_v5(uuid_ns_url(), 'https://example.com/resource/42');
```

Version `2.0` provides the nil UUID, DNS/URL/OID/X.500 namespace constants, random UUID v4, time-based v1 and v1mc, and name-based v3 and v5 functions. It requires FreeBSD and links against `libmd`; it is not a portable implementation for Linux or other platforms.

### Obsolete status

The owner describes the project as obsolete and archived the repository in 2023. The last source change predates modern PostgreSQL releases. Prefer the UUID functions supplied by your current PostgreSQL version, or a maintained `uuid-ossp` package when its additional algorithms are required. Keep `uuid-freebsd` only for compatible legacy FreeBSD installations.
