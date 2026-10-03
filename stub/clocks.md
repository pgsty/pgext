## Usage

Sources:

- [README.md](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/README.md)
- [schema/sql/01-types.sql](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/sql/01-types.sql)
- [schema/clocks/clocks.control](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/clocks/clocks.control)
- [schema/clocks/clocks--1.0.sql](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/clocks/clocks--1.0.sql)
- [schema/clocks/clocks.c](https://github.com/nuno-faria/crdv/blob/15e58bb9379e77a77de56517ff5bf6a514ff56d4/schema/clocks/clocks.c)

`clocks` 1.0 supplies C helpers for CRDV vector clocks and hybrid logical clocks. CRDV is a research prototype tested on PostgreSQL 16. This component does not provide replication or create the prerequisite types itself.

### Core workflow

```sql
-- In a fresh database, matching the CRDV bootstrap:
CREATE TYPE hlc AS (physical_time bigint, logical_time int);
CREATE DOMAIN vclock AS bigint[];
CREATE EXTENSION clocks;
SELECT vclock_lte(ARRAY[1,2]::vclock, ARRAY[2,3]::vclock);
SELECT vclock_max(ARRAY[1,3]::vclock, ARRAY[2,2]::vclock);
SELECT next_hlc(ROW(0,0)::hlc);
```

### Prerequisites and functions

Use the existing `hlc` and `vclock` definitions when initializing a CRDV database; do not redefine live types. They must be visible during extension creation. `vclock_lte` checks component-wise ordering, `vclock_max` returns component-wise maxima, and `next_hlc` advances the hybrid clock using the current time and prior logical state. Keep clock dimensions and ordering conventions consistent with the application.

Installation requires a superuser, and no shared preload is declared. Full CRDV cluster initialization and logical-replication settings belong to the surrounding CRDV system, not to these three functions. Preserve its type definitions when backing up or restoring the extension.
