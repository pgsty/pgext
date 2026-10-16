## Usage

Sources:

- [README.md](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/README.md)
- [docs/user-guide.md](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/docs/user-guide.md)
- [sql/age_gds--1.0.sql](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/sql/age_gds--1.0.sql)
- [age_gds.control](https://github.com/Dn-Media-Group-Public/age_gds/blob/4b9210f3511c53fd8f967f9995efacbda18d15a2/age_gds.control)

`age_gds` adds NetworKit graph analytics to Apache AGE. The installed SQL extension version is 1.0; the separate `age_gds.version()` implementation diagnostic is documented as 0.1.0. Project an existing AGE graph, run algorithms and optionally write results back.

### Core Workflow

```sql
CREATE EXTENSION age;
LOAD 'age';
SET search_path = ag_catalog, "$user", public;
CREATE EXTENSION age_gds;
SELECT age_gds.graph_project(
  name => 'social_projection', age_graph => 'social',
  node_query => $$MATCH (n:Person) RETURN id(n) AS id$$,
  edge_query => $$MATCH (a:Person)-[:KNOWS]->(b:Person)
                 RETURN id(a) AS src, id(b) AS dst, 1.0 AS weight$$);
SELECT * FROM age_gds.pageRank_stream(graph_name => 'social_projection')
ORDER BY score DESC;
```

### Operational Boundaries

Install Apache AGE first and load `age` in the session. Installation is superuser-only. The source requires PostgreSQL 14 or later, an appropriate AGE build, OpenMP and the bundled NetworKit library; compatibility still depends on the AGE/server combination. There is no documented requirement to preload `age_gds` globally.

`graph_project` materializes a named projection from node and edge Cypher queries. `pageRank_stream`, community detection and shortest-path functions return rows; corresponding write functions modify the AGE graph. `graph_list`, `graph_exists`, `graph_reindex` and `graph_drop` manage projections. Ordinary snapshots become stale when source data changes; rebuild them or use the documented live-index mode.

Projection files default to the server data directory and are distinct from AGE tables. Account for them in backup and restore procedures. `age_gds.session_cache_mb` bounds the cache per backend; `age_gds.data_dir` changes the storage directory and requires superuser privileges. Large projections and algorithms consume memory and CPU.
