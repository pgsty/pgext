## Usage

Sources:

- [Archived upstream repository and README](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/README.adoc)
- [Source control file](https://github.com/sim51/neo4j-fdw/blob/1a864186dcd3d3c4e95ebdfe8d8ff14bd527be66/neo4j-fdw.control)
- [Latest formal PGXN release](https://pgxn.org/dist/neo4j-fdw/)

`neo4j-fdw` is a Python integration that uses Multicorn to expose the result of a Neo4j Cypher query as a PostgreSQL foreign table. It can push basic comparison and pattern conditions into a generated Cypher query.

After installing the `neo4jPg` Python package, its matching Neo4j Python driver, and Multicorn, define a server and a foreign table:

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

SELECT * FROM neo4j_movie;
```

Cypher expressions used for a foreign table must return a collection and alias every returned field. Credentials are foreign-server options, so restrict catalog visibility and privileges accordingly.

### Archived status and versions

The repository was archived in 2023. The latest formal PGXN release is `4.3.1`, while the final source control file says `4.3.2-dev`, matching this catalog's source snapshot. Upstream's documented compatibility is limited to old Multicorn, Python, Neo4j, and PostgreSQL combinations, with the stated PostgreSQL range ending at 12. Treat it as legacy software and validate the complete dependency stack before use on any maintained PostgreSQL release.
