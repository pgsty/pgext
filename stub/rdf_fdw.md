
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION rdf_fdw;
> CREATE SERVER dbpedia FOREIGN DATA WRAPPER rdf_fdw
>   OPTIONS (endpoint 'https://dbpedia.org/sparql');
> ```
>
> Source: [README](https://github.com/jimjonesbr/rdf_fdw)

`rdf_fdw` is a foreign data wrapper for RDF triplestores exposed through SPARQL endpoints. It lets PostgreSQL query RDF data with SQL, supports SQL clause pushdown, adds an `rdfnode` type for RDF terms, and includes SPARQL 1.1 function support.

## Server Setup

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

Proxy credentials belong in a user mapping.

## Foreign Tables

`rdf_fdw` works by declaring foreign tables that embed SPARQL queries and map result variables to PostgreSQL columns. The README also highlights native RDF node handling through the custom `rdfnode` type.

## Pushdown and DML

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

## Helper Functions

The README lists utility functions including:

- `rdf_fdw_version()`
- `rdf_fdw_settings()`
- `rdf_fdw_clone_table()`

It also documents broader SPARQL function coverage, including aggregates, string functions, numeric functions, date/time functions, hash functions, and custom functions.

## Notes

The current README warns that retrieved RDF data is loaded into memory before conversion for PostgreSQL, so large result sets require adequate PostgreSQL memory. It also documents PostgreSQL 9.5+ as the supported baseline.
