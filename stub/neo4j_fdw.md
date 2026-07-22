## Usage

Sources:

- [Official upstream documentation](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/Readme)
- [Version 1.0 SQL objects](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/neo4j_fdw--1.0.sql)
- [Version 1.0 foreign-table examples](https://github.com/nuko-yokohama/neo4j_fdw/tree/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/test)

`neo4j_fdw` is a C foreign data wrapper that exposes the rows returned by a Neo4j Cypher request as a PostgreSQL foreign table. It targets Neo4j's historical HTTP Cypher endpoint and should not be confused with similarly named Python or Multicorn projects.

### Core Workflow

Create a server with the endpoint URL, then define a foreign table whose `query` option is a JSON request. The Cypher `RETURN` list must match the PostgreSQL columns by position and compatible type.

```sql
CREATE EXTENSION neo4j_fdw;

CREATE SERVER neo4j_srv
FOREIGN DATA WRAPPER neo4j_fdw
OPTIONS (url 'http://127.0.0.1:7474/db/data/cypher');

CREATE FOREIGN TABLE neo4j_people (
    name text,
    gender text,
    location text
)
SERVER neo4j_srv
OPTIONS (
    query '{"query":"START n=node(*) RETURN n.name AS name, n.gender? AS gender, n.location? AS location"}'
);

SELECT name, location FROM neo4j_people;
```

### Objects and Boundaries

- `neo4j_fdw` is the foreign data wrapper; the server option is `url` and the foreign-table option is `query`.
- `exec_cypher(text, text)` sends an endpoint URL and a JSON request directly and returns Neo4j's JSON response.
- The extension links against `libcurl` and `json-c`. Version `1.0` documentation does not define authentication, TLS verification, timeout, write support, or planner pushdown guarantees.

The examples use Neo4j's obsolete `START` syntax and legacy `/db/data/cypher` endpoint, so modern Neo4j releases may be incompatible. Treat query text and URLs as administrator-controlled input, restrict outbound network access, and verify authentication, escaping, NULL/type conversion, local versus remote filtering, cancellation, and failure behavior against the exact Neo4j release before use.
