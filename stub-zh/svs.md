## 用法

来源：

- [README](https://api.pgxn.org/src/svs/svs-0.1.0/README.md)
- [Control](https://api.pgxn.org/src/svs/svs-0.1.0/svs.control)
- [Install SQL](https://api.pgxn.org/src/svs/svs-0.1.0/sql/svs--0.1.0.sql)
- [Server setup](https://api.pgxn.org/src/svs/svs-0.1.0/docs/build_guide/README.md)
- [User guide](https://api.pgxn.org/src/svs/svs-0.1.0/docs/USER_GUIDE.md)

`svs` 为 pgvector 增加 Vamana 近似近邻索引，并可选用 LeanVec 与 LVQ 压缩。此 0.1.0 发行要求 Linux x86-64、18 系列的 PostgreSQL 18.1 或更新小版本、pgvector 0.8.2 或更高版本，以及 Intel SVS C API。压缩索引要求 Intel CPU；不支持 ARM、macOS 或 Windows。

### 服务器与数据库配置

预加载 `svs`，启用逻辑 WAL 并重启。每个 Vamana 索引需要一个逻辑复制槽；`max_replication_slots` 应覆盖这些索引及其他使用者。未启用逻辑 WAL 时启动器会停止。在 `svs.launcher_database` 指定的启动器数据库中安装依赖，并通过 `vamana_databases` 登记目标库；默认启动器库为 `postgres`。创建扩展需要超级用户。

```ini
shared_preload_libraries = 'svs'
wal_level = logical
max_replication_slots = 10
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION svs;
CREATE DATABASE svs_demo;
INSERT INTO vamana_databases (datname) VALUES ('svs_demo');
```

### 建立索引与查询

连接已登记的数据库，在其中创建两个扩展后建立索引。查询应使用距离排序表达式与 `LIMIT`。示例提供完整的三维查询向量。此索引方法也支持 `halfvec`；操作符类分别以 `<->`、`<#>` 和 `<=>` 对应 L2、内积与余弦距离。内积以负值返回，便于升序扫描。

```sql
CREATE EXTENSION vector;
CREATE EXTENSION svs;
CREATE TABLE items (id bigint GENERATED ALWAYS AS IDENTITY, embedding vector(3));
INSERT INTO items (embedding) VALUES ('[1,0,0]'), ('[0,1,0]'), ('[0,0,1]');
CREATE INDEX items_vamana ON items USING vamana (embedding vector_l2_ops);
SELECT id FROM items ORDER BY embedding <-> '[1,0,0]'::vector LIMIT 2;
```

### 调优与监控

索引选项包括 `graph_degree`、`build_window_size` 和 `compression_type`（0：不压缩，1：LeanVec，2：LVQ）。查询搜索范围由 `svs.search_window_size` GUC 控制，同名索引选项不是运行时回退值。构建内存由 `svs.max_build_memory` 控制，不受 `maintenance_work_mem` 控制；驻留内存与搜索临时内存有独立预算。通过 `pg_stat_vamana_worker` 和 `pg_stat_vamana_worker_slot` 视图检查状态；调用 `svs_restart_worker` 需要启动器登记表的更新权限。

### 运维边界

索引由每库后台工作进程持有，崩溃后通过逻辑 WAL 重放恢复。复制槽可能保留 WAL，应监控保留量和资源上限。暂停登记数据库与永久移除不同。永久移除时，先在目标库调用 `svs_teardown_database()`，再删除启动器登记；此函数会检查所有权并删除 Vamana 索引。不要在仍保留索引时删除登记。并发写入可能阻塞搜索，ANN 结果为近似结果。
