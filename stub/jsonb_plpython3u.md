

## Usage

> [jsonb_plpython3u: Transform between jsonb and PL/Python3](https://www.postgresql.org/docs/current/datatype-json.html)

Provides a transform for the `jsonb` type for PL/Python3U. When loaded, `jsonb` values are automatically converted to Python dicts, lists, and scalars, and vice versa.

```sql
CREATE EXTENSION jsonb_plpython3u;

CREATE FUNCTION jsonb_keys_sorted(val jsonb) RETURNS text[]
LANGUAGE plpython3u TRANSFORM FOR TYPE jsonb AS $$
  # val is now a Python dict
  return sorted(val.keys())
$$;

CREATE FUNCTION build_jsonb(name text, age integer) RETURNS jsonb
LANGUAGE plpython3u TRANSFORM FOR TYPE jsonb AS $$
  return {"name": name, "age": age}
$$;

SELECT jsonb_keys_sorted('{"b": 2, "a": 1, "c": 3}'::jsonb);
SELECT build_jsonb('Alice', 30);
```
