## Usage

Sources:

- [Quria README at the reviewed commit](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/README.md)
- [quria.control at the reviewed commit](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/quria.control)
- [Cargo package metadata](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/Cargo.toml)
- [Upstream demonstration SQL](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/demo.sql)
- [Bootstrap and preload behavior](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/sql/_bootstrap.sql)
- [Public functions and hooks](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/lib.rs)
- [Full-text operator implementation](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/fulltext/operators.rs)
- [On-disk inverted-index implementation](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/fulltext/index/inverted/inverted.rs)
- [Vector module at the reviewed commit](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/vector/hnsw/mod.rs)

`quria` is a version `0.0.0` pgrx prototype for BM25-style full-text search. It defines the `quria.fulltext` and `quria.query` types, a `quria_fts` index access method, the `~>` match operator, and `quria.ft_score`.

### Prototype Full-Text Search

After superuser installation, reconnect so the database-level session preload setting created by the bootstrap takes effect:

```sql
CREATE EXTENSION quria;

CREATE TABLE docs (body quria.fulltext);
INSERT INTO docs VALUES ('tepid text turtle'), ('fast green turtle');
CREATE INDEX docs_body_quria ON docs USING quria_fts (body);

SELECT body,
       quria.ft_score(ctid, '{"query":"turtle","column":"body"}'::quria.query) AS score
FROM docs
WHERE body ~> '{"query":"turtle","column":"body"}'::quria.query
ORDER BY score DESC;
```

### Caveats

- The README advertises HNSW vector search, but the reviewed vector module contains no implementation. Only the experimental full-text path has substantive code.
- Bootstrap changes `session_preload_libraries` for the whole database and grants all privileges on schema `quria` to `PUBLIC`. Review and narrow both changes before any shared test system.
- Index data is serialized under the PostgreSQL operating-system user's local `.quria/fulltext` directory, outside ordinary PostgreSQL relation storage and WAL. Standard backup, replication, crash recovery, and transactional guarantees do not automatically cover those files.
- The source is marked WIP, has no release tags, uses old internal index APIs, and contains many panic paths. Cargo features cover PostgreSQL 11 through 15 with 13 as default; they are not a production compatibility statement.
