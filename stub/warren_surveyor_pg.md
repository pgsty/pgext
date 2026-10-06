## Usage

Sources:

- [warren-surveyor-pg/README.md](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/README.md)
- [warren-surveyor-pg/Cargo.toml](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/Cargo.toml)
- [warren-surveyor-pg/src/lib.rs](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/src/lib.rs)
- [warren-surveyor-pg/warren_surveyor_pg.control](https://github.com/soulstompp/warren-surveyor/blob/59f4d0f475cd6a8ed6811be945b9738c35932eca/warren-surveyor-pg/warren_surveyor_pg.control)

`warren_surveyor_pg` 0.1.0 provides the `surveyor` index access method. Its marker indexes hold no user data and are never scanned by query plans; planner hooks inspect existing indexes to improve estimates. It also installs the bundled query-respelling hook.

### Core Workflow

```sql
CREATE EXTENSION warren_surveyor_pg;
LOAD 'warren_surveyor_pg';
CREATE TABLE survey_demo (id bigint PRIMARY KEY, value text);
CREATE INDEX ON survey_demo USING surveyor (id);
ANALYZE survey_demo;
EXPLAIN (ANALYZE, BUFFERS) SELECT * FROM survey_demo WHERE id = 42;
```

### Operational Boundaries

Create a marker index on the table’s unique key, keep the real B-tree/GIN/GiST indexes, and analyze the table. Load the library before planning in every participating session; the example uses `LOAD` for the current session. Session preload can enable it for later connections without a server-wide restart. Installation requires a superuser.

The README targets PostgreSQL 18; Cargo also exposes a PostgreSQL 19 build feature, which is not a validated compatibility promise. Measurements consume planning I/O and CPU and may not improve a workload. Inspect both estimated/actual rows and planning/execution buffers. `warren_surveyor_pg.planning_read_limit` caps reads; when measurement stops, the original estimates remain. Recursive queries can be executed during planning under the documented budget.

The extension rewrites only query shapes it considers equivalent. Compare result sets and plans on representative workloads before enabling it broadly. This is an initial source release, separate from `pg_warrendex`.
