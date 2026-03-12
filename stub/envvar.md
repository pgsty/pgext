


## Usage

> [envvar: Access environment variables from PostgreSQL](https://github.com/theory/pg-envvar)

Provides a single function to read environment variables from the database server.

### Function

```sql
get_env(name TEXT) RETURNS TEXT
```

Returns the value of the specified environment variable set on the database server, or NULL if not set.

### Examples

```sql
SELECT get_env('PGTZ');
-- UTC

SELECT get_env('HOME');
-- /var/lib/postgresql

SELECT get_env('PATH');
-- /usr/local/sbin:/usr/local/bin:...

SELECT get_env('NONEXISTENT');
-- NULL
```

### Schema Support

You can install the extension into a specific schema:

```sql
CREATE SCHEMA env;
CREATE EXTENSION envvar SCHEMA env;

SELECT env.get_env('PGTZ');
```
