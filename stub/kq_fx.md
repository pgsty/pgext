## Usage

Sources:

- [Official README](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/README.md)
- [Official Rust implementation](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/src/lib.rs)
- [Official extension control file](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/kq_fx.control)

`kq_fx` loads currency exchange-rate histories into PostgreSQL shared memory for fast date-based lookup. Version `1.0.1` requires preload and restart, and its fixed shared cache is cluster-wide rather than isolated per database.

### Core Workflow

Provide the expected source tables, then preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'kq_fx'
```

```sql
CREATE EXTENSION kq_fx;

SELECT kq_fx_check_db();
SELECT kq_fx_populate_cache();
SELECT kq_fx_get_rate(1, 2, DATE '2026-01-31');
SELECT kq_fx_get_rate_xuid('USD', 'EUR', DATE '2026-01-31');
```

The default queries read `plan.currency(id,xuid)` and `plan.fx_rate(currency_id,to_currency_id,date,rate)`. Superuser-settable GUCs `kq.currency.q1_validation`, `kq.currency.q2_get_currencies_xuid`, and `kq.currency.q3_get_currency_entries` can replace those queries.

### Cache Semantics and Limits

`kq_fx_invalidate_cache()` clears the cache, `kq_fx_populate_cache()` reloads it, and `kq_fx_display_cache()` reports entries. Lookup also populates lazily. Equal currencies return 1.0; a date before the first observation returns NULL; dates between observations use the latest rate on or before the date; dates after the final observation use that final rate.

The reviewed implementation fixes capacity at 64 currencies, 1024 currency pairs, and 512 dated entries per pair, with currency XUIDs limited to 16 bytes. Its default entry query keeps only the newest 512 rows for each pair.

Shared memory is not keyed by database. The first session that populates the cache determines the data visible to sessions in other databases, so do not use different exchange-rate data sets within one PostgreSQL cluster. Restrict GUC changes and cache-management functions, validate capacity before loading, and explicitly invalidate after source-table changes.
