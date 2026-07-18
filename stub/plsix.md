## Usage

Sources:

- [README at the reviewed commit](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/README.md)
- [Extension control file](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/plsix.control)
- [Language installation SQL](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/plsix--2.sql)
- [Language handler implementation](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/plsix.c)

`plsix` is an untrusted PostgreSQL procedural-language handler for Perl 6, the language now known as Raku. It derives from PL/sh: a function body is copied to a temporary executable script, arguments are passed to the child process, and standard output is converted to the declared PostgreSQL return type.

### Stored Function

The first line selects the installed interpreter. This historical project expects a `perl6` executable.

```sql
CREATE EXTENSION plsix;

CREATE FUNCTION plsix_echo(integer)
RETURNS text
LANGUAGE plsix
AS $$
#!/usr/bin/env perl6
say "argument is " ~ @ARGS[0];
$$;

SELECT plsix_echo(7);
```

An empty standard-output stream returns SQL `NULL`; standard error and nonzero child status are reported as errors. The handler also supports `DO`, trigger, and event-trigger entry points. Trigger context is exported through inherited `PLSH_TG_*` environment variables.

### Caveats

- This is an untrusted language that executes arbitrary operating-system code with the PostgreSQL server account's privileges. Only superusers can create functions in it; audit every function body and interpreter path.
- The handler has no SPI integration. A script that needs database access must open a separate client connection, and recursive `psql` calls have independent transaction behavior.
- Trigger scripts receive context but cannot return a modified row through this interface. Do not use `plsix` where row-rewriting trigger semantics are required.
- The README retains copied PL/sh names and installation examples, while the actual extension and language names are `plsix`. Use the version `2` install SQL as the API authority.
- The source was last updated in 2019, targets the old `perl6` executable name, and publishes no current PostgreSQL or Raku compatibility matrix. Prefer a maintained procedural language for production code.
