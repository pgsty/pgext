## Usage

Sources:

- [Official upstream README](https://github.com/joshuajerin/pg_linearalgebra/blob/63a4ca281720241ace072df5f07b972e6a6598ac/README.md)
- [Official extension control file (pg_linearAlgebra.control)](https://github.com/joshuajerin/pg_linearalgebra/blob/63a4ca281720241ace072df5f07b972e6a6598ac/pg_linearAlgebra.control)
- [Official extension SQL (pg_linearAlgebra--1.0.sql)](https://github.com/joshuajerin/pg_linearalgebra/blob/63a4ca281720241ace072df5f07b972e6a6598ac/sql/pg_linearAlgebra--1.0.sql)

`pg_linearAlgebra` — A PostgreSQL extension for basic linear algebra operations. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "pg_linearAlgebra";

SELECT mAdd('[[1.0, 2.0], [3.0, 4.0]]', '[[5.0, 6.0], [7.0, 8.0]]', 2, 2);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mAdd(matrix1 text, matrix2 text, rows integer, cols integer)` is an extension function and returns `text`.
- `mMultiply(matrix1 text, matrix2 text, rows integer, cols integer)` is an extension function and returns `text`.
- `mSubtract(matrix1 text, matrix2 text, rows integer, cols integer)` is an extension function and returns `text`.
- `mSvd(matrix text, rows integer, cols integer)` is an extension function and returns `text`.
- `mTranspose(matrix text, rows integer, cols integer)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
