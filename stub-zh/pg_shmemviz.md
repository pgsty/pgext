## 用法

来源：

- [README.md](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/README.md)
- [extension/pg_shmemviz.control](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/extension/pg_shmemviz.control)
- [extension/pg_shmemviz--0.1.sql](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/extension/pg_shmemviz--0.1.sql)
- [LICENSE](https://github.com/bdrouvot/pg_shmemviz/blob/423608b9ddf775d90fe5bd311753af9f93851d9d/LICENSE)

`pg_shmemviz` 捕获 PostgreSQL 共享内存，供离线浏览器查看器分析。扩展版本为 0.1，README 将诊断工具版本标为 0.1.0-beta.1，并仅验证了 PostgreSQL 20devel。只应在可丢弃或隔离的开发实例中使用。

### 捕获与查看

需要与运行服务器准确匹配的开发文件，以及未剥离且含 DWARF 调试信息的服务器可执行文件。Python CLI 在 macOS 上需要支持 Python 的 LLDB，其他系统需要支持 Python 的 GDB，还需本地超级用户数据库连接。

```sql
CREATE EXTENSION pg_shmemviz;
```

```sh
pg_shmemviz capture --pg-config /opt/postgresql/bin/pg_config --dbname postgres /tmp/pg-memory-snapshot
pg_shmemviz serve /tmp/pg-memory-snapshot
```

### 对象与边界

`pg_shmemviz_snapshot(text)` 和 `pg_shmemviz_snapshot(text, text)` 将快照写入服务端路径，默认撤销公共执行权限。CLI 将捕获结果与可执行文件的结构信息合并。`--numa-scope` 选择 NUMA 检查范围，`pg_buffercache` 和 `--buffercache-details` 可补充缓冲区信息。查看器默认绑定回环地址。

无需预加载。捕获按顺序无锁复制活动内存，相关字段可能对应不同时间点，多字节值也可能撕裂。快照可能包含敏感数据并占用大量磁盘空间；NUMA 检查可能触发缺页并影响放置。调试器读取可执行文件元数据，不附加到服务器进程；查看器不会修改服务器中被捕获的内存。
