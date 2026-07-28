## 用法

来源：

- [官方上游 README](https://gitlab.com/nfiesta/nfiesta_target_data/-/blob/main/README.md)
- [官方扩展控制文件](https://gitlab.com/nfiesta/nfiesta_target_data/-/blob/main/nfiesta_target_data.control)
- [官方项目页面](https://gitlab.com/nfiesta/nfiesta_target_data)

`nfiesta_target_data` — 这是用于处理目标数据的 PostgreSQL 扩展，可以在图层级别聚合局部密度贡献。请在相应的空间数据或地理空间工作流中使用它。在安装扩展及其依赖项并验证它们之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION nfiesta_target_data;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `2.33.2`。
- 请首先安装并验证确认的扩展依赖项：`plpgsql`、`plpython3u`、`nfiesta_sdesign`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
