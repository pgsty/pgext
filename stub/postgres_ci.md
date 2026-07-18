## Usage

Sources:

- [Pinned extension control file](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/postgres_ci.control)
- [Complete version 0.3 installation SQL](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/postgres_ci--0.3.sql)
- [Pinned PUBLIC privilege grants](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/packages.sql)
- [Pinned user creation function](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/functions/users/add.sql)
- [Pinned password reset function](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/functions/password/reset.sql)
- [Pinned GitHub secret getter](https://github.com/postgres-ci/core/blob/a45eab1ad37e96eba05b5eed24cba9cc71926647/src/functions/project/get_github_secret.sql)

postgres_ci version 0.3 is the database backend of an abandoned Postgres-CI service. It stores application users and sessions, projects and webhook secrets, Git commits, builds, build parts and tests, worker queues, and notifications. It requires PL/pgSQL, pgcrypto, and pg_trgm and creates multiple hard-coded schemas.

### Isolated installation audit

Install only in a disposable, isolated database and inspect the definer surface before calling any service API:

```sql
CREATE EXTENSION postgres_ci CASCADE;

SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_ci';

SELECT n.nspname, p.proname, p.prosecdef
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
WHERE n.nspname IN (
  'auth', 'hook', 'users', 'build', 'project',
  'password', 'notification', 'postgres_ci'
)
ORDER BY n.nspname, p.proname;

REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA
  auth, hook, users, build, project,
  password, notification, postgres_ci
FROM PUBLIC;
```

The application functions implement login, webhook ingestion, project mutation, and worker fetch/accept flows. They are an inseparable historical service API, not general CI primitives.

### Critical authentication and privilege flaws

Nearly every service function is SECURITY DEFINER without a fixed safe search_path. The installation explicitly grants PUBLIC usage on all service schemas and PUBLIC execution on all functions. Any database role can therefore reach privileged routines unless those grants are revoked. The user-creation routine accepts an application is_superuser flag without authorizing the caller; password.reset can replace another user's password; project.get and project.get_github_secret return stored webhook secrets.

Passwords use a fast salted SHA-1 digest rather than a modern password KDF. GitHub webhook secrets are stored as plaintext. Sessions are kept in an unlogged table, can disappear after a crash, and use a one-hour sliding expiry. Error messages distinguish unknown users from invalid passwords, enabling account enumeration.

Do not expose this extension to untrusted roles or a network-facing application. Revoking PUBLIC is only the first step: every definer function needs a fixed pg_catalog-first search_path, explicit caller authorization, least-privilege ownership, secret encryption and rotation, a modern password migration, rate limiting, audit logging, and adversarial tests. A secure replacement is safer than retrofitting this code.

The extension is non-relocatable and creates generic auth, hook, users, build, project, password, and notification schemas. Its last code commit was in 2016, the README has no usable operating documentation, and the catalog lifecycle is abandoned. Preserve it only to read or export legacy state; do not start a new CI service on it.
