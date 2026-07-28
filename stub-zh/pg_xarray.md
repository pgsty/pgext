## 用法

来源：

- [官方上游 README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/README.md)
- [官方扩展控制文件 (pg_xarray.control)](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_xarray/pg_xarray.control)
- [官方实现源代码](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_xarray/src/lib.rs)

`pg_xarray` — 分类和查询层，用于分块科学数组（NetCDF、Zarr、HDF5、GRIB、COG、SELAFIN、MED、CGNS、FITS）。使用它来进行相应的分析或存储工作流。上游明确表示该项目尚未准备好投入生产。

### 核心工作流

```sql
CREATE EXTENSION pg_xarray;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `chunk_count` 是一个扩展函数。
- `list_datasets()` 是一个扩展函数。
- `register_chunk` 是一个扩展函数。
- `register_dataset` 是一个扩展函数。
- `register_file` 是一个扩展函数。
- `register_mesh` 是一个扩展函数。
- `register_mesh_cell` 是一个扩展函数。
- `register_mesh_node` 是一个扩展函数。
- `register_mesh_version` 是一个扩展函数。
- `register_variable` 是一个扩展函数。
- `xarray_to_glb` 是一个扩展函数。
- `xarray_to_png` 是一个扩展函数。

### 要求与注意事项

- 分类记录版本 `0.2.0`。
- 先安装确认的扩展依赖项：`postgis`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 上游明确表示该项目尚未准备好投入生产。
- 上游将该项目描述为概念验证。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
