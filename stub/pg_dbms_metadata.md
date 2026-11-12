## Usage

Sources:

- [Version 1.1.0 README](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/README.md)
- [SQL definitions](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/sql/pg_dbms_metadata--1.1.0.sql)
- [Upgrade script](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/updates/pg_dbms_metadata--1.0.0--1.1.0.sql)
- [Release notes](https://github.com/HexaCluster/pg_dbms_metadata/releases/tag/v1.1.0)

`pg_dbms_metadata` 1.1.0 extracts PostgreSQL object DDL through an Oracle-style API. It is a SQL/PL/pgSQL extension with no preload requirement. It implements the GET functions, not Oracle FETCH or XML interfaces.

### Enabling

```sql
CREATE EXTENSION pg_dbms_metadata;
```

### GET_DDL

Extract DDL of database objects. Supported types: TABLE, VIEW, SEQUENCE, PROCEDURE, FUNCTION, TRIGGER, INDEX, CONSTRAINT, CHECK_CONSTRAINT, REF_CONSTRAINT, TYPE, ENUM.

```sql
SELECT dbms_metadata.get_ddl('TABLE', 'employees', 'public');
SELECT dbms_metadata.get_ddl('VIEW', 'active_users', 'myschema');
SELECT dbms_metadata.get_ddl('FUNCTION', 'calculate_tax', 'public');
SELECT dbms_metadata.get_ddl('INDEX', 'idx_name', 'public');

-- Schema is optional; uses search_path when omitted
SELECT dbms_metadata.get_ddl('SEQUENCE', 'my_seq');
```

### GET_DEPENDENT_DDL

Extract DDL of all dependent objects of a specified type for a base object. Supported: SEQUENCE, TRIGGER, CONSTRAINT, REF_CONSTRAINT, INDEX, ENUM.

```sql
SELECT dbms_metadata.get_dependent_ddl('CONSTRAINT', 'employees', 'public');
SELECT dbms_metadata.get_dependent_ddl('INDEX', 'orders', 'sales');
SELECT dbms_metadata.get_dependent_ddl('TRIGGER', 'accounts');
```

### GET_GRANTED_DDL

Extract SQL statements to recreate granted privileges and roles. Supported: ROLE_GRANT.

```sql
SELECT dbms_metadata.get_granted_ddl('ROLE_GRANT', 'user_test');
```

### SET_TRANSFORM_PARAM

Customize DDL output with session-level transform parameters:

```sql
-- Append SQL terminator (;) to each DDL statement
CALL dbms_metadata.set_transform_param('SQLTERMINATOR', true);

-- Exclude constraints from TABLE DDL
CALL dbms_metadata.set_transform_param('CONSTRAINTS', false);

-- Exclude foreign keys from TABLE DDL
CALL dbms_metadata.set_transform_param('REF_CONSTRAINTS', false);

-- Exclude partitioning clauses
CALL dbms_metadata.set_transform_param('PARTITIONING', false);

-- Exclude storage parameters
CALL dbms_metadata.set_transform_param('STORAGE', false);

-- Reset all params to defaults
CALL dbms_metadata.set_transform_param('DEFAULT', true);
```

### 1.1.0 Behavior and Upgrade

ENUM extraction is now supported and explicitly supplied schemas qualify enum names. For overloaded functions or procedures, an unqualified routine name returns every matching overload; append argument types to select one:

```sql
SELECT dbms_metadata.get_ddl('FUNCTION', 'my_function(char, char)', 'my_schema');
ALTER EXTENSION pg_dbms_metadata UPDATE TO '1.1.0';
```

1.1.0 also fixes quoted table comments and boolean transform defaults. Install the new extension files before the SQL update. Upstream documents PostgreSQL 9.1–18 and tests this release's upgrade path on 12–19; package support is separate.

DDL visibility is not limited to objects owned by the caller. When a schema is omitted, lookup uses `search_path` and requires schema USAGE; TRIGGER and REF_CONSTRAINT calls require an explicit schema. Restrict API access when object definitions are sensitive. Session transform settings affect subsequent extraction calls and can be reset with the DEFAULT transform; this API does not implement a full database backup.
