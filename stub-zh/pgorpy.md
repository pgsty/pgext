## 用法

来源：

- [官方上游 README](https://github.com/pgrouting/pgorpy/blob/9ec1b516e1358d6fbf05dfaf0da93f10dd1d766f/README.md)
- [官方扩展控制文件 (pgorpy.control)](https://github.com/pgrouting/pgorpy/blob/9ec1b516e1358d6fbf05dfaf0da93f10dd1d766f/sql/pg_controls/pgorpy.control)

`pgorpy` — OR-tools Python 从数据库可访问。当 SQL 需要这些特殊函数或聚合时使用它。在安装扩展依赖并验证它们之前，请勿集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION pgorpy;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.0.0`。
- 首先安装确认的扩展依赖项：`plpython3u`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
