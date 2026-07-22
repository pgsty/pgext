## Usage

Sources:

- [Official usage documentation at the reviewed commit](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/doc/multicdr_fdw.md)
- [Official README at the reviewed commit](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/README.md)
- [Version 1.2.2 SQL objects](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/sql/multicdr_fdw.sql)
- [FDW implementation](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/src/multicdr_fdw.c)
- [Official PGXN distribution](https://pgxn.org/dist/multicdr_fdw/)

`multicdr_fdw` is a read-only foreign data wrapper for scanning fixed-position call-detail-record files from one server-local directory as a SQL table. It can select files by regular expression, map non-contiguous fields to columns, expose the source filename, and prune files using dates encoded in filenames. Version 1.2.2 is old code documented for PostgreSQL 9.1 and 9.2.

### Core Workflow

The PostgreSQL operating-system account must be able to read the directory and files:

```sql
CREATE EXTENSION multicdr_fdw;
CREATE SERVER cdr_files FOREIGN DATA WRAPPER multicdr_fdw;

CREATE FOREIGN TABLE call_records (
  calling_number text,
  called_number text,
  duration_seconds integer
) SERVER cdr_files
OPTIONS (
  directory '/var/lib/postgresql/cdr',
  pattern '.*\.cdr$',
  posfields '0,20,40',
  rowminlen '41'
);

SELECT * FROM call_records;
```

`directory`, `pattern`, and `posfields` define the non-recursive file scan and the zero-based starting positions of fixed-width fields. Empty fields become SQL NULL. The reviewed reader supports only `text` and four-byte `integer` data columns; validate width and overflow in source files.

### Table Options

`mapfields` maps table columns to selected CDR fields and preserves table-column order. `filefield` names a text column that receives the file path. `rowminlen` and `rowmaxlen` discard records outside the accepted byte-length range. Without `mapfields`, fields map sequentially.

`dateformat` associates a capture group from `pattern` with a filename date format. `dateminfield` and `datemaxfield` name synthetic timestamp columns used by recognized equality predicates so the FDW can skip files outside a requested date boundary. This is filename pruning, not parsing a timestamp field from each record.

### Security and Compatibility Boundaries

The server process opens local paths supplied in foreign-table options. Restrict who can create or alter the server and tables, and who can select them; otherwise the wrapper can become a server-file disclosure path. Scans are not recursive and provide no remote authentication, write support, transactionally stable file snapshot, or protection against files changing during a query.

The implementation targets obsolete FDW and backend APIs, has no current-major compatibility matrix, and performs its own fixed-width parsing and type conversion. Port, compile, and regression-test it on the exact PostgreSQL build before use, including malformed lines, concurrent file rotation, date pruning, permissions, encodings, and oversized records.
