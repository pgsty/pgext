## Usage

Sources:

- [Official extension control file (pgerror.control)](https://api.pgxn.org/src/pgerror/pgerror-0.2.1/pgerror.control)
- [Official extension SQL (pgerror.in.sql)](https://api.pgxn.org/src/pgerror/pgerror-0.2.1/sql/pgerror.in.sql)

`pgerror` — Tools for better error handling. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgerror;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `error_data(sqlstate text = '' , message text = '' , hint text = '' , detail text = '' , context text = '' , schema_name text = '' , table_name text = '' , column_name text = '' , constraint_name text = '' , type_name text = '')` is an extension function and returns `error_data`.
- `pg_temp.raise_error_internal()` is an extension function and returns `void`.
- `raise(error error_data , level text = 'EXCEPTION')` is an extension function and returns `void`.
- `raise(message text , level text = 'EXCEPTION' , sqlstate text = NULL , hint text = NULL , detail text = NULL , schema_name text = NULL , table_name text = NULL , column_name text = NULL , constraint_name text = NULL , type_name text = NULL)` is an extension function and returns `void`.
- `try(code text , OUT row_count int , OUT error error_data)` is an extension function and returns `record`.
- `try_cursor(query text , cursor_name text DEFAULT NULL , OUT result refcursor , OUT error error_data)` is an extension function and returns `record`.
- `try_into(code text , INOUT result anyelement , strict boolean DEFAULT false , OUT error error_data)` is an extension function and returns `record`.
- `error_data` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
