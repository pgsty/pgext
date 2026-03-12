

## 用法

> [bool_plperl: bool 与 PL/Perl 之间的类型转换](https://www.postgresql.org/docs/current/plperl.html)

为 PL/Perl 提供 `bool` 类型的转换支持。加载后，PostgreSQL 的 `boolean` 值会自动转换为 Perl 原生布尔表示（而非字符串），反之亦然。

```sql
CREATE EXTENSION bool_plperl;

CREATE FUNCTION check_flag(val boolean) RETURNS text
LANGUAGE plperl TRANSFORM FOR TYPE boolean AS $$
  # val 是 Perl 原生布尔值（1 或 undef），而非字符串
  if ($_[0]) {
    return "flag is set";
  }
  return "flag is not set";
$$;

SELECT check_flag(true);   -- flag is set
SELECT check_flag(false);  -- flag is not set
```
