## Usage

Sources:

- [jsonbc--1.0.sql](https://github.com/akorotkov/jsonbc/blob/c59545546211f17cabb419d9424e33774308f242/jsonbc--1.0.sql)
- [Makefile](https://github.com/akorotkov/jsonbc/blob/c59545546211f17cabb419d9424e33774308f242/Makefile)
- [jsonbc.control](https://github.com/akorotkov/jsonbc/blob/c59545546211f17cabb419d9424e33774308f242/jsonbc.control)

`jsonbc` is a historical compressed JSONB-like type. Its versioned SQL provides a key dictionary, JSON-style operators and B-tree/hash/GIN operator classes.

### Core Workflow

```sql
CREATE EXTENSION jsonbc;
SELECT '{"name":"Alice"}'::jsonbc;
SELECT '{"name":"Alice"}'::jsonbc->>'name';
```

### Operational Boundaries

`get_id_by_name` and `get_name_by_id` access the extension’s key dictionary. JSON object/array access, containment, existence and expansion functions operate on the custom type. Treat its dictionary and stored values as one persistent data format; do not remove dictionary data independently.

Installation requires a superuser and the compiled C library. No preload is declared. The repository has no current support documentation, major-version matrix or license declaration in the reviewed files. This is source archaeology, not a validated substitute for native JSONB; test installation, round trips, dumps and restores before placing data in the type.
