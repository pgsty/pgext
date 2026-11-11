## 用法

来源：

- [教学示例文档](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/README.md)
- [扩展控制文件](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/traceheap.control)
- [SQL 定义](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/traceheap--1.0.sql)
- [可执行教学测试](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/test.sql)

`traceheap` 是 BioseqDB 教程提供的 PostgreSQL 17 教学表访问方法。`1.0` 版本保留堆存储，同时统计部分表访问回调。它不是 BioseqDB 的列式引擎，也不适合生产存储。源码明确拒绝其他 PostgreSQL 主版本。

### 观察扫描

在可丢弃的 PostgreSQL 17 环境中安装库和扩展文件。创建访问方法需要超级用户，无需共享预加载。

```sql
CREATE EXTENSION traceheap;
SET max_parallel_workers_per_gather = 0;

CREATE TABLE specimens (
    id integer NOT NULL,
    label text NOT NULL,
    read_count integer NOT NULL
) USING traceheap;

INSERT INTO specimens VALUES (1, 'soil', 120), (2, 'river', 900);
SELECT traceheap.reset();
SELECT id, label FROM specimens WHERE read_count >= 400 ORDER BY id;
SELECT traceheap.stats();
```

`traceheap.stats()` 返回当前后端的 JSON 计数器，`traceheap.reset()` 将其清零。计数器不受事务回滚保护、不持久化、不按表区分，也不会汇总并行工作进程。被回滚的写入仍可能增加计数，采样及委托执行的扫描路径也未被全部统计。

### 教学支持范围

上游测试覆盖顺序扫描、插入、`COPY`、更新与删除、事务回滚、`ANALYZE`、普通 `VACUUM`、堆 TOAST 和 `TRUNCATE`。明确不支持用户索引、依赖索引的主键或唯一约束、索引验证，以及包括 `VACUUM FULL` 在内的聚簇重写操作。

扩展拥有 `traceheap` 模式且不可迁移。实验只应使用可丢弃的数据表；某项操作通过测试，并不表示所有堆操作都受支持。扩展示例前应先阅读上游测试脚本。
