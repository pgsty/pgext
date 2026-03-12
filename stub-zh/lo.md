

## 用法

> [lo: 大对象维护](https://www.postgresql.org/docs/current/lo.html)

`lo` 扩展提供数据类型和触发器函数用于管理 PostgreSQL 大对象，防止在更新或删除引用时产生孤立对象。

### 数据类型

`lo` 类型是 `oid` 的域，用于标识持有大对象引用的列。这对于 ODBC 驱动程序兼容性特别有用。

```sql
CREATE TABLE image (
    title  text,
    raster lo      -- 大对象引用列
);
```

### 触发器函数

`lo_manage()` 触发器在更新或删除行时自动调用 `lo_unlink()` 删除关联的大对象：

```sql
CREATE TRIGGER t_raster
    BEFORE UPDATE OR DELETE ON image
    FOR EACH ROW EXECUTE FUNCTION lo_manage(raster);
```

对于多个 `lo` 列，为每个列创建单独的触发器：

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

仅在列更新时限制触发器：

```sql
CREATE TRIGGER t_raster
    BEFORE UPDATE OF raster OR DELETE ON image
    FOR EACH ROW EXECUTE FUNCTION lo_manage(raster);
```

### 限制

- `DROP TABLE` 和 `TRUNCATE` 不会触发行级触发器，因此大对象会变成孤立的。请在删除表之前先执行 `DELETE FROM table`。
- 触发器假设每个大对象仅被一个列/行引用。
- 使用 `vacuumlo` 工具清理任何孤立的大对象。

该扩展是受信任的，具有 `CREATE` 权限的非超级用户可以安装。
