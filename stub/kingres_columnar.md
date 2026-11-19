## Usage

Sources:

- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_columnar/README.md)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_columnar/kingres_columnar.control)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_columnar/kingres_columnar--0.1.sql)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_columnar` 0.1 is an experimental insert-only columnar table access method. The reviewed source explicitly warns against storing important data: unresolved defects affect serializable transactions and write buffers. It is catalogued as a preview, not validated production storage.

### Create a test table

The source targets PostgreSQL 18/19. Preload and restart every writer, standby and restore node because the library registers a WAL resource manager. Installation requires a superuser.

```ini
shared_preload_libraries = 'kingres_columnar'
```

```sql
CREATE EXTENSION kingres_columnar;
CREATE TABLE public.events (ts timestamptz, value float8) USING columnar;
SELECT * FROM kingres.columnar_info('public.events');
SELECT * FROM kingres.columnar_check(false);
```

### Storage and operations

The `columnar` access method compresses stripes with `none`, `pglz`, `lz4` or `zstd`, and supports serial projected scans and min/max skipping. It has no index support and does not implement ordinary row updates/deletes. `kingres.columnar_set_options` affects future stripes; it is owner-only and nontransactional. `kingres.columnar_compact` merges small stripes; `kingres.columnar_stripes` and `kingres.columnar_options` inspect them. Read utilities require SELECT.

`kingres.columnar_check(false)` reports damage; repair mode and `kingres.columnar_reconcile` are superuser operations. The experimental WAL resource-manager ID is 128 and must be reserved before a release. Once columnar WAL exists, removing the library can prevent crash/archive recovery or restore. Without preload, some existing data can be read but writes are refused. Format changes within this development line require close revision matching.
