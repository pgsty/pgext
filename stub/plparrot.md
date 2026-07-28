## Usage

Sources:

- [Official project page](https://pgxn.org/dist/plparrot/)
- [Official upstream README](https://github.com/leto/plparrot/blob/5adaf4be2d00d8ca0aee06a6dbd1cec21ceff12a/README.md)
- [Official language-definition SQL](https://github.com/leto/plparrot/blob/5adaf4be2d00d8ca0aee06a6dbd1cec21ceff12a/plparrot.sql.in)

`plparrot` 0.4.0 embeds the Parrot virtual machine in PostgreSQL. Its legacy install script creates trusted and untrusted Parrot/PIR language names and, when Rakudo is available, Perl 6 language handlers.

### Core Workflow

Build the module against the target PostgreSQL and Parrot installations, then run its installed SQL script in the intended database:

```sh
psql -d appdb -f "$(pg_config --sharedir)/contrib/plparrot.sql"
```

The distribution predates modern extension packaging; do not use `CREATE EXTENSION plparrot`.

### Installed Languages

- `plparrot` and its alias `plpir` are created as trusted languages.
- `plparrotu` and `plpiru` are the untrusted variants.
- `plperl6` and `plperl6u` are created when the corresponding Rakudo support is built.

### Requirements and Caveats

- The reviewed control, registry, or catalog evidence identifies version `0.4.0`.
- The upstream test workflow expects `plpgsql`; runtime language support requires compatible Parrot and optionally Rakudo libraries.
- The upstream distribution uses a legacy or non-control installation layout; do not assume modern `ALTER EXTENSION UPDATE` behavior.
- Release 0.4.0 dates from 2011, and the reviewed repository's latest commit is historical. Verify whether its VM, Perl 6 naming, C API, and security assumptions are viable before attempting a modern build.
