## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/README.md)
- [pg_pure.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure.control)
- [pg_pureu.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu.control)
- [pg_pure--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure--0.1.sql)
- [pg_pureu--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu--0.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/CHANGELOG.md)

`pg_pure` 是可安装的受信任 PL/Pure 语言入口。分发版本 0.1.1 携带 SQL 版本 0.1，但在 libpure 具备沙箱之前，处理函数会主动拒绝执行。因此它是实验性占位实现，并非可用的受信任解释器；目前能够运行的是配套的 `pg_pureu`。

### 启用方式与当前边界

```sql
CREATE EXTENSION pg_pure;
SELECT extversion FROM pg_extension WHERE extname = 'pg_pure';
```

控制文件将扩展标记为受信任、不可迁移，创建扩展本身也不要求超级用户。该目录属性不会绕过处理函数的执行限制。不要为了绕过限制而把非受信任的配套语言交给普通角色使用。

### 运行条件

需要 PostgreSQL 17–19，以及上游适合嵌入宿主进程的 libpure 分支。两个语言共用 `pg_pure` 动态库。libpure 支持原生 `extern` 调用；受信任执行所需的沙箱尚未实现，未来必须限制原生调用、内联 C、求值与任意模块加载。

预加载 `pg_pure` 是可选项；通过 `shared_preload_libraries` 配置时需要重启服务器。`pg_pure.using` 选择预载模块，默认是 `math`。不预加载时，各后端会分别创建解释器。

### 部署限制

不要让 libpure 与 PostgreSQL 的 LLVM JIT 在同一后端加载不同主版本的 LLVM：上游报告这会导致后端退出时崩溃，进而可能触发整个集群重启。0.1 版本受支持的部署方式是使用不带 LLVM JIT 的 PostgreSQL；LLVM 主版本一致也是上游记录的配置。本次分发不代表已经具备生产可用性或受信任沙箱。
