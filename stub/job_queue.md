## Usage

Sources:

- [Official extension control file](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/job_queue.control)
- [Official upstream documentation](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/README.md)

`job_queue` — Proof-of-concept database-side asynchronous job queue using dynamic PostgreSQL background workers.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "job_queue";
SELECT extversion
FROM pg_extension
WHERE extname = 'job_queue';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
