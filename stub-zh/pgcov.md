## 用法

来源：

- [上游 README](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/README.md)
- [版本 1.0 SQL API](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/pgcov--1.0.sql)
- [覆盖率实现](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/pgcov.c)

`pgcov` 版本 `1.0` 是一个实验性 PL/pgSQL 语句覆盖率收集器。它通过预加载钩子与共享内存，把测试后端的执行报告发送到一个监听会话。

### 核心流程

```conf
shared_preload_libraries = 'pgcov'
```

重启 PostgreSQL、安装扩展，然后让一个专用会话运行阻塞式监听器，同时在其他会话执行测试。

```sql
CREATE EXTENSION pgcov;
SELECT pgcov_listen();
```

测试负载完成后取消监听器，并在同一个后端中检查已收集的调用与逐行计数：

```sql
SELECT * FROM pgcov_called_functions();
SELECT * FROM pgcov_fn_line_coverage('public.my_function(integer)');
SELECT pgcov_fn_line_coverage_src('public.my_function(integer)');
```

同时只允许一个活动监听器。收集状态属于监听器所在后端；会话结束后状态会丢失。本次核对的实现会让 `pgcov_called_functions` 返回的 `coverage` 列保持 NULL，因此应从逐行结果计算覆盖率，不要依赖该字段。

### 运行注意事项

上游称该项目仍在开发中，目录也将其标为废弃。它依赖 PostgreSQL 内部 PL/pgSQL 执行钩子和旧式共享内存 API，不能假定兼容现代版本。只能在隔离测试集群中使用：预加载与插桩会影响整台服务器，崩溃或取消可能中断收集。在信任报告前，应验证函数签名、嵌套调用、异常、并行测试与监听器清理。
