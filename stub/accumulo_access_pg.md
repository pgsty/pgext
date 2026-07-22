## Usage

Sources:

- [Official README](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/README.md)
- [Official extension control file](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/accumulo_access_pg.control)
- [Official Rust implementation](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/src/lib.rs)
- [Official package metadata](https://github.com/larsw/accumulo-access-pg/blob/8e201122665eac9ff2a233727c057eef5b9eee3b/Cargo.toml)

`accumulo_access_pg` 0.1.5 parses and evaluates Apache Accumulo-style authorization expressions inside PostgreSQL. Its main use is applying expressions such as `label1 & (label2 | label3)` in a row-level security policy. The extension evaluates labels supplied by the caller; it does not authenticate users or obtain trusted authorization claims by itself.

### Core Workflow

Store an authorization expression with each protected row and pass a trusted comma-separated label set to `sec_authz_check`. A policy can reject rows when the expression or token list is null or empty.

```sql
CREATE EXTENSION accumulo_access_pg;

CREATE TABLE protected_notes (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    body text NOT NULL,
    authz_expr text NOT NULL
);

ALTER TABLE protected_notes ENABLE ROW LEVEL SECURITY;

CREATE POLICY protected_notes_read ON protected_notes
USING (
    sec_authz_check(
        authz_expr,
        current_setting('app.authorizations', true)
    )
);

SET app.authorizations = 'finance,eu';
SELECT * FROM protected_notes;
```

Do not let an untrusted database user freely set the value used by the policy. The upstream example uses a session GUC for illustration, but a user who can choose `app.authorizations` can grant themselves labels. Derive claims through a controlled login/session setup path or a security-reviewed helper.

### Functions

- `sec_authz_check(expression, tokens)` evaluates one expression against comma-separated tokens; malformed expressions raise an error.
- `sec_expr_as_json(expression)` and `sec_expr_as_json_string(expression)` expose the parsed expression for inspection.
- `sec_authz_cache_stats()` returns the `hits`, `misses`, and `size` fields of the parser/evaluation cache.
- `sec_authz_clear_cache()` clears that cache and returns success as boolean.

Expression syntax and escaping follow the upstream `accumulo-access` Rust library. Test quoted labels, Unicode, whitespace, invalid syntax, and the exact semantics required by the application.

### Installation and Security Boundaries

The pinned 0.1.5 package enables only the pgrx PostgreSQL 15 feature. Its control file marks installation as superuser-only and non-relocatable. Build and test a matching artifact before using another PostgreSQL major version.

Row-level security depends on every participating role, policy, bypass privilege, and claim source—not only on `sec_authz_check`. Validate behavior for table owners, superusers, roles with `BYPASSRLS`, pooled connections, reset/reuse of session state, and malformed claims. Cache statistics are operational aids, not an authorization audit trail.
