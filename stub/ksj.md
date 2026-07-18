## Usage

Sources:

- [Official extension control file](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj.control)

`ksj` — Japanese kanji-numeral integer type with arithmetic, aggregates, casts, comparison, and B-tree indexing

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "ksj";
SELECT extversion
FROM pg_extension
WHERE extname = 'ksj';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
