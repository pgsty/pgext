## Usage

Sources:

- [README](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/README.md)
- [Control file](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/chronoturtle.control)
- [Makefile](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/Makefile)
- [sql/functions/chronoturtle.sql](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/sql/functions/chronoturtle.sql)

`chronoturtle` replaces ordinary tables with branch-aware views backed by append-only row history. Upstream explicitly describes 0.1.0 as an unmaintained feasibility study, tested only on PostgreSQL 16.

### Core Workflow

```sql
CREATE EXTENSION chronoturtle;
SET search_path TO chronoturtle, public;
CREATE TABLE public.contacts (id serial PRIMARY KEY, email text NOT NULL);
INSERT INTO public.contacts(email) VALUES ('alice@example.com');
SELECT migrate_to_vie('public.contacts');
SELECT set_commit_author('editor@example.com');
SELECT set_current_branch('feature/contact', 'main');
UPDATE public.contacts SET email = 'alice@example.org' WHERE id = 1;
SELECT set_current_branch('main');
SELECT * FROM public.contacts;
```

### Objects and Maintenance

`migrate_to_vie` converts named tables; `migrate_schemas_to_vie` handles schemas and orders foreign-key dependencies. `set_current_branch` selects a branch, while `set_commit_author` records the author of subsequent changes. The `chronoturtle` schema contains public functions, `_vie` holds metadata, and `_backing` retains every row version. `add_unique_constraint` and `drop_unique_constraint` manage supported view constraints.

History is not automatically pruned. `drop_vie_table` defaults to keeping all backing rows, including historical versions; it is not an export of only the visible branch snapshot. Its second argument can request data deletion.

### Limitations

This is a pure SQL extension with no shared library or preload. Test the conversion on a copy with an appropriately privileged table owner. Merge, rebase, checkout, and diff operations are not implemented. Schema changes, quoted mixed-case identifiers, and `TRUNCATE` on converted views are unsupported; there are no upgrade scripts.
