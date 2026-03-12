

## Usage

> [ltree_plpython3u: Transform between ltree and PL/Python3](https://www.postgresql.org/docs/current/ltree.html)

Provides a transform for the `ltree` type for PL/Python3U. When loaded, `ltree` values are automatically converted to Python strings and vice versa.

```sql
CREATE EXTENSION ltree_plpython3u;

CREATE FUNCTION ltree_depth(val ltree) RETURNS integer
LANGUAGE plpython3u TRANSFORM FOR TYPE ltree AS $$
  # val is now a Python string like "a.b.c"
  return len(val.split('.'))
$$;

CREATE FUNCTION ltree_leaf(val ltree) RETURNS text
LANGUAGE plpython3u TRANSFORM FOR TYPE ltree AS $$
  return val.split('.')[-1]
$$;

SELECT ltree_depth('top.science.astronomy'::ltree);  -- 3
SELECT ltree_leaf('top.science.astronomy'::ltree);    -- astronomy
```
