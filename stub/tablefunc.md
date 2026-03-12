

## Usage

> [tablefunc: Functions that manipulate whole tables, including crosstab](https://www.postgresql.org/docs/current/tablefunc.html)

The `tablefunc` module provides functions that return tables (multiple rows).
It is a trusted extension available to non-superusers with `CREATE` privilege.

```sql
CREATE EXTENSION tablefunc;
```

### normal_rand -- Generate Random Values

Produces a set of normally distributed random values (Gaussian distribution):

```sql
SELECT * FROM normal_rand(1000, 5, 3);
-- numvals: number of values, mean: 5, stddev: 3
```

### crosstab(text) -- Single-Parameter Pivot

Pivots data from long format to wide format. The SQL must return `row_name`, `category`, and `value` columns:

```sql
CREATE TABLE ct(id SERIAL, rowid TEXT, attribute TEXT, value TEXT);
INSERT INTO ct(rowid, attribute, value) VALUES
  ('test1','att1','val1'), ('test1','att2','val2'),
  ('test1','att3','val3'), ('test2','att1','val5'),
  ('test2','att2','val6'), ('test2','att3','val7');

SELECT *
FROM crosstab(
  'SELECT rowid, attribute, value FROM ct ORDER BY 1,2'
) AS ct(row_name text, category_1 text, category_2 text, category_3 text);

 row_name | category_1 | category_2 | category_3
----------+------------+------------+------------
 test1    | val1       | val2       | val3
 test2    | val5       | val6       | val7
```

The input query should always use `ORDER BY 1,2` to ensure proper grouping.
Extra output columns beyond the available values are filled with nulls.

### crosstab(text, text) -- Two-Parameter Pivot with Categories

Handles cases where some groups may not have data for all categories:

```sql
CREATE TABLE sales(year int, month int, qty int);
INSERT INTO sales VALUES(2007,1,1000),(2007,2,1500),(2007,7,500),
  (2007,11,1500),(2007,12,2000),(2008,1,1000);

SELECT * FROM crosstab(
  'SELECT year, month, qty FROM sales ORDER BY 1',
  'SELECT m FROM generate_series(1,12) m'
) AS (
  year int, "Jan" int, "Feb" int, "Mar" int, "Apr" int,
  "May" int, "Jun" int, "Jul" int, "Aug" int,
  "Sep" int, "Oct" int, "Nov" int, "Dec" int
);

 year | Jan  | Feb  | Mar | Apr | May | Jun | Jul | Aug | Sep | Oct | Nov  | Dec
------+------+------+-----+-----+-----+-----+-----+-----+-----+-----+------+------
 2007 | 1000 | 1500 |     |     |     |     | 500 |     |     |     | 1500 | 2000
 2008 | 1000 |      |     |     |     |     |     |     |     |     |      |
```

The source SQL may include "extra" columns between `row_name` and `category`/`value`.

### crosstab2, crosstab3, crosstab4 -- Predefined Wrappers

Pre-built wrappers that avoid writing the FROM clause (text input/output only):

```sql
SELECT * FROM crosstab3(
  'SELECT rowid, attribute, value FROM ct ORDER BY 1,2'
);
```

Create custom wrappers for other types:

```sql
CREATE TYPE my_crosstab_float8_5_cols AS (
    my_row_name text, c1 float8, c2 float8, c3 float8, c4 float8, c5 float8
);

CREATE OR REPLACE FUNCTION crosstab_float8_5_cols(text)
    RETURNS setof my_crosstab_float8_5_cols
    AS '$libdir/tablefunc','crosstab' LANGUAGE C STABLE STRICT;
```

### connectby -- Hierarchical Tree Display

Displays hierarchical data stored in a table with key and parent-key fields:

```sql
CREATE TABLE connectby_tree(keyid text, parent_keyid text, pos int);
INSERT INTO connectby_tree VALUES
  ('row1',NULL,0), ('row2','row1',0), ('row3','row1',0),
  ('row4','row2',1), ('row5','row2',0), ('row6','row4',0),
  ('row7','row3',0), ('row8','row6',0), ('row9','row5',0);

-- With branch display and ordering
SELECT * FROM connectby('connectby_tree', 'keyid', 'parent_keyid', 'pos', 'row2', 0, '~')
  AS t(keyid text, parent_keyid text, level int, branch text, pos int);

 keyid | parent_keyid | level |       branch        | pos
-------+--------------+-------+---------------------+-----
 row2  |              |     0 | row2                |   1
 row5  | row2         |     1 | row2~row5           |   2
 row9  | row5         |     2 | row2~row5~row9      |   3
 row4  | row2         |     1 | row2~row4           |   4
 row6  | row4         |     2 | row2~row4~row6      |   5
 row8  | row6         |     3 | row2~row4~row6~row8 |   6
```

Parameters: `connectby(relname, keyid_fld, parent_keyid_fld [, orderby_fld], start_with, max_depth [, branch_delim])`

- `start_with`: key value of the row to start at (text)
- `max_depth`: maximum depth to descend (0 = unlimited)
- `branch_delim`: separator for branch output (optional, default `~`)
- Index the parent-key field for performance on large tables
