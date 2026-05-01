## Usage

Sources: [PGXN metadata](https://api.pgxn.org/dist/pg_bikram_sambat.json), [PGXN source tree](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/), [type SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/types.sql), [function SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/functions.sql), [operator SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/operators.sql), [cast SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/casts.sql), [regression examples](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/tests/pg_regress/sql/003_functions.sql), [TODO](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/todos.txt)

`pg_bikram_sambat` adds a `bs_date` type for Bikram Sambat dates plus conversion, formatting, comparison, and btree indexing support. Install it as a normal PostgreSQL extension:

```sql
CREATE EXTENSION pg_bikram_sambat;
```

### Date Type

`bs_date` stores a Bikram Sambat date and displays it as `YYYY-MM-DD`. Text input accepts year/month/day values separated with `/`, `-`, or `.`; the input parser also accepts day-first strings when the year appears in the last field.

```sql
SELECT '2057/10/19'::bs_date;
SELECT CAST('2057-10-19' AS bs_date);
SELECT '19.10.2057'::bs_date;
```

Use it in tables like any other PostgreSQL type:

```sql
CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  ad_date date,
  bs bs_date NOT NULL
);

INSERT INTO events (ad_date, bs)
VALUES
  ('2001-02-01', '2057/10/19'),
  ('1972-02-17', '2028/11/05');
```

### Conversion Functions

`ad_to_bs(date)` converts a Gregorian `date` to `bs_date`:

```sql
SELECT ad_to_bs('2001-02-01'::date);  -- 2057-10-19
SELECT ad_to_bs('1972-02-17'::date);  -- 2028-11-05
```

`current_bs_date()` returns the current transaction timestamp converted to `bs_date`, so repeated calls inside the same transaction are stable:

```sql
SELECT current_bs_date();
SELECT pg_typeof(current_bs_date());  -- bs_date
```

Version `0.1.0` does not expose a SQL `bs_to_ad()` function or direct `bs_date` to `date` cast; the upstream TODO file lists those as future work.

### Formatting

The extension overloads PostgreSQL `to_char` for `bs_date`:

```sql
SELECT to_char('2057/10/19'::bs_date, 'YYYY-MM-DD');
SELECT to_char('2057/10/19'::bs_date, 'DD/MM/YYYY');
SELECT to_char('2057/10/19'::bs_date, 'Month DD, YYYY');
SELECT to_char('2057/10/19'::bs_date, 'Day, DD Month YYYY');
```

Supported date-format tokens are `YYYY`, `YY`, `Month`, `Mon`, `MM`, `Day`, `Dy`, and `DD`. Month and weekday names follow the casing of the format token, so `MONTH`, `Month`, and `month` produce upper-case, title-case, and lower-case English names.

Pass `dev` as the third argument for Devanagari digits, month names, and weekday names:

```sql
SELECT to_char('2057/10/19'::bs_date, 'YYYY-MM-DD', 'dev');
SELECT to_char('2057/10/19'::bs_date, 'Day, DD Month YYYY', 'dev');
```

### Operators And Indexes

`bs_date` supports the comparison operators `=`, `<>`, `>`, `>=`, `<`, and `<=`. The default btree operator class `bs_date_ops` enables ordinary btree indexes, range predicates, and ordering:

```sql
CREATE INDEX events_bs_idx ON events (bs);

SELECT * FROM events WHERE bs >= '2057/01/01' ORDER BY bs;
SELECT * FROM events WHERE bs BETWEEN '2056/01/01' AND '2058/12/12';
```

### Caveats

The packaged conversion dataset covers BS years `2000` through `2100`, with `1943-04-14` AD as the reference date for `2000-01-01` BS. Dates before the reference date or beyond the mapped BS range raise PostgreSQL errors. The extension defines an implicit cast from `text` to `bs_date`, but it does not define casts from arbitrary numeric types.
