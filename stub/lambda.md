## Usage

Sources:

- [Official extension control file (lambda.control)](https://api.pgxn.org/src/pg_lambda/pg_lambda-1.0.5/lambda.control)
- [Official extension SQL (lambda.sql)](https://api.pgxn.org/src/pg_lambda/pg_lambda-1.0.5/sql/lambda.sql)

`lambda` — Lambda function for Postgres. Use it when database code must run in or interoperate with this procedural language. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION lambda;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `lambda(code text , VARIADIC inputs anyarray)` is an extension function and returns `anyelement`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.5`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
