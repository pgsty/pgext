## Usage

Sources:

- [README.md](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/README.md)
- [sql/lance_fdw--0.2.sql](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/sql/lance_fdw--0.2.sql)
- [LICENSE](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/LICENSE)
- [lance_fdw.control](https://github.com/MisterRaindrop/cloudberry-lance/blob/0663d3062150896d9deb9073c4caecc5b1e911b7/lance_fdw.control)

`lance_fdw` 0.2 reads Lance datasets from Apache Cloudberry segments in parallel. It is a Cloudberry-specific wrapper, with no verified vanilla PostgreSQL compatibility claim. The wrapper depends on the Lance C binding and supports a defined subset of Arrow types.

### Core Workflow

```sql
CREATE EXTENSION lance_fdw;
CREATE SERVER lance_local FOREIGN DATA WRAPPER lance_fdw
  OPTIONS (base_uri '/data/lance');
IMPORT FOREIGN SCHEMA lance LIMIT TO ("events.lance")
  FROM SERVER lance_local INTO public;
SELECT count(*) FROM public."events.lance";
```

### Operational Boundaries

`CREATE SERVER`, user mappings, foreign tables and `IMPORT FOREIGN SCHEMA` expose the data. Scans are read-only; supported predicates and nearest-neighbour operations can be pushed down when semantics agree. Unsupported types are rejected, and timestamps, collations and numeric values have documented pushdown limits. Dataset paths and credentials must be usable from the participating segments. After matching native files are installed, use `ALTER EXTENSION lance_fdw UPDATE` for the 0.1 to 0.2 upgrade; 0.2 records vector-distance operators already present in some development 0.1 builds. Native installation requires a superuser.
