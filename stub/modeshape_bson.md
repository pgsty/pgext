## Usage

Sources:

- [Project README](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/README.md)
- [Extension control file](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/modeshape_bson.control)
- [Version 1.0 SQL API](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/modeshape_bson--1.0.sql)
- [libbson conversion implementation](https://github.com/dalbani/pg-modeshape-bson/blob/82166389b332bfec0469232fc3b305aeec83d0d8/modeshape_bson.c)

`modeshape_bson` 1.0 converts ModeShape BSON stored as `bytea` to canonical Extended JSON text and converts JSON text back to BSON. It was designed to place a PostgreSQL JSONB table behind a writable proxy table expected by ModeShape.

### Test a conversion

```sql
CREATE EXTENSION modeshape_bson;

SELECT modeshape_bson_version();

WITH encoded AS (
  SELECT json_text_to_modeshape_bson('{"name":"node","count":2}') AS value
)
SELECT modeshape_bson_to_json_text(value)
FROM encoded;
```

The returned JSON is libbson's canonical Extended JSON, so textual formatting and representations may differ after a round trip even when values remain equivalent. Validate every BSON type used by the application, including dates, binary values, numeric widths, object identifiers, arrays, nulls, and nested documents.

### Legacy format and proxy-table risk

Upstream documents a ModeShape `repository:info` record whose `$date` value uses BSON type 9 instead of the usually observed ISO string and notes uncertainty about writing it. Treat that record and all uncommon types as migration blockers until byte-for-byte and semantic tests pass.

The source last changed in 2018 and vendors an old Mongo C driver/libbson tree. It publishes no current PostgreSQL, ModeShape, or libbson compatibility matrix. The suggested writable-view design uses rules to redirect all reads and writes, which adds concurrency, privilege, trigger, and error-propagation complexity. Use a least-privileged owner, avoid credentials in stored configuration examples, retain raw source bytes during migration, and prefer a maintained application-side converter unless this native code is rebuilt and audited for the exact platform.
