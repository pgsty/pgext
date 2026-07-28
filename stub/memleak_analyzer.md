## Usage

Sources:

- [Official upstream README](https://github.com/vostralis/pg-memleak-analyzer/blob/f629e05e3fe7f0f67b26ab35974a2d476cc4c631/README.md)
- [Official extension control file (memleak_analyzer.control)](https://github.com/vostralis/pg-memleak-analyzer/blob/f629e05e3fe7f0f67b26ab35974a2d476cc4c631/memleak_analyzer.control)
- [Official extension SQL (memleak_analyzer--1.0.sql)](https://github.com/vostralis/pg-memleak-analyzer/blob/f629e05e3fe7f0f67b26ab35974a2d476cc4c631/memleak_analyzer--1.0.sql)

`memleak_analyzer` — pg-memleak-analyzer is a diagnostic tool and a PostgreSQL extension designed to analyze and localize logical memory leaks inside PostgreSQL backend sessions and background workers. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION memleak_analyzer;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `memleak_analyzer.analyze_bgw(target_pid INTEGER, observation_interval INTEGER)` is an extension function and returns `TABLE`.
- `memleak_analyzer.analyze_query(query text)` is an extension function and returns `TABLE`.
- `memleak_analyzer.get_bgw_snapshot(target_pid INTEGER)` is an extension function and returns `TABLE`.
- `memleak_analyzer` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
