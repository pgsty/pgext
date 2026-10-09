## 用法

来源：

- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)
- [modules/vexec/vexec--1.0.sql](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec/vexec--1.0.sql)
- [modules/vexec/vexec.control](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec/vexec.control)

`vexec` 在 PostgreSQL 19 内按批次规划和执行查询，可以使用 PostgreSQL 规划器或另行安装的 ORCA 集成。不支持的表达式回退到 PostgreSQL 求值；Cloudberry 的存储和节点间通信集成另有分支要求。

### 核心用法

```ini
shared_preload_libraries = 'vexec'
vexec.mode = auto
```

```sql
EXPLAIN (VEXEC) SELECT count(*) FROM pg_catalog.pg_class;
```

### 运行边界

将库追加到现有预加载列表并重启。`vexec.mode` 默认为 off；auto 使用成本模型，force 更积极地选择支持的向量化路径。普通向量化执行不要求创建 SQL 扩展；由超级用户创建后提供 `vexec.sources`、`vexec.sinks`、`vexec.type_layouts`、`vexec.kernels`、`vexec.kernel_packs`、`vexec.declared_calls` 和 `vexec.ingest_stream`，最后一个函数是 Flight 批量导入的前提。

这是面向 PostgreSQL 19 的早期源码版本。采用前应与关闭模式时的结果和计划比较。ORCA 需要兼容的 gp_orca 模块；PAX、ao_column 和分布式执行需要单独的 Cloudberry 移植版本。安装基础扩展不会提供这些可选集成。
