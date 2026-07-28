## Usage

Sources:

- [Official upstream README](https://github.com/mbpcore/orgpasscheck/blob/4246a0f2f390a8a8f056758c09f05af5042ad23c/orgpasscheck_v5/README.md)
- [Official extension control file (orgpasscheck.control)](https://github.com/mbpcore/orgpasscheck/blob/4246a0f2f390a8a8f056758c09f05af5042ad23c/orgpasscheck_v5/orgpasscheck.control)
- [Official extension SQL (orgpasscheck--5.0.sql)](https://github.com/mbpcore/orgpasscheck/blob/4246a0f2f390a8a8f056758c09f05af5042ad23c/orgpasscheck_v5/orgpasscheck--5.0.sql)

`orgpasscheck` — **Enterprise password policy enforcement extension for PostgreSQL 16+**. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION orgpasscheck;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `orgpasscheck.add_blacklist(p_pattern TEXT, p_reason TEXT DEFAULT NULL)` is an extension function and returns `VOID`.
- `orgpasscheck.add_expiry_exemption(p_username TEXT, p_reason TEXT DEFAULT NULL, p_expires_at TIMESTAMPTZ DEFAULT NULL)` is an extension function and returns `VOID`.
- `orgpasscheck.change_password(p_username TEXT, p_password TEXT, p_expiry_days INTEGER DEFAULT NULL)` is an extension function and returns `VOID`.
- `orgpasscheck.create_user(p_username TEXT, p_password TEXT, p_login BOOLEAN DEFAULT true, p_superuser BOOLEAN DEFAULT false, p_expiry_days INTEGER DEFAULT NULL)` is an extension function and returns `VOID`.
- `orgpasscheck.list_expiry_exemptions()` is an extension function and returns `TABLE`.
- `orgpasscheck.purge_audit_log(p_older_than INTERVAL DEFAULT '1 year')` is an extension function and returns `INTEGER`.
- `orgpasscheck.purge_old_history()` is an extension function and returns `INTEGER`.
- `orgpasscheck.purge_user_history(p_username TEXT)` is an extension function and returns `INTEGER`.
- `orgpasscheck.record_password_history(p_username TEXT, p_password TEXT)` is an extension function and returns `VOID`.
- `orgpasscheck.remove_blacklist(p_pattern TEXT)` is an extension function and returns `VOID`.
- `orgpasscheck.remove_expiry_exemption(p_username TEXT)` is an extension function and returns `VOID`.
- `orgpasscheck.verify_password_hash(p_password TEXT, p_salt TEXT, p_stored TEXT)` is an extension function and returns `BOOLEAN`.
- `orgpasscheck.expired_passwords` is an extension-defined view.
- `orgpasscheck.policy_summary` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `5.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
