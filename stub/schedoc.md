


## Usage

> [schedoc: Schema documentation via COMMENT on PostgreSQL objects](https://github.com/ZeroGachis/pg_schedoc)

Builds automatic documentation from `COMMENT` on PostgreSQL objects. Requires the `ddl_historization` extension.

### Setup

```sql
CREATE EXTENSION schedoc CASCADE;
SELECT schedoc_start();
```

### Adding Column Documentation

Set comments on columns in JSON format with predefined fields:

```sql
COMMENT ON COLUMN my_table.id IS '{"status": "private"}';
COMMENT ON COLUMN my_table.email IS '{"status": "public"}';
COMMENT ON COLUMN my_table.name IS '{"status": "internal"}';
```

### Querying Documentation

Query the parsed column comments:

```sql
SELECT * FROM schedoc_column_comments;
```

Result:

```
 databasename | tablename | columnname | status
--------------+-----------+------------+---------
 mydb         | my_table  | id         | private
 mydb         | my_table  | email      | public
 mydb         | my_table  | name       | internal
```

### Use Case

Cross-reference column metadata with other systems (e.g., Django `db_comment`, DBT docs) to define data contracts between developers and data analysts.
