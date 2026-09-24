## Usage

Sources:

- [README](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/README.md)
- [Control](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgauth/pgauth.control)
- [SQL](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgauth/pgauth--0.1.0.sql)

`pgauth` adds password hashing, signed JWT sessions and SQL helpers for row-level-security policies. This 0.1.0 PGEverything source snapshot targets its PostgreSQL 16 deployment. Its claims live in a mutable session setting, so use a trusted application gateway: it is not an identity boundary against clients allowed to issue arbitrary SQL.

### Core Workflow

Install `pgcrypto`, `pgjwt` and `pgauth`, with PL/pgSQL available. Initialize the private signing secret as the extension owner; the PGEverything container provisions it from `PGE_JWT_SECRET`. Use a fresh private value, not the container's development default. None of these helpers requires a pgauth preload library.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pgjwt;
CREATE EXTENSION pgauth;
INSERT INTO auth.secret (value)
VALUES (encode(gen_random_bytes(32), 'hex'));

SELECT auth.register('alice@example.com', :'user_password');
BEGIN;
SELECT auth.authenticate(auth.login('alice@example.com', :'user_password'));
SELECT auth.uid(), auth.role();
COMMIT;
```

The psql variable `user_password` must be supplied securely. Authentication claims are transaction-local: authenticate and perform protected work inside the same transaction. The example only initializes a new database; do not replace a live signing secret casually.

### SQL Surface

- `auth.register(text,text,text)` hashes a password and returns a user UUID. Its optional third argument stores a role claim; it does not create a PostgreSQL role.
- `auth.login(text,text)` returns a JWT with a one-hour expiry, or NULL for bad credentials.
- `auth.authenticate(text)` validates the signature/expiry and sets `auth.claims`; it returns a boolean. `auth.uid()` and `auth.role()` read those claims.
- `auth.users` and `auth.secret` are extension configuration tables whose data is included in logical dumps. Protect those dumps as authentication data.

### RLS and Security Boundaries

Define application-table policies using `auth.uid()` and use a non-superuser application role; table owners and other RLS-bypassing roles are not constrained in the ordinary way. Restrict public registration and its caller-supplied role argument before using role claims for authorization. The extension grants public API execution but restricts direct access to the secret accessor/table.

Untrusted SQL can change `auth.claims` directly, so the accessors alone cannot authenticate such a client. Security-definer paths also include `public`; keep that schema outside untrusted write access. There is no refresh-token API. The extension creates the fixed `auth` schema and has no explicit upstream license declaration.
