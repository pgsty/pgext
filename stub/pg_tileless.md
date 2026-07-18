## Usage

Sources:

- [Version 0.1 SQL objects](https://github.com/TilelessMap/pg_tileless/blob/e951bafabb65aaacce9fd0bf1a40eece7e9beb68/pg_tileless--0.1.sql)
- [SQLite writer implementation](https://github.com/TilelessMap/pg_tileless/blob/e951bafabb65aaacce9fd0bf1a40eece7e9beb68/sqlite_writer.c)
- [Extension control file](https://github.com/TilelessMap/pg_tileless/blob/e951bafabb65aaacce9fd0bf1a40eece7e9beb68/pg_tileless.control)

`pg_tileless` is an abandoned prototype for packing PostGIS geometries into TWKB datasets and writing them into SQLite. It creates `tileless` and `tmp` schemas and depends on `plpgsql`, `postgis`, and `postgis_sfcgal`.

```sql
CREATE EXTENSION pg_tileless CASCADE;
SELECT tileless.pack_twkb_linestring(
  'public.road', 'geom', 'road_id', '',
  '/srv/export/roads.sqlite', 'roads'
);
```

The polygon and linestring packers construct dynamic SQL, create and drop working tables, simplify and subdivide geometries, and call `tileless.TWKB_Write2SQLite(...)` to write a server-local file. Arguments such as table, column, prefix, query, and output path are concatenated into SQL or passed to native file code; never expose these functions to untrusted callers.

The reviewed README contains no operational documentation, and the project is marked abandoned. Run only under a tightly restricted role on disposable data. Confirm filesystem permissions, path confinement, SQL-injection resistance, geometry validity, precision loss, cleanup after failure, SQLite consistency, and PostgreSQL/PostGIS/SFCGAL ABI compatibility before any use.
