## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/ketteq-neon/kq_cx/blob/585f574defae0a5aa15fa746814e9adc3719fa5c/README.md)
- [Extension control file](https://github.com/ketteq-neon/kq_cx/blob/585f574defae0a5aa15fa746814e9adc3719fa5c/kq_cx.control)
- [Cargo feature matrix](https://github.com/ketteq-neon/kq_cx/blob/585f574defae0a5aa15fa746814e9adc3719fa5c/Cargo.toml)

`kq_cx` keeps application-specific calendar slices in PostgreSQL shared memory and performs date arithmetic against that cache. Add `kq_cx` to `shared_preload_libraries`, restart PostgreSQL, and only then install it in the schema containing the required ketteQ calendar tables.

```sql
CREATE EXTENSION kq_cx;
SELECT kq_invalidate_calendar_cache();
SELECT kq_add_days('2008-01-15', 1, 'quarter');
```

`kq_add_days_by_id` selects a calendar by numeric slice-type ID, while `kq_add_days` selects it by name. `kq_invalidate_calendar_cache` clears and reloads the shared cache after the underlying tables change.

The README documents PostgreSQL 15–17; the current Cargo feature matrix also defines PostgreSQL 18. The extension is not standalone: its source expects a specific application schema and populated calendar data, so validate those table contracts before deployment. Cache invalidation affects shared state and should be coordinated with application writes.
