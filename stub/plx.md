## Usage

Sources:

- [PGXN plx 2.0.1 README](https://pgxn.org/dist/plx/2.0.1/README.html)
- [plx 2.0.1 user guide](https://api.pgxn.org/src/plx/plx-2.0.1/doc/USERGUIDE.md)
- [plx 2.0.1 compatibility notes](https://api.pgxn.org/src/plx/plx-2.0.1/doc/COMPATIBILITY.md)
- [plx 2.0.1 changelog](https://api.pgxn.org/src/plx/plx-2.0.1/CHANGELOG.md)
- [plx control file](https://api.pgxn.org/src/plx/plx-2.0.1/plx.control)

`plx` provides familiar procedural-language dialects that transpile to ordinary PL/pgSQL when `CREATE FUNCTION` runs. PostgreSQL stores and executes the generated PL/pgSQL with its built-in trusted handler; no Ruby, PHP, JavaScript, Python, Go, COBOL, Oracle, or SQL Server runtime is loaded into the backend.

```sql
CREATE EXTENSION plx;
```

### Available Dialects

| Language | Surface syntax |
|---|---|
| `plxruby` | Ruby |
| `plxphp` | PHP |
| `plxjs` | JavaScript |
| `plxts` | TypeScript annotations over the JavaScript dialect |
| `plxpython3` | Python 3 |
| `plxgo` | Go |
| `plxcobol` | ISO COBOL |
| `plxplsql` | Oracle PL/SQL |
| `plxtsql` | Transact-SQL |

All dialects target the same PL/pgSQL statement surface, including assignments, conditionals, loops, query iteration, dynamic SQL, cursors, exceptions, triggers, and set-returning functions.

### Create a Function

Choose a dialect in the `LANGUAGE` clause while keeping the function signature in PostgreSQL types:

```sql
CREATE FUNCTION grade(score integer)
RETURNS text
LANGUAGE plxruby
AS $$
  grade #:: text
  if score >= 90
    grade = "A"
  elsif score >= 80
    grade = "B"
  else
    grade = "F"
  end
  return grade
$$;

SELECT grade(85);
```

Translation happens once, when the function is created. The executable body stored in `pg_proc.prosrc` is regular PL/pgSQL, so it can be dumped, reviewed, and run without a separate interpreter.

### Inspect and Debug Generated Code

```sql
SELECT pg_get_functiondef('grade(integer)'::regprocedure);
SELECT prosrc
FROM pg_proc
WHERE oid = 'grade(integer)'::regprocedure;

SELECT plx_source('grade(integer)'::regprocedure);
```

Runtime error line numbers refer to generated PL/pgSQL. `plx_source()` recovers the original embedded dialect body; use it together with `pg_get_functiondef()` when correlating an error with the source.

### SQL and String Building

Expressions retain PostgreSQL SQL semantics rather than emulating a complete source-language runtime. Use each dialect's query/execute form for SQL and explicit PostgreSQL types for non-literal expressions. The `plx_strbuild` expanded-object helper accelerates repeated string appends on PostgreSQL 18:

```sql
CREATE FUNCTION labels(n integer)
RETURNS text
LANGUAGE plxjs
AS $$
  let out: text = "";
  for (let i = 1; i <= n; i++) {
    out += `item-${i},`;
  }
  return out;
$$;
```

The builder remains correct on PostgreSQL 13-17, but its in-place optimization requires PostgreSQL 18.

### Boundaries and Caveats

- plx implements syntax surfaces, not the source languages' runtimes: there are no gems, Python modules, JavaScript imports, Go goroutines, PHP classes, Oracle packages, or SQL Server transaction commands.
- Functions run in PL/pgSQL's trusted sandbox, with no direct filesystem, network, arbitrary native-code, or transaction-control access.
- Parameters and return types must be PostgreSQL types. Type inference for locals is limited; explicitly declare types for calls and compound expressions.
- SQL uses three-valued logic and PostgreSQL numeric/string semantics. Source-language truthiness and string concatenation with `+` are not reproduced.
- Locals are hoisted into one PL/pgSQL `DECLARE` block, so block-local scope and redeclaration with a different type are unavailable.
- Distribution 2.0.1 installs extension version 2.0.0; 2.0.1 changes PGXN packaging only and has no SQL upgrade of its own. Version 2.0.0 changes interpolation in `plxruby`, `plxphp`, `plxjs`, `plxts`, `plxpython3`, and `plxgo`: a NULL operand now makes the whole interpolated string NULL instead of silently becoming an empty string. Run `ALTER EXTENSION plx UPDATE TO '2.0.0'`, then redeploy affected functions because existing generated PL/pgSQL in `pg_proc.prosrc` is not rewritten by the extension update.
