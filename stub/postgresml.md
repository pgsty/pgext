## Usage

Sources:

- [Official README](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/README.md)
- [Extension control file](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/postgresml.control)
- [Installed extension SQL](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/postgresml--1.0.sql)
- [Separate naive Bayes demonstration](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/naive_bayes.sql)
- [Build manifest](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/Makefile)

`postgresml` 1.0 is not a general machine-learning platform. The installed extension exposes one helper, `ts_lexemes(tsvector)`, while the repository's binary naive Bayes spam classifier is a separate demonstration script that is not installed by `CREATE EXTENSION`.

### Core Workflow

```sql
CREATE EXTENSION postgresml;

SELECT lexeme, n
FROM ts_lexemes(to_tsvector('english', 'red fox red dog'))
ORDER BY lexeme;
```

`ts_lexemes` returns one row per lexeme in the input `tsvector`. Its `n` field is the number of stored positional entries for that lexeme. It is useful for expanding a text-search vector into relational rows, but positional counts can be absent or altered when a vector was stripped or constructed differently; do not assume `n` is always the original raw-token frequency.

### Demonstration Script Boundary

The README's `test_naive_bayes(0.3)` workflow belongs to `naive_bayes.sql`, not the extension install script listed by the Makefile. That demonstration creates generic tables and functions, expects the `tablefunc` `crosstab` facility, uses psql file-loading commands, and expects an external SMS Spam Collection data file that is not present in the pinned repository. It also chooses a random training split, so results are not deterministic.

Review and adapt the demonstration explicitly rather than assuming it appears after extension installation. Use controlled schemas, versioned data, a deterministic split, held-out evaluation, and application-specific privacy and retention controls for any experiment. The C helper accesses PostgreSQL full-text-search internals and the pinned source dates from 2013, so current-major compatibility is not established; compile, test, and verify dump/restore against the exact server major. No preload is required.
