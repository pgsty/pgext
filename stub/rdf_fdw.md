


## Usage

Sources:

- [PGXN rdf_fdw 2.6.0](https://pgxn.org/dist/rdf_fdw/2.6.0/)
- [rdf_fdw README](https://github.com/jimjonesbr/rdf_fdw)
- [rdf_fdw CHANGELOG](https://github.com/jimjonesbr/rdf_fdw/blob/master/CHANGELOG.md)
- [rdf_fdw control file](https://pgxn.org/dist/rdf_fdw/2.6.0/)
- [Local package metadata](../db/extension.csv)

`rdf_fdw` is a PostgreSQL foreign data wrapper for querying RDF triplestores over SPARQL endpoints. It exposes SPARQL result variables as foreign-table columns, supports pushdown for common SQL clauses, includes a native `rdfnode` type for RDF terms, provides SPARQL 1.1 helper functions, and can perform SPARQL `INSERT`, `UPDATE`, and `DELETE` through writable foreign tables.

v2.6.0 adds Bearer-token authentication through `USER MAPPING`, a `max_response_size` server option to cap HTTP response bodies, BCE date/timestamp cast handling, and many `rdfnode` parser/comparison fixes. v2.5 added `request_timeout` and `readonly` options.

### Create the Extension

```sql
CREATE EXTENSION IF NOT EXISTS rdf_fdw;

SELECT rdf_fdw_version();
SELECT * FROM rdf_fdw_settings();
```

To install or update to the exact SQL version:

```sql
CREATE EXTENSION rdf_fdw WITH VERSION '2.6';
ALTER EXTENSION rdf_fdw UPDATE TO '2.6';
```

### Register a SPARQL Endpoint

```sql
CREATE SERVER dbpedia
FOREIGN DATA WRAPPER rdf_fdw
OPTIONS (
  endpoint          'https://dbpedia.org/sparql',
  enable_pushdown   'true',
  request_timeout   '60',
  max_response_size '104857600',
  readonly          'true'
);
```

Useful server options include:

- `endpoint`: SPARQL endpoint URL; required.
- `batch_size`: number of rows per SPARQL UPDATE batch.
- `enable_pushdown`: enables SQL-to-SPARQL pushdown.
- `format`: expected SPARQL result MIME type.
- `http_proxy`: proxy URL; proxy credentials belong in `USER MAPPING`.
- `connect_timeout`: connection timeout.
- `request_timeout`: complete HTTP request timeout.
- `max_response_size`: maximum response body size in bytes; `0` means unlimited.
- `readonly`: prevents `INSERT`, `UPDATE`, and `DELETE` before requests reach the endpoint.
- `request_redirect` and `request_max_redirect`: redirect behavior.

Use `max_response_size` for public or untrusted endpoints because `rdf_fdw` loads retrieved RDF data into memory before converting it for PostgreSQL.

### User Mapping

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (
  user 'sparql_user',
  password 'secret'
);
```

v2.6.0 adds Bearer-token authentication:

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (
  token 'eyJhbGciOi...'
);
```

Proxy credentials are also `USER MAPPING` options:

```sql
CREATE USER MAPPING FOR app_user
SERVER dbpedia
OPTIONS (
  proxy_user 'proxy-user',
  proxy_password 'proxy-secret'
);
```

### Foreign Tables with rdfnode Columns

Declare foreign-table columns as `rdfnode` to preserve RDF terms, IRIs, blank nodes, language tags, and XSD datatypes.

```sql
CREATE FOREIGN TABLE dbpedia_films (
  film rdfnode OPTIONS (variable '?film'),
  name rdfnode OPTIONS (variable '?name'),
  year rdfnode OPTIONS (variable '?year')
)
SERVER dbpedia
OPTIONS (
  sparql $$
    SELECT ?film ?name ?year
    WHERE {
      ?film a dbo:Film ;
            rdfs:label ?name ;
            dbo:releaseDate ?year .
      FILTER (lang(?name) = 'en')
    }
  $$
);
```

Native PostgreSQL column types are deprecated for RDF values in v2.6.0. Existing native-typed tables continue to work, but they emit warnings and lose RDF term details.

### Querying and Pushdown

```sql
SELECT film, sparql.lex(name) AS title
FROM dbpedia_films
WHERE name = '"The Matrix"@en'::rdfnode
ORDER BY year
LIMIT 10;

EXPLAIN (VERBOSE, COSTS OFF)
SELECT *
FROM dbpedia_films
WHERE film = '<http://dbpedia.org/resource/The_Matrix>'::rdfnode;
```

`rdf_fdw` can push down `WHERE`, `LIMIT`, `ORDER BY`, `DISTINCT`, and supported comparisons/functions. Use `EXPLAIN VERBOSE` to inspect the generated remote SPARQL.

### Prefix Management

`rdf_fdw` provides catalog tables and helper functions under the `sparql` schema for reusable SPARQL prefixes:

```sql
SELECT sparql.add_context('default', 'Default SPARQL prefix context');
SELECT sparql.add_prefix('default', 'rdf',  'http://www.w3.org/1999/02/22-rdf-syntax-ns#');
SELECT sparql.add_prefix('default', 'rdfs', 'http://www.w3.org/2000/01/rdf-schema#');
SELECT sparql.add_prefix('default', 'xsd',  'http://www.w3.org/2001/XMLSchema#');
```

### Data Modification

Writable foreign tables can translate PostgreSQL `INSERT`, `UPDATE`, and `DELETE` into SPARQL UPDATE requests when the foreign table has the required SPARQL update pattern.

```sql
ALTER FOREIGN TABLE dbpedia_films OPTIONS (ADD readonly 'false');

INSERT INTO dbpedia_films(film, name)
VALUES (
  '<http://example.org/film/1>'::rdfnode,
  '"Example Film"@en'::rdfnode
);
```

Use `readonly = true` at the server or table level when an endpoint should never receive writes.

### Clone a Foreign Table

```sql
CALL rdf_fdw_clone_table(
  foreign_table := 'dbpedia_films',
  target_table  := 'dbpedia_films_local',
  fetch_size    := 1000,
  create_table  := true
);
```

`rdf_fdw_clone_table()` copies data from a foreign table into a local table in batches. v2.5 fixed several round-trip issues for RDF terms during cloning.

### SPARQL Functions

The `sparql` schema implements many SPARQL 1.1 functions and aggregates, including:

- aggregates such as `sparql.sum`, `sparql.avg`, `sparql.min`, `sparql.max`, `sparql.group_concat`, and `sparql.sample`
- RDF term helpers such as `sparql.isiri`, `sparql.isblank`, `sparql.isliteral`, `sparql.datatype`, `sparql.iri`, `sparql.strdt`, and `sparql.strlang`
- string functions such as `sparql.strlen`, `sparql.substr`, `sparql.ucase`, `sparql.lcase`, `sparql.contains`, and `sparql.replace`
- numeric, date/time, hash, and custom convenience functions

### Caveats

- PostgreSQL 9.5+ is the upstream baseline, but Pigsty packages target modern PostgreSQL majors listed in local metadata.
- Retrieved RDF data is accumulated in memory before conversion. Set `max_response_size`, use `LIMIT`, and keep remote result sets bounded.
- Prefer `rdfnode` columns. Native PostgreSQL typed columns are deprecated for RDF terms and will lose IRI/language/datatype information.
- Store secrets in `USER MAPPING`; do not put proxy credentials or endpoint tokens into `SERVER` options.
- Public SPARQL endpoints can be slow or rate-limited. Use `connect_timeout`, `request_timeout`, retries, and local materialization when needed.
