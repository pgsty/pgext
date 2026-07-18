## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/jmckulk/pg_string_extension/blob/aaa0fa3d4c052ed8e35091c0eeb39cce1cbbf3a7/README.md)
- [Version 0.1 SQL objects](https://github.com/jmckulk/pg_string_extension/blob/aaa0fa3d4c052ed8e35091c0eeb39cce1cbbf3a7/pg_string_extension--0.1.sql)
- [C implementation](https://github.com/jmckulk/pg_string_extension/blob/aaa0fa3d4c052ed8e35091c0eeb39cce1cbbf3a7/pg_string_extension.c)

`pg_string_extension` is an abandoned educational C extension providing `echo_text(text)`, `is_palindrome(text)`, `sort_chars(text)`, and `is_anagram(text, text)`.

```sql
CREATE EXTENSION pg_string_extension;
SELECT is_palindrome('level'),
       sort_chars('postgres'),
       is_anagram('silent', 'listen');
```

The reviewed documentation does not define Unicode normalization, locale, case folding, whitespace, punctuation, or collation semantics. Therefore results should be treated as byte- or implementation-level demonstrations, not linguistically correct text processing. Test multibyte UTF-8, empty strings, combining characters, non-ASCII case, and long inputs before any reuse.

Version 0.1 has no current compatibility matrix, maintenance statement, or security review. Prefer PostgreSQL core string functions or a maintained language-aware library. If the sample functions are used in persisted generated values, indexes, or constraints, pin the binary and verify exact semantics across upgrades; a behavior change requires rebuilding dependent objects.
