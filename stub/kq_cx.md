## Usage

Sources:

- [Official README](https://github.com/ketteq-neon/kq_cx/blob/main/README.md)
- [Extension control file](https://github.com/ketteq-neon/kq_cx/blob/main/kq_cx.control)
- [SQL API and shared-memory implementation](https://github.com/ketteq-neon/kq_cx/blob/main/src/lib.rs)
- [Calendar arithmetic implementation](https://github.com/ketteq-neon/kq_cx/blob/main/src/math.rs)
- [Supported build features](https://github.com/ketteq-neon/kq_cx/blob/main/Cargo.toml)

`kq_cx` caches business-calendar dates in PostgreSQL shared memory and provides fast working-day arithmetic. It is intended for KetteQ schemas, but its four SQL-query GUCs can adapt validation and loading to another compatible calendar model. It requires preload, a server restart, superuser installation, and careful database-wide ownership of the shared cache.

### Core Workflow

Add the library to `shared_preload_libraries`, restart PostgreSQL, create the extension, then load or inspect the cache:

```conf
shared_preload_libraries = 'kq_cx'
```

```sql
CREATE EXTENSION kq_cx;

SELECT kq_cx_populate_cache();
SELECT * FROM kq_cx_cache_info();
SELECT kq_cx_add_days(DATE '2026-07-22', 5, 42);
SELECT kq_cx_add_days_xuid(DATE '2026-07-22', -3, 'US-WORKDAYS');
```

By default the loader expects `plan.calendar`, `plan.calendar_date`, and `plan.data_date`. Customize the authoritative SQL only through `kq.calendar.q_schema_validation`, `kq.calendar.q1_get_calendar_min_max_id`, `kq.calendar.q2_get_calendars_entry_count`, and `kq.calendar.q3_get_calendar_entries` before populating the cache.

### Main Objects

- `kq_cx_add_days(date, integer, bigint)` uses a numeric calendar identifier.
- `kq_cx_add_days_xuid(date, integer, text)` uses an external identifier.
- `kq_cx_populate_cache()` loads calendar data; `kq_cx_invalidate_cache()` forces a later reload.
- `kq_cx_cache_info()`, `kq_cx_info()`, `kq_cx_display_cache()`, and `kq_cx_display_page_map()` expose diagnostic state.

The current SQL API names above come from the implementation; some README examples use older names. An unknown calendar returns `NULL` with a warning. Arithmetic outside the cached range returns sentinel dates `1970-01-01` or `2199-01-01`, not `NULL` or an error.

### Limits and Caveats

The fixed shared-memory layout supports at most 64 calendars, 8192 entries per calendar, 512 page-map entries, and 32 bytes for an external calendar identifier. Validate source data before loading it.

The cache and its filled flag are postmaster-wide while the loading queries run through one database connection. In a multi-database cluster, designate one compatible database as the cache authority and coordinate invalidation. The arithmetic functions are declared immutable even though administrators can reload shared state, so do not place them in persistent expression indexes or generated columns whose correctness depends on an unchanging calendar.
