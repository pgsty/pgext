---
title: "babelfishpg_tsql"
linkTitle: "babelfishpg_tsql"
description: "SQL Server Transact SQL compatibility"
weight: 9310
categories: ["SIM"]
width: full
---

[**babelfish**](https://babelfishpg.org/) : SQL Server Transact SQL compatibility


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9310** | {{< badge content="babelfishpg_tsql" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_tsql" "babelfish" >}} | `5.4.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "babelfishpg_common" >}} {{< ext "uuid-ossp" >}} |
|    **Need By**    | {{< ext "babelfishpg_tds" >}} |
|   **See Also**    | {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "tds_fdw" >}} {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "db_migrator" >}} |
|    **Siblings**   | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tds" >}} {{< ext "babelfishpg_money" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `babelfish` | `babelfishpg_common`, `uuid-ossp` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4.0` | {{< bg "18" "babelfish-18" "green" >}} {{< bg "17" "babelfish-17" "green" >}} {{< bg "16" "babelfish-16" "red" >}} {{< bg "15" "babelfish-15" "red" >}} {{< bg "14" "babelfish-14" "red" >}} | `babelfish-$v` | `antlr4-runtime413` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4.0` | {{< bg "18" "babelfish-18" "green" >}} {{< bg "17" "babelfish-17" "green" >}} {{< bg "16" "babelfish-16" "red" >}} {{< bg "15" "babelfish-15" "red" >}} {{< bg "14" "babelfish-14" "red" >}} | `babelfish-$v` | `libantlr4-runtime413` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "PIGSTY 6.0.0" "babelfish-18 : FORK 1" >}}      |      {{< bg "PIGSTY 5.4.0" "babelfish-17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish-14 : FORK 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://babelfishpg.org/" title="Repository" icon="link" subtitle="babelfishpg.org/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="babelfish-17-17.7-5.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg babelfish;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install babelfish;		# install via package name, for the active PG version
pig install babelfishpg_tsql;		# install by extension name, for the current active PG version

pig install babelfishpg_tsql -v 18;   # install for PG 18
pig install babelfishpg_tsql -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_tsql CASCADE; -- requires babelfishpg_common, uuid-ossp
```

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
