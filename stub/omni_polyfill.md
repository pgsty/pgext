


## Usage

> [omni_polyfill: Postgres API polyfills](https://docs.omnigres.org/omni_polyfill/polyfills/)

The `omni_polyfill` extension provides polyfill implementations for PostgreSQL functions that may not be available in older versions.

### Configuration

To use polyfills, set `search_path` to list `omni_polyfill` **before** `pg_catalog`:

```sql
SET search_path TO '$user', public, omni_polyfill, pg_catalog;
```

This ordering is critical for proper function resolution and polyfill precedence.

### Available Polyfills

- **`trim_array`** -- Array trimming functionality
- **UUIDv7** -- UUID version 7 generation and related utilities
