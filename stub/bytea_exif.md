## Usage

Sources:

- [Official upstream README](https://github.com/mkgrgis/bytea_exif/blob/aec959fbca8422a69f3aa159668c5ed7a2af3845/README.md)
- [Official extension control file (bytea_exif.control)](https://github.com/mkgrgis/bytea_exif/blob/aec959fbca8422a69f3aa159668c5ed7a2af3845/bytea_exif.control)
- [Official extension SQL (bytea_exif--1.0.sql)](https://github.com/mkgrgis/bytea_exif/blob/aec959fbca8422a69f3aa159668c5ed7a2af3845/bytea_exif--1.0.sql)

`bytea_exif` — EXIF data reader from binary bytea data for PostgreSQL ======================================================. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION bytea_exif;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bytea_exif_libexif_version()` is an extension function and returns `int`.
- `bytea_exif_version()` is an extension function and returns `int`.
- `bytea_get_exif_dest_point(data bytea)` is an extension function and returns `text`.
- `bytea_get_exif_gps_local_timestamp(data bytea)` is an extension function and returns `timestamp`.
- `bytea_get_exif_gps_utc_timestamp(data bytea)` is an extension function and returns `timestamptz`.
- `bytea_get_exif_json(data bytea)` is an extension function and returns `json`.
- `bytea_get_exif_point(data bytea)` is an extension function and returns `text`.
- `bytea_get_exif_tag_value(data bytea, tag text)` is an extension function and returns `text`.
- `bytea_get_exif_user_comment(data bytea)` is an extension function and returns `text`.
- `bytea_get_mime_type(data bytea)` is an extension function and returns `text`.
- `bytea_has_exif(data bytea)` is an extension function.
- `bytea_has_exif_ifd(data bytea, ifd text)` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
