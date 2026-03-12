

## 用法

> [hstore_plperlu: hstore 与不受信 PL/Perl 之间的类型转换](https://www.postgresql.org/docs/current/hstore.html)

为不受信 PL/Perl 提供 `hstore` 类型的转换支持。加载后，`hstore` 值会自动转换为 Perl 哈希，反之亦然。

```sql
CREATE EXTENSION hstore_plperlu;

CREATE FUNCTION hstore_to_json_u(val hstore) RETURNS text
LANGUAGE plperlu TRANSFORM FOR TYPE hstore AS $$
  use JSON;
  # val 现在是一个 Perl 哈希引用
  return encode_json($val);
$$;

CREATE FUNCTION make_hstore_u(key text, value text) RETURNS hstore
LANGUAGE plperlu TRANSFORM FOR TYPE hstore AS $$
  my ($k, $v) = @_;
  return { $k => $v };
$$;

SELECT hstore_to_json_u('a=>1, b=>2'::hstore);
SELECT make_hstore_u('color', 'blue');
```
