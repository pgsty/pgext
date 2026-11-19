## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/pg_python.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/pg_python--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/pg_python.c)

`pg_python` `1.0` 提供通过 PL/Python 执行注册的 Python 脚本。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pg_python;
SELECT python_eval('return args["a"] + args["b"]', '{"a":2,"b":3}');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

python_register 与 python_exec 管理脚本版本并执行，参数以 JSON 对象传入。执行使用进程内的不受信任 plpython3u 运行时。环境注册表和 pg_python.executable 仅为元数据，不提供进程隔离或包安装器。运行器、注册表与包装函数只应向可信用户开放：Python 可访问服务器资源，外部副作用不会随 SQL 回滚。

此模块对应上游旧名 `alohadb_python`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
