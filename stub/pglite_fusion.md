
--------

## Usage

> https://github.com/frectonz/pglite-fusion/blob/main/README.md


Here's some demo usage.

```sql
-- Load PG extension
CREATE EXTENSION pglite_fusion;

-- Create a table with an SQLite column
CREATE TABLE people (
                        name     TEXT NOT NULL,
                        database SQLITE DEFAULT init_sqlite('CREATE TABLE todos (task TEXT)')
);

-- Insert a row into the people table
INSERT INTO people VALUES ('frectonz');

-- Create a todo for "frectonz"
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''solve multitenancy'')'
               )
WHERE name = 'frectonz';

-- Create a todo for "frectonz"
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''buy milk'')'
               )
WHERE name = 'frectonz';

-- Fetch frectonz's info
SELECT
    name,
    (
        SELECT json_agg(get_sqlite_text(sqlite_row, 0))
        FROM query_sqlite(
                database,
                'SELECT * FROM todos'
             )
    ) AS todos
FROM
    people
WHERE
    name = 'frectonz';
```


--------

## API Doc


### `empty_sqlite`

Creates an empty SQLite database and returns it as a binary object. This can be used to initialize an empty SQLite database in a PostgreSQL column.

#### Example Usage:

```sql
SELECT empty_sqlite();
```

-----

### `query_sqlite`

Executes a SQL query on a SQLite database stored as a binary object and returns the result as a table of JSON-encoded rows. This function is useful for querying SQLite databases stored in PostgreSQL columns.

#### Parameters:

- `sqlite`: The SQLite database to query, stored as a binary object.
- `query`: The SQL query string to execute on the SQLite database.

#### Example Usage:

```sql
SELECT * FROM query_sqlite(
        database,
        'SELECT * FROM todos'
              );
```

-----

### `execute_sqlite`

Executes a SQL statement (such as `INSERT`, `UPDATE`, or `DELETE`) on a SQLite database stored as a binary object. The updated SQLite database is returned as a binary object, allowing further operations on it.

#### Parameters:

- `sqlite`: The SQLite database to execute the SQL query on, stored as a binary object.
- `query`: The SQL statement to execute on the SQLite database.

##### Example Usage:

```sql
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''solve multitenancy'')'
               )
WHERE name = 'frectonz';
```

-----

### `init_sqlite`

Creates an SQLite database with an initialization query already applied on it. This can be used to initialize a SQLite database with the expected tables already created.

#### Parameters:

- `query`: The SQL statement to execute on the SQLite database.

##### Example Usage:

```sql

CREATE TABLE people (
                        name     TEXT NOT NULL,
                        database SQLITE DEFAULT init_sqlite('CREATE TABLE todos (task TEXT)')
);
```

-----

### `get_sqlite_text`
Extracts a text value from a specific column in a row returned by `query_sqlite`. Use this function to retrieve text values from query results.

#### Parameters:

- `sqlite_row`: A row from the results of `query_sqlite`.
- `index`: The index of the column to extract from the row.

#### Example Usage:

```sql
SELECT get_sqlite_text(sqlite_row, 0)
FROM query_sqlite(database, 'SELECT * FROM todos');
```

----

### `get_sqlite_integer`

Extracts an integer value from a specific column in a row returned by `query_sqlite`. Use this function to retrieve integer values from query results.

#### Parameters:

- `sqlite_row`: A row from the results of `query_sqlite`.
- `index`: The index of the column to extract from the row.

#### Example Usage:

```sql
SELECT get_sqlite_integer(sqlite_row, 1)
FROM query_sqlite(database, 'SELECT * FROM todos');
```

----

### `get_sqlite_real`

Extracts a real (floating-point) value from a specific column in a row returned by `query_sqlite`. Use this function to retrieve real number values from query results.

#### Parameters:

- `sqlite_row`: A row from the results of `query_sqlite`.
- `index`: The index of the column to extract from the row.

#### Example Usage:

```sql
SELECT get_sqlite_real(sqlite_row, 2)
FROM query_sqlite(database, 'SELECT * FROM todos');
```