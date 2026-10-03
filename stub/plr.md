## Usage

Sources:

- [8.4.8.7 README](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/README.md)
- [Versioned user guide](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/userguide.md)
- [Control file](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/plr.control)
- [Version 8.4.8.7 SQL](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/plr--8.4.8.7.sql)

`plr` enables writing PostgreSQL functions in the R programming language, providing full access to R's statistical and data analysis capabilities.

```sql
CREATE EXTENSION plr;
```

### Create Functions

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (arg1 > arg2)
  return(arg1)
else
  return(arg2)
' LANGUAGE plr STRICT;

SELECT r_max(10, 20);  -- 20
```

With named arguments:

```sql
CREATE OR REPLACE FUNCTION sd(vals float8[]) RETURNS float AS '
sd(vals)
' LANGUAGE plr STRICT;

SELECT sd(ARRAY[1.0, 2.0, 3.0, 4.0, 5.0]);
```

### Argument Handling

- Unnamed arguments are available as `arg1`, `arg2`, ...; an explicitly named parameter replaces its corresponding `argN` variable.
- Scalar SQL NULL becomes R NULL; null array elements become R `NA`. A `STRICT` function skips calls with a whole NULL argument, not arrays containing null elements.
- Composite types (rows) are passed as R data.frames
- One-dimensional arrays become R vectors; two-dimensional arrays become matrices and three-dimensional arrays become R arrays. Higher dimensions are unsupported.

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (is.null(arg1) && is.null(arg2))
  return(NULL)
if (is.null(arg1))
  return(arg2)
if (is.null(arg2))
  return(arg1)
if (arg1 > arg2)
  return(arg1)
return(arg2)
' LANGUAGE plr;
```

### Database Access via SPI

```sql
CREATE OR REPLACE FUNCTION test_spi(text) RETURNS SETOF record AS '
pg.spi.exec(arg1)
' LANGUAGE plr;

SELECT * FROM test_spi('SELECT oid, typname FROM pg_type LIMIT 5')
  AS t(oid oid, typname name);
```

Initialize the type OID variables in the same connection before calling a function that prepares a parameterized query:

```sql
SELECT load_r_typenames();

CREATE OR REPLACE FUNCTION lookup_type(type_name text)
RETURNS SETOF record AS $$
sp <- pg.spi.prepare(
  'SELECT oid, typname FROM pg_type WHERE typname = $1',
  c(NAMEOID)
)
pg.spi.execp(sp, list(type_name))
$$ LANGUAGE plr;

SELECT * FROM lookup_type('text') AS t(oid oid, typname name);
```

### Set-Returning Functions

Return an R vector for a set of scalar values:

```sql
CREATE OR REPLACE FUNCTION get_numbers(n int) RETURNS SETOF integer AS '
1:n
' LANGUAGE plr;

SELECT * FROM get_numbers(5);
```

### Window Functions

```sql
CREATE OR REPLACE FUNCTION r_regr_slope(float8, float8, int)
RETURNS float8 AS '
slope <- NA
y <- farg1
x <- farg2
if (fnumrows == arg3 + 1L)
  try(slope <- lm(y ~ x)$coefficients[2])
return(slope)
' LANGUAGE plr WINDOW;
```

Window functions receive `farg1..fargN` (vectors of values in the window frame), `fnumrows` (frame size), and `prownum` (current row position in partition).

### Global Variables

Persist data across function calls using R's global environment:

```sql
CREATE OR REPLACE FUNCTION set_state(key text, val text) RETURNS void AS '
assign(key, val, env=.GlobalEnv)
' LANGUAGE plr;
```

### Useful Support Functions

```sql
SELECT load_r_typenames();  -- Load type OID variables
SELECT * FROM r_typenames(); -- List available type OIDs
SELECT plr_version();        -- PL/R version
```

### Trigger Functions

PL/R supports trigger functions with access to `pg.tg.name`, `pg.tg.relname`, `pg.tg.when`, `pg.tg.level`, `pg.tg.op`, `pg.tg.new`, and `pg.tg.old`.

### Runtime and Privileges

This page follows PL/R 8.4.8.7. PL/R is an untrusted language: creating its functions requires a superuser, and R code runs with the PostgreSQL operating-system user's access to files and processes. Review function bodies and EXECUTE grants accordingly. The R shared library must be available, and upstream requires `R_HOME` in the PostgreSQL server process environment before startup on Unix systems. Adding it only to an interactive client shell does not configure the server.

SQL scalar NULL converts to R NULL, while null elements within an array convert to R NA; declaring a function STRICT prevents calls with null arguments. R global state belongs to a backend process, not to all sessions or to a durable database table. Upstream revokes PUBLIC execution of environment-changing helpers including `plr_set_rhome(text)`; use an administrator-managed runtime rather than exposing these helpers to application roles. Normal usage does not require shared preload.
