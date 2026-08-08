## Usage

Sources:

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [hstore_plruby v1.0 control file](https://github.com/commandprompt/plruby/blob/v2.5.0/hstore_plruby/hstore_plruby.control)
- [hstore_plruby v1.0 extension SQL](https://github.com/commandprompt/plruby/blob/v2.5.0/hstore_plruby/hstore_plruby--1.0.sql)

`hstore_plruby` installs a PostgreSQL transform between `hstore` and Ruby `Hash` values for the `plruby` language. Keys become Ruby strings and values become strings or `nil`; a compatible Ruby hash can be returned directly as `hstore`.

### Install and Use the Transform

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION plruby;
CREATE EXTENSION hstore_plruby;

CREATE FUNCTION ruby_add_hstore_key(hstore)
RETURNS hstore
LANGUAGE plruby
TRANSFORM FOR TYPE hstore
AS $$
  value = args[0]
  value['processed'] = 'yes'
  value
$$;

SELECT ruby_add_hstore_key('id=>42'::hstore);
```

The transform is used only by functions that declare `TRANSFORM FOR TYPE hstore`.

### Objects and Caveats

- `hstore_to_plruby(internal)` implements SQL-to-Ruby conversion.
- `plruby_to_hstore(internal)` implements Ruby-to-SQL conversion.
- The extension version is `1.0`, it requires both `hstore` and `plruby`, and it is relocatable.
- `hstore` is a flat string-to-string-or-NULL map. It does not preserve nested Ruby hashes, arrays, or typed numeric values; use `jsonb_plruby` when those shapes matter.
- PL/Ruby remains untrusted. Installing this transform does not sandbox Ruby code or reduce the privileges required to create PL/Ruby functions.
