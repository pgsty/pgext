

## 用法

> [jsonb_plpython3u: jsonb 与 PL/Python3 之间的类型转换](https://www.postgresql.org/docs/current/datatype-json.html)

为 PL/Python3U 提供 `jsonb` 类型的转换支持。加载后，`jsonb` 值会自动转换为 Python 字典、列表和标量，反之亦然。

```sql
CREATE EXTENSION jsonb_plpython3u;

CREATE FUNCTION jsonb_keys_sorted(val jsonb) RETURNS text[]
LANGUAGE plpython3u TRANSFORM FOR TYPE jsonb AS $$
  # val 现在是一个 Python 字典
  return sorted(val.keys())
$$;

CREATE FUNCTION build_jsonb(name text, age integer) RETURNS jsonb
LANGUAGE plpython3u TRANSFORM FOR TYPE jsonb AS $$
  return {"name": name, "age": age}
$$;

SELECT jsonb_keys_sorted('{"b": 2, "a": 1, "c": 3}'::jsonb);
SELECT build_jsonb('Alice', 30);
```
