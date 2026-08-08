## Usage

Sources:

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [jsonb_plruby v1.0 control file](https://github.com/commandprompt/plruby/blob/v2.5.0/jsonb_plruby/jsonb_plruby.control)
- [jsonb_plruby v1.0 extension SQL](https://github.com/commandprompt/plruby/blob/v2.5.0/jsonb_plruby/jsonb_plruby--1.0.sql)

`jsonb_plruby` installs a PostgreSQL transform between `jsonb` and native Ruby values for the `plruby` language. A transformed `jsonb` argument becomes a Ruby `Hash`, `Array`, `String`, `Integer`, `Float`, `true`, `false`, or `nil`; compatible Ruby values can be returned directly as `jsonb`.

### Install and Use the Transform

```sql
CREATE EXTENSION plruby;
CREATE EXTENSION jsonb_plruby;

CREATE FUNCTION ruby_mark_processed(jsonb)
RETURNS jsonb
LANGUAGE plruby
TRANSFORM FOR TYPE jsonb
AS $$
  value = args[0]
  value['processed'] = true
  value
$$;

SELECT ruby_mark_processed('{"id": 42}'::jsonb);
```

The transform is used only by functions that declare `TRANSFORM FOR TYPE jsonb`. Other PL/Ruby functions keep the language's ordinary JSONB conversion behavior.

### Objects and Caveats

- `jsonb_to_plruby(internal)` implements SQL-to-Ruby conversion.
- `plruby_to_jsonb(internal)` implements Ruby-to-SQL conversion.
- The extension version is `1.0`, it requires `plruby`, and it is relocatable.
- Ruby `Hash` keys returned to PostgreSQL must be valid JSON object keys, and numeric/special values must be representable by PostgreSQL `jsonb`. Test nested values and numeric limits explicitly.
- PL/Ruby remains untrusted. Installing this transform does not sandbox Ruby code or reduce the privileges required to create PL/Ruby functions.
