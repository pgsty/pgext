


## Usage

> [pgpcre: Perl-compatible regular expressions (PCRE) for PostgreSQL](https://github.com/petere/pgpcre)

Provides a `pcre` data type and operators/functions for PCRE pattern matching.

### Basic Matching

```sql
SELECT 'foo' ~ pcre 'fo+';        -- true
SELECT 'bar' !~ pcre 'fo+';       -- true
SELECT 'foo' =~ pcre 'fo+';       -- true (Perl-style operator)
```

Reverse operand order:

```sql
SELECT pcre 'fo+' ~ 'foo';        -- true
SELECT pcre 'fo+' ~ ANY(ARRAY['foo', 'bar']);
```

### Case-Insensitive Matching

```sql
SELECT 'FOO' ~ pcre '(?i)fo+';    -- true
```

### Extract Matched String

```sql
SELECT pcre_match('fo+', 'foobar');    -- 'foo'
SELECT pcre_match('fo+', 'barbar');    -- NULL
```

### Extract Captured Substrings

```sql
SELECT pcre_captured_substrings('(fo+)(b..)', 'foobar');
-- ARRAY['foo','bar']

SELECT pcre_captured_substrings('(a|(z))(bc)', 'abc');
-- ARRAY['a',NULL,'bc']
```

### Storing Regular Expressions

The `pcre` type can be stored in table columns. The binary form contains the compiled regex, tied to the PCRE library version. After a PCRE library upgrade, recompile stored values:

```sql
UPDATE my_table SET pcre_col = pcre_col::text::pcre;
```
