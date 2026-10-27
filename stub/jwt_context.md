## Usage

Sources:

- [README.md](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/README.md)
- [jwt_context.control](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/jwt_context.control)
- [sql/jwt_context--1.0.0.sql](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/sql/jwt_context--1.0.0.sql)
- [jwt_context.c](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/jwt_context.c)
- [.github/workflows/ci.yml](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/.github/workflows/ci.yml)

`jwt_context` 1.0.0 is an experimental signed session context for SQL and row-security policies. It verifies HS256 or ES256 signatures before accepting a JSON context. The reviewed CI matrix covers PostgreSQL 14–17. Install as a superuser; the extension is relocatable.

### Read a Verified Context

```sql
CREATE EXTENSION jwt_context;
SELECT jwt();
SELECT jwt('TENANT_ID');
RESET jwt.context;
```

A trusted administrator sets `jwt.secret` for HS256 or `jwt.pubkey` for ES256; both are superuser-only settings and hold base64url-encoded key material. The application sets `jwt.context` to a valid token from its trusted signer. `jwt()` returns the JSONB claims and `jwt(key)` reads one claim as text. The helpers `jwt_sign`, `jwt_verify`, `jwt_generate_keys` and `jwt_generate_secret` are provided for experimentation; protect signing secrets and do not expose signing privileges to application users.

### Security and Connection Pools

Applications must connect as unprivileged roles: a superuser or a user allowed to replace the verification key can forge a context. Signature validity does not establish application authorization. The prototype does not enforce token expiration or revocation; design those checks outside it before relying on a token. A leaked signed token can be replayed.

The context is session state. Reset it before a pooled connection is handed to another user; use RESET or RESET ALL as appropriate, and test the pool's actual reset behavior. Review the upstream threat model before using this experimental module for row security.
