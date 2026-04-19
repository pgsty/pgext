
## Usage

> Sources: [README](https://github.com/jimjonesbr/rdf_fdw/blob/main/README.md), [v2.4 release](https://github.com/jimjonesbr/rdf_fdw/releases/tag/v2.4)

`rdf_fdw` is a foreign data wrapper for RDF triplestores exposed through SPARQL endpoints. It lets PostgreSQL query RDF data with SQL, supports SQL clause pushdown, adds an `rdfnode` type for RDF terms, and includes SPARQL 1.1 function support.

### Register a SPARQL Endpoint

Register a SPARQL endpoint with `CREATE SERVER`:

```sql
CREATE SERVER dbpedia
FOREIGN DATA WRAPPER rdf_fdw
OPTIONS (endpoint 'https://dbpedia.org/sparql');
```

The README documents server options such as:

- `endpoint` (required)
- `batch_size`
- `enable_pushdown`
- `format`
- `http_proxy`
- `connect_timeout`

In `v2.4`, proxy credentials were moved out of `SERVER` options and into `USER MAPPING` for security.

### User Mapping and Foreign Tables

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (user 'admin', password 'secret');
```

`rdf_fdw` works by declaring foreign tables that embed SPARQL queries and map result variables to PostgreSQL columns. The README also highlights native RDF node handling through the custom `rdfnode` type.

```sql
CREATE FOREIGN TABLE film (
  film_id text OPTIONS (variable '?film', nodetype 'iri'),
  name text OPTIONS (variable '?name', nodetype 'literal', literal_type 'xsd:string')
)
SERVER dbpedia
OPTIONS (sparql 'SELECT ?film ?name WHERE { ?film dbp:name ?name }');
```

### Pushdown, DML, and Helpers

The upstream docs specifically call out pushdown for:

- `WHERE`
- `LIMIT`
- `ORDER BY`
- `DISTINCT`

They also document data modification support:

- `INSERT`
- `UPDATE`
- `DELETE`

Batching for SPARQL UPDATE traffic is controlled with the `batch_size` option.

The README lists utility functions including:

- `rdf_fdw_version()`
- `rdf_fdw_settings()`
- `rdf_fdw_clone_table()`

It also documents broader SPARQL function coverage, including aggregates, string functions, numeric functions, date/time functions, hash functions, and custom functions.

### Caveats

The current README warns that retrieved RDF data is loaded into memory before conversion for PostgreSQL, so large result sets require adequate PostgreSQL memory. It also documents PostgreSQL 9.5+ as the supported baseline.
