

## 用法

> [hstore_plpython3u: hstore 与 PL/Python3 之间的类型转换](https://www.postgresql.org/docs/current/hstore.html)

为 PL/Python3U 提供 `hstore` 类型的转换支持。加载后，`hstore` 值会自动转换为 Python 字典，反之亦然。

```sql
CREATE EXTENSION hstore_plpython3u;

CREATE FUNCTION hstore_to_pairs(val hstore) RETURNS text
LANGUAGE plpython3u TRANSFORM FOR TYPE hstore AS $$
  # val 现在是一个 Python 字典
  return ', '.join(f'{k}={v}' for k, v in sorted(val.items()))
$$;

CREATE FUNCTION make_hstore(key text, value text) RETURNS hstore
LANGUAGE plpython3u TRANSFORM FOR TYPE hstore AS $$
  return {key: value}
$$;

SELECT hstore_to_pairs('a=>1, b=>2'::hstore);
SELECT make_hstore('color', 'blue');
```
