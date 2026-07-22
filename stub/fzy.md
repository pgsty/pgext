## Usage

Sources:

- [Official README](https://github.com/romgrk/pg_fzy/blob/0c0aab00004c738e86b031d618c215d900f1f5f4/README.md)
- [Extension SQL for 0.0.1](https://github.com/romgrk/pg_fzy/blob/0c0aab00004c738e86b031d618c215d900f1f5f4/fzy--0.0.1.sql)

`fzy` exposes the fzy fuzzy-matching score to SQL. Use it to rank a modest set of candidate strings against an abbreviated or partially typed query; it is a scoring function, not an index or a complete search engine.

### Core Workflow

Install the extension in the database, then calculate a score for each candidate and sort the result. The second argument must be converted to PostgreSQL's C-string pseudo-type.

```sql
CREATE EXTENSION fzy;

SELECT name,
       fzy('Stev', name::cstring) AS score
FROM users
ORDER BY score DESC;
```

Higher scores are better matches. Version 0.0.1 implements only the case-insensitive form of the fzy algorithm.

### SQL Object

- `fzy(cstring, cstring) RETURNS integer` computes the match score.

The function is declared strict, so a null argument produces null. It is also declared volatile by the extension SQL, which may limit planner optimizations. Because no index operator or access method is provided, PostgreSQL evaluates the function for every candidate row selected by the query.
