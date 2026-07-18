## Usage

Sources:

- [README at the reviewed commit](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/README.md)
- [Extension control file](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/pg_twkb.control)
- [SQL API and indexed-tile implementation](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/pg_twkb--0.1.sql)
- [TWKB C entry points](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/pg_twkb.c)
- [Server-side file writers](https://github.com/nicklasaven/pg_twkb/blob/345cd979c9dafa804556d2aabdc872ca194d44f8/lwout_twkb.c)

`pg_twkb` is an unfinished PostGIS-dependent toolkit for constructing indexed Tiny Well-Known Binary data sets. `TWKB_IndexedTiles` recursively partitions a geometry table into tiles and combines tile IDs with TWKB payloads; `TWKB_Collect` combines already encoded payload arrays.

### Build an Indexed Payload

For a trusted table with an integer identifier and PostGIS geometry column:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_twkb;

SELECT TWKB_IndexedTiles(
  'public.features',
  'geom',
  'id',
  5,
  4326,
  1000
);
```

Other installed functions include `TWKB_MakeSquarebox`, `TWKB_DivideBox`, `TWKB_getTileId`, overloaded `TWKB_Write2File`, and `TWKB_Write2SQLite`.

### Caveats

- Upstream says the repository exists for personal versioning and that its functionality still needs documentation before it is usable. Version `0.1` is old experimental code, not a supported interchange format or production tool.
- `TWKB_Write2File` and `TWKB_Write2SQLite` write paths visible to the PostgreSQL server operating-system account. The install SQL grants no special protection; revoke `EXECUTE` from untrusted roles and never accept caller-controlled paths.
- Both `TWKB_Write2File` overloads are incorrectly declared `IMMUTABLE` despite external file side effects. Do not use them in expressions where PostgreSQL may fold, duplicate, or reorder calls.
- Indexed-tile functions construct dynamic SQL from table and column-name text. Supply only trusted identifiers and review permissions on the referenced relation.
- The control file requires `postgis`, while the implementation uses historical PostGIS and PostgreSQL C APIs and has no current compatibility matrix. Test the exact build and validate consumers against the generated custom container.
