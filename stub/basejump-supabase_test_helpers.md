## Usage

Sources:

- [Official README](https://github.com/usebasejump/supabase-test-helpers/blob/d90a51f197d49f0a2af155462669cf1dc0218534/README.md)
- [Extension SQL for version 0.0.6](https://github.com/usebasejump/supabase-test-helpers/blob/d90a51f197d49f0a2af155462669cf1dc0218534/supabase_test_helpers--0.0.6.sql)
- [database.dev package page](https://database.dev/basejump/supabase_test_helpers)

`basejump-supabase_test_helpers` is the database.dev package name for Basejump's Supabase test-helper extension. It creates test-only functions for users, JWT/role impersonation, row-level-security assertions, and frozen time. The upstream project explicitly recommends activating it only inside a transactional pgTAP test, not persistently in production.

### Test-Transaction Workflow

Install the package through dbdev, then create the provider-named extension inside a transaction that will be rolled back:

```sql
SELECT dbdev.install('basejump-supabase_test_helpers');

BEGIN;
CREATE EXTENSION "basejump-supabase_test_helpers";

SELECT plan(2);

CREATE TABLE public.test_accounts (id bigint PRIMARY KEY);
ALTER TABLE public.test_accounts ENABLE ROW LEVEL SECURITY;

SELECT tests.rls_enabled('public', 'test_accounts');
SELECT tests.create_supabase_user('test_owner', 'owner@example.test');

SELECT * FROM finish();
ROLLBACK;
```

The repository's native control and SQL files use `supabase_test_helpers`, while database.dev exposes the namespaced `basejump-supabase_test_helpers` package. Follow the installation mechanism's reported extension name rather than renaming control files manually.

### Important Helpers

- `tests.create_supabase_user` inserts a synthetic row in `auth.users` and returns its UUID; `tests.get_supabase_user` and `tests.get_supabase_uid` retrieve it by a memorable test identifier.
- `tests.authenticate_as`, `tests.authenticate_as_service_role`, and `tests.clear_authentication` change the transaction-local role and `request.jwt.claims` setting to exercise Supabase authorization paths.
- `tests.rls_enabled` returns pgTAP assertions for one table or every table in a schema.
- `tests.freeze_time` and `tests.unfreeze_time` cooperate with `test_overrides.now` to make time-dependent tests deterministic.

When a security-definer function fixes its own search path, frozen time only works if the test temporarily puts `test_overrides` before `pg_catalog` in that function's path. Do this only inside a test transaction and let rollback restore the catalog settings.

### Prerequisites and Safety

The extension requires `pgtap` and assumes a Supabase layout with `auth.users`, the `anon`, `authenticated`, and `service_role` roles, plus the UUID helper used by its SQL. Several helpers are security-definer functions and deliberately impersonate privileged application roles. Run tests in an isolated database or disposable branch and never against production identities.

The repository contains a 0.0.6 install script while its checked-in control file still names 0.0.5 as the default. Pin the requested version through the package manager and verify the installed version in `pg_extension`. Rollback removes objects created in the transaction, but external side effects performed by application code under test may not be transactional.
