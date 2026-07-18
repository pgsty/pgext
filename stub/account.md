## Usage

Sources:

- [bizniff@account registry page](https://database.dev/bizniff/account)
- [database.dev package installation guide](https://database.dev/docs/install-a-package)

`account` is a pure-SQL Trusted Language Extension for Supabase-oriented team, permission, invitation, and billing support. Registry package `bizniff@account` version `0.1.2` creates a fixed `account` schema containing team and membership tables, configuration, row-level security policies, and supporting functions and triggers.

### Generate and inspect the migration

Install `pg_tle` as required by database.dev, then use the official `dbdev` client to generate a migration. Review that migration before applying it.

```shell
dbdev add -o ./migrations -s extensions -v 0.1.2 package -n "bizniff@account"
```

After the generated migration has been applied, the published SQL exposes configuration helpers such as:

```sql
SELECT account.get_config();
SELECT account.is_set('enable_account_teams');
```

The package assumes a Supabase-compatible environment, including `auth.users`, `auth.uid()`, the `anon`, `authenticated`, and `service_role` roles, and `extensions.uuid_generate_v4()`. It also installs an `auth.users` trigger and grants privileges on its objects. Do not apply it to a generic PostgreSQL database without first reconciling those dependencies and reviewing its RLS and privilege model.
