## Usage

Sources:

- [docs/auth.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/auth.rst)
- [auth/auth--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/auth/auth--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [auth/auth.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/auth/auth.control)

`auth` is Oliphant’s privilege-management extension. It inspects, copies or moves authorizations and can save table grants before a table is replaced.

### Core Workflow

```sql
CREATE EXTENSION auth;
SELECT * FROM role_auths(current_user::name);
```

### Operational Boundaries

These are legacy pure-SQL/PLpgSQL utilities originally documented for PostgreSQL 9.x. Current-major compatibility is unverified. Installation uses the control defaults and needs administrator privileges; no shared library or preload is required. `copy_role_auths` adds missing grants to a destination role; `move_role_auths` also removes them from the source. The caller still needs grant/revoke authority, usually a superuser. Object ownership is not transferred. `store_table_auths` writes `stored_table_auths`; `restore_table_auths` replays the saved grants. Inspect saved rows before restoring, especially after table replacement or manual edits.
