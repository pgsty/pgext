## Usage

Sources:

- [Project README](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/README.md)
- [Version 1.0 installation SQL](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/pg_strverscmp--1.0.sql)
- [C comparison wrapper](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/pg_strverscmp.c)
- [Build definition](https://github.com/singularita/pg_strverscmp/blob/266395a35003dea9763bf25a584c4771f1a21552/Makefile)

`pg_strverscmp` 1.0 provides a natural, version-like ordering for text, so numeric runs compare by value rather than purely character by character. It exposes `pg_strverscmp(text,text)`, six comparison operators, and the `strverscmp_ops` B-tree operator class.

### Natural ordering

```sql
CREATE EXTENSION pg_strverscmp;

SELECT value
FROM (VALUES ('item2'), ('item10'), ('item1')) AS v(value)
ORDER BY value USING +<;
```

The operators are `+<`, `+<=`, `+=`, `+<>`, `+>=`, and `+>`. To support indexed searches with the same semantics, declare a text index with `strverscmp_ops`; ordinary text indexes use PostgreSQL's normal collation and operator class instead.

### Important limitations

The build links against `libunistring`. Upstream describes the comparator as Unicode-aware and generally case-insensitive, but explicitly warns that it is not binary-safe around a NUL byte, does not use PostgreSQL collation facilities, and is inefficient because it repeatedly invokes Unicode collation while comparing characters. Treat this as a specialized ordering, not a drop-in replacement for locale-aware text equality. Keep query operators and index operator classes consistent, and verify ordering after library or platform changes.
