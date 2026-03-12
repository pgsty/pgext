

## Usage

> [pg_bulkload: pg_bulkload is a high speed data loading utility for PostgreSQL](https://github.com/ossc-db/pg_bulkload)

A high-speed data loading tool for PostgreSQL that bypasses shared buffers for massive data loads, with built-in ETL features for input validation and data transformation.

### Basic Usage

Load data using a control file:

```bash
pg_bulkload sample_csv.ctl
```

Output:

```
NOTICE: BULK LOAD START
NOTICE: BULK LOAD END
    0 Rows skipped.
    8 Rows successfully loaded.
    0 Rows not loaded due to parse errors.
    0 Rows not loaded due to duplicate errors.
    0 Rows replaced with new rows.
```

### Control File Example

```ini
# sample_csv.ctl
OUTPUT = my_table
INPUT = /path/to/data.csv
TYPE = CSV
DELIMITER = ,
QUOTE = "\""
ESCAPE = "\""
NULL = ""
SKIP = 1              # skip header row
PARSE_ERRORS = 100    # allow up to 100 parse errors
DUPLICATE_ERRORS = 0  # reject on duplicate key errors
ON_DUPLICATE_KEEP = NEW  # or OLD
TRUNCATE = NO
```

### Loading Modes

- **DIRECT**: Bypasses shared buffers, writes directly to data files (fastest)
- **PARALLEL**: Uses multiple processes for loading
- **CSV/BINARY/FIXED**: Supports various input formats

### SQL Interface

```sql
-- Load data from within SQL
SELECT pg_bulkload(
    'OUTPUT = my_table, INPUT = /path/to/data.csv, TYPE = CSV'
);
```

### Key Features

- Bypasses PostgreSQL shared buffers for maximum throughput
- Input data validation with configurable error thresholds
- Duplicate key handling (keep new, keep old, or reject)
- CSV, fixed-length, and binary input formats
- Skip rows, filter functions for data transformation
- Parallel loading support

### Documentation

Full documentation: http://ossc-db.github.io/pg_bulkload/index.html
