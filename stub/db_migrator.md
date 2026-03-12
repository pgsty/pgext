

## Usage

> [db_migrator: Tools to migrate other databases to PostgreSQL](https://github.com/cybertec-postgresql/db_migrator)

A framework for migrating databases from other data sources to PostgreSQL using foreign data wrappers and source-specific plugins.

### Enabling

```sql
CREATE EXTENSION db_migrator;
```

### Available Plugins

- **ora_migrator** - Oracle migration
- **mysql_migrator** - MySQL/MariaDB migration
- **mssql_migrator** - Microsoft SQL Server migration

### Complete Migration Example (Oracle)

```sql
-- Setup (as superuser)
CREATE EXTENSION oracle_fdw;
CREATE SERVER oracle FOREIGN DATA WRAPPER oracle_fdw
    OPTIONS (dbserver '//dbserver.mydomain.com/ORADB');
GRANT USAGE ON FOREIGN SERVER oracle TO migrator;
CREATE USER MAPPING FOR migrator SERVER oracle
    OPTIONS (user 'orauser', password 'orapwd');

-- Migrate (as migrator user)
CREATE EXTENSION ora_migrator;

SELECT db_migrate(
    plugin => 'ora_migrator',
    server => 'oracle',
    only_schemas => '{TESTSCHEMA1,TESTSCHEMA2}'
);
```

### Step-by-Step Migration

For more control, execute migration in stages:

```sql
-- 1. Create staging schemas and snapshot metadata
SELECT db_migrate_prepare(
    plugin => 'ora_migrator',
    server => 'oracle',
    only_schemas => '{SCHEMA1}'
);

-- 2. Review and modify staging data
-- Edit pgsql_stage tables to customize type mappings, rename objects, etc.
UPDATE pgsql_stage.tables SET migrate = TRUE WHERE ...;

-- 3. Create schemas and migrate data
SELECT db_migrate_mkforeign(plugin => 'ora_migrator', server => 'oracle');
SELECT db_migrate_tables(plugin => 'ora_migrator');

-- 4. Create constraints and indexes
SELECT db_migrate_constraints(plugin => 'ora_migrator');
SELECT db_migrate_indexes(plugin => 'ora_migrator');

-- 5. Cleanup
SELECT db_migrate_finish();
```

### Key Features

- Migrates tables, sequences, indexes, constraints, views, functions
- Data type mapping from source to PostgreSQL types (customizable)
- Continues on errors, reporting which objects failed
- Uses FDW staging schema for metadata inspection before migration
- Schema and object name translation via plugin functions
