## Usage

Sources:

- [README](https://api.pgxn.org/src/gsscode/gsscode-1.1.3/README.md)
- [Control file](https://api.pgxn.org/src/gsscode/gsscode-1.1.3/gsscode.control)
- [SQL](https://api.pgxn.org/src/gsscode/gsscode-1.1.3/gsscode--1.1.3.sql)

`gsscode` packs a nine-character UK ONS/GSS geography code into 32 bits. Version 1.1.3 includes the corrected prefix-search semantics: a prefix predicate is not equality and must not occupy the btree equality strategy.

### Core Workflow

```sql
CREATE EXTENSION gsscode;
CREATE TABLE areas (code gsscode PRIMARY KEY, label text);
INSERT INTO areas VALUES ('E01000001', 'Example area');
SELECT code, country(code), gss_type(code), area(code)
FROM areas
WHERE code >= gsscode_range_lower('E01')
  AND code < gsscode_range_upper('E01');
```

### Matching and Registry

`%` and `!%` remain Boolean prefix filters, including array forms, but no longer accelerate prefix matching through a btree index by themselves. Use `gsscode_range_lower` and `gsscode_range_upper` for a genuine half-open range. Invalid prefix lengths raise an error.

`is_valid_gss`, `country`, `gss_type`, and `area` inspect lexical form or packed components. `description` and `type_info` consult the private `gsscode_types` registry; `isnan` recognizes reserved area codes. Text-compatible regular expressions and `left` are available but do not automatically use the base type’s btree index.

### Upgrade and Boundaries

Version 1.1.3 renames `is_valid(text)` to `is_valid_gss(text)` to avoid collisions with other extensions. Update SQL callers after upgrading; the extension update preserves the function identity and dependencies. Version 1.1.2 also adds text equality and an assignment cast from text.

Existing 1.0.0 installations should run `ALTER EXTENSION gsscode UPDATE` after installing current files to remove the unsound operator-family entry. Replacing the library alone does not perform this catalog repair.

Valid input shape does not prove that ONS assigned the code. Registry refresh belongs to `gsscode_ons_refresh`, whose control version remains 1.0.0. The core extension is relocatable and needs no preload; the release does not declare a PostgreSQL-major matrix.
