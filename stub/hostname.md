## Usage

Sources:

- [Official README](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/README.md)
- [SQL declaration](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/sql/hostname.sql)
- [C implementation](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/src/hostname.c)
- [PGXN distribution](https://pgxn.org/dist/hostname/)

`hostname` provides one function, `hostname() RETURNS text`, that returns the operating-system host name visible to the PostgreSQL server process. It identifies the database server or container, not the client host and not necessarily a DNS-qualified or cluster-unique name.

### Core Workflow

```sql
CREATE EXTENSION hostname;

SELECT hostname();
```

The C function calls the POSIX `gethostname` API and returns NULL if that call fails. The extension is relocatable and requires a POSIX platform with `<unistd.h>`.

### Interpretation and Versioning

In containers, Kubernetes pods, failover clusters, and cloned hosts, the returned value can change across restarts or role transitions. Do not use it as a durable node identifier, security boundary, or routing source without comparing it to infrastructure metadata. The function is declared `IMMUTABLE`, so PostgreSQL may fold repeated calls within a planned expression even though an administrator can change the operating-system hostname.

The SQL extension version remains 1.0.0, while the upstream package/PGXN distribution is 1.0.4 at the pinned source revision. That is a packaging release difference rather than a new SQL extension version. Current upstream documentation supports PostgreSQL 9.0 or later and documents PostgreSQL 18 extension-path settings for custom install prefixes.
