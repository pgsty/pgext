## Usage

Sources:

- [readme.md](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [packages/measurements/sql/measurements--0.0.1.sql](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/measurements/sql/measurements--0.0.1.sql)
- [packages/measurements/Makefile](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/measurements/Makefile)
- [packages/measurements/measurements.control](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/measurements/measurements.control)

`measurements` installs the seeded `measurements.quantities` reference table. Its columns describe quantity names, labels, units and explanatory text.

### Core Workflow

```sql
CREATE EXTENSION "measurements" CASCADE;
SELECT name, unit, description FROM measurements.quantities WHERE name = 'Length';
```

### Operational Boundaries

This historical SQL distribution supplies reference rows, not a unit-conversion engine or dimensional type system. Inspect seed values before using them as authoritative domain data; some descriptive entries are malformed. No current PostgreSQL-major matrix or license declaration was found in this source revision.

Install its declared dependencies first: `plpgsql`, `uuid-ossp`, `launchql-extension-verify`. This is SQL/PLpgSQL code with no own shared library or preload. The control permits non-superuser installation but is not marked trusted; dependency and schema privileges still apply.
