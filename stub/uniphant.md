## Usage

Sources:

- [Pinned upstream README](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/README.md)
- [Version 1.5 application SQL](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/uniphant--1.5.sql)
- [Pinned extension control file](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/uniphant.control)
- [Pinned WebAuthn verification function](https://github.com/truthly/uniphant/blob/cf76fd388ebb06977582df7fb3d5e8ec5c9a2e44/FUNCTIONS/api/verify_assertion.sql)

uniphant version 1.5 is a full-stack demonstration of PostgreSQL, PostgREST, WebAuthn, nginx, and browser clients. Its fixed public and api schemas implement users, credentials, roles, resources, permissions, access tokens, WebAuthn sign-up/sign-in, and generated OpenAPI metadata.

### Demo-only installation

The expected database roles must exist, and the webauthn dependency is installed through CASCADE:

```sql
CREATE ROLE web_anon NOLOGIN;
CREATE ROLE postgrest NOLOGIN;

CREATE EXTENSION uniphant
WITH SCHEMA public
CASCADE;

SELECT api.openapi_swagger();
```

The complete demonstration also requires PostgREST, a trusted reverse proxy, browser WebAuthn support, TLS for non-local use, and the related PostgreSQL cryptographic dependencies. The README's deployment recipe targets Ubuntu 20.04, PostgreSQL 13, PostgREST 7.0.1, and old request-header conventions.

### Security-sensitive side effects

Installation grants web_anon usage on api and webauthn, grants web_anon to postgrest, and grants web_anon SELECT on all existing tables in both api and public. It creates many SECURITY DEFINER functions. The first successful non-anonymous sign-up receives the admin role automatically. These are deliberate demo shortcuts with major consequences in an existing database.

Do not install uniphant into a general-purpose or production public schema. A modern deployment must review every grant and definer function, remove first-user bootstrap behavior, verify fixed search paths, adapt PostgREST request settings, enforce trusted proxy headers and TLS, protect access-token storage, and retest current WebAuthn semantics. Use an empty disposable database for evaluation and compare the generated object/privilege inventory before and after installation.
