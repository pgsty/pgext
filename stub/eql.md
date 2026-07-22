## Usage

Sources:

- [Official `cipherstash@eql` 2.2.1 package page](https://database.dev/cipherstash/eql)
- [Official EQL documentation](https://cipherstash.com/docs)
- [Current upstream 3.0.2 release](https://github.com/cipherstash/encrypt-query-language/releases/tag/eql-3.0.2)

`eql` (packaged as `cipherstash@eql`) provides PostgreSQL types, functions, operators, and index configuration for searchable encrypted data. The cataloged package version is 2.2.1 and uses the `eql_v2` API. EQL supplies the database-side representation and query surface; CipherStash Proxy or Protect.js must encrypt and decrypt values outside PostgreSQL.

### Core Workflow

Generate and apply the pinned dbdev migration, then confirm the installed EQL API version:

```sh
dbdev add -o ./migrations -s extensions -v 2.2.1 package -n "cipherstash@eql"
psql -f ./migrations/*.sql
```

```sql
SELECT eql_v2.version();
```

Define an encrypted column with the persistent public type, register its plaintext type, and optionally add a searchable index configuration:

```sql
CREATE TABLE users (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  encrypted_email public.eql_v2_encrypted
);

SELECT eql_v2.add_column('users', 'encrypted_email', 'text');
SELECT eql_v2.add_search_config(
  'users', 'encrypted_email', 'unique', 'text'
);
```

Write and query this column through a configured CipherStash client. Do not insert plaintext directly into `public.eql_v2_encrypted` and expect EQL to encrypt it.

### Important Objects

- `public.eql_v2_encrypted` is the encrypted column type, backed by validated `jsonb` data.
- `public.eql_v2_configuration` stores searchable-encryption configuration; `public.eql_v2_configuration_state` is its state type.
- The `eql_v2` schema contains the functions and operators used to configure and query encrypted data.
- `eql_v2.add_column(...)` registers an encrypted column and its logical plaintext type.
- `eql_v2.add_search_config(...)` adds a search configuration such as a unique/equality index.
- `eql_v2.version()` reports the installed API version.

### Privileges, Upgrades, and Version Boundary

Installation needs `CREATE` on the database and public schema. The migration role also needs `ALTER` on target tables and read/write access to `public.eql_v2_configuration`; runtime roles normally need `USAGE` on `eql_v2`, function execution, configuration read access, and ordinary table privileges.

Dropping `public.eql_v2_encrypted` can remove dependent application columns. EQL deliberately leaves persistent public types and configuration data outside the replaceable `eql_v2` schema, so treat uninstall and major upgrades as data migrations.

Upstream GitHub release 3.0.2 uses a substantially different `eql_v3` domain-based API and a release SQL installer. The upstream README warns that dbdev publication can lag GitHub. Do not mix v3 examples with the cataloged `cipherstash@eql` 2.2.1 package; pin the distribution and follow its matching documentation.
