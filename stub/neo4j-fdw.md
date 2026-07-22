## Usage

Sources:

- [Archived official README at the reviewed commit](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/README.adoc)
- [Bundled SQL at the reviewed commit](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/sql/neo4j-fdw.sql)
- [Multicorn wrapper implementation](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/neo4jPg/neo4jfdw.py)
- [Extension control file](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/neo4j-fdw.control)
- [Official PGXN distribution](https://pgxn.org/dist/neo4j-fdw/)

`neo4j-fdw` is an archived Python integration that uses Multicorn to expose a Neo4j Cypher result as a PostgreSQL foreign table. It is not a native FDW library: the PostgreSQL server process imports the separately installed `neo4jPg` Python package, the Neo4j Python driver, and Multicorn.

### Core Workflow

After installing a mutually compatible Python, Multicorn, and Neo4j-driver stack, create the Multicorn server and map one aliased Cypher result to one foreign table:

```sql
CREATE EXTENSION multicorn;

CREATE SERVER multicorn_neo4j
FOREIGN DATA WRAPPER multicorn
OPTIONS (
  wrapper 'neo4jPg.neo4jfdw.Neo4jForeignDataWrapper',
  url 'neo4j://neo4j.example:7687',
  user 'neo4j',
  password 'replace-me'
);

CREATE FOREIGN TABLE neo4j_movie (movie text)
SERVER multicorn_neo4j
OPTIONS (
  cypher 'MATCH (n:Movie) RETURN n.title AS movie',
  database 'neo4j'
);

SELECT * FROM neo4j_movie WHERE movie ILIKE '%matrix%';
```

The Cypher expression must return a collection and alias every returned field to a foreign-table column. Simple comparison and pattern qualifiers can be rewritten into Cypher; optional strict `/*WHERE{...}*/` placeholders tell the wrapper where to insert predicates more efficiently.

### Packaging and Compatibility Boundaries

The canonical project name is `neo4j-fdw`, but the reviewed source is not a reliable modern extension bundle. Its control file reports `4.3.2-dev`, while the repository installs an unversioned SQL file rather than a matching versioned extension script. That SQL also requests the legacy `plpythonu` extension but declares functions in `plpython3u`. Validate and patch packaging instead of assuming that `CREATE EXTENSION "neo4j-fdw"` succeeds.

The FDW table path above needs Multicorn and the Python package; the optional direct `cypher()` SQL helpers additionally need untrusted PL/Python. Upstream's stated PostgreSQL compatibility ends at 12, and the repository was archived in 2023. Credentials are stored as foreign-server options. Restrict catalog and server privileges, use a dedicated Neo4j account and encrypted transport, and test query generation against the exact Neo4j driver/server pair.
