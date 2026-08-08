## Usage

Sources:

- [pg_readme 0.7.1 README](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/README.md)
- [Test-extension control file](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme_test_extension/pg_readme_test_extension.control)
- [Test-extension SQL fixture](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme_test_extension/pg_readme_test_extension--forever.sql)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_readme_test_extension)

`pg_readme_test_extension` is the bundled integration-test fixture for `pg_readme`. It installs commented domains, types, tables, views, routines, triggers, and processing instructions so upstream can verify `pg_extension_readme()`. It is not an application feature or a production dependency.

### Exercise the Fixture

```sql
CREATE EXTENSION pg_readme CASCADE;
CREATE EXTENSION pg_readme_test_extension;

SELECT pg_extension_readme('pg_readme_test_extension'::name);
```

Use the output to test or demonstrate the generator, then remove the fixture from a disposable database:

```sql
DROP EXTENSION pg_readme_test_extension;
```

### Boundary and Caveats

- The upstream distribution version is 0.7.1, but this fixture's control version is deliberately the literal `forever`.
- The fixture ships with `pg_readme`; current Pigsty DEB packages are 0.7.1 and RPM packages remain 0.7.0. The extension version stays `forever` on both sides.
- It is relocatable, does not require `hstore` itself, and creates sample objects with generic names. Install it only where those objects cannot collide with real application schemas.
- Its SQL surface exists to cover generator behavior and can change as tests evolve. Do not build application code against the fixture objects.
