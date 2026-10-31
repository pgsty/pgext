## 用法

来源：

- [v1.5.3 README](https://github.com/aws/pg_tle/blob/v1.5.3/README.md)
- [v1.5.3 发布](https://github.com/aws/pg_tle/releases/tag/v1.5.3)
- [1.5.2 至 1.5.3 SQL](https://github.com/aws/pg_tle/blob/v1.5.3/pg_tle--1.5.2--1.5.3.sql)
- [v1.5.3 控制模板](https://github.com/aws/pg_tle/blob/v1.5.3/pg_tle.control.in)

`pg_tle` 1.5.3 管理以 PostgreSQL 可信语言编写的扩展。原生组件安装需要共享预加载和服务器重启，之后的 TLE 定义存储在数据库中。上游文档支持 PostgreSQL 12 至 19 Beta 4。

### 核心用法

```ini
shared_preload_libraries = 'pg_tle'
```

```sql
CREATE EXTENSION pg_tle;
SELECT pgtle.available_extensions();
```

### 运行边界

`pgtle.install_extension()` 存储扩展定义，`CREATE EXTENSION` 在数据库中启用它。`pgtle.install_update_path()` 与 `pgtle.set_default_version()` 管理版本，`pgtle.available_extensions()` 列出定义。1.5.3 新增 `pgtle.set_extension_schema()` 设置受管扩展的模式元数据，并收紧标识符校验。TLE 管理由 `pgtle_admin` 限制，已安装扩展的执行仍遵循 PostgreSQL 语言和对象权限。删除已启用扩展与卸载其 TLE 定义是两个独立操作。仅使用可信扩展代码，并审查依赖与升级 SQL。
