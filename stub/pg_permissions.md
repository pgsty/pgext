## Usage

Sources:

- [README.md](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/README.md)
- [pg_permissions.control](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/pg_permissions.control)
- [pg_permissions--1.5--1.6.sql](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/pg_permissions--1.5--1.6.sql)

`pg_permissions` compares actual privileges with a desired baseline and can generate or apply corrective grants. Version 1.6 requires PostgreSQL 15 or later. It is pure SQL, requires no preload, and is not relocatable; choose the installation schema initially. The installer needs CREATE privileges there.

### Define and Inspect a Baseline

The application role and schema in this example must already exist:

```sql
CREATE EXTENSION pg_permissions SCHEMA public;
INSERT INTO public.permission_target
  (role_name, permissions, object_type, schema_name)
VALUES ('appuser', '{SELECT,INSERT,UPDATE,DELETE}', 'TABLE', 'appschema'),
       ('appuser', '{USAGE}', 'SCHEMA', 'appschema');
SELECT * FROM public.permission_diffs();
SELECT * FROM public.permission_diff_summary ORDER BY objects DESC;
SELECT public.permission_apply();
```

A missing flag of true means a required privilege is absent; false means an extra privilege. A NULL object name covers all objects of the specified type in that schema. Audit scope is limited to targeted objects, managed roles and roles inheriting from them. Use `all_permissions` for a broader review. Superusers are omitted because they implicitly hold all privileges.

### Review Before Applying

`permission_fix_sql()` returns corrective SQL. `permission_apply()` defaults to a dry run; `permission_apply(false)` executes it using the caller's privileges. The writable `granted` column on permission views also executes GRANT or REVOKE. Restrict both baseline writes and permission-changing access.

The object views cover databases, schemas, tables, views, columns, functions and sequences. Their `source` column distinguishes ownership, direct grants, PUBLIC and inherited-role privileges. `PUBLIC` can be an explicit target grantee. `permission_target_snapshot(role, schema)` adopts current grants as the baseline; this changes the desired state and should be deliberate.

`default_acl_target`, `default_acl_permissions` and `default_acl_diffs()` separately handle ALTER DEFAULT PRIVILEGES. A default privilege affects future objects; it does not repair existing grants. Baseline tables are readable by PUBLIC but writable only by the extension owner and explicitly authorized roles. `role_dependencies()` inspects dependencies; `role_drop()` can reassign ownership and drop roles and warrants a separate review.

### Upgrade Boundary

Follow the version-by-version update scripts. For an installed 1.4 database:

```sql
ALTER EXTENSION pg_permissions UPDATE TO '1.5';
ALTER EXTENSION pg_permissions UPDATE TO '1.6';
```

Version 1.6 restores the missing 1.2-to-1.3 path. Its 1.5-to-1.6 migration rewrites the baseline: empty permission lists are removed, COLUMN entries without a column become table-level entries, and duplicates are merged. Back up and review the desired state before upgrading. MAINTAIN privileges are rejected before PostgreSQL 17.
