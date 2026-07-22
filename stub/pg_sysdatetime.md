## Usage

Sources:

- [Official README](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/README.md)
- [Version 1.0 SQL functions](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/pg_sysdatetime--1.0.sql)
- [Extension implementation](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/pg_sysdatetime.c)

`pg_sysdatetime` provides SQL Server-style current-time functions for PostgreSQL 9.4 and older, primarily to improve timestamp precision on Windows. PostgreSQL 9.5 added native use of the precise Windows clock where available, so modern PostgreSQL installations do not need this deprecated extension.

### Functions

```sql
CREATE EXTENSION pg_sysdatetime;

SELECT sysutcdatetime(),
       sysdatetime(),
       sysdatetimeoffset();
```

`sysutcdatetime()` returns UTC as `timestamp without time zone`. `sysdatetime()` returns local wall-clock time, also without a time-zone marker, using the session `TimeZone`. `sysdatetimeoffset()` returns the current instant as `timestamp with time zone`. Do not compare the first two values without accounting for their different implied zones.

### Windows Timer Setting

On old Windows systems, a superuser can request a finer system timer interval:

```conf
pg_sysdatetime.adjust_timer_resolution = on
```

The setting can be reloaded and can also be changed by a superuser per session. It affects the operating system's timer behavior and power use, so measure the system-wide effect before enabling it.

The extension does not call `GetSystemTimePreciseAsFileTime`. Before Windows 8 and Windows Server 2012, the underlying clock can still advance only once per millisecond; extra displayed digits therefore do not imply extra accuracy. For migration, replace these names with small SQL wrappers around PostgreSQL's native clock functions, and verify time-zone and volatility semantics expected by the application.
