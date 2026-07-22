## Usage

Sources:

- [Official README](https://github.com/polydbms/pg_sheet_fdw/blob/3850b6b84196eb9c1e93b7957f0e786a6f5b084b/README.md)
- [Extension SQL](https://github.com/polydbms/pg_sheet_fdw/blob/3850b6b84196eb9c1e93b7957f0e786a6f5b084b/pg_sheet_fdw--0.1.2.sql)
- [FDW implementation](https://github.com/polydbms/pg_sheet_fdw/blob/3850b6b84196eb9c1e93b7957f0e786a6f5b084b/src/pg_sheet_fdw.c)

`pg_sheet_fdw` 0.1.2 is a read-only foreign data wrapper for local Microsoft Excel `.xlsx` worksheets. It maps one sheet to a foreign table and converts cells into the PostgreSQL column types declared by the user. The spreadsheet remains a server-side file; this is not a Google Sheets or remote-object connector.

### Core Workflow

```sql
CREATE EXTENSION pg_sheet_fdw;
CREATE SERVER sheets FOREIGN DATA WRAPPER pg_sheet_fdw;

CREATE FOREIGN TABLE monthly_sales (
    sold_on date,
    sku text,
    quantity integer,
    amount numeric
)
SERVER sheets
OPTIONS (
    filepath '/srv/imports/sales.xlsx',
    sheetname 'January',
    skiprows '1',
    numberofthreads '4',
    batchsize '1000'
);

SELECT * FROM monthly_sales WHERE quantity > 0;
```

`filepath` is mandatory and should be absolute. `sheetname` defaults to the first sheet, `skiprows` defaults to 0, and `numberofthreads` lets the parser choose automatically when omitted (explicit values are limited to 1–10). If `batchsize` is absent, the implementation derives a value intended to split the sheet into about 101 batches, with a minimum of 1,000 rows.

### Data and Execution Boundaries

The PostgreSQL server operating-system user must be able to read the workbook and every directory in its path. Define columns in worksheet order with types matching the actual cells. The upstream examples show that some mismatches can be converted or saturated rather than rejected, so validate imports and cast into a staging table before relying on values.

The FDW handler implements scan callbacks only; it has no insert, update, or delete callbacks. Planner/executor support is basic, and the sheet parser may read substantially more data than a SQL filter returns. Large workbooks consume backend CPU and memory and can change independently of a transaction snapshot. The README reports tests on PostgreSQL 13 and CI on PostgreSQL 16; validate both build compatibility and workbook behavior on any other major version.
