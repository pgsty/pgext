## Usage

Sources:

- [Official PL/V8 dialect documentation](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#dialects)
- [Official CoffeeScript example](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#coffeescript-example)
- [Official build rules](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/Makefile)

`plcoffee` is the optional CoffeeScript procedural-language dialect shipped by PL/V8. It compiles CoffeeScript function bodies for execution by the same V8-backed handler used by PL/V8. Version `3.1.10` is a legacy dialect: use it to maintain existing functions, not as the default language for new database code.

### Core Workflow

The server package must include the optional `plcoffee.control` and versioned SQL files. A superuser creates the language, after which permitted users can define functions in it:

```sql
CREATE EXTENSION plcoffee;

CREATE FUNCTION plcoffee_test(keys text[], vals text[])
RETURNS text AS $$
  return JSON.stringify(keys.reduce(((o, key, idx) ->
    o[key] = vals[idx]; return o), {}), {})
$$ LANGUAGE plcoffee IMMUTABLE STRICT;

SELECT plcoffee_test(ARRAY['name', 'age'], ARRAY['Tom', '29']);
```

The result is a JSON string assembled by the CoffeeScript body.

### Objects and Runtime

`CREATE EXTENSION plcoffee` installs the trusted language `plcoffee` plus its call, inline, and validator handler functions in `pg_catalog`. The control and SQL files are generated as part of the PL/V8 build and load the shared `$libdir/plv8` module; there is no separate CoffeeScript server library and no preload setting.

### Caveats

PL/V8 can be built with `DISABLE_DIALECT`, in which case the `plcoffee` extension files do not exist. Confirm the package includes the dialect and matches the installed PL/V8 binary/version. Extension creation is superuser-only even though the resulting procedural language is marked trusted.

CoffeeScript and its PL/V8 dialect are legacy surfaces. JavaScript execution consumes memory and CPU inside a PostgreSQL backend, so restrict who may define functions, review volatility/strictness declarations, and test upgrades with the exact PostgreSQL, PL/V8, V8, and CoffeeScript build combination.
