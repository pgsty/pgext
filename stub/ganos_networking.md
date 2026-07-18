## Usage

Sources:

- [Alibaba Cloud GanosBase Networking path model](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/path-model)
- [Alibaba Cloud RDS for PostgreSQL supported extensions](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_networking` is Alibaba Cloud's GanosBase Networking extension for road and traffic-network analysis on ApsaraDB RDS for PostgreSQL. It provides pgRouting-compatible path functions, including `pgr_dijkstra`, `pgr_astar`, and `pgr_trsp`, for shortest, fastest, heuristic, and turn-restricted path calculations.

### Minimal Path Query

Install the provider-supplied extension in an RDS database. `CASCADE` lets PostgreSQL install provider-declared dependencies.

```sql
CREATE EXTENSION ganos_networking WITH SCHEMA public CASCADE;

SELECT *
FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edge_table',
  2,
  3
);
```

The edge query supplies an edge identifier, source and target node identifiers, and traversal costs. A negative `cost` makes the forward direction impassable; a negative `reverse_cost` makes the reverse direction impassable. Path results include sequence, node, edge, per-edge cost, and accumulated cost columns.

### Caveats

- This is a proprietary Alibaba Cloud provider extension, not a separately installable open-source package. Use the RDS plugin marketplace or a privileged `CREATE EXTENSION` command on a supported instance.
- Availability and version depend on the RDS PostgreSQL major/minor engine and edition. Check the live support table before deployment; catalog version `7.0` is not a promise that every instance offers that version.
- The SQL text passed to path functions identifies the graph and cost model. Validate its result shape and restrict who may supply dynamic SQL.
- A* additionally requires coordinate columns, while turn-restricted routing requires restriction data. Consult the provider's path-model reference for the exact signature used by your engine version.
