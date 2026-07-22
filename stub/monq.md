## Usage

Sources:

- [Official README](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/README.md)
- [Version 0.1 SQL API](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/monq--0.1.sql)
- [Extension control file](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/monq.control)

`monq` translates a subset of MongoDB query syntax into JsQuery expressions and evaluates them against PostgreSQL `jsonb`. It exposes the `mquery` type and a commutative `<=>` match operator; it is a compatibility aid, not a MongoDB server or a complete implementation of MongoDB query semantics.

### Core Workflow

Install the required `jsquery` extension first, cast the query text to `mquery`, and match it against a JSON document.

```sql
CREATE EXTENSION jsquery;
CREATE EXTENSION monq;

SELECT '{"a":["ssl","security","pattern"]}'::jsonb <=>
       '{a: {$all: ["ssl", "security"]}}'::mquery;
```

Both `jsonb <=> mquery` and `mquery <=> jsonb` are declared. Supported families include comparison operators, logical operators, `$exists`, `$type`, `$text`, `$all`, `$elemMatch`, and `$size`.

### Compatibility Limits

The README explicitly excludes `$mod`, `$regex`, `$where`, bitwise, geospatial, comment, and projection operators. Results are constrained by JsQuery behavior, so do not assume MongoDB coercion, missing-field, array, collation, or null semantics without a focused test corpus. The SQL API defines no index operator class for `<=>`; confirm plans on realistic data rather than assuming these predicates are index-backed.

The project describes itself as still under development and targets PostgreSQL 9.4-era APIs. Pin `monq` and `jsquery` together, validate parser behavior with untrusted query text, and compare every application query against the intended MongoDB result before migration.
