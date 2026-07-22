## Usage

Sources:

- [Official PGXN documentation](https://pgxn.org/dist/pair/0.1.8/README.html)
- [Official extension SQL](https://api.pgxn.org/src/pair/pair-0.1.8/sql/pair.sql)

`pair` adds a small ordered key/value composite type. Both fields are `text`; its main use is passing an ordered variadic list of pairs to a function without adopting a larger container such as `hstore` or JSON.

### Core Workflow

Construct a value with `pair()` or the `~>` operator, then access its `k` and `v` fields:

```sql
CREATE EXTENSION pair;

SELECT pair('environment', 'production');
SELECT 'retries' ~> 3;
SELECT ('region' ~> 'us-east-1').k,
       ('region' ~> 'us-east-1').v;
```

The overloads accept `text` and polymorphic `anyelement` arguments, but the resulting composite still stores each value as `text`.

### Variadic Use

The type is especially useful as a variadic function parameter:

```sql
CREATE FUNCTION show_pairs(VARIADIC items pair[])
RETURNS TABLE (key text, value text)
LANGUAGE sql
AS $$
  SELECT (item).k, (item).v
  FROM unnest(items) AS item
$$;

SELECT * FROM show_pairs(
  'environment' ~> 'production',
  'retries' ~> 3
);
```

### Important Objects and Notes

- Type: `pair`, with fields `k text` and `v text`.
- Constructors: `pair(anyelement, text)`, `pair(text, anyelement)`, `pair(anyelement, anyelement)`, and `pair(text, text)`.
- Operator: matching `~>` overloads that call the constructors.

`pair` has no runtime dependency beyond PostgreSQL, is relocatable, and needs no preload or restart. PGXN distribution 0.1.8 still declares the provided SQL extension version as 0.1.2, matching the catalog version; the distribution update is not an extension-schema version bump.
