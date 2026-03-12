

## 用法

> [jsonb_plperl: jsonb 与 PL/Perl 之间的类型转换](https://www.postgresql.org/docs/current/datatype-json.html)

为 PL/Perl 提供 `jsonb` 类型的转换支持。加载后，`jsonb` 值会自动转换为 Perl 数据结构（哈希、数组、标量），反之亦然。

```sql
CREATE EXTENSION jsonb_plperl;

CREATE FUNCTION jsonb_keys_sorted(val jsonb) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE jsonb AS $$
  # val 现在是一个 Perl 哈希/数组引用
  return join(', ', sort keys %$val);
$$;

CREATE FUNCTION build_jsonb(name text, age integer) RETURNS jsonb
LANGUAGE plperl TRANSFORM FOR TYPE jsonb AS $$
  my ($name, $age) = @_;
  return { name => $name, age => $age };
$$;

SELECT jsonb_keys_sorted('{"b": 2, "a": 1, "c": 3}'::jsonb);
SELECT build_jsonb('Alice', 30);
```
