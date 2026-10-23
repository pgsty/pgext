## 用法

来源：

- [官方资料](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/README.md)
- [官方资料](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/Cargo.toml)
- [官方资料](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/crates/pg_warren/pg_warren.control)
- [官方资料](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/sql/pg_warren--0.4.0.sql)
- [官方资料](https://github.com/ExaDev/pg_warren/blob/0a14927e2fa6336439133c386ea7abb27e60a74b/docs/limits.md)

`pg_warren` 从普通 SQL 表构建版本化邻接快照，用于图遍历与图计算。源码修订 0a14927e 声明开发版本 0.4.0。README 中提到的 0.3.0 发布引用无法解析，因此本文说明的是固定修订的开发源码，不代表已经确认的稳定发行版。

### 核心流程

该可信扩展支持 PostgreSQL 15–18，在数据库上拥有 `CREATE` 权限的角色可以安装，无需预加载。

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
```

遍历读取不可变版本，而非源表的当前记录。源数据变化后需要刷新。版本构建与切换参与事务，可用 `warren.info` 和 `warren.diff` 检查状态与差异。

### 计算与内存

`warren.pagerank`、`warren.components` 和 `warren.shortest_path` 提供常见算法。`warren.spmv`、`warren.gather` 与向量接口可以将图计算组合到 SQL 中。标识符类型必须与图一致；没有标识符参数的函数可能需要显式提供结果类型示例值。

缓存属于每个后端进程，连接池会放大内存占用。`warren.max_cache_memory`、`warren.max_compute_memory` 与 `warren.max_build_memory` 分别约束不同分配，不构成一个全局内存上限。`warren.threads` 控制计算并行度。源码尚处于 1.0 之前，不提供自动增量维护或后端之间的共享缓存。

### 访问、备份与升级

图的源查询以图所有者身份执行，并使用保存的 `search_path`。应显式限定对象模式，只信任由可靠角色可写的模式。公共函数会检查图访问权限，安装扩展并不自动授予私有目录表的直接访问权。

应以超级用户转储，才能包含图目录、版本和向量；普通角色转储可能因私有表权限失败。恢复前先建立所需所有者和读取者角色，并遵循普通恢复流程，保留目录检查。从源表重建图是另一种恢复选择，不等同于保留历史版本。安装新库文件后，应先更新扩展 SQL 再调用函数。删除扩展会删除其中保存的图。
