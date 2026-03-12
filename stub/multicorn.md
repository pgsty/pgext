


## Usage

> [multicorn: Fetch foreign data in Python in your PostgreSQL server](https://github.com/pgsql-io/multicorn2)

Multicorn2 allows you to write Foreign Data Wrappers in Python. You implement a Python class that inherits from `multicorn.ForeignDataWrapper`, and Multicorn handles bridging it to PostgreSQL's FDW interface.

### Define a Python FDW Class

Create a Python module (e.g., `myfdw.py`) accessible to the PostgreSQL process:

```python
from multicorn import ForeignDataWrapper

class MyFDW(ForeignDataWrapper):
    def __init__(self, options, columns):
        super().__init__(options, columns)
        self.options = options
        self.columns = columns

    def execute(self, quals, columns):
        """Yield rows as dictionaries. quals contains WHERE pushdown info."""
        yield {"id": 1, "name": "example"}

    def insert(self, new_values):
        """Handle INSERT operations."""
        pass

    def update(self, old_values, new_values):
        """Handle UPDATE operations."""
        pass

    def delete(self, old_values):
        """Handle DELETE operations."""
        pass
```

### Create Server and Foreign Table

```sql
CREATE EXTENSION multicorn;

CREATE SERVER multicorn_srv FOREIGN DATA WRAPPER multicorn
  OPTIONS (wrapper 'myfdw.MyFDW');

CREATE FOREIGN TABLE my_table (
  id integer,
  name text
)
SERVER multicorn_srv
OPTIONS (
  option1 'value1'
);

SELECT * FROM my_table;
```

The `wrapper` option specifies the fully qualified Python class name. Any additional options are passed to the class constructor's `options` parameter.

### Built-in FDW Examples

Multicorn ships with several example FDW implementations that can be used directly or as reference:

- **CsvFdw** -- read CSV files
- **ProcessFdw** -- execute system commands and parse output
- **GCalFdw** -- access Google Calendar
- **ImapFdw** -- query IMAP mailboxes
- **RssFdw** -- read RSS/Atom feeds

```sql
CREATE SERVER csv_srv FOREIGN DATA WRAPPER multicorn
  OPTIONS (wrapper 'multicorn.csvfdw.CsvFdw');

CREATE FOREIGN TABLE csvtest (
  col1 text,
  col2 text
)
SERVER csv_srv
OPTIONS (
  filename '/tmp/data.csv',
  skip_header '1',
  delimiter ','
);
```
