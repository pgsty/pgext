

## 用法

> [documentdb_distributed: DocumentDB 的多节点 API 接口](https://github.com/documentdb/documentdb)

`documentdb_distributed` 扩展为 PostgreSQL 上的 DocumentDB 提供多节点分布式能力。它扩展了核心 DocumentDB 功能以支持跨多个 PostgreSQL 节点的水平扩展。

### 概述

该扩展与 `documentdb_core` 配合使用，提供分布式文档数据库操作。它支持：

- 将文档集合分片到多个节点
- 分布式查询执行，用于 MongoDB 兼容操作
- 大规模文档工作负载的水平扩展

### 前置条件

`documentdb_distributed` 扩展需要：

- 已安装并配置 `documentdb_core` 扩展
- 多节点 PostgreSQL 集群（通常使用 Citus 进行分布）

### 启用

```sql
CREATE EXTENSION documentdb_distributed;
```

分布式层透明地将 CRUD 操作和聚合管道路由到集群节点，同时保持 MongoDB 线协议兼容性。
