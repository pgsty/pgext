## 用法

来源：

- [README.md](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/README.md)
- [src/pg/am.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/am.rs)
- [src/pg/catalog.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/catalog.rs)
- [src/pg/api.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/api.rs)
- [src/pg/worker.rs](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/src/pg/worker.rs)
- [Cargo.toml](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/Cargo.toml)
- [pghouse.control](https://github.com/myyrakle/pghouse/blob/a8f8b5e553586b29426d1eff1f1369e51eb3bb6d/pghouse.control)

`pghouse` 0.1.0 是面向 PostgreSQL 18 的原型，将基于堆表的表访问方法与异步维护的列存副本文件结合。普通 SQL 仍读取委托的堆表，显式扫描辅助函数才读取列存副本。它尚不是具备完整持久性保障的列存引擎。

### 核心用法

```sql
CREATE EXTENSION pghouse;
CREATE TABLE pghouse_demo (id bigint PRIMARY KEY, value text) USING pghouse;
SELECT pghouse_register_table('public.pghouse_demo', 'id');
INSERT INTO pghouse_demo VALUES (1, 'hello');
SELECT pghouse_run_maintenance_once();
SELECT pghouse_explain_scan('public.pghouse_demo');
SELECT pghouse_scan_rows('public.pghouse_demo');
```

### 运行边界

控制文件要求超级用户安装且禁止重定位；扩展会创建 `pghouse` 模式，无需共享预加载。`pghouse_register_table` 按表及主键列安装 DML 捕获，可配置粒度、压缩、缓冲区阈值与存储路径。`pghouse_run_maintenance_once` 处理待执行维护，`pghouse_launch_worker` 可在当前数据库启动动态工作进程。

`pghouse_scan_rows` 以 JSONB 返回副本中的行，可设置投影、键范围、上限与快照代次。`pghouse_explain_scan` 给出扫描计划，`pghouse_schedule_merge` 将合并加入队列。副本可见性取决于维护进度，不能将这些辅助函数等同于普通事务性表读取。

默认文件根目录为 `$PGDATA/pghouse/db_<db_oid>/rel_<table_oid>/`。粒度文件未接入 WAL，不应假定 PostgreSQL 复制、崩溃恢复或常规数据库备份会保护这些文件。应将权威数据保留在堆表，并在隔离的评估数据库试用。源码声明 pgrx 0.17，仅提供 PostgreSQL 18 特性；这不代表已提供 Pigsty 软件包。
