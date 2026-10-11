## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_grammar_guard/pg_grammar_guard-0.4.8/README.md)
- [pg_grammar_guard.control](https://api.pgxn.org/src/pg_grammar_guard/pg_grammar_guard-0.4.8/pg_grammar_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_grammar_guard/pg_grammar_guard-0.4.8/CHANGELOG.md)

`pg_grammar_guard` 0.4.8 creates GBNF or JSON Schema from catalog identifiers and watches approved grammars for drift. It is pure SQL and requires `pg_living_assertions` 0.5.9 or later.

### Core Workflow

```sql
CREATE EXTENSION pg_grammar_guard CASCADE;
CREATE TABLE public.grammar_demo (id integer, label text);
SELECT grammar_guard.grammar_for_json(ARRAY[
  ROW('column', 'enum',
      grammar_guard.catalog_columns('public.grammar_demo'), true)
]::grammar_guard.grammar_field[]);
```

### Operational Boundaries

`grammar_guard.catalog_tables()` and `grammar_guard.catalog_columns()` expose catalog names; `grammar_guard.grammar_for_json()` assembles typed fields, nested structures and bounded arrays. `grammar_guard.watch()` records the SQL that regenerates a grammar, and `grammar_guard.check_grammar()` compares it with the approved state in `living_assertions.status`. The generating query resolves with its author search path and executes as that author. Input NULLs, empty/duplicate field names, empty pivots and invalid bounds are rejected; `max_items` is capped at 1000. Version 0.4.8 changes canonical fingerprints, and the ordered catalog list can make an existing watch fail once: inspect and re-approve it. A grammar controls allowed tokens, not semantic correctness or authorization. Register trusted SQL only. No preload is required. Current Apache-2.0 licensing begins at 0.4.4; earlier releases retain their original license.
