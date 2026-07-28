## 用法

来源：

- [官方扩展控制文件](https://gitlab.com/nfiesta/nfiesta_gisdata_old/-/blob/main/extension/nfiesta_gisdata.control)
- [官方项目页面](https://gitlab.com/nfiesta/nfiesta_gisdata_old)

`nfiesta_gisdata` — nfi gisdata 数据库。使用它来进行相应的空间数据或地理空间工作流。在安装扩展及其依赖项并验证之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION nfiesta_gisdata;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审核后的控制文件声明默认版本为 `3.0.7`。
- 首先安装确认的扩展依赖项：`plpgsql`, `postgis`, `postgis_raster`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
