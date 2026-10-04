## Usage

Sources:

- [0.6.39 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.39)
- [pgrdf.control](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/pgrdf.control)
- [README.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/README.md)
- [Cargo.toml](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/Cargo.toml)
- [guide/01-install.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/01-install.md)
- [guide/tour.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/tour.md)
- [guide/05-graphs.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/05-graphs.md)
- [guide/06-validation-recipes.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/06-validation-recipes.md)

`pgrdf` 0.6.39 stores RDF graphs in PostgreSQL with Turtle/TriG/N-Quads ingestion, SPARQL queries and updates, SHACL validation, and RDFS/OWL reasoning. Preload `pgrdf` and restart before creating the extension as a superuser.

### Core Workflow

```ini
shared_preload_libraries = 'pgrdf'
```

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.add_graph('http://example.org/people');
SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> . ex:alice ex:name "Alice" .',
  pgrdf.graph_id('http://example.org/people'));
SELECT * FROM pgrdf.sparql('SELECT ?s ?p ?o WHERE { ?s ?p ?o } LIMIT 10');
SELECT * FROM pgrdf.surface();
```

### Operational Boundaries

The fixed `pgrdf` schema provides `add_graph`, `graph_id`, `parse_turtle`, `load_turtle`, `sparql`, `materialize`, `validate`, `stats` and `surface`. File loaders read server-side paths; grant access deliberately. `can_clear_graphs()` reports graph-clear privileges. Non-owner graph writes require the documented underlying table grants; graph locks still apply. `shacl_capability()` reports the supported validation surface, and path-depth/truncation settings bound traversal. Source features cover PostgreSQL 14–18; published Linux binaries target PG18. Version 0.6.39 allocates graph IDs from a non-reusing sequence; the IRI is the graph identity, and gaps are normal. After installing the matching library, run `ALTER EXTENSION pgrdf UPDATE` in each database. A newer library rejects old SQL with SQLSTATE 55000. Do not use the broken 0.6.35–0.6.38 PGXN source archives. Preserve graph/dictionary data together in verified backups.
