## 用法

来源：

- [官方 README](https://github.com/tensorchord/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/README.md)
- [SQL 可见函数定义](https://github.com/tensorchord/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/src/lib.rs)
- [资源限制实现](https://github.com/tensorchord/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/src/rlimit.rs)

`pgnice` 0.0.0 允许会话查看并修改当前 PostgreSQL 后端进程的操作系统调度优先级与资源限制。它适合降低特定工作负载的优先级，或限制专用连接池中的后端。设置作用于进程，而不是事务或数据库角色。

### 核心流程

创建扩展，查看当前后端状态，再应用宿主操作系统与 PostgreSQL 服务账号允许的限制：

```sql
CREATE EXTENSION pgnice;

SELECT pgnice.get_backend_nice();
SELECT pgnice.get_backend_ionice();
SELECT pgnice.get_backend_rlimit('nofile');

SELECT pgnice.set_backend_nice(15);
SELECT pgnice.set_backend_ionice('B', 7);
SELECT pgnice.set_backend_rlimit('nofile', 4096);
```

`set_backend_rlimit` 会把软限制和硬限制设为同一个值。支持的资源名随操作系统而异，常见值包括 `cpu`、`data`、`fsize`、`nofile`、`nproc`、`rss`、`stack` 与 `as`。

### 函数索引

- `pgnice.get_backend_nice()` 与 `pgnice.set_backend_nice(integer)` 读取或设置 CPU 调度 nice 值。
- `pgnice.get_backend_ionice()` 与 `pgnice.set_backend_ionice(char, integer)` 读取或设置 I/O 类别及级别；类别为 `R`（实时）、`B`（尽力）与 `I`（空闲）。
- `pgnice.get_backend_rlimit(text)` 返回当前软、硬资源限制。
- `pgnice.set_backend_rlimit(text, bigint)` 为该后端同时设置两项资源限制。

### 运维与安全注意事项

内核会执行权限检查。普通 PostgreSQL 服务账号通常可以降低自身优先级或硬限制，却不能随后恢复，因此设置可能在该后端的整个生命周期内不可逆。使用会话池或事务池时，被修改的进程可能交给其他客户端复用；应拆分连接池，并显式初始化每个池的后端。I/O 优先级是 Linux 专用能力，资源名也因平台而异。

修改地址空间、打开文件数、CPU 时间或进程数等限制，可能终止后端或导致其不稳定。应撤销 public 执行权限，只允许受控工作负载策略所需的函数和值。
