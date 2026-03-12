


## Usage

> [hstore_pllua: hstore transform for PL/Lua](https://github.com/pllua/pllua)

The `hstore_pllua` extension provides a transform that allows `hstore` values to be passed directly to and from PL/Lua functions as native Lua tables, instead of opaque datum objects.

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION pllua;
CREATE EXTENSION hstore_pllua;
```

### Using hstore in PL/Lua

With the transform installed, `hstore` arguments are automatically converted to Lua tables, and Lua tables can be returned as `hstore` values:

```sql
CREATE FUNCTION process_hstore(h hstore) RETURNS hstore LANGUAGE pllua AS $$
  -- h is a Lua table with string keys and string values
  h["new_key"] = "new_value"
  h["modified"] = "true"
  return h
$$;

SELECT process_hstore('name => "Alice", age => "30"'::hstore);
```

### Iterate Over hstore Keys

```sql
CREATE FUNCTION hstore_keys(h hstore) RETURNS SETOF text LANGUAGE pllua AS $$
  for k, v in pairs(h) do
    coroutine.yield(k)
  end
$$;
```

Without this transform, `hstore` values would need to be manually converted using the `hstore` type functions.
