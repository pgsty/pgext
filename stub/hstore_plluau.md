


## Usage

> [hstore_plluau: hstore transform for PL/Lua untrusted](https://github.com/pllua/pllua)

The `hstore_plluau` extension provides a transform that allows `hstore` values to be passed directly to and from untrusted PL/Lua (`plluau`) functions as native Lua tables.

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION plluau;
CREATE EXTENSION hstore_plluau;
```

### Using hstore in plluau

With the transform installed, `hstore` arguments are automatically converted to Lua tables:

```sql
CREATE FUNCTION process_hstore(h hstore) RETURNS hstore LANGUAGE plluau AS $$
  -- h is a Lua table with string keys and string values
  h["processed"] = "true"
  -- Can also do unrestricted operations (file I/O, etc.)
  return h
$$;

SELECT process_hstore('key1 => "val1", key2 => "val2"'::hstore);
```

This is the untrusted counterpart to `hstore_pllua`. The behavior is identical -- `hstore` values become Lua tables and vice versa -- but functions run in the unrestricted `plluau` environment.
