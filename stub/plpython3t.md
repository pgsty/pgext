## Usage

Sources:

- [Official extension SQL](https://github.com/pykello/plpythont/blob/3b05e9ffbe405d7768bec8fc3a0764642ceff125/plpython3t--1.0.sql)
- [Official extension control file](https://github.com/pykello/plpythont/blob/3b05e9ffbe405d7768bec8fc3a0764642ceff125/plpython3t.control)

`plpython3t` registers PostgreSQL's existing PL/Python 3 handlers as a trusted procedural language. It does not add a sandbox or a restricted Python runtime; it changes the database privilege boundary around the same server-side Python interpreter.

### Core Workflow

The PL/Python 3 runtime and handler functions must already be installed, normally through `plpython3u`, before the wrapper language is created.

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION plpython3t;

CREATE FUNCTION py_add(a integer, b integer)
RETURNS integer
LANGUAGE plpython3t
AS $$
return a + b
$$;

SELECT py_add(2, 3);
```

### Installed Object

The extension runs `CREATE TRUSTED PROCEDURAL LANGUAGE plpython3t` with `plpython3_call_handler`, `plpython3_inline_handler`, and `plpython3_validator`. Its control file points at the normal `plpython3` shared library. There are no additional SQL functions, GUCs, or preload requirements.

### Security Boundary

PostgreSQL marks the standard PL/Python language untrusted because Python code can access server files, processes, libraries, and the operating-system environment with the database server's privileges. `plpython3t` merely labels that same handler stack trusted, allowing ordinary roles with the relevant database and schema privileges to create functions in it.

Installing this extension therefore grants a much stronger capability than its small SQL file suggests. Use it only in a tightly controlled database where every role able to create `plpython3t` functions is trusted at the operating-system level. Do not treat the mruby-style sandboxing of unrelated extensions as applicable here.
