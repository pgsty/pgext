## Usage

Sources:

- [Official PL/V8 dialect documentation](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#dialects)
- [Official LiveScript example](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#livescript-example)
- [Official build rules](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/Makefile)

`plls` is the optional LiveScript procedural-language dialect shipped by PL/V8. It compiles LiveScript function bodies for the V8-backed PL/V8 handler. Version `3.1.10` is a legacy dialect intended mainly for compatibility with existing database functions.

### Core Workflow

The server package must include the optional `plls.control` and versioned SQL files. A superuser creates the language, after which permitted users can define LiveScript functions:

```sql
CREATE EXTENSION plls;

CREATE FUNCTION plls_test(keys text[], vals text[])
RETURNS text AS $$
  return JSON.stringify { [key, vals[idx]] for key, idx in keys }
$$ LANGUAGE plls IMMUTABLE STRICT;

SELECT plls_test(ARRAY['name', 'age'], ARRAY['Tom', '29']);
```

The result is a JSON string assembled by the LiveScript body.

### Objects and Runtime

`CREATE EXTENSION plls` installs the trusted language `plls` plus its call, inline, and validator handler functions in `pg_catalog`. Its generated control/SQL files load `$libdir/plv8`; LiveScript does not have a separate PostgreSQL shared library and needs no preload setting.

### Caveats

PL/V8 builds made with `DISABLE_DIALECT` omit the `plls` files. Confirm the installed package includes this dialect and matches the PL/V8 binary/version. A superuser must create the extension even though the resulting language is marked trusted.

LiveScript and this PL/V8 integration are legacy surfaces. Executed code consumes memory and CPU inside the database backend. Restrict function-definition privileges, review volatility and strictness declarations, and regression-test the exact PostgreSQL/PL/V8/V8/LiveScript combination before upgrades.
