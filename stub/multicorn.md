


## Usage

Sources: [README](https://github.com/pgsql-io/multicorn2/blob/v3.2/README.md), [CHANGELOG](https://github.com/pgsql-io/multicorn2/blob/v3.2/CHANGELOG)

Multicorn2 allows you to write Foreign Data Wrappers in Python. You implement a Python class that inherits from `multicorn.ForeignDataWrapper`, and Multicorn handles bridging it to PostgreSQL's FDW interface. Version 3.2 is tested with PostgreSQL 14-18 and Python 3.9-3.13, though upstream recommends Python 3.10-3.12 for near-term distro compatibility.

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

### Version Notes

Multicorn 3.2 adds basic OFFSET/LIMIT pushdown and LDAP paging support, and fixes LDAP right-parenthesis escaping. Upstream 3.1 added PostgreSQL 18 and Python 3.13 support while dropping PostgreSQL versions before 14.

### Caveats

Multicorn2 and PL/Python are incompatible in the same PostgreSQL database on Python 3.12 due to CPython limitations. They can be installed on the same system, but avoid enabling both in one database.
