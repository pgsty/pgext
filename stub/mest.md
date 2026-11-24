## Usage

Sources:

- [MEST README](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/README.md)
- [Control](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/mest.control)
- [SQL 1.0](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/mest--1.0.sql)

`mest` 1.0 adds multi-entry GiST and SP-GiST index access methods. One source value can contribute several index entries, useful for multiranges and paths. Upstream documents PostgreSQL 17 and 18.

### Core Workflow

Install the library for the server's PostgreSQL major, then enable it with a superuser:

```sql
CREATE EXTENSION mest;

CREATE TABLE time_windows (id bigint, spans tstzmultirange);
CREATE INDEX time_windows_mgist ON time_windows USING mgist (spans);
CREATE INDEX time_windows_mspgist ON time_windows USING mspgist (spans);
```

Choose the appropriate index after comparing the query plan and index size for the actual data; the example shows both available methods.

### Objects and Companion Extensions

- `mgist` and `mspgist` are index access methods.
- `multirange_mrtree_ops` and `multirange_mquadtree_ops` provide default classes for `anymultirange`.
- `path_mrtree_ops` and `path_mquadtree_ops` support `path`.
- `ranges()` has overloads for integer, date and timestamp multiranges.
- `postgis_mest` adds geometry classes; `mobilitydb_mest` adds classes for MobilityDB types. Install those extensions separately when needed.

### Operational Boundaries

The base control file is relocatable and declares no extension dependencies. Its installation script creates access methods and C functions; installation is privileged. The base extension has no documented preload or restart requirement. Companion extensions have their own library and preload prerequisites.

Multiple entries per row change index size and maintenance work. Ordinary queries must use operators supported by the selected class; installing the extension does not make every predicate indexable.
