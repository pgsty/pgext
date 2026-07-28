## Usage

Sources:

- [Official upstream README](https://github.com/beyondoss/queue/blob/731d6039952814955a77d1bc065f50291deca781/README.md)
- [Official extension control file (beyond_queue.control)](https://github.com/beyondoss/queue/blob/731d6039952814955a77d1bc065f50291deca781/beyond-queue-extension/beyond_queue.control)
- [Official implementation source](https://github.com/beyondoss/queue/blob/731d6039952814955a77d1bc065f50291deca781/beyond-queue-extension/src/lib.rs)

`beyond_queue` — A faster fork of pgmq for https://beyond.dev queues. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION beyond_queue;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
