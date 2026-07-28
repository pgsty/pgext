## Usage

Sources:

- [Official upstream README](https://github.com/bisoftbilgi/bisoft-postgresql-toolkit/blob/9a49035d2d370bdd8daefe7e8bdf67d29e0f648d/password_profile/README.md)
- [Official extension control file (password_profile.control)](https://github.com/bisoftbilgi/bisoft-postgresql-toolkit/blob/9a49035d2d370bdd8daefe7e8bdf67d29e0f648d/password_profile/password_profile.control)
- [Official implementation source](https://github.com/bisoftbilgi/bisoft-postgresql-toolkit/blob/9a49035d2d370bdd8daefe7e8bdf67d29e0f648d/password_profile/src/lib.rs)

`password_profile` — Enterprise‑grade password policy and authentication hardening for PostgreSQL. Built with Rust + pgrx (v0.16.1). Tested on PostgreSQL 16-18. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION password_profile;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_to_blacklist` is an extension function.
- `check_password` is an extension function.
- `check_password_expiry` is an extension function.
- `check_user_access` is an extension function.
- `clear_login_attempts` is an extension function.
- `get_lock_cache_stats()` is an extension function.
- `get_password_stats` is an extension function.
- `init_login_attempts_table()` is an extension function.
- `is_user_locked` is an extension function.
- `load_blacklist_from_file` is an extension function.
- `record_failed_login` is an extension function.
- `record_password_change` is an extension function.
- `remove_from_blacklist` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
