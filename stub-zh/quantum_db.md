## 用法

来源：

- [官方资料](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/README.md)
- [官方资料](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/pg_extension/README.md)
- [官方资料](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/pg_extension/quantum_db.control)
- [官方资料](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/pg_extension/sql/quantum_db--0.1.1.sql)

`quantum_db` 0.1.1 是连接本地张量网络与量子侧车的 PostgreSQL 16 研究型桥接模块。PostgreSQL 仍是权威执行方，会重新检查候选行谓词。项目明确没有许可证；源码公开不代表已授予复用许可。

### 启动本地侧车

侧车需要 Python 3.11 或更高版本。真实量子硬件为可选能力，另需服务商访问权限；普通 PostgreSQL 集成不能证明已经运行硬件任务。

```bash
qdb-sidecar --socket /tmp/qdb.sock --state outputs/qdb_jobs.sqlite3
```

配置 PostgreSQL，并重启以在所有后端加载规划器钩子：

```ini
shared_preload_libraries = 'quantum_db'
quantum_db.socket_path = '/tmp/qdb.sock'
```

由超级用户在每个目标数据库中安装，扩展依赖标准过程语言：

```sql
CREATE EXTENSION quantum_db;
SELECT qdb_explain('SELECT 1');
```

### 数据集与任务接口

`qdb_register_dataset` 向侧车注册关系、主键列和索引列。在 `qdb_unregister_dataset` 取消注册前，已注册关系会拒绝插入、更新、删除与截断。不能在未安排写入限制的情况下注册持续写入的应用表。

`qdb_submit`、`qdb_status`、`qdb_result` 和 `qdb_cancel` 管理异步任务。`qdb` 模式存储数据集、产物和任务审计记录。SQL 文本及选定的数据集信息会跨越本地 Unix socket 边界，应保护 socket、侧车状态与云端凭据。部分 SQL 包装函数使用定义者权限，向不可信角色开放前应检查执行授权。

### 故障与升级边界

0.1.1 增加产物校验和检查与 SQL 任务审计镜像。尚未升级的旧表结构会保留原生扫描。安装匹配文件、重启使用旧库的进程，再逐库升级：

```sql
ALTER EXTENSION quantum_db UPDATE TO '0.1.1';
```

侧车是独立运维依赖，其 SQLite 状态不在 PostgreSQL 单库备份中。应验证服务就绪状态、数据集快照一致性及回退行为。研究遥测或候选生成不能证明完整查询获得量子加速。
