## 用法

来源：

- [DocumentDB Extended RUM README](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_extended_rum/README.md)
- [`documentdb_extended_rum`控制文件](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)
- [访问方法SQL定义](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_extended_rum/sql/documentdb_extended_rum--0.106-0.sql)
- [DocumentDB v0.114-0变更日志](https://github.com/documentdb/documentdb/blob/v0.114-0/CHANGELOG.md)

`documentdb_extended_rum`是DocumentDB的扩展RUM索引访问方法。它是由DocumentDB索引层选择的实现组件，而不是通用的应用程序索引或替代安装`documentdb`。

### 配置与安装

该库只能从`shared_preload_libraries`初始化。在基础DocumentDB库之后预加载它并重启PostgreSQL：

```conf
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum'
documentdb.alternate_index_handler_name = 'extended_rum'
```

然后使用相同的版本安装扩展，就像安装基础堆栈一样：

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_extended_rum;
```

DocumentDB部署工具通常会管理此配置。现有数据库应遵循特定于版本的升级程序而不是随意切换索引处理器。

### 重要对象

- `documentdb_extended_rum`是该扩展注册的索引访问方法。
- `documentdb_extended_rum_catalog`包含用于DocumentDB的BSON操作符家族和类。
- `documentdb.alternate_index_handler_name = 'extended_rum'`指示DocumentDB索引层使用适配器。
- 实现是一个RUM分支，其磁盘布局和内容设计保持与上游RUM向后兼容的同时，改变查询和易变路径以适应文档工作负载。

### 运行边界

安装和升级此组件时，请确保与`documentdb`和`documentdb_core`二进制文件版本匹配。除非遵循上游开发指导，否则不要直接使用其内部操作符类构建索引；通过DocumentDB API创建和管理索引以保持元数据的一致性。

v0.114-0变更日志描述了RUM WAL页面重用标记和目标发布树修剪功能，但它们是受控特性的，默认情况下未启用。这些特性不是此版本的默认用户可见功能。
