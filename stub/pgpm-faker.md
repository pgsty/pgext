## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/faker/README.md)
- [Official extension control file (pgpm-faker.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/faker/pgpm-faker.control)
- [Official extension SQL (pgpm-faker--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/faker/sql/pgpm-faker--0.15.5.sql)

`pgpm-faker` — @pgpm/faker provides a comprehensive set of fake data generation functions directly in PostgreSQL. Perfect for seeding test databases, creating demo data, and development environments. All functions are implemented in pure plpgsql and return realistic-looking data without external dependencies. Use it when SQL needs these specialized functions or aggregates. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-faker";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `faker.address(state text DEFAULT NULL, city text DEFAULT NULL)` is an extension function and returns `text`.
- `faker.attachment(mime text DEFAULT NULL)` is an extension function and returns `attachment`.
- `faker.birthdate(min int DEFAULT 1, max int DEFAULT 100)` is an extension function and returns `date`.
- `faker.boolean()` is an extension function and returns `boolean`.
- `faker.business()` is an extension function and returns `text`.
- `faker.city(state text DEFAULT NULL)` is an extension function and returns `text`.
- `faker.date(min int DEFAULT 1, max int DEFAULT 100, future boolean DEFAULT false)` is an extension function and returns `date`.
- `faker.email()` is an extension function and returns `text`.
- `faker.ext(mime text DEFAULT faker.mime())` is an extension function and returns `text`.
- `faker.file(mime text DEFAULT NULL)` is an extension function and returns `text`.
- `faker.float(min double precision DEFAULT 0, max double precision DEFAULT 100)` is an extension function and returns `double`.
- `faker.fullname(gender text DEFAULT NULL)` is an extension function and returns `text`.
- `faker.gender(gender text DEFAULT NULL)` is an extension function and returns `text`.
- `faker.hostname()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.15.5`.
- Install the confirmed extension dependencies first: `citext`, `pgcrypto`, `plpgsql`, `pgpm-types`, `pgpm-verify`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
