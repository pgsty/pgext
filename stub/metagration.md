## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/metagration/metagration-1.0.4/README.md)
- [Official extension SQL (metagration.sql)](https://api.pgxn.org/src/metagration/metagration-1.0.4/sql/metagration.sql)

`metagration` — Migrating logically replicated PostgreSQL databases is a delicate dance of applying the right script at the right time, and enduring possible downtime making sure replicas are correctly up to date. Consider the *warnings from the documentation*:. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Important Objects

- `metagration.check_script_trigger()` is an extension function and returns `TRIGGER`.
- `metagration.current_revision()` is an extension function and returns `bigint`.
- `metagration.export(replace_scripts boolean=false, transactional boolean=false, run_migrations boolean=false)` is an extension function and returns `text`.
- `metagration.new_script(up_script text, down_script text=null, up_declare text=null, down_declare text=null, args jsonb='{}', use_schema text='metagration_scripts', comment text=null)` is an extension function and returns `bigint`.
- `metagration.next_revision(from_revision bigint=null)` is an extension function and returns `bigint`.
- `metagration.previous_revision(from_revision bigint=null)` is an extension function and returns `bigint`.
- `metagration.run` is an extension procedure.
- `metagration.run_down` is an extension procedure.
- `metagration.run_up` is an extension procedure.
- `metagration.log` is a table installed or managed by the extension.
- `metagration.script` is a table installed or managed by the extension.
- `metagration` is a schema created by the extension.
- `metagration_scripts` is a schema created by the extension.

### Requirements and Caveats

- The catalog records version `1.0.4`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
