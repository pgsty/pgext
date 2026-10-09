## 用法

来源：

- [modules/vexec_pgvector/README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_pgvector/README.md)
- [modules/vexec_pgvector/vexec_pgvector--1.0.sql](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_pgvector/vexec_pgvector--1.0.sql)
- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)
- [modules/vexec_pgvector/vexec_pgvector.control](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_pgvector/vexec_pgvector.control)

`vexec_pgvector` 是 PostgreSQL 19 上 `vexec` 向量化执行器的配套模块。它声明底层扩展函数的行为，让批量执行复用这些函数。

### 核心用法

```ini
shared_preload_libraries = 'vexec,vexec_pgvector'
vexec.mode = auto
```

```sql
SELECT * FROM vexec_pgvector.declarations();
SELECT * FROM vexec.declared_calls();
```

### 运行边界

在每个参与节点预加载 `vexec` 和 `vexec_pgvector` 并重启，预加载顺序不限。通过 `vexec.mode` 启用向量化执行。`CREATE EXTENSION vexec_pgvector` 只安装可选的检查接口，依赖底层 `vector` 扩展，本身不会激活钩子。`vexec.declared_calls()` 显示已绑定的调用。不可迁移的 SQL 对象由超级用户创建。

仅绑定 pgvector 0.8.6 和 0.8.7，其他版本不会获得该包的加速。它通过 fmgr 调用 pgvector 函数并检查维度是否相同；会报错的行回退到 PostgreSQL 求值。它不增加新的向量索引方法。
