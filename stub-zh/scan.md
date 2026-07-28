## 用法

来源：

- [官方上游 README](https://github.com/jnidzwetzki/pg-dev-container/blob/d7cd5db9481e8fbc7f2d88139e40564463657b10/README.md)
- [官方扩展控制文件 (scan.control)](https://github.com/jnidzwetzki/pg-dev-container/blob/d7cd5db9481e8fbc7f2d88139e40564463657b10/src/extensions/06_scan/scan.control)
- [官方扩展 SQL (scan--1.0.sql)](https://github.com/jnidzwetzki/pg-dev-container/blob/d7cd5db9481e8fbc7f2d88139e40564463657b10/src/extensions/06_scan/scan--1.0.sql)

`scan` — Visual Studio Code - 开发容器 - PostgreSQL。当应用程序需要此特定数据库功能时使用它。在目标 PostgreSQL 构建中使用链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION scan;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `full_table_scan(REGCLASS)` 是一个扩展函数，返回 `VOID`。
- `get_attribute_type(tablename REGCLASS, attrname TEXT)` 是一个扩展函数，返回 `OID`。
- `table_scan_and_sort_attribute(tablename REGCLASS, attrname TEXT)` 是一个扩展函数，返回 `VOID`。
- `table_scan_with_index(tablename REGCLASS, indexname REGCLASS)` 是一个扩展函数，返回 `VOID`。
- `table_scan_with_scankeys(REGCLASS)` 是一个扩展函数，返回 `VOID`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
