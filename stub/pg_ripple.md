## Usage

Sources:

- [Official installation guide](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/docs/src/getting-started/installation.md)
- [Official five-minute walkthrough](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/docs/src/getting-started/hello-world.md)
- [Official extension control file](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/pg_ripple.control)

`pg_ripple` is a PostgreSQL 18 RDF knowledge-graph engine with SPARQL querying, SHACL validation, Datalog reasoning, federation, and graph analytics. It stores graph data inside PostgreSQL; the reviewed release is `0.128.0` and requires superuser installation.

### Core Workflow

Preloading is required for the HTAP merge worker and shared-memory dictionary cache. Add the library, restart PostgreSQL, create the extension, then load RDF and query it.

```sql
ALTER SYSTEM SET shared_preload_libraries = 'pg_ripple';
-- Restart PostgreSQL before continuing.

CREATE EXTENSION pg_ripple;

SELECT pg_ripple.register_prefix('ex', 'http://example.org/');
SELECT pg_ripple.register_prefix('foaf', 'http://xmlns.com/foaf/0.1/');

SELECT pg_ripple.load_turtle('
  @prefix ex: <http://example.org/> .
  @prefix foaf: <http://xmlns.com/foaf/0.1/> .
  ex:alice foaf:name "Alice" .
  ex:bob foaf:name "Bob" .
');

SELECT *
FROM pg_ripple.sparql('
  PREFIX foaf: <http://xmlns.com/foaf/0.1/>
  SELECT ?person ?name WHERE {
    ?person foaf:name ?name .
  }
');
```

`load_turtle` returns the number of loaded triples. `sparql` returns one JSONB object per result row, keyed by query variables.

### Main SQL Surface

- `register_prefix` manages short namespace prefixes.
- `insert_triple`, `delete_triple`, and `triple_count` provide direct triple operations and basic inspection.
- `load_turtle` loads Turtle documents; named-graph APIs separate tenants or provenance sources.
- `sparql` and its explain/cursor variants execute graph queries from SQL.
- SHACL, Datalog, inference, vector/hybrid search, subscriptions, export, and administration APIs are documented as separate feature groups.

### Operational Boundary

This source line targets PostgreSQL `18`; do not assume compatibility with earlier majors. The extension creates internal schemas and predicate-specific storage, including delta, main, and tombstone relations. Use its public functions rather than modifying internal objects. Basic calls can load the library on demand, but production HTAP and shared-memory features need `shared_preload_libraries` and a restart.

Treat this broad, rapidly evolving API as versioned. Pin the extension and any HTTP or relay companions, validate backup/restore and upgrades, and review the security and outbound-federation settings before exposing endpoints or remote sources.
