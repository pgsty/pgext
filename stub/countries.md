## Usage

Sources:

- [Official database.dev package page](https://database.dev/raminder/countries)
- [Official database.dev installation guide](https://database.dev/docs/install-a-package)

`countries` is a database.dev Trusted Language Extension that creates a small lookup table named `countries`. Despite the package description, the published version `0.0.4` SQL contains only six rows, so it must not be treated as a complete or standards-based country dataset.

### Core Workflow

Install the `pg_tle` prerequisite, use the dbdev CLI to generate a migration for the exact package version, review that generated SQL, and apply it through your normal migration process.

```bash
dbdev add -o ./migrations -s public -v 0.0.4 package -n "raminder@countries"
```

After the migration creates the extension, inspect the table before relying on it.

```sql
TABLE public.countries ORDER BY name;
```

The published `0.0.4` install script defines one column and these values: Afghanistan, Albania, Algeria, Zambia, Zimbabwe, and India.

```sql
SELECT count(*) AS row_count
FROM public.countries;
```

### Important Objects

- `public.countries`: a table with a single `name text PRIMARY KEY` column.
- `pg_extension_config_dump`: the install script registers `public.countries` as extension configuration data so its contents are included by extension-aware dumps.

### Operational Notes

This package is distributed through database.dev and requires `pg_tle`; it is not a conventional native extension archive. Version `0.0.4` hard-codes the `public.countries` name even if the package is generated with another target schema. The six supplied names have no ISO codes, regions, localized names, lifecycle status, or update mechanism. If an application needs an authoritative country list, use a maintained ISO-derived dataset instead. Treat edits as application data and test dump, restore, and extension-upgrade behavior before changing the seeded rows.
