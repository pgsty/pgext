

## 用法

> [ltree_plpython3u: ltree 与 PL/Python3 之间的类型转换](https://www.postgresql.org/docs/current/ltree.html)

为 PL/Python3U 提供 `ltree` 类型的转换支持。加载后，`ltree` 值会自动转换为 Python 字符串，反之亦然。

```sql
CREATE EXTENSION ltree_plpython3u;

CREATE FUNCTION ltree_depth(val ltree) RETURNS integer
LANGUAGE plpython3u TRANSFORM FOR TYPE ltree AS $$
  # val 现在是一个 Python 字符串，如 "a.b.c"
  return len(val.split('.'))
$$;

CREATE FUNCTION ltree_leaf(val ltree) RETURNS text
LANGUAGE plpython3u TRANSFORM FOR TYPE ltree AS $$
  return val.split('.')[-1]
$$;

SELECT ltree_depth('top.science.astronomy'::ltree);  -- 3
SELECT ltree_leaf('top.science.astronomy'::ltree);    -- astronomy
```
