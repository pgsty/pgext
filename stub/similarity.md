## Usage

Sources:

- [Version 1.0 README](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/README.md)
- [Version 1.0 SQL declarations](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity--1.0.sql)
- [PostgreSQL and ICU wrapper](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity_pg.c)
- [ICU build dependency](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/Makefile)
- [Extension control file](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity.control)

`similarity` 1.0 compares two strings with an fstrcmp-style approximate matching algorithm and returns a normalized `double precision` score from 0 for entirely different strings to 1 for identical strings. Use it for fuzzy candidate comparison, not linguistic search or indexed nearest-neighbor retrieval.

### Core Workflow

```sql
CREATE EXTENSION similarity;

SELECT similarity('similarity', 'distinction');

SELECT name, similarity(name, 'postgresql') AS score
FROM projects
ORDER BY score DESC
LIMIT 10;
```

The extension creates only two functions:

- `similarity(text,text)` performs a complete comparison, equivalent to a zero cutoff.
- `similarity(text,text,double precision)` accepts a minimum score and can stop early once the candidate cannot meet that cutoff.

### Cutoff Semantics

The three-argument function's result is exact only when it reaches the requested cutoff. If it returns below the cutoff, upstream explicitly says the numeric value is invalid, though it is guaranteed to be below that threshold. Use the cutoff form as a yes/no pruning optimization or with a current best score; use `similarity(text,text)` when ranking all returned numeric values.

Both functions are declared `STRICT` and `IMMUTABLE`. They do not define operators, an index operator class, a selectivity estimator, or a search index. A query that calls the function for every row remains a CPU and memory scan unless another predicate narrows the candidate set.

### Encoding, Dependency, and Compatibility Boundaries

The C wrapper converts PostgreSQL text with ICU's UTF-8 conversion routine and links to `libicu`. Use a `UTF8` database and test multilingual, combining-character, empty, and very long inputs. The implementation does not robustly surface ICU conversion errors, and similarity is based on its character sequence algorithm rather than locale collation or semantic language rules.

Version 1.0 has no release series or current PostgreSQL compatibility matrix. Keep the ICU runtime compatible with the binary, and validate builds plus regression results on the exact PostgreSQL major. Bound input length for untrusted callers because approximate string comparison work and memory grow with both strings.
