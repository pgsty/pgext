## Usage

Source: [README](https://github.com/darold/external_file/blob/master/README.md), [Release v1.2](https://github.com/darold/external_file/releases/tag/v1.2)

`external_file` stores file locators as `(directory, filename)` pairs and accesses server-side files through PostgreSQL `lo_*` helpers rather than direct file reads.

### Basic workflow

```sql
CREATE EXTENSION external_file;

INSERT INTO directories(directory_name, directory_path)
VALUES ('temporary', '/tmp/');

INSERT INTO directory_roles(directory_name, directory_role, directory_read, directory_write)
VALUES ('temporary', 'app_reader', true, false);

SELECT writeEfile('\x48656c6c6f0a', ('temporary', 'hello.txt'));
SELECT readEfile(efilename('temporary', 'hello.txt'));
SELECT copyEfile(('temporary', 'hello.txt'), ('temporary', 'hello-copy.txt'));
```

### Core objects

- `directories`: maps an alias to an on-server directory path.
- `directory_roles`: grants read/write rights on that alias to roles.
- `efilename(directory, filename)`: constructs an `efile` locator.
- `readEfile(efile)`: reads the target file into `bytea`.
- `writeEfile(bytea, efile)`: writes `bytea` to the target file.
- `copyEfile(src, dest)`: copies one external file to another.
- `getEfilePath(efile, need_read, need_write)`: resolves the full path and checks access.

### Caveats

- Creating the extension requires a PostgreSQL superuser.
- Upstream creates all objects in the `external_file` schema by default.
- The PostgreSQL OS user still needs filesystem read/write permission on the target directory.
- Filenames must not contain `/` or `\`; access is intentionally mediated through the directory tables.
