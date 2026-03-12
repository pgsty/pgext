


## Usage

> [plr: load R interpreter and execute R script from within a database](https://github.com/postgres-plr/plr)

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

- Arguments are available as `arg1`, `arg2`, ... or by named parameter
- NULL arguments become R `NA` values (unless function is `STRICT`)
- Composite types (rows) are passed as R data.frames
- Arrays are passed as R vectors

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

Prepared statements:

```sql
-- Prepare
sp <<- pg.spi.prepare('SELECT * FROM pg_type WHERE typname = $1', c(NAMEOID))
-- Execute
pg.spi.execp(sp, list('text'))
```

### Set-Returning Functions

Return a data.frame for set-returning functions:

```sql
CREATE OR REPLACE FUNCTION get_numbers(n int) RETURNS SETOF integer AS '
1:arg1
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
assign(arg1, arg2, env=.GlobalEnv)
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
