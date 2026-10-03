## Usage

Sources:

- [cc/postgres/anon_func.control](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/anon_func.control)
- [cc/postgres/README.md](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/README.md)
- [cc/postgres/anon_func--1.0.0.sql](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/anon_func--1.0.0.sql)
- [cc/postgres/install_extension.sh](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/install_extension.sh)

`anon_func` exposes anonymous count, sum, mean, variance and related aggregates from Google differential-privacy. The extension SQL version is 1.0.0; it is distinct from the repository-wide library release number.

### Core Workflow

```sql
CREATE EXTENSION anon_func;
SELECT anon_count(value, 1.0)
FROM (VALUES (1), (2), (3)) AS sample(value);
```

### Operational Boundaries

The official setup targets PostgreSQL 12 and requires the compiled C++ library; no wider compatibility is asserted here. Installation is privileged and no preload is specified. Epsilon and optional bounds are literal aggregate parameters. Each input entry is assumed to belong to a distinct user, so callers must bound contributions before aggregation. Empty groups return NULL and are explicitly not differentially private in this implementation. Repeated queries consume privacy budget; these wrappers do not by themselves establish an end-to-end privacy policy.
