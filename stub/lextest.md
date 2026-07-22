## Usage

Sources:

- [Extension control file at the reviewed commit](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.control)
- [Version 1.0 SQL objects](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest--1.0.sql)
- [Lexer implementation](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.c)

`lextest` exposes PostgreSQL's internal core SQL scanner for parser development and diagnostics. It is not a general text-search tokenizer: numeric token codes and payload rules come directly from PostgreSQL grammar internals and can change across major versions.

### Core Workflow

Use `lexer(cstring)` to inspect the token stream, or `lexer_test(cstring)` to exercise the scanner and discard the result.

```sql
CREATE EXTENSION lextest;

SELECT *
FROM lexer('SELECT 1 + $2 FROM public.demo');

SELECT lexer_test('UPDATE demo SET value = 42 WHERE id = 7');
```

`lexer(cstring)` returns `pos`, `token`, and `value`. `pos` is a byte location in the scanner input, not a character index. `token` is PostgreSQL's internal numeric grammar token. `value` is populated only for token classes for which this wrapper extracts a string or numeric payload; otherwise it is `NULL`. The function materializes all rows before returning them.

### Compatibility Boundary

The C code calls non-stable backend scanner functions and includes parser-internal headers. Build it separately for every exact PostgreSQL major version and compiler environment; a successful build for one server does not establish compatibility with another. Numeric token values are implementation details and should not be stored as a durable application protocol. Restrict `lextest` to extension-development or parser-debugging environments.
