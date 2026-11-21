## Usage

Sources:

- [README.md](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/README.md)
- [bytea_unzip.control](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/bytea_unzip.control)
- [bytea_unzip--1.0.sql](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/bytea_unzip--1.0.sql)
- [Makefile](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/Makefile)
- [License](https://github.com/mkgrgis/bytea_unzip/blob/11e9854146e461f3aa1712ce8e3211bdc25a5783/License)

`bytea_unzip` reads ZIP members from a bytea value and decomposes an OpenDocument spreadsheet table into rows of text arrays. SQL version 1.0 uses libzip and libxml2. Upstream documents PostgreSQL 16 or later and Linux/POSIX environments; the functions operate on supplied binary values.

### Core Workflow

Install as a superuser. The control file permits relocation and does not require preload.

```sql
CREATE EXTENSION bytea_unzip;
CREATE TABLE archives (id bigint PRIMARY KEY, payload bytea);
SELECT id, bytea_is_zip(payload) FROM archives;
SELECT a.id, f.*
FROM archives AS a
CROSS JOIN LATERAL bytea_unzip_list_files(a.payload) AS f;
SELECT id, convert_from(bytea_unzip(payload, 'mimetype'), 'UTF8')
FROM archives;
```

Populate the table with actual archive bytes through your application. `bytea_is_zip()` checks whether libzip recognizes an archive, not whether it contains useful files. `bytea_unzip()` returns the named member as bytea, returns NULL when that member is absent, and reports an error for invalid ZIP input. The listing returns file names, uncompressed and compressed sizes, and modification timestamps.

### OpenDocument Tables

```sql
WITH sheets AS (
  SELECT unnest(xpath(
    '//office:body/office:spreadsheet/table:table',
    convert_from(bytea_unzip(payload, 'content.xml'), 'UTF8')::xml,
    ARRAY[
      ARRAY['office', 'urn:oasis:names:tc:opendocument:xmlns:office:1.0'],
      ARRAY['table', 'urn:oasis:names:tc:opendocument:xmlns:table:1.0']
    ]
  )) AS sheet
  FROM archives
)
SELECT decompose_ods_table_node(sheet) FROM sheets;
```

The input to `decompose_ods_table_node()` is one XML table node, and the result is a set of text arrays. `bytea_unzip_version()` and `bytea_unzip_libzip_version()` expose extension and library version numbers. NULL arguments produce NULL through strict functions.

### Boundaries

The SQL surface contains functions, not a foreign-data wrapper. The README's HTTP download example uses the separate `http` extension; that is optional and is not a dependency of these functions. Its unrelated spatial-result wording does not describe the verified installation SQL. Keep libzip and libxml2 available to the server and manage archive inputs within the application's resource limits.
