## Usage

Sources:

- [Extension control file](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security.control)
- [Extension SQL](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security--1.0.0.sql)
- [C implementation](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security.c)

`postgres_security` 1.0.0 is a minimal type-development example, not a security feature. It creates a schema-local base type named `postgres_security.text` whose storage and input/output routines mirror PostgreSQL's built-in text type. It provides no encryption, masking, validation, access control, authentication, or auditing.

### Inspection Workflow

Install only in an isolated development database and always schema-qualify the custom type:

```sql
CREATE EXTENSION postgres_security;

CREATE TABLE security_type_demo (
  value postgres_security.text
);

INSERT INTO security_type_demo VALUES ('example');
SELECT pg_typeof(value), value FROM security_type_demo;
```

Do not put the `postgres_security` schema before `pg_catalog` in `search_path`; unqualified `text` could then resolve to the custom type unexpectedly.

### Object Surface

The install script creates the `postgres_security` schema, the custom `text` base type, and four C I/O functions: `textin`, `textout`, `textrecv`, and `textsend`. It creates no comparison operators, casts, operator classes, constraints, or security policies. The custom type has a distinct type OID and does not automatically inherit built-in text functions or operator behavior.

### Operational Notes

The control file says the extension is relocatable, but its SQL hard-codes creation and use of the `postgres_security` schema; do not rely on relocation. Installation can also conflict with a pre-existing schema of that name. Avoid this extension in application schemas, ORMs, or security designs. If encountered in an existing database, inventory dependent columns and migrate them to `pg_catalog.text` under controlled casts or data export/import rather than assuming binary or operator compatibility.
