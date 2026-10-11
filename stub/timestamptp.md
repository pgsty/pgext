## Usage

Sources:

- [README.md](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/README.md)
- [timestamptp--1.0.sql](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/timestamptp--1.0.sql)
- [timestamptp.c](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/timestamptp.c)
- [timestamptp.control](https://github.com/brownan/timestamptp/blob/6663d316cef644075624be31602add416f83060c/timestamptp.control)

`timestamptp` stores an absolute timestamp together with its original UTC offset and prints using that offset. It preserves an offset, not a named zone with future daylight-saving rules.

### Core Workflow

```sql
CREATE EXTENSION timestamptp;
SELECT '2026-10-10 09:00:00+08'::timestamptp;
SELECT '2026-10-10 09:00:00+08'::timestamptp + interval '1 hour';
```

### Operational Boundaries

`make_timestamptp(timestamptz, integer)` and `make_timestamptp(timestamptz, text)` provide constructors. Casting to `timestamptz` discards the saved offset; casting back uses the session time zone. Equality compares instants while ignoring offsets, and interval addition preserves the saved offset. The type is relocatable and uses a native library. Upstream does not publish a PostgreSQL-major support matrix or a clear license declaration. Installation uses the control-file superuser default.
