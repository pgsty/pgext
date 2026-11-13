## 用法

来源：

- [DocumentDB Extended RUM README](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/README.md)
- [`documentdb_extended_rum` 控制文件](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)
- [访问方法SQL定义](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/sql/documentdb_extended_rum--0.106-0.sql)
- [DocumentDB v1.0-0变更日志](https://github.com/documentdb/documentdb/blob/v1.0-0/CHANGELOG.md)
- [DocumentDB 1.2 release](https://github.com/documentdb/documentdb/releases/tag/v1.2-0)
- [1.2 control](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)
- [1.1 to 1.2 SQL](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb_extended_rum/sql/documentdb_extended_rum--1.1-0--1.2-0.sql)
- [Official upgrade procedure](https://github.com/documentdb/documentdb/blob/v1.2-0/packaging/README.md)

`documentdb_extended_rum` 是DocumentDB的扩展RUM索引访问方法。它是由DocumentDB索引层选择的实现组件，而不是通用的应用程序索引或替代安装 `documentdb`。

### 配置与安装

该库只能从 `shared_preload_libraries` 初始化。在基础DocumentDB库之后预加载它并重启PostgreSQL：

```ini
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum'
documentdb.rum_library_load_option = 'require_documentdb_extended_rum'
```

然后使用相同的版本安装扩展，就像安装基础堆栈一样：

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_extended_rum;
```

DocumentDB部署工具通常会管理此配置。现有数据库应遵循特定于版本的升级程序而不是随意切换索引处理器。

### 重要对象

- `documentdb_extended_rum` 是该扩展注册的索引访问方法。
- `documentdb_extended_rum_catalog` 包含用于DocumentDB的BSON操作符家族和类。
- `documentdb.rum_library_load_option = 'require_documentdb_extended_rum'` 指示DocumentDB索引层使用适配器。
- 实现是一个RUM分支，其磁盘布局和内容设计保持与上游RUM向后兼容的同时，改变查询和易变路径以适应文档工作负载。

### 运行边界

安装和升级此组件时，请确保与 `documentdb` 和 `documentdb_core` 二进制文件版本匹配。除非遵循上游开发指导，否则不要直接使用其内部操作符类构建索引；通过DocumentDB API创建和管理索引以保持元数据的一致性。

在 1.0-0 中，`documentdb.rum_library_load_option` 继续默认使用 `require_documentdb_extended_rum`。本版本修复了清理时并发根页分裂的处理，恢复清理延迟与中断检查，并默认启用定向倒排树剪枝与空条目叶页回收。应应用匹配的升级 SQL，并通过受支持的 DocumentDB 操作验证现有索引。

### DocumentDB 1.2 变化

1.2 源码为此组件提供 SQL 版本 `1.2-0`。备份后，应与匹配的 DocumentDB API/core 一起升级；不能安全地将升级后的数据降回 1.0。此版本修复了大型索引项导致最右叶页分裂时溢出的问题。真空清理期间回收孤立的空 posting-tree 叶页，另由 `documentdb_rum.enable_vacuum_cleanup_posting_tree_pruning` 控制，默认关闭；不能因为升级了扩展就认为这项可选清理已经启用。
