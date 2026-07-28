## Usage

Sources:

- [Official upstream README](https://github.com/polyaklaci/pg_argon2id/blob/df949026ea01e9978092b24582ba6f7cb3897a6d/README.md)
- [Official extension control file (pg_argon2id.control)](https://github.com/polyaklaci/pg_argon2id/blob/df949026ea01e9978092b24582ba6f7cb3897a6d/pg_argon2id.control)
- [Official extension SQL (pg_argon2id--1.0.sql)](https://github.com/polyaklaci/pg_argon2id/blob/df949026ea01e9978092b24582ba6f7cb3897a6d/sql/pg_argon2id--1.0.sql)

`pg_argon2id` — A PostgreSQL extension that provides secure password hashing using the Argon2id algorithm. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_argon2id;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `argon2id_hash(password text, salt_length integer DEFAULT 16, time_cost integer DEFAULT 3, memory_cost integer DEFAULT 262144, parallelism integer DEFAULT 4)` is an extension function and returns `text`.
- `argon2id_verify(password text, hash text)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
