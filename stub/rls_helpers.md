## Usage

Sources:

- [Official database.dev package page](https://database.dev/mansueli/rls_helpers)

`rls_helpers` provides small PL/pgSQL procedures for assuming an authenticated or anonymous Supabase-style identity while testing row-level security policies. Use it in development and test sessions; changing the session identity can alter every RLS-protected query executed on the same connection.

### Core Workflow

The official package workflow generates a versioned migration with the `dbdev` CLI. Review that migration before applying it to the target database:

```bash
dbdev add -o ./migrations -s extensions -v 1.0.0 package -n "mansueli@rls_helpers"
```

Run an RLS test inside an explicit transaction, choose the simulated identity, execute the application query, and restore the session afterward:

```sql
BEGIN;

CALL auth.login_as_user('rodrigo@contoso.com');
SELECT * FROM private.user_documents;

CALL auth.logout();
ROLLBACK;
```

To exercise policies for unauthenticated requests, use the anonymous helper:

```sql
BEGIN;
CALL auth.login_as_anon();
SELECT * FROM public.visible_documents;
CALL auth.logout();
ROLLBACK;
```

### User-Facing Procedures

- `auth.login_as_user(text)` selects the identity identified by an email address for subsequent policy checks.
- `auth.login_as_anon()` switches the session to the anonymous test context.
- `auth.logout()` clears the simulated login context.

### Operational Notes

The upstream page documents these helpers as a way to test or simulate RLS, not as an application authentication system. Keep tests on isolated connections, bracket them with a transaction, and always call `auth.logout()` or roll back before returning a pooled connection. The package page documents version `1.0.0`; inspect the generated migration and verify its role, schema, and privilege assumptions against the Supabase/PostgreSQL environment under test.
