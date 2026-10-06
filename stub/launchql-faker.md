## Usage

Sources:

- [packages/faker/readme.md](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/readme.md)
- [packages/faker/sql/launchql-faker--0.1.0.sql](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/sql/launchql-faker--0.1.0.sql)
- [packages/faker/Makefile](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/Makefile)
- [packages/faker/launchql-faker.control](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/launchql-faker.control)

`launchql-faker` supplies the `faker` schema, a dictionary table and random generators for names, addresses, text, identifiers and numeric values.

### Core Workflow

```sql
CREATE EXTENSION "launchql-faker" CASCADE;
SELECT faker.fullname(), faker.email(), faker.integer(1, 100);
```

### Operational Boundaries

`faker.name`, `faker.fullname`, `faker.email`, `faker.address`, `faker.sentence`, `faker.integer` and `faker.uuid` provide representative generators. Values use bundled dictionary data and randomness; do not treat them as verified contact data or use generated passwords and tokens as a security contract. This historical version does not declare a current PostgreSQL-major matrix.

Install its declared dependencies first: `citext`, `pgcrypto`, `plpgsql`, `uuid-ossp`, `launchql-ext-types`. This is SQL/PLpgSQL code with no own shared library or preload. The control permits non-superuser installation but is not marked trusted; dependency and schema privileges still apply.
