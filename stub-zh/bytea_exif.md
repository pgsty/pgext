## 用法

来源：

- [官方上游 README](https://github.com/mkgrgis/bytea_exif/blob/aec959fbca8422a69f3aa159668c5ed7a2af3845/README.md)
- [官方扩展控制文件 (bytea_exif.control)](https://github.com/mkgrgis/bytea_exif/blob/aec959fbca8422a69f3aa159668c5ed7a2af3845/bytea_exif.control)
- [官方扩展 SQL (bytea_exif--1.0.sql)](https://github.com/mkgrgis/bytea_exif/blob/aec959fbca8422a69f3aa159668c5ed7a2af3845/bytea_exif--1.0.sql)

`bytea_exif` — 用于从二进制 bytea 数据读取 EXIF 数据的 PostgreSQL 扩展 ======================================================. 请使用它来配合相应的 SQL 或数据库实用程序工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION bytea_exif;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，验证安装版本和返回值，然后再将其集成到应用程序 SQL 中。

### 重要对象

- `bytea_exif_libexif_version()` 是一个扩展函数，返回 `int`。
- `bytea_exif_version()` 是一个扩展函数，返回 `int`。
- `bytea_get_exif_dest_point(data bytea)` 是一个扩展函数，返回 `text`。
- `bytea_get_exif_gps_local_timestamp(data bytea)` 是一个扩展函数，返回 `timestamp`。
- `bytea_get_exif_gps_utc_timestamp(data bytea)` 是一个扩展函数，返回 `timestamptz`。
- `bytea_get_exif_json(data bytea)` 是一个扩展函数，返回 `json`。
- `bytea_get_exif_point(data bytea)` 是一个扩展函数，返回 `text`。
- `bytea_get_exif_tag_value(data bytea, tag text)` 是一个扩展函数，返回 `text`。
- `bytea_get_exif_user_comment(data bytea)` 是一个扩展函数，返回 `text`。
- `bytea_get_mime_type(data bytea)` 是一个扩展函数，返回 `text`。
- `bytea_has_exif(data bytea)` 是一个扩展函数。
- `bytea_has_exif_ifd(data bytea, ifd text)` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，与固定源进行比对。
