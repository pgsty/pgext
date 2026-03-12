

## Usage

> [hstore_plpython3u: Transform between hstore and PL/Python3](https://www.postgresql.org/docs/current/hstore.html)

Provides a transform for the `hstore` type for PL/Python3U. When loaded, `hstore` values are automatically converted to Python dicts and vice versa.

```sql
CREATE EXTENSION hstore_plpython3u;

CREATE FUNCTION hstore_to_pairs(val hstore) RETURNS text
LANGUAGE plpython3u TRANSFORM FOR TYPE hstore AS $$
  # val is now a Python dict
  return ', '.join(f'{k}={v}' for k, v in sorted(val.items()))
$$;

CREATE FUNCTION make_hstore(key text, value text) RETURNS hstore
LANGUAGE plpython3u TRANSFORM FOR TYPE hstore AS $$
  return {key: value}
$$;

SELECT hstore_to_pairs('a=>1, b=>2'::hstore);
SELECT make_hstore('color', 'blue');
```
