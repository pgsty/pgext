## Usage

Sources:

- [Official README](https://github.com/rhodiumtoad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/README.md)
- [Extension SQL for 2.0](https://github.com/rhodiumtoad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/uuid-freebsd--2.0.sql)
- [FreeBSD UUID implementation](https://github.com/rhodiumtoad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/uuid-freebsd.c)

`uuid-freebsd` 2.0 provides UUID generation functions for PostgreSQL on FreeBSD. It supplies random, time-based, and namespace-based UUIDs plus the standard namespace constants; it is a platform-specific alternative for systems where this package is built against FreeBSD facilities.

### Core Workflow

```sql
CREATE EXTENSION "uuid-freebsd";

SELECT uuid_generate_v4();
SELECT uuid_generate_v5(uuid_ns_url(), 'https://example.com/orders/42');
```

Version 4 generates a new random UUID. Versions 3 and 5 deterministically combine a namespace UUID and name, making them suitable when the same logical name must always map to the same identifier.

### Functions

- `uuid_generate_v1()` creates a time-based UUID using the host node identifier.
- `uuid_generate_v1mc()` creates a time-based UUID with a random multicast node identifier.
- `uuid_generate_v3(namespace, name)` creates an MD5 namespace UUID.
- `uuid_generate_v4()` creates a random UUID.
- `uuid_generate_v5(namespace, name)` creates a SHA-1 namespace UUID.
- `uuid_nil()`, `uuid_ns_dns()`, `uuid_ns_url()`, `uuid_ns_oid()`, and `uuid_ns_x500()` return standard constants.

The extension is relocatable and needs neither preloading nor restart. Time-based version 1 values may reveal stable host/time characteristics; prefer version 4 when no deterministic or time-based identity is required, and use `uuid_generate_v1mc()` when version 1 ordering is needed without the real node identifier. The project is unmaintained, so validate the package on the exact FreeBSD and PostgreSQL versions in use.
