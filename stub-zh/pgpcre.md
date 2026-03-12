

## 用法

> [pgpcre: PostgreSQL 的 Perl 兼容正则表达式（PCRE）支持](https://github.com/petere/pgpcre)

提供 `pcre` 数据类型以及用于 PCRE 模式匹配的运算符和函数。

### 基本匹配

```sql
SELECT 'foo' ~ pcre 'fo+';        -- true
SELECT 'bar' !~ pcre 'fo+';       -- true
SELECT 'foo' =~ pcre 'fo+';       -- true（Perl 风格运算符）
```

反转操作数顺序：

```sql
SELECT pcre 'fo+' ~ 'foo';        -- true
SELECT pcre 'fo+' ~ ANY(ARRAY['foo', 'bar']);
```

### 大小写不敏感匹配

```sql
SELECT 'FOO' ~ pcre '(?i)fo+';    -- true
```

### 提取匹配的字符串

```sql
SELECT pcre_match('fo+', 'foobar');    -- 'foo'
SELECT pcre_match('fo+', 'barbar');    -- NULL
```

### 提取捕获的子串

```sql
SELECT pcre_captured_substrings('(fo+)(b..)', 'foobar');
-- ARRAY['foo','bar']

SELECT pcre_captured_substrings('(a|(z))(bc)', 'abc');
-- ARRAY['a',NULL,'bc']
```

### 存储正则表达式

`pcre` 类型可以存储在表的列中。二进制形式包含已编译的正则表达式，与 PCRE 库版本关联。PCRE 库升级后，需要重新编译存储的值：

```sql
UPDATE my_table SET pcre_col = pcre_col::text::pcre;
```
