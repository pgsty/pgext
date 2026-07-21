## Usage

Sources:

- [pgRDF 0.6.20 README](https://github.com/styk-tv/pgRDF/blob/v0.6.20/README.md)
- [pgRDF 0.6.20 user guide](https://github.com/styk-tv/pgRDF/tree/v0.6.20/guide)
- [pgRDF 0.6.20 changelog](https://github.com/styk-tv/pgRDF/blob/v0.6.20/CHANGELOG.md)
- [pgRDF 0.6.20 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.20)

`pgRDF` stores RDF data inside PostgreSQL and exposes SQL-callable helpers for Turtle/TriG/N-Quads loading, SPARQL query/update, named graphs, SHACL validation, and RDFS/OWL 2 RL materialization.

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.version();
```

### Preload And PostgreSQL Version Caveat

pgRDF 0.6.20 supports PostgreSQL 14-18 and moves from `pgrx` 0.16.1 to 0.19.1. Upstream describes 0.6.20 as a build/runtime migration with no schema or query-surface change; PostgreSQL 19 remains a tracked follow-up.

`pgrdf` must be present in `shared_preload_libraries` before PostgreSQL starts. Without preload, upstream documents that the shared-memory dictionary and plan-cache atomics are not initialized and the first pgRDF call can fail.

```conf
shared_preload_libraries = 'pgrdf'
```

Restart PostgreSQL after changing this setting, then verify:

```sql
SHOW shared_preload_libraries;

SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> . ex:t a ex:T .',
  1::bigint,
  'http://example.org/'
);
```

### Load RDF

Use `parse_turtle` for inline Turtle payloads and `load_turtle` for server-side files. Graph ids are `bigint` values; named graph helpers map ids to IRIs. Version 0.6.x adds a parallel bulk loader path via `load_turtle(..., bulk_load => true)`.

```sql
SELECT pgrdf.add_graph(100::bigint, 'http://example.org/graph/main');

SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> .
   ex:alice ex:knows ex:bob .
   ex:alice ex:name "Alice" .',
  100::bigint,
  'http://example.org/graph/main'
);

SELECT pgrdf.load_turtle('/srv/rdf/foaf.ttl', 100::bigint);
SELECT pgrdf.load_turtle('/srv/rdf/bulk.ttl', 100::bigint, bulk_load => true);
SELECT pgrdf.count_quads(100::bigint);
```

Related ingest and graph-management functions documented upstream include `parse_trig`, `parse_nquads`, `add_graph`, `drop`, `clear`, `copy`, `move_graph`, `graph_id`, and `graph_iri`.

### Carve Graph Slices

The 0.6.x series adds `carve_graph` overloads for copying a predicate-defined slice or a bounded neighbourhood into another graph without decoding and re-encoding the shared dictionary:

```sql
SELECT pgrdf.carve_graph(
  100::bigint,
  'http://example.org/type'::text,
  200::bigint
);

SELECT pgrdf.carve_graph(
  100::bigint,
  ARRAY['http://example.org/alice', 'http://example.org/bob']::text[],
  201::bigint,
  2
);
```

The neighbourhood form uses `max_hops` as a graph-distance boundary. Set `pgrdf.on_path_truncation` to `warn` or `error` when a truncated property-path walk must not be silently accepted.

### Query With SPARQL

`pgrdf.sparql(text)` returns SPARQL results as SQL rows. The upstream v0.5 surface includes `SELECT` and `ASK`, filters, ordering, limits, `OPTIONAL`, `UNION`, `MINUS`, aggregates, `VALUES`, `BIND`, `CONSTRUCT`, `DESCRIBE`, named-graph `GRAPH` clauses, and property paths.

```sql
SELECT *
FROM pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   SELECT ?person ?name
   WHERE {
     ?person ex:name ?name .
     FILTER(REGEX(?name, "^A", "i"))
   }
   ORDER BY ?name
   LIMIT 20'
);
```

Named-graph queries can bind graph IRIs:

```sql
SELECT *
FROM pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   SELECT ?g (COUNT(*) AS ?n)
   WHERE { GRAPH ?g { ?s ex:name ?name } }
   GROUP BY ?g
   ORDER BY ?g'
);
```

### Update Graphs

The upstream v0.5 surface includes SPARQL Update forms such as `INSERT DATA`, `DELETE DATA`, `INSERT/DELETE WHERE`, `DELETE+INSERT WHERE`, and graph lifecycle statements.

```sql
SELECT pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   INSERT DATA {
     GRAPH <http://example.org/graph/main> {
       ex:bob ex:name "Bob" .
     }
   }'
);
```

### Reasoning And Validation

Use `pgrdf.materialize(graph_id, profile)` to write inferred triples for `rdfs` or `owl-rl` profiles. Materialization is intended to be repeatable; upstream documents that previous inferred rows are dropped before writing the new closure.

```sql
SELECT pgrdf.parse_turtle(
  '@prefix ex:   <http://example.com/> .
   @prefix rdf:  <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
   @prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
   ex:Engineer rdfs:subClassOf ex:Person .
   ex:Person   rdfs:subClassOf ex:Agent .
   ex:alice    rdf:type        ex:Engineer .',
  100::bigint
);

SELECT pgrdf.materialize(100::bigint, 'owl-rl');

SELECT *
FROM pgrdf.sparql(
  'PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
   PREFIX ex:  <http://example.com/>
   SELECT ?class WHERE { ex:alice rdf:type ?class }'
);
```

Use `pgrdf.validate(data, shapes, mode)` for SHACL validation; upstream documents JSONB `sh:ValidationReport` output and native SHACL Core support. SHACL-SPARQL constraint execution is documented upstream as gated by its RDF library dependency, so treat `mode => 'sparql'` as an advanced surface to verify against the exact installed version.

### Operational Helpers

Useful introspection and cache-management helpers documented upstream include:

| Function | Use |
|----------|-----|
| `pgrdf.stats()` | Inspect runtime counters and cache state |
| `pgrdf.shmem_reset()` | Reset shared-memory dictionary/cache state |
| `pgrdf.plan_cache_clear()` | Clear prepared SPARQL plan cache |
| `pgrdf.sparql_parse(text)` | Inspect parsed SPARQL without executing it |

The `pgrdf.path_max_depth` setting guards property-path expansion depth, while `pgrdf.on_path_truncation = count | warn | error` controls how callers learn that the guard was reached.

### Version Notes

The releases between 0.6.4 and 0.6.20 materially improve large RDF ingestion and query correctness: streaming/windowed bulk ingestion, a staged multi-backend loader, safe repeated loads into populated dictionaries, graph-carving helpers, dictionary inclusion in `pg_dump`, SPARQL expression/aggregate additions, and fail-closed path-truncation handling. The 0.6.20 release itself changes only the build/runtime layer to `pgrx` 0.19.1 and adds PostgreSQL 18 support.

For very large fresh N-Triples loads, upstream documents `pgrdf.load_turtle_staged_run` as the resumable, phase-oriented path. It commits parsing, dictionary, resolution, and index phases separately and is operationally different from the transactional `load_turtle()` call; validate staging tablespaces, disk headroom, and recovery procedures before using it for production-scale imports.
