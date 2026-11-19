## 用法

来源：

- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/README.md)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/docs/USER_MANUAL.md)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/RELEASE_NOTES_v1.0.0.md)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/include/common.h)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/Makefile)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/pg_fastload.control.in)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/pg_fastload.sql.in)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/pg_fastload.c)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/export.c)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/bin/recovery.c)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/LICENSE)

`pg_fastload` `1.0.0` 是具有独立扩展身份的 pg_bulkload 分支，提供并行解析、批量加载和文件导出。已检查的发行源码支持 PostgreSQL 15–19；上游完整二进制验证使用 x86_64 Linux 上的 PostgreSQL 18.6。控制文件和版本化安装 SQL 由模板生成，对象位于固定的 `pgfastload` 模式。

### 加载与导出

安装及加载、导出函数都要求数据库超级用户。服务端输入文件必须允许 PostgreSQL 操作系统用户读取，输出目录必须允许其写入。仅授予模式访问权限不能绕过函数内的超级用户检查。

对于不含表头的 CSV 文件，下面的示例使用正常的缓冲区和 WAL 写入路径：

```sql
CREATE EXTENSION pg_fastload;
CREATE TABLE load_target (id integer, label text);
SELECT * FROM pgfastload.pg_fastload(ARRAY[
  'INPUT=/srv/import/data.csv',
  'OUTPUT=public.load_target',
  'TYPE=CSV',
  'WRITER=BUFFERED',
  'PARALLEL=ON',
  'PARALLEL_WORKERS=4'
]);
```

无需配置 `shared_preload_libraries`。并行加载使用动态后台工作进程，`max_worker_processes` 必须容纳解析进程、一个读取进程及集群其他工作进程。修改该服务端参数需要重启；单独安装扩展无需重启。

命令行通过控制文件接受相同选项：`pg_fastload control.ctl`。SQL 加载函数返回行数、拒绝记录数与耗时。导出要求尚不存在的服务端绝对输出路径：

```sql
SELECT pgfastload.pg_fastload_export(ARRAY[
  'TABLE=public.load_target',
  'OUTPUT=/srv/export/load_target.csv',
  'TYPE=CSV',
  'HEADER=ON',
  'PARALLEL=OFF'
]);
SELECT * FROM pgfastload.pg_fastload_progress;
```

`TABLE=` 导出可在共享 MVCC 快照下并行运行，输出顺序不确定。`QUERY=` 接受一条只读查询，始终使用串行路径。长时间导出可能延迟清理旧行。进度视图会暴露数据库、关系和文件名称，应限制监控授权；底层 `fl_progress_rows()` 函数已撤销公共执行权限。

### 写入与恢复边界

默认的 `WRITER=DIRECT` 绕过正常缓冲路径写数据文件，在 `$PGDATA/pg_fastload/` 中保留加载状态。直接加载中断后，目标表可能无法使用。应演练文档中的离线恢复流程：停止数据库服务后运行 `pg_fastload -r -D "$PGDATA"`；恢复程序会检查服务进程确已退出。不要通过删除状态文件跳过恢复。需要正常缓冲区和 WAL 路径时选择 `WRITER=BUFFERED`。

`PARALLEL=ON` 并行处理读取和解析，插入仍位于调用者事务内。独立的 `WRITER=PARALLEL` 模式通过另一连接和事务插入，与调用者不具备事务原子性，也不会与并行读取组合运行。

### 约束与索引

这是高权限批量加载工具，不能直接替代普通应用插入。并行加载期间会独占锁定目标表。过滤器、编码覆盖、约束检查、允许解析错误以及若干输入模式等不支持的条件会回退到串行执行并给出诊断。`CHECK_CONSTRAINTS` 默认为 `NO`，应明确设置并验证所需的校验行为。并行加载不保持行顺序，也不提供逐行容错。

`SKIP_INDEX=ON` 在加载事务中删除二级索引并输出重建语句，不会自动重建。依附约束的索引会被拒绝删除，该选项不能与 `DUP_ERRORS` 组合。独立唯一索引可能被删除，随之失去重复值保护。应保存输出的重建语句，并在恢复业务使用前核验索引。
