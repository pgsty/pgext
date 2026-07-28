## Usage

Sources:

- [Official upstream README](https://github.com/ffeller/mastermobilitydb/blob/60004d6b8f5ac491dcbad17f3a7f5bf0aab69e00/readme.txt)
- [Official extension control file (mastermobilitydb.control)](https://github.com/ffeller/mastermobilitydb/blob/60004d6b8f5ac491dcbad17f3a7f5bf0aab69e00/mastermobilitydb.control)
- [Official extension SQL (mastermobilitydb--1.0.0.sql)](https://github.com/ffeller/mastermobilitydb/blob/60004d6b8f5ac491dcbad17f3a7f5bf0aab69e00/mastermobilitydb--1.0.0.sql)

`mastermobilitydb` — Schema, types, and utilities for managing moving-object trajectories with MobilityDB. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION mastermobilitydb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ASPECT_ATTRIBUTE_COUNT()` is an extension function and returns `INTEGER`.
- `ASPECT_ATTRIBUTE_FIND_ALL()` is an extension function and returns `SETOF`.
- `ASPECT_ATTRIBUTE_FIND_BY_ID(IN P_ASPECT_ID INTEGER, IN P_ATTRIBUTE_ID INTEGER)` is an extension function and returns `SETOF`.
- `ASPECT_COUNT()` is an extension function and returns `INTEGER`.
- `ASPECT_FIND_ALL()` is an extension function and returns `SETOF`.
- `ASPECT_FIND_BY_ID(IN P_ASPECT_ID INTEGER)` is an extension function and returns `SETOF`.
- `ASPECT_FIND_BY_NAME(IN P_DESCRIPTION VARCHAR(50))` is an extension function and returns `SETOF`.
- `ASPECT_TYPE_COUNT()` is an extension function and returns `INTEGER`.
- `ASPECT_TYPE_FIND_ALL()` is an extension function and returns `SETOF`.
- `ASPECT_TYPE_FIND_BY_ID(IN P_ASPECT_TYPE_ID INTEGER)` is an extension function and returns `SETOF`.
- `ASPECT_TYPE_FIND_BY_NAME(IN P_DESCRIPTION VARCHAR(50))` is an extension function and returns `SETOF`.
- `ATTRIBUTE_COUNT()` is an extension function and returns `INTEGER`.
- `ATTRIBUTE_FIND_ALL()` is an extension function and returns `SETOF`.
- `ATTRIBUTE_FIND_BY_ID(IN P_ATTRIBUTE_ID INTEGER)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Install the confirmed extension dependencies first: `file_fdw`, `postgis`, `mobilitydb`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
