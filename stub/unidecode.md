## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/alexkuz/pg_unidecode/blob/9cb9868332773b4551837eb4eaf48f2903d80fd2/README.md)
- [Version 0.0.6 SQL objects](https://github.com/alexkuz/pg_unidecode/blob/9cb9868332773b4551837eb4eaf48f2903d80fd2/unidecode--0.0.6.sql)
- [C implementation](https://github.com/alexkuz/pg_unidecode/blob/9cb9868332773b4551837eb4eaf48f2903d80fd2/unidecode.c)

`unidecode` transliterates Unicode text into an approximate ASCII representation using tables derived from the Python Unidecode project.

```sql
CREATE EXTENSION unidecode;
SELECT unidecode('Français, Русский, 漢語');
```

Transliteration is lossy, language-agnostic, and not reversible. Multiple source strings can map to the same ASCII output, culturally preferred spellings can differ, and results can change when tables change. Never use it alone for identity, uniqueness, authentication, legal names, or security canonicalization; preserve the original text and enforce identifiers independently.

Upstream explicitly describes the C code as early development and says not to use it in production. Version 0.0.6 has no current compatibility or memory-safety statement. Prefer maintained application libraries. If used for search hints in a legacy database, pin the exact tables, test every supported script and malformed UTF-8 boundary, bound input length, and rebuild generated columns and indexes after any upgrade.
