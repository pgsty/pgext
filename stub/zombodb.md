## Usage

Sources:

- [Upstream README at version 3000.2.8](https://github.com/zombodb/zombodb/blob/11d5cdfbd2b3fd238acab0ce785cf374e8f6a909/README.md)
- [Extension control file](https://github.com/zombodb/zombodb/blob/11d5cdfbd2b3fd238acab0ce785cf374e8f6a909/zombodb.control)
- [Operational design notes](https://github.com/zombodb/zombodb/blob/11d5cdfbd2b3fd238acab0ce785cf374e8f6a909/THINGS-TO-KNOW.md)
- [Archived upstream repository](https://github.com/zombodb/zombodb)

`zombodb` implements a PostgreSQL index access method backed by Elasticsearch. It synchronously manages the remote index and exposes the `==>` search operator while preserving PostgreSQL transaction visibility:

```sql
CREATE EXTENSION zombodb;

CREATE TABLE products (
  id bigserial PRIMARY KEY,
  name text NOT NULL,
  description zdb.fulltext
);

CREATE INDEX products_zdb_idx
  ON products
  USING zombodb ((products.*))
  WITH (url = 'http://elasticsearch.example.net:9200/');

SELECT *
FROM products
WHERE products ==> 'name:keyboard';
```

### Archived status and operational risk

The repository is archived. Its final control version `3000.2.8` documents PostgreSQL 12 through 15 and Elasticsearch 7.10 or later. It supports only one ZomboDB index per table, not partial indexes or `CREATE INDEX CONCURRENTLY`. Writes synchronously round-trip to Elasticsearch, so network or Elasticsearch failures abort the PostgreSQL transaction. Never modify ZomboDB-managed indexes externally, and account for the remote copy storing whole rows plus analyzed terms. Treat this as legacy infrastructure and plan migration before upgrading either dependency.
