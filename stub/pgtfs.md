## Usage

Sources:

- [Official upstream README](https://github.com/adrianprelipcean/pgtfs/blob/8e2512ce6dd95db1f74b513e366a588fc6ae8ceb/readme.md)
- [Official extension control file (pgtfs.control)](https://github.com/adrianprelipcean/pgtfs/blob/8e2512ce6dd95db1f74b513e366a588fc6ae8ceb/pgtfs.control)
- [Official extension SQL (pgtfs--0.0.1.sql)](https://github.com/adrianprelipcean/pgtfs/blob/8e2512ce6dd95db1f74b513e366a588fc6ae8ceb/sql/pgtfs--0.0.1.sql)

`pgtfs` — PGTFS is a PostgreSQL extension designed to facilitate routing on top of the GTFS (General Transit Feed Specification) format. It provides functionality to query and analyze public transportation data stored in a PostgreSQL database using GTFS feeds. Use it for the corresponding spatial data or geospatial workflow. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION pgtfs;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgtfs_csa(origin TEXT, destination TEXT, departure_time DOUBLE PRECISION, network TEXT)` is an extension function and returns `TABLE`.
- `pgtfs_csa(origin TEXT, destination TEXT, departure_time DOUBLE PRECISION, network TEXT, minimize_transfers BOOLEAN DEFAULT FALSE)` is an extension function and returns `TABLE`.
- `pgtfs_raptor(origin TEXT, destination TEXT, departure_time DOUBLE PRECISION, network TEXT, max_rounds INT DEFAULT 5)` is an extension function and returns `TABLE`.
- `pgtfs_version()` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.4`.
- The control file marks the extension as relocatable.
- Upstream explicitly says the project is not production-ready.
- Upstream describes the project as a work in progress.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
