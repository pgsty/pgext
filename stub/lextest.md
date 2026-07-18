## Usage

Sources:

- [Extension control file at the reviewed commit](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.control)
- [Version 1.0 SQL definitions](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest--1.0.sql)
- [Lexer implementation](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.c)

`lextest` is a tiny development and benchmarking extension that invokes PostgreSQL's core SQL lexer. `lexer(cstring)` returns each token's `pos`, numeric `token`, and decoded `value`; `lexer_test(cstring)` scans the input without returning the token stream.

```sql
CREATE EXTENSION lextest;

SELECT *
FROM lexer('SELECT 1 + 2');

SELECT lexer_test('SELECT 1 + 2');
```

### Caveats

This is a parser-development aid, not an application text-analysis API. It includes PostgreSQL internal parser headers and depends on private token definitions, so it may require source changes for a different major release. Upstream publishes no README, release tags, compatibility matrix, or license file.
