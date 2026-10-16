## Usage

Sources:

- [README.md](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/README.md)
- [pg_mic--1.0.sql](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/pg_mic--1.0.sql)
- [SECURITY.md](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/SECURITY.md)
- [pg_mic_compat.h](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/pg_mic_compat.h)
- [pg_mic.control](https://github.com/jconway/pg_mic/blob/8a8a3079e942d398deb6b4de160cc41089744f09/pg_mic.control)

`pg_mic` adds Biba integrity checks to roles, relations and functions. HIGH, MEDIUM and LOW labels restrict reading lower-integrity data and writing higher-integrity data. This controls integrity, not confidentiality; SQL privileges still apply. The control/SQL version is 1.0, while upstream security documentation describes beta status.

### Core Workflow

```ini
shared_preload_libraries = 'pg_mic'
```

```sql
CREATE EXTENSION pg_mic;
CREATE ROLE integrity_reader;
CREATE TABLE verified_facts (id integer, fact text);
SELECT mic_promote_table('verified_facts', 'high');
GRANT SELECT ON verified_facts TO integrity_reader;
SELECT mic_get_session_level();
SELECT mic_get_table_level('verified_facts');
SET ROLE integrity_reader;
SELECT * FROM verified_facts;
RESET ROLE;
```

### Operational Boundaries

Add `pg_mic` to `shared_preload_libraries`, restart PostgreSQL, and create the extension in each database as a superuser. Upstream documents PostgreSQL 14–19 and the development branch; this source revision contains version guards for those server APIs.

A superuser has HIGH integrity, while an unlabeled role or user object defaults to LOW. New tables and functions receive their creator’s level. SELECT requires the subject’s level to be no higher than the object’s; INSERT and DELETE require it to be no lower; UPDATE applies both rules. Built-in functions and system catalog reads have exemptions. Existing unlabeled objects can therefore become inaccessible to HIGH sessions until reviewed and relabeled.

`mic_get_session_level`, `mic_get_role_level` and `mic_get_table_level` inspect labels. Superusers administer labels with `mic_set_role_level`, `mic_promote_table`, `mic_promote_object`, `mic_promote_extension` and `mic_initialize_labels`. `mic_labels` and `mic_role_labels` expose label views. Only promote reviewed data and code.

`mic.mode` defaults to `enforcing`; `permissive` warns and allows violations, and `disabled` skips checks. Mode changes need a configuration reload. `mic.debug_audit` enables decision logging. RLS checks are conservative and can reject paths PostgreSQL would otherwise allow. Function demotion invalidates plans, but function pointers cached outside the plan cache can retain old behavior until the session ends; reconnect affected long-lived sessions after a demotion. New server utility commands require explicit upstream coverage.
