## Usage

Sources:

- [Official PGXN documentation](https://pgxn.org/dist/test_factory/0.5.0/doc/test_factory.html)
- [Official README](https://github.com/BlueTreble/test_factory/blob/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054/README.md)
- [Version 0.5.0 extension SQL](https://github.com/BlueTreble/test_factory/blob/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054/sql/test_factory.sql)
- [Extension control files](https://github.com/BlueTreble/test_factory/tree/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054)

`test_factory` 0.5.0 registers named recipes for relational test data, creates a recipe's rows on first request, and caches a persistent copy for later tests. It is designed to avoid repeatedly rebuilding deep test-data dependencies while keeping references readable.

### Register and Retrieve a Test Set

Installation creates roles and fixed schemas, so load the extension as a superuser. Register each recipe with SQL that returns rows shaped exactly like the target table:

```sql
CREATE EXTENSION test_factory;

CREATE TABLE customer (
  customer_id bigint GENERATED ALWAYS AS IDENTITY,
  email text NOT NULL
);

SELECT tf.register(
  table_name => 'customer',
  test_sets => ARRAY[
    ROW(
      'base',
      'INSERT INTO customer(email) VALUES (''base@example.org'') RETURNING *'
    )::tf.test_set
  ]
);

SELECT * FROM tf.get(NULL::customer, 'base');
```

The first `tf.get` executes the registered SQL and caches a copy. Later calls with the same table type and set name return that copy without rerunning the recipe. Recipes can call `tf.get` for other tables, which expresses dependencies without requiring registration order.

### Main Objects

- `tf.test_set` is `(set_name text, insert_sql text)`.
- `tf.register(table_name text, test_sets tf.test_set[])` adds or replaces recipes with matching names for an existing table.
- `tf.get(table_type anyelement, set_name text) RETURNS SETOF anyelement` identifies the result row type through a typed NULL such as `NULL::customer`.
- Optional extension `test_factory_pgtap` requires `pgtap` and adds `tf.tap(table_name, set_name)` as an `isnt_empty` convenience assertion.

The extension stores recipe definitions in `_tf._test_factory` and cached row copies in `_test_factory_test_data`, owned through the `test_factory__owner` role. Its configuration tables are included in extension-aware dumps.

### Cache, Privilege, and Cleanup Boundaries

Cached copies are never removed automatically. Changes to the original application rows are not reflected by later `tf.get` calls, and registering a new list does not remove older set names omitted from that call. Plan explicit cleanup or recreate the test database between independent suites.

Recipe SQL is dynamic SQL executed when a caller first requests a set. The caller needs privileges for its statements, and recipe text must be trusted code rather than user input. A recipe must return rows compatible with the target table, normally through `RETURNING *`; schema drift can otherwise make retrieval fail or preserve stale cached shapes.

Use this extension only in controlled test databases or carefully isolated test-data schemas. Its persistent roles, schemas, executable recipes, and cached data are inappropriate for an unreviewed production installation.
