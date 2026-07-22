## Usage

Sources:

- [Official format_x documentation](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/doc/format_x.md)
- [Version 0.0.1 SQL objects](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/sql/format_x--0.0.1.sql)
- [Official PGXN distribution page](https://pgxn.org/dist/format_x/)

`format_x` extends PostgreSQL's `format()` with named and chained lookups into composite values, `json`, `jsonb`, and `hstore`. Use it to format structured values without extracting every field first.

### Named Lookups

```sql
CREATE EXTENSION format_x;

SELECT format_x(
  'Hello %(name)s from %(address.city)s',
  '{"name":"Ada","address":{"city":"London"}}'::jsonb
);
```

A lookup has the form `%(key)s`; dots traverse nested values. Positional argument selection can be combined with a lookup, such as `%2(name)s`. The standard `%s`, `%I`, and `%L` conversions, position fields, widths, and precision are also supported.

```sql
SELECT format_x(
  'SELECT * FROM %1(name)I WHERE code = %1(code)L',
  '{"name":"places","code":"GB"}'::jsonb
);
```

### Inputs and Caveats

The installed overloads are `format_x(text)` and `format_x(text, VARIADIC "any")`. Composite, JSON, and JSONB lookups work directly. `hstore` is only needed when an `hstore` value is passed; it is not a declared extension dependency.

Missing keys, unsupported input types, and some `NULL` lookups raise errors. `%I` quotes an identifier and rejects `NULL`; `%L` quotes a literal using `quote_nullable` behavior. Quoting fields does not make an arbitrary format template safe: keep the template trusted and use parameters when executing generated SQL. This version uses PostgreSQL internal formatting and type APIs, so verify the build on the exact PostgreSQL release before deployment.
