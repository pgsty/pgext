## 用法

来源：

- [官方说明](https://github.com/ExaDev/pg_warren/blob/v0.8.0/README.md)
- [控制文件](https://github.com/ExaDev/pg_warren/blob/v0.8.0/crates/pg_warren/pg_warren.control)
- [SQL 接口](https://github.com/ExaDev/pg_warren/blob/v0.8.0/sql/pg_warren--0.8.0.sql)
- [限制与共享缓存](https://github.com/ExaDev/pg_warren/blob/v0.8.0/docs/limits.md)

`pg_warren` 0.8.0 从普通 PostgreSQL 表构建版本化邻接快照，用于图遍历与图计算。它支持 PostgreSQL 15–18，图版本、向量及其变更都参与 PostgreSQL 事务。遍历读取已构建的快照；源表发生变化后需要显式刷新。

### 核心流程

这是可信扩展，在数据库上拥有 `CREATE` 权限的角色可以安装。基本图由节点查询与边查询定义：

```sql
CREATE EXTENSION pg_warren;
CREATE TABLE people (id bigint PRIMARY KEY, name text);
CREATE TABLE follows (src bigint, dst bigint);
INSERT INTO people VALUES (1, 'Ada'), (2, 'Bob'), (3, 'Cy');
INSERT INTO follows VALUES (1, 2), (2, 3);
SELECT warren.create_graph('social',
  'SELECT id FROM public.people', 'SELECT src, dst FROM public.follows');
SELECT warren.refresh('social');
SELECT * FROM warren.bfs('social', ARRAY[1::bigint], 3);
SELECT warren.pagerank('social', 'rank');
SELECT * FROM warren.vector_topk('social', 'rank', 3, NULL::bigint);
```

图会保存源查询。修改源表后执行 `warren.refresh`，使用 `warren.info` 查看版本，使用 `warren.diff` 比较差异。版本构建与切换会随外层事务一起回滚。

### 遍历与计算

| 接口 | 用途 |
| --- | --- |
| `warren.bfs`、`warren.shortest_path` | 广度优先遍历与最短路径 |
| `warren.pagerank`、`warren.components` | 节点排名与连通分量 |
| `warren.leiden` | 社区发现 |
| `warren.spmv`、`warren.gather` | 将图计算组合到 SQL 中 |
| `warren.vector_topk` | 读取已存向量中最大的值 |
| `warren.shared_cache`、`warren.backends` | 检查共享条目与后端缓存使用情况 |

标识符类型必须与图一致。没有节点标识符参数的函数可能需要 `NULL::bigint` 这样的类型示例值。0.8.0 的 PageRank 与 Leiden 使用原生计算内核，将结果写入命名向量。

### 内存与预加载

PostgreSQL 17–18 会在首次使用时初始化共享缓存。PostgreSQL 15–16 要在后端间共享图数据，需要配置以下参数并重启服务器：

```ini
shared_preload_libraries = 'pg_warren'
```

PostgreSQL 15–16 未预加载时仍可执行图操作，但使用各后端的私有缓存。如果已有预加载列表，应将该库追加到列表中。

`warren.shared_cache_memory` 限制整个集群的共享缓存，重新加载配置后生效。`warren.max_cache_memory`、`warren.max_compute_memory` 与 `warren.max_build_memory` 分别限制私有缓存、计算与构建分配，不是同一个总内存预算。`warren.threads` 控制计算并行度。Linux 上的共享条目还需要足够的 `/dev/shm` 空间，放不下的条目会退回私有缓存。

### 访问、备份与升级

源查询以图所有者身份执行，并使用保存的 `search_path`。应显式限定源对象模式，并使用受控模式。公共函数检查图访问权限；安装扩展不会授予私有目录表的直接访问权。

应以超级用户转储，包含图目录、版本与向量，并在恢复前建立所需的所有者和读取者角色。从源表重建图不能保留历史图版本。安装新库文件后，在每个使用该扩展的数据库执行 `ALTER EXTENSION pg_warren UPDATE`。删除扩展会删除已存图。该项目尚未到 1.0，不会增量维护图快照。
