## Usage

Sources:

- [Official upstream README](https://github.com/ivanyu/pg_kafka/blob/0d589ece33258b7f751c0095267971eef0924b87/README.md)
- [Official extension control file (pg_kafka.control)](https://github.com/ivanyu/pg_kafka/blob/0d589ece33258b7f751c0095267971eef0924b87/pg_kafka.control)
- [Official implementation source](https://github.com/ivanyu/pg_kafka/blob/0d589ece33258b7f751c0095267971eef0924b87/src/lib.rs)

`pg_kafka` — PostgreSQL foreign data wrapper for Kafka. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_kafka;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
