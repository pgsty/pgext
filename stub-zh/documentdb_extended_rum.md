

## 用法

> [documentdb_extended_rum: DocumentDB 扩展 RUM 索引访问方法](https://github.com/documentdb/documentdb)

`documentdb_extended_rum` 扩展为 PostgreSQL 上的 DocumentDB 提供增强的 RUM（递归联合归并）索引访问方法，改善基于文档的工作负载查询性能。

### 概述

该扩展扩展了 RUM 索引类型，以更好地支持 DocumentDB 中的 BSON 文档索引。它提供了以下优化的索引访问方法：

- 文档字段的全文搜索
- BSON 数据的复合索引操作
- 索引文档属性的高效范围查询和排序

### 前置条件

需要安装 `documentdb_core`。

### 启用

```sql
CREATE EXTENSION documentdb_extended_rum;
```

扩展的 RUM 索引在文档查询模式适合时由 DocumentDB 查询规划器自动使用。
