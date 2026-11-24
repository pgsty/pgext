## 用法

来源：

- [README 0.1.12](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/README.md)
- [Control](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/pg_tci.control)
- [SQL 0.1.12](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/sql/pg_tci--0.1.12.sql)
- [Operations](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/docs/operations.md)
- [Limitations](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/docs/limitations.md)

`pg_tci` 0.1.12 是面向 PostgreSQL 18 的外部预览版本，用于统计某个时刻被多少时间区间覆盖。它通过显式函数调用工作，不是索引访问方法，也不会自动改写查询。

### 基本流程

安装需要超级用户。区间左闭右开，结束时间为空表示尚未结束。

```sql
CREATE EXTENSION pg_tci;
CREATE TABLE sessions (
    id bigserial PRIMARY KEY,
    started_at timestamptz NOT NULL,
    ended_at timestamptz
);
SELECT tci.create('sessions'::regclass, 'started_at', 'ended_at');
SELECT tci.count_at('sessions'::regclass, now());
SELECT * FROM tci.stats('sessions'::regclass);
```

`tci.create_filtered`、`tci.count_at_filtered` 支持预先登记的等值过滤，不支持任意谓词。`tci.create_partitioned`、`tci.count_at_partitioned` 对叶子分区执行操作；新增挂接分区后应重新登记。

### 维护与权限

使用 `tci.health`、`tci.seal_bounded`、`tci.compact_pairwise` 和 `tci.maintain` 监控增量尾部、封装数据并合并运行段：

```sql
SELECT * FROM tci.health('sessions'::regclass);
SELECT tci.seal_bounded('sessions'::regclass, 8192);
SELECT * FROM tci.maintain('sessions'::regclass);
```

手动维护或外部调度无需预加载。可选后台进程要求将 `pg_tci` 加入 `shared_preload_libraries`，启用 `pg_tci.enable_maintenance_worker` 并重启服务器；通过 `pg_tci.maintenance_databases` 选择数据库。

读取接口检查表访问权限。压缩结构的计数会拒绝可能绕过行级安全的请求；使用 `tci.count_at_safe` 或 `tci.count_at_filtered_safe` 按当前策略进行原生计数。管理函数未向 PUBLIC 开放执行权限；初始化时向表所有者授权，DBA 也可以使用 `tci.grant_admin`。

### 限制与恢复

每张表仅支持一对开始、结束时间戳。写入会增加增量记录，需要持续维护；高并发写入和合并可能主导查询延迟。此预览版本不是生产 GA 版本，也不代表形式化的 MVCC 正确性证明。

结构变更或恢复后，无效索引会回退到原生计数，直至完成修复。按文档执行重建，例如 `tci.rebuild_invalid`，并检查健康状态后再依赖压缩结构的结果。
