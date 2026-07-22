## Usage

Sources:

- [Official README and limitations](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/README.md)
- [Table access method implementation](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/src/tam/mod.rs)
- [Index access method implementation](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/src/iam/mod.rs)
- [FoundationDB health function](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/src/health.rs)

`pgfdb` 0.0.2 is an experimental pgrx extension that implements PostgreSQL table and index access methods backed by FoundationDB. It is a prototype for distributed storage and transaction research, not production middleware. Upstream explicitly warns not to connect it even to an existing production FoundationDB cluster.

### Supported Trial Workflow

Upstream distributes the project as a Docker image and expects a FoundationDB 7.3 cluster file mounted into the container. After the prepared environment starts, verify connectivity and explicitly choose the custom access methods:

```sql
CREATE EXTENSION pgfdb;

SELECT fdb_is_healthy();

CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid(),
    name text
) USING pgfdb_table;

CREATE INDEX users_id_fdb ON users USING pgfdb (id);

INSERT INTO users (name) VALUES ('Alice'), ('Bob');

SET enable_seqscan = off;
SELECT * FROM users WHERE id = '00000000-0000-0000-0000-000000000000';
```

`pgfdb_table` is the table access method. `pgfdb` is the current source's index access method and supplies operator classes for a limited set of integer, floating-point, text, and UUID comparisons. The README contains an older `pgfdb_idx` spelling; the pinned implementation registers `pgfdb`, so verify the objects present in the exact image used.

### Prototype Limitations

DDL and PostgreSQL catalogs are not persisted in FoundationDB. A fresh PostgreSQL instance therefore loses the schema even if FoundationDB still holds data, preventing the advertised stateless multi-node model in this revision. Cost estimation is incomplete, so the planner may not choose custom indexes without test-only settings such as `enable_seqscan=off`.

FoundationDB's five-second read/write transaction limit also constrains PostgreSQL transactions, making long statements and OLAP workloads unsuitable. Primary-key constraints are not supported by the custom index access method, and only a limited set of data types can be indexed. Performance is explicitly not representative.

Use only disposable data and isolated FoundationDB clusters. The extension initializes the FoundationDB client and a transaction callback in every backend that loads it; ordinary package installation is insufficient without a valid cluster file and compatible FoundationDB native library. Validate the Docker image, server major version, extension source revision, and failure recovery as one tightly coupled test environment.
