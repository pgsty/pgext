## Usage

Sources:

- [Official usage documentation](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/doc/pg_cache_tree.md)
- [Version 1.0.0 SQL implementation](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/sql/pg_cache_tree--1.0.0.sql)
- [Official regression workflow](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/test/sql/base.sql)

`pg_cache_tree` 1.0.0 maintains materialized ancestor and descendant ID arrays for an adjacency-list tree. Row triggers update the arrays after inserts, deletes, or changes to a node ID or parent ID. It is useful when reads of complete ancestry or descendants dominate, at the cost of write amplification and denormalized state.

### Prepare a Tree Table

The documented and safest layout uses the default column names:

```sql
CREATE EXTENSION pg_cache_tree;

CREATE TABLE business_unit (
  id integer PRIMARY KEY,
  "parentId" integer REFERENCES business_unit(id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  "parentsCache" integer[] NOT NULL DEFAULT '{}',
  "childrenCache" integer[] NOT NULL DEFAULT '{}'
);

SELECT ct_create_trigger('business_unit');
```

Inserts and moves then update both directions of the cache:

```sql
INSERT INTO business_unit (id, "parentId") VALUES
  (1, NULL), (2, 1), (20, 2), (21, 2);

SELECT id, "parentId", "parentsCache", "childrenCache"
FROM business_unit
ORDER BY id;
```

Use `ct_drop_trigger('business_unit')` to remove the two generated triggers.

### Main Objects

- `ct_create_trigger(regclass, id, parent, parents, children)`: creates INSERT/DELETE and UPDATE triggers.
- `ct_trigger_update_cache()`: trigger function that maintains ancestor and descendant arrays.
- `ct_drop_trigger(regclass, parent)`: drops triggers created by `ct_create_trigger`.
- `ct_insert_array` and `ct_subtract_array`: internal array helpers; the `-#` operator exposes subtraction.
- `ct_assign` and `ct_copy_value`: dynamic composite-field helpers used by the trigger.

### Integrity and Version Boundaries

Version 1.0.0 has implementation defects behind its configurable-column interface. The trigger still reads `NEW."parentId"` literally, so a custom parent column can fail even when supplied to `ct_create_trigger`. The scalar `ct_insert_array(anyarray, anyelement, anyelement)` overload calls an unrelated `extra_modules.t_cpc_insert_array` function. Use the documented default columns, avoid that helper overload, and regression-test any local fix.

The extension does not prevent cycles. A self-referencing foreign key guarantees only that a parent row exists, not that the graph remains a tree. Add an application or constraint-trigger cycle check before enabling cache maintenance.

Moving or deleting a high-level node updates many ancestor and descendant rows in the same transaction. Size write/WAL/lock impact with realistic depth and fan-out. Disabling triggers, failed bulk loads, or direct cache-column edits can leave arrays stale, and upstream provides no rebuild command; keep a canonical `id`/`parentId` source and a tested reconciliation procedure.
