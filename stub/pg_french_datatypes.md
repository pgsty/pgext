## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_french_datatypes/pg_french_datatypes-0.1.1/README)
- [Official extension SQL (pg_french_datatypes.sql)](https://api.pgxn.org/src/pg_french_datatypes/pg_french_datatypes-0.1.1/sql/pg_french_datatypes.sql)

`pg_french_datatypes` — This small projects aims at including french-centric data type, such as :. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Important Objects

- `jour` is an extension-defined type.
- `mois` is an extension-defined type.
- `code_postal_fr` is an extension-defined domain.
- `numero_securite_sociale_fr` is an extension-defined domain.

### Requirements and Caveats

- The catalog records version `0.1.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
