## Usage

Sources:

- [Official README](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/README.md)
- [Official base extension SQL](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/init.sql)
- [Official RUM support SQL](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/rum_support.sql)
- [Official extension control file](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/tsvector2.control)

`tsvector2` is an extended full-text-search vector type designed for better compression and to remove the original `tsvector` 1 MB size limit. Version `1.0` mirrors common PostgreSQL text-search operations and supports B-tree, GIN, GiST, and optional RUM indexes.

### Core Workflow

Store generated `tsvector2` values and index them with the supplied GIN operator class:

```sql
CREATE EXTENSION tsvector2;

CREATE TABLE article (
  id bigint PRIMARY KEY,
  body text,
  search tsvector2
);

UPDATE article
SET search = to_tsvector2('english'::regconfig, body);

CREATE INDEX article_search_idx
  ON article USING gin (search gin_tsvector2_ops);

SELECT id
FROM article
WHERE search @@ plainto_tsquery('english', 'database search');
```

The type supports comparison, concatenation with `||`, matching with `@@`, ranking, weighting, filtering, lexeme deletion, and conversion from text, arrays, JSON, and JSONB.

### Important Objects and Indexes

Extension-specific entry points include `to_tsvector2`, `array_to_tsvector2`, `tsvector2_to_array`, `tsvector2_stat`, `jsonb_to_tsvector2`, `json_to_tsvector2`, `tsvector2_update_trigger`, and `tsvector2_update_trigger_column`. Common names such as `strip`, `unnest`, `length`, `setweight`, `ts_rank`, `ts_rank_cd`, `ts_delete`, and `ts_filter` are overloaded for the new type.

The default classes are `bt_tsvector2_ops`, `gin_tsvector2_ops`, and `gist_tsvector2_ops`. If the RUM extension already exists when `tsvector2` is created, installation adds RUM distance, score, and operator classes through `install_rum_for_tsvector2()`; install RUM first when that support is required.

The extension is fixed to the `public` schema and is not relocatable. Upstream describes PostgreSQL 9.6 or newer, and notes that JSON conversion behavior differs between PostgreSQL 10 and 11 in the same way as core text search. Test application queries, generated columns or triggers, and index rebuild procedures on the exact server version before replacing core `tsvector` storage.
