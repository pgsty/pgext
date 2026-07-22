## Usage

Sources:

- [Official documentation](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/doc/user_info.md)
- [Official extension SQL](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/sql/user_info--0.0.1.sql)
- [Official extension control file](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/user_info.control)

`user_info` version `0.0.1` provides catalog-reporting functions for role membership, owned objects, and directly granted privileges. The project is archived and its SQL targets old PostgreSQL catalogs, so use it as historical diagnostic code rather than a current authorization authority.

### Core Workflow

Functions accept a role OID, role name, or, where available, no argument for the current user:

```sql
CREATE EXTENSION user_info;

SELECT * FROM user_objects('app_user');
SELECT * FROM granted_roles('app_user');
SELECT * FROM granted_roles_pretty('app_user');
SELECT * FROM accessible_objects('app_user');
```

`user_objects(...)` unions several system-catalog and information-schema reports, so a single table can appear as a relation, table, row type, and array type. Duplicates across object categories are expected rather than automatically deduplicated.

### Reporting Limits

`granted_roles(...)` recursively returns membership paths and levels. `granted_roles_pretty(...)` formats the same hierarchy for people. `accessible_objects(...)` reports privileges explicitly granted directly to the selected role.

The privilege report does not include default privileges or permissions inherited through other roles, and therefore does not calculate effective access. It also relies on catalog layouts and object classes from the extension's era; installation or queries may fail on modern PostgreSQL, and newer object types are not covered. Validate every result against current PostgreSQL catalog and privilege functions before using it for review, and do not use it alone for security or compliance decisions.
