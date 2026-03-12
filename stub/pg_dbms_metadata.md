

## Usage

> [pg_dbms_metadata: Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL](https://github.com/HexaCluster/pg_dbms_metadata)

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
