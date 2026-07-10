


## Usage

> Sources: [pgRDF upstream README](https://github.com/styk-tv/pgRDF/blob/v0.6.4/README.md), [pgRDF user guide](https://github.com/styk-tv/pgRDF/tree/v0.6.4/guide), [v0.6.4 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.4).

`pgRDF` stores RDF data inside PostgreSQL and exposes SQL-callable helpers for Turtle/TriG/N-Quads loading, SPARQL query/update, named graphs, SHACL validation, and RDFS/OWL 2 RL materialization.

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.version();
```

### Preload And PostgreSQL Version Caveat

Upstream documents PostgreSQL 14-17 support and defers PostgreSQL 18 while pgRDF remains pinned to `pgrx` 0.16.

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

The `pgrdf.path_max_depth` setting guards property-path expansion depth.

### Version Notes

`pgrdf` 0.6.4 improves the deferred-index bulk-load path: for fresh bulk loads above `pgrdf.bulk_defer_index_min`, `load_turtle(..., bulk_load => true)` also defers the dictionary `unique_term` constraint, then rebuilds and validates it in the same transaction. PostgreSQL 18 remains deferred upstream while pgRDF is pinned to `pgrx` 0.16.
