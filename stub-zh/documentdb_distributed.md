## 用法

来源：

- [`documentdb_distributed` v0.114-0 控制文件](https://github.com/documentdb/documentdb/blob/v0.114-0/internal/pg_documentdb_distributed/documentdb_distributed.control)
- [DocumentDB 包范围](https://github.com/documentdb/documentdb/blob/v0.114-0/packaging/README.md)
- [官方预加载助手](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)
- [DocumentDB v0.114-0 更新日志](https://github.com/documentdb/documentdb/blob/v0.114-0/CHANGELOG.md)

`documentdb_distributed` 是 DocumentDB 的内部多节点执行层。它将公共 `documentdb` API 与 Citus 集成；它不是一个独立的文档 API，并不会增加单独的客户端工作流程。

### 先决条件和安装

所有节点需要匹配版本的 Citus、`documentdb_core`、`documentdb` 和 `documentdb_distributed`。官方助手会将库按预加载顺序放置；在更改后，请重启每个节点：

```conf
shared_preload_libraries = 'citus, pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_distributed'
```

配置好 Citus 拓扑和基础 DocumentDB 堆栈之后，以超级用户身份安装分布式组件：

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_distributed;

SELECT extname, extversion
FROM pg_extension
WHERE extname IN ('citus', 'documentdb_core', 'documentdb', 'documentdb_distributed');
```

安装完成后可以使用正常的 DocumentDB 网关或 `documentdb_api` 函数。分布式的集合放置、分片拓扑、工作进程可用性以及元数据一致性必须作为集群部署的一部分进行管理。

### 重要边界

- 控制文件需要 `citus`、`documentdb_core` 和 `documentdb`，并标记该扩展为超级用户专用且不可重定位。
- `documentdb_distributed` 向现有的 DocumentDB 命令提供分布式规划和执行支持；它不是公共 API 扩展的替代品。
- 在尝试升级扩展之前，发布版本应在协调器和工作进程之间保持同步。
- 备份、恢复、故障转移以及滚动升级流程必须包括 Citus 元数据和 DocumentDB 数据。

上游 `packaging/README.md` 明确指出标准包不包含 `internal/pg_documentdb_distributed` 组件。在添加到配置之前，请确认该分发实际包含了此扩展。版本 0.114-0 包含了一个后端特性标志下的分片 `$sample` 优化修复；不应将其视为无条件行为。
