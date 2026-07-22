## 用法

来源：

- [官方 README](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/README.md)
- [SQL API 与元数据](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/src/lib.rs)
- [原始堆加载器](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/src/loader.rs)
- [表模式辅助代码](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/src/schema.rs)
- [扩展控制文件](https://github.com/askyx/pg_tpch/blob/5684779c89b29b4674604c447a30ec07aaebf307/pg_tpch.control)

`pg_tpch` 1.0.0 是一个 pgrx 扩展，用于创建八张 TPC-H 表、通过 `tpchgen-rs` 在进程内生成数据，并取得 22 条具体基准查询。它通过专用原始堆文件加载器获得速度：绕过常规 PostgreSQL 元组插入，再重建索引。只能在可丢弃的基准数据库中使用，不能指向生产关系。

### 创建并加载数据集

在专用模式中运行流程。传入真会在表定义内嵌的主键之外，再创建扩展提供的二级索引。

```sql
CREATE EXTENSION pg_tpch;

CREATE SCHEMA benchmark;
SET search_path = benchmark, public;

SELECT create_tpch_tables(true);
SELECT * FROM tpch_dbgen(0.01);

SELECT query AS q1 FROM tpch_queries(1) \gset
:q1

SELECT cleanup_tpch_data();
```

如果安装了 `dblink`，`tpch_dbgen(double precision)` 可以打开本地连接并行加载各表；否则串行加载。调用者与本地认证配置必须允许这些连接。单表函数 `generate_region`、`generate_nation`、`generate_part`、`generate_supplier`、`generate_customer`、`generate_partsupp`、`generate_orders` 和 `generate_lineitem` 会加载一张已存在的表，并报告行数、堆写入时间与重建索引时间。

### 管理与元数据

- `create_tpch_tables(boolean)`：在当前模式创建标准表；默认省略二级索引。
- `create_tpch_indexes()`：创建元数据中存储的索引定义。
- `cleanup_tpch_data()`：截断八张表，保留定义与索引。
- `drop_tpch_tables()`：删除八张表。
- `tpch_queries(integer)`：返回一条查询；参数为 NULL 时返回全部查询。
- `tpch.tpch_table_metadata`：可编辑的完整表 DDL 与二级索引数组，供后续创建调用使用。
- `tpch.tpch_query_metadata`：内置具体查询文本。

修改元数据会在目标模式执行其中存储的 DDL。应把更新权和扩展函数限制给基准所有者，不能让不可信用户注入定义。

### 原始加载器安全边界

加载器取得 `AccessExclusiveLock`，自行构造堆页，由服务器进程打开关系物理段文件并直接写页。它只记录有限的新页 WAL 镜像路径，绕过普通缓冲区管理器插入与索引维护，随后执行 `REINDEX TABLE`。它还根据主机内存/CPU 选择激进的会话级重建索引设置，包括让 `maintenance_work_mem` 使用最多约三分之一的检测内存，并设置 `synchronous_commit = off`。

这些内部 API 与物理布局绑定 PostgreSQL 大版本。后端崩溃、操作系统崩溃、备库、备份、校验和、存储或索引交互都可能不同于普通 SQL/COPY 语义。即使现有贵重表的名称与列恰好匹配 TPC-H，也绝不能对其运行生成器。应使用精确匹配大版本的制品，隔离数据库与表空间，不运行并发负载，在克隆上验证校验和与恢复，并在基准后丢弃集群。性能数字只是上游测试环境的观察，不能保证胜过另一台主机上正确配置的 `COPY`。
