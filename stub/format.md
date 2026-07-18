## Usage

Sources:

- [Project README](https://github.com/melanieplageman/format/blob/837b6dcc957d2190105c819181eb19c0b1ad9b76/README.md)
- [Extension control file](https://github.com/melanieplageman/format/blob/837b6dcc957d2190105c819181eb19c0b1ad9b76/format.control)
- [SQL and C implementation](https://github.com/melanieplageman/format/tree/837b6dcc957d2190105c819181eb19c0b1ad9b76/sql)
- [Regression examples](https://github.com/melanieplageman/format/blob/837b6dcc957d2190105c819181eb19c0b1ad9b76/test/sql/base.sql)

`format` 0.0.1 is a deprecated C extension that overloads `format(text, hstore)` to substitute named values from an `hstore`. Upstream directs users to its successor, `format_x`; prefer that successor or PostgreSQL's built-in formatting facilities for new code.

### Named substitution

Install the required `hstore` dependency before this extension:

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION format;

SELECT format(
  '%(name)I = %(value)L',
  hstore(ARRAY['name', 'status', 'value', 'ready'])
);
```

Placeholders use `%(key)s`, `%(key)I`, or `%(key)L` for plain strings, SQL identifiers, and SQL literals. The formatter also implements width and alignment flags. Identifier and literal modes should be used when constructing SQL; plain string substitution is not SQL escaping.

### Compatibility and migration

The function is strict and immutable. A missing key or a null value can emit a warning; a null identifier is an error, while literal formatting can produce `NULL`. Exercise every template with missing and null inputs before relying on it.

The upstream repository is archived in practice and explicitly marks this implementation deprecated. Its overloaded name can also be confused with PostgreSQL's built-in `format()` function during review and migration. Do not introduce a new dependency on version 0.0.1; migrate existing templates to `format_x` or explicit built-in formatting, and compare escaping and null behavior before removing the extension.
