

## 用法

> [jsonb_plperlu: jsonb 与不受信 PL/Perl 之间的类型转换](https://www.postgresql.org/docs/current/datatype-json.html)

为不受信 PL/Perl 提供 `jsonb` 类型的转换支持。加载后，`jsonb` 值会自动转换为 Perl 数据结构（哈希、数组、标量），反之亦然。

```sql
CREATE EXTENSION jsonb_plperlu;

CREATE FUNCTION process_json_u(val jsonb) RETURNS text
LANGUAGE plperlu TRANSFORM FOR TYPE jsonb AS $$
  use JSON;
  # val 已经是 Perl 数据结构，重新编码并排序
  return encode_json($val);
$$;

CREATE FUNCTION build_jsonb_u(name text, age integer) RETURNS jsonb
LANGUAGE plperlu TRANSFORM FOR TYPE jsonb AS $$
  my ($name, $age) = @_;
  return { name => $name, age => $age };
$$;

SELECT process_json_u('{"b": 2, "a": 1}'::jsonb);
SELECT build_jsonb_u('Alice', 30);
```
