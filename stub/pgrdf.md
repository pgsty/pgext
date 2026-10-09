## Usage

Sources:

- [0.6.40 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.40)
- [pgrdf.control](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/pgrdf.control)
- [README.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/README.md)
- [Cargo.toml](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/Cargo.toml)
- [guide/01-install.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/01-install.md)
- [guide/tour.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/tour.md)
- [guide/05-graphs.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/05-graphs.md)
- [guide/06-validation-recipes.md](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/guide/06-validation-recipes.md)
- [0.6.39 to 0.6.40 upgrade](https://github.com/styk-tv/pgRDF/blob/eefe44a459407c0d391b44baebe8905baaf7f9de/sql/pgrdf--0.6.39--0.6.40.sql)

`pgrdf` 0.6.40 stores RDF graphs in PostgreSQL with Turtle/TriG/N-Quads ingestion, SPARQL queries and updates, SHACL validation, and RDFS/OWL reasoning. Preload `pgrdf` and restart before creating the extension as a superuser.

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

### Ownership Repair in 0.6.40

The 0.6.39 upgrade could leave `_pgrdf_graph_id_seq` owned by the upgrade executor after storage tables had been transferred to another role, preventing new graph creation. Version 0.6.40 aligns that sequence with the owner of `_pgrdf_quads`; granting sequence usage alone is insufficient. `pgrdf.ownership_drift()` reports mismatched storage owners and proposed repair commands. Review them before applying role changes. Upgrade SQL in each database after installing matching binaries; versions 0.6.34–0.6.38 can upgrade directly to 0.6.40.
