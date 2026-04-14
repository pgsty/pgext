## Usage

- Source: [GitHub repo](https://github.com/darold/external_file), [README](https://github.com/darold/external_file/blob/master/README.md)
- `external_file` provides server-side access to external files through a PostgreSQL extension, similar to Oracle BFILE-style locators.

```sql
CREATE EXTENSION external_file;
```

The README states that objects are created in the `external_file` schema by default and that creating the extension requires a PostgreSQL superuser.

## Core Workflow

The extension uses a directory alias plus a file name to identify an external file. The upstream README shows this sequence:

```sql
INSERT INTO directories(directory_name, directory_path)
VALUES ('temporary', '/tmp/');

INSERT INTO directory_roles(directory_name, directory_role, directory_read, directory_write)
VALUES ('temporary', 'a_role', true, false);

SELECT writeEfile('\x48656c6c6f2c0a0a596f75206172652072656164696e67206120746578742066696c652e0a0a526567617264732c0a',
                  ('temporary', 'blahblah.txt'));
SELECT readefile(the_file) FROM efile_test;
SELECT copyefile(('temporary', 'blahblah.txt'), ('temporary', 'copy_blahblah.txt'));
```

The main exported helpers are `efilename`, `readEfile`, `writeEfile`, `copyEfile`, and `getEfilePath`.

## Notes

The extension does not read files directly from the server filesystem. It uses the server-side `lo_*` family and enforces access through directory and role tables.
