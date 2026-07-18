## Usage

Sources:

- [Upstream behavior and warning](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/README.md)
- [Event-trigger implementation](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.c)
- [Extension control file](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.control)

`implicit_schema` installs an event trigger that creates a missing schema when a table or view is created with a schema-qualified name. It was built as a compatibility hack for an application that generated timestamped schema names.

```sql
CREATE EXTENSION implicit_schema;
CREATE TABLE fizz.fuzz (id integer);
SELECT nspname FROM pg_namespace WHERE nspname = 'fizz';
```

Upstream labels the behavior a bad idea. It silently turns an object-creation attempt into additional DDL and therefore changes privilege, audit, locking, and error expectations. The reviewed C code constructs a `CREATE SCHEMA` command from the requested name without safely escaping embedded identifier quotes. Do not grant it to untrusted users; prefer explicit schema provisioning, and use only after adversarial identifier and privilege tests.
