

## 用法

> [hstore_plperl: hstore 与 PL/Perl 之间的类型转换](https://www.postgresql.org/docs/current/hstore.html)

为 PL/Perl 提供 `hstore` 类型的转换支持。加载后，`hstore` 值会自动转换为 Perl 哈希，反之亦然。

```sql
CREATE EXTENSION hstore_plperl;

CREATE FUNCTION hstore_to_text(val hstore) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE hstore AS $$
  # val 现在是一个 Perl 哈希引用
  my $result = '';
  for my $key (sort keys %$val) {
    $result .= "$key => $val->{$key}\n";
  }
  return $result;
$$;

CREATE FUNCTION make_hstore(key text, value text) RETURNS hstore
LANGUAGE plperl TRANSFORM FOR TYPE hstore AS $$
  my ($k, $v) = @_;
  return { $k => $v };
$$;

SELECT hstore_to_text('a=>1, b=>2'::hstore);
SELECT make_hstore('color', 'blue');
```
