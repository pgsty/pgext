## Usage

Sources:

- [Official README](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/README.md)
- [Version 1.0 SQL function](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/pgsynck--1.0.sql)
- [Parser implementation](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/pgsynck.c)

`pgsynck` splits a text string into SQL statements, runs each statement through PostgreSQL's raw parser without executing it, and returns one diagnostic row per statement. It is useful for an early syntax screen, not for proving that SQL is executable.

### Check a Script

```sql
CREATE EXTENSION pgsynck;

SELECT sql, cursorpos, sqlerrcode, message, hint
FROM pgsynck($sql$
  SELECT 1;
  SELECT 1 FROM;
  INSERT INTO target_table(id) VALUES (1);
$sql$);
```

For a statement accepted by the raw parser, `cursorpos` and `sqlerrcode` are zero and `message` and `hint` are empty strings. A syntax error returns the parser cursor position, PostgreSQL's internally encoded integer error code, primary message, and optional hint. `sqlerrcode` is not the usual five-character SQLSTATE text.

### Validation Boundary

The parser does not resolve table or column names, check permissions, infer parameter types, run rewrite/planning, or execute functions and constraints. Thus a nonexistent relation or a type error can pass. Run later validation in a transaction against the correct schema and role when semantic assurance is required.

Version 1.0 implements its own statement splitter. It handles ordinary single/double quotes, dollar-quoted strings, and comments, but it predates many current PostgreSQL releases and is not a full replacement for the server lexer. Test nested comments, escape-string settings, tagged dollar quotes, procedural bodies, empty statements, and cursor positions before accepting untrusted scripts. Bound input size and statement count to avoid using a database backend as an unrestricted parsing service.

The control-file comment says the extension executes shell commands, but the reviewed C and SQL implementation only calls the PostgreSQL raw parser. Treat that mismatch as another sign that the abandoned release needs local source review and exact-version testing.
