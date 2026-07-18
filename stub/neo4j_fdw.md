## Usage

Sources:

- [Upstream usage description](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/Readme)
- [Version 1.0 SQL objects](https://github.com/nuko-yokohama/neo4j_fdw/blob/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/neo4j_fdw--1.0.sql)
- [Regression examples](https://github.com/nuko-yokohama/neo4j_fdw/tree/36dc2011de7adc8f68e419f0d4b1bd4858d3bd07/test)

`neo4j_fdw` is a C foreign data wrapper for exposing Neo4j Cypher results as PostgreSQL foreign tables. A server supplies the Neo4j URL, while foreign-table columns carry the Cypher query and expected result columns; the query's `RETURN` list must match the PostgreSQL definition.

```sql
CREATE EXTENSION neo4j_fdw;
CREATE SERVER neo4j_srv FOREIGN DATA WRAPPER neo4j_fdw
OPTIONS (url 'http://127.0.0.1:7474');
```

The extension links `libcurl` and `json-c`, and also exposes an ad-hoc JSON query path. It is an old source-only implementation, distinct from similarly named Python/Multicorn projects. Keep credentials and Cypher text out of untrusted options, enforce network timeouts and TLS externally, and validate types, NULLs, escaping, pushdown, and remote failure behavior before use.
