## Usage

Sources:

- [PGXN 0.4.1](https://pgxn.org/dist/pg_grammar_guard/0.4.1/)
- [License](https://github.com/Manuelreyesbravo/pg_grammar_guard/blob/2d55b1b58c24913195dcb9de2126dc8d3114c4fb/LICENSE)
- [Changelog](https://github.com/Manuelreyesbravo/pg_grammar_guard/blob/2d55b1b58c24913195dcb9de2126dc8d3114c4fb/CHANGELOG.md)

`pg_grammar_guard` generates GBNF or JSON Schema from catalog identifiers and detects drift from approved grammars. It is pure SQL, needs no preload and is packaged for PostgreSQL 14–18. Its control file requires `pg_living_assertions`.

### Generate a grammar

```sql
CREATE EXTENSION pg_grammar_guard CASCADE;
CREATE TABLE public.grammar_demo (id integer, label text);
SELECT grammar_guard.grammar_for_json(ARRAY[
  ROW('column', 'enum',
      grammar_guard.catalog_columns('public.grammar_demo'), true)
]::grammar_guard.grammar_field[]);
```

The catalog helpers enumerate actual tables, columns and enum labels. Generation supports nested objects and bounded arrays; open-ended values such as arbitrary SQL or file paths need separate validation.

### Detect changes

`grammar_guard.watch()` stores the query that rebuilds a grammar, and `grammar_guard.check_grammar()` evaluates it against the live catalog. Results and their age are available through `living_assertions.status`.

Register baseline SQL only through trusted administrators. A grammar limits valid identifiers; it cannot prove that a selected table, join or answer is semantically correct. Baselines from the old 0.2 series lack the original generating query and require manual re-approval when upgrading.

### License Change in 0.4.4

Version 0.4.4 adopts Apache-2.0. Previously published versions through 0.4.3 retain the PostgreSQL License. This release changes no database objects; its upgrade only records the version.
