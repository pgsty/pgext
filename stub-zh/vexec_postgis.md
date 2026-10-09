## 用法

来源：

- [modules/vexec_postgis/README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_postgis/README.md)
- [modules/vexec_postgis/vexec_postgis--1.0.sql](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_postgis/vexec_postgis--1.0.sql)
- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)
- [modules/vexec_postgis/vexec_postgis.control](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_postgis/vexec_postgis.control)

`vexec_postgis` 是 PostgreSQL 19 上 `vexec` 向量化执行器的配套模块。它声明底层扩展函数的行为，让批量执行复用这些函数。

### 核心用法

```ini
shared_preload_libraries = 'vexec,vexec_postgis'
vexec.mode = auto
```

```sql
SELECT * FROM vexec_postgis.declarations();
SELECT * FROM vexec.declared_calls();
```

### 运行边界

在每个参与节点预加载 `vexec` 和 `vexec_postgis` 并重启，预加载顺序不限。通过 `vexec.mode` 启用向量化执行。`CREATE EXTENSION vexec_postgis` 只安装可选的检查接口，依赖底层 `postgis` 扩展，本身不会激活钩子。`vexec.declared_calls()` 显示已绑定的调用。不可迁移的 SQL 对象由超级用户创建。

仅绑定 PostGIS 3.7.0rc2，其他版本不会获得该包的加速。包围盒预过滤和点类型检查只处理受支持的情况，其余行调用原 PostGIS 函数。它不增加空间存储或新的索引方法。
