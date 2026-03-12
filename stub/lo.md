


## Usage

> [lo: Large Object maintenance](https://www.postgresql.org/docs/current/lo.html)

The `lo` extension provides a data type and trigger function for managing PostgreSQL Large Objects, preventing orphaned objects when references are updated or deleted.

### Data Type

The `lo` type is a domain over `oid`, used to identify columns that hold Large Object references. This is especially useful for ODBC driver compatibility.

```sql
CREATE TABLE image (
    title  text,
    raster lo      -- large object reference column
);
```

### Trigger Function

The `lo_manage()` trigger automatically calls `lo_unlink()` to delete the associated Large Object when a row is updated or deleted:

```sql
CREATE TRIGGER t_raster
    BEFORE UPDATE OR DELETE ON image
    FOR EACH ROW EXECUTE FUNCTION lo_manage(raster);
```

For multiple `lo` columns, create a separate trigger for each:

```sql
CREATE TABLE gallery (
    title     text,
    thumbnail lo,
    fullsize  lo
);

CREATE TRIGGER t_thumbnail
    BEFORE UPDATE OR DELETE ON gallery
    FOR EACH ROW EXECUTE FUNCTION lo_manage(thumbnail);

CREATE TRIGGER t_fullsize
    BEFORE UPDATE OR DELETE ON gallery
    FOR EACH ROW EXECUTE FUNCTION lo_manage(fullsize);
```

To restrict the trigger to column updates only:

```sql
CREATE TRIGGER t_raster
    BEFORE UPDATE OF raster OR DELETE ON image
    FOR EACH ROW EXECUTE FUNCTION lo_manage(raster);
```

### Limitations

- `DROP TABLE` and `TRUNCATE` do not fire row-level triggers, so Large Objects will be orphaned. Run `DELETE FROM table` before dropping.
- The trigger assumes each Large Object is referenced by only one column/row.
- Use the `vacuumlo` utility to clean up any orphaned Large Objects.

The extension is trusted and can be installed by non-superusers with `CREATE` privilege.
