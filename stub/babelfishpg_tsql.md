## Usage

Sources:

- [Babelfish extensions BABEL_5_4_0 README](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/BABEL_5_4_0/README.md)
- [Installation guide](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/BABEL_5_4_0/INSTALLING.md.tmpl)
- [`babelfishpg_tsql` control file](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/BABEL_5_4_0/contrib/babelfishpg_tsql/babelfishpg_tsql.control.in)
- [Babelfish limitations](https://babelfishpg.org/docs/limitations/limitations-of-babelfish/)
- [Handling T-SQL](https://babelfishpg.org/docs/usage/handling-tsql/)

`babelfishpg_tsql` implements the T-SQL language and SQL Server-compatible catalog behavior used by Babelfish. It is one component of a Babelfish database, not a compatibility layer that can be added to stock PostgreSQL by itself: the complete stack requires the Babelfish-patched PostgreSQL engine plus the common, TDS, and T-SQL extensions.

### Core Workflow

Configure the TDS protocol extension for preload and restart the Babelfish server:

```conf
shared_preload_libraries = 'babelfishpg_tds'
```

Create the TDS extension with `CASCADE` so its extension dependencies, including `babelfishpg_tsql`, are installed. Choose the migration mode before initialization.

```sql
CREATE EXTENSION IF NOT EXISTS babelfishpg_tds CASCADE;

ALTER SYSTEM SET babelfishpg_tsql.database_name = 'babelfish_db';
ALTER SYSTEM SET babelfishpg_tsql.migration_mode = 'multi-db';

CALL sys.initialize_babelfish('babelfish_user');
```

After configuration is reloaded as directed by the installation guide, SQL Server clients connect to the TDS listener, commonly on port 1433, and issue T-SQL in the logical databases created by Babelfish.

### Component and Object Index

- `babelfishpg_tsql` supplies the T-SQL parser, procedural language, system objects, compatibility functions, and T-SQL configuration variables.
- `babelfishpg_tds` supplies the Tabular Data Stream listener and is the normal installation entry point.
- `babelfishpg_common` supplies shared data types and functions. It and `uuid-ossp` are declared dependencies of `babelfishpg_tsql`.
- `babelfishpg_money` supplies money-related compatibility objects used by the stack.
- `sys.initialize_babelfish(login_name)` provisions the Babelfish catalogs and initial login.
- `sys.sp_babelfish_configure` controls documented compatibility escape hatches.
- `babelfishpg_tsql.database_name` identifies the physical PostgreSQL database hosting Babelfish.
- `babelfishpg_tsql.migration_mode` selects `single-db` or `multi-db` logical-database mapping.

### Operational Boundaries

Installation requires superuser privileges and a Babelfish build matched to the extension release. Do not install `babelfishpg_tsql` alone and expect TDS connectivity. The migration mode is a provisioning decision and is not intended to be changed after the database is initialized.

Babelfish implements a substantial but incomplete SQL Server surface. Validate application syntax, data types, system-catalog assumptions, drivers, and escape-hatch settings against the official limitations before migration. PostgreSQL and T-SQL connections can observe different naming and transaction semantics.

The catalog change from 5.5.0 to 5.4.0 is a version correction to the official `BABEL_5_4_0` release line, not evidence of a new feature or an automatic downgrade procedure.
