


## Usage

> [plprql: Use PRQL in PostgreSQL - Pipelined Relational Query Language](https://github.com/kaspermarstal/plprql)

`plprql` enables writing PostgreSQL functions using [PRQL](https://prql-lang.org/) (Pipelined Relational Query Language), a modern language that compiles to SQL with a pipeline syntax.

```sql
CREATE EXTENSION plprql;
```

### Create Functions with PRQL

```sql
CREATE FUNCTION match_stats(int)
RETURNS TABLE(player text, kd_ratio float) AS $$
  from matches
  filter match_id == $1
  group player (
    aggregate {
      total_kills = sum kills,
      total_deaths = sum deaths
    }
  )
  filter total_deaths > 0
  derive kd_ratio = total_kills / total_deaths
  select { player, kd_ratio }
$$ LANGUAGE plprql;

SELECT * FROM match_stats(42);
```

### Execute PRQL Directly

```sql
SELECT prql('from matches | filter player == "Player1"')
  AS (id int, match_id int, player text, kills int)
  LIMIT 2;
```

### Use Cursors for Large Results

```sql
SELECT prql('from matches | filter player == "Player1"', 'cursor_name');
FETCH 2 FROM cursor_name;
```

### Inspect Compiled SQL

```sql
SELECT prql_to_sql('from matches | filter player == "Player1"');
```

### PRQL Syntax Overview

PRQL uses pipeline transformations:

```
from employees                    # data source
filter department == "Engineering" # row filtering
derive monthly_salary = salary / 12 # computed columns
sort {-monthly_salary}            # sorting (- for descending)
select {name, monthly_salary}     # column projection
take 10                           # limit rows
```

### Limitations

PRQL only supports `SELECT` statements. For `INSERT`, `UPDATE`, and `DELETE`, use SQL or PL/pgSQL.
