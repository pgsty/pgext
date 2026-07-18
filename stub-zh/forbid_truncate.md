## 用法

来源：

- [已复核 commit 的 forbid_truncate README](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/README.md)
- [已复核 commit 的 forbid_truncate 钩子实现](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.c)
- [已复核 commit 的 forbid_truncate control 文件](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.control)

`forbid_truncate` 安装一个 ProcessUtility 钩子，拒绝所有 TRUNCATE 语句。该库必须随 PostgreSQL 启动时加载。安装二进制后，把它加入服务器配置并重启 PostgreSQL：

```ini
shared_preload_libraries = 'forbid_truncate'
```

在需要纳管的每个数据库中登记扩展，再用一次性测试表验证保护；最后一条语句应以 “forbid truncate” 报错。

```sql
CREATE EXTENSION forbid_truncate;

CREATE TABLE truncate_guard_demo (id integer);
TRUNCATE TABLE truncate_guard_demo;
```

### 注意事项

- 修改 `shared_preload_libraries` 后必须重启服务器。只创建扩展而没有预加载该库，不会在既有后端进程中激活钩子。
- 实现会阻止包括超级用户在内的所有用户，并且没有按角色、按表或维护操作的绕过机制。
- 该钩子直接调用 PostgreSQL 标准工具命令处理器，而不继续调用先前安装的钩子。应与集群所用的其他 ProcessUtility 钩子扩展组合测试。
- 此保护只覆盖 TRUNCATE；DELETE、DROP TABLE、分区维护及其他数据删除路径不受影响。
