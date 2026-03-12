


## Usage

> [sparql: SPARQL query support for PostgreSQL](https://github.com/lacanoid/pgsparql)

Query SPARQL endpoints (e.g., DBpedia/Virtuoso) from PostgreSQL. SPARQL queries are compiled into Postgres views for use in SQL.

### Get Properties of a Resource

```sql
SELECT * FROM sparql.get_properties('dbpedia', 'http://dbpedia.org/resource/Johann_Sebastian_Bach');
```

### Get References to a Resource

```sql
SELECT * FROM sparql.get_references('dbpedia', 'http://dbpedia.org/resource/Johann_Sebastian_Bach');
```

### Compile SPARQL Query into SQL View

```sql
SELECT sparql.compile_query(endpoint, identifier, sparql_query [, grouping]);
```

Parameters:
- `endpoint` -- default SPARQL endpoint name
- `identifier` -- SQL identifier for the created function and view
- `sparql_query` -- the SPARQL query to compile
- `grouping` -- optional array of identifiers to group by (non-grouped columns are aggregated into arrays)

### Example

```sql
SELECT sparql.compile_query('dbpedia', 'ludwig_van', $$
  SELECT ?predicate, ?object
  WHERE {
    <http://dbpedia.org/resource/Ludwig_van_Beethoven> ?predicate ?object.
  }
$$, '{predicate}');

-- Now query via the created view
SELECT * FROM ludwig_van;
```

This creates a function `ludwig_van()` and a view `ludwig_van` that queries the SPARQL endpoint and returns results as a SQL table.
