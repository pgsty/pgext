## Usage

Sources:

- [Official README](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/README.md)
- [Official extension SQL](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/sql/pg_vartype--0.4.sql)
- [Official comparison implementation](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/src/pg_vartype.c)

`pg_vartype` provides a `vartype` column type that can hold integer, floating-point, Boolean, string, date, or timestamp values. It predates PostgreSQL `jsonb`; upstream notes that native JSON storage is a better fit for many current applications.

### Core Workflow

```sql
CREATE EXTENSION pg_vartype;

CREATE TABLE attributes (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  value vartype
);

INSERT INTO attributes (value) VALUES
  (42::bigint::vartype),
  ('blue'::text::vartype),
  (DATE '2026-01-01'::vartype),
  (false::boolean::vartype);

SELECT id, value, vartype_type(value), vartype_len(value)
FROM attributes;
```

When using an untyped literal, a string must contain double quotes, for example `'"blue"'::vartype`; otherwise a token such as `'false'::vartype` is parsed as a Boolean. Casting from `text` is less ambiguous for application parameters.

### Types, Casts, and Operators

- `vartype_type` is an enum describing the stored category; the `vartype_type(vartype)` and `vartype_len(vartype)` helpers expose category and stored length.
- Explicit casts cover common integer and floating types, `boolean`, `text`/character types, `date`, `timestamp`, and `timestamptz` in both directions where defined.
- Comparison operators and the default `vartype_ops` B-tree operator class allow ordering and indexing.

Values from different category groups are ordered by the extension's internal precedence. Integer and floating values are intended to compare numerically, and date/timestamp values are compared through conversion; these rules are extension-specific rather than PostgreSQL's ordinary single-type semantics.

### Critical Comparison Caveat

The reviewed `0.4` C comparator handles mixed numeric operands asymmetrically: one integer-versus-float branch converts the float to an integer, while the reverse branch compares as floating point. Fractional values can therefore produce inconsistent comparison results and may violate the ordering contract required by `vartype_ops`.

Do not use mixed integer/float `vartype` values in B-tree indexes, unique constraints, ordered aggregates, or range predicates until this behavior is fixed or conclusively reproduced as safe on the chosen build. Keeping indexed values within one category avoids that particular path but does not replace validation.

The extension is non-relocatable and the reviewed codebase was last changed in 2017, with documentation centered on PostgreSQL 9.6. Test input parsing, casts, time-zone behavior, comparison laws, dump/restore, and the exact PostgreSQL version before production use.
