## 用法

来源：

- [锁定提交的上游 README](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/README.md)
- [0.1 版 SQL API](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/pg_worker_pool--0.1.sql)
- [锁定提交的工作进程与队列实现](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/pg_worker_pool.c)
- [锁定提交的扩展控制文件](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/pg_worker_pool.control)

pg_worker_pool 0.1 创建按需启动的具名后台工作进程，执行存储在 worker_pool.jobs 中的 SQL。工作进程会在提交或启动它的事务提交后开始，使用提交者的 PostgreSQL 角色连接提交者的数据库，排空其具名队列，并在没有等待行时退出。

### 预加载与提交

库必须在 postmaster 启动期间分配共享内存。在创建扩展前先将它加入配置，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_worker_pool'
```

```sql
CREATE EXTENSION pg_worker_pool;

CALL worker_pool.submit(
  'indexes',
  'CREATE INDEX my_table_name_idx ON my_table(name)'
);

CALL worker_pool.launch('indexes');

SELECT id, worker_name, query_text, status
FROM worker_pool.jobs
ORDER BY id;
```

编译后的池有八个集群级槽位。max_worker_processes 必须为它们和其他扩展预留容量。改变池上限需要编辑 MAX_WORKERS 并重新构建。

### 注入、身份与恢复风险

C 代码使用百分号-s 格式直接把工作进程名和查询文本插入内部 SQL。它不使用 SPI 参数，也不引用字面量。含普通单引号的作业 SQL 就可能提交失败，精心构造的工作进程名或查询文本则能改变内部 INSERT 或队列选择语句。不得向不可信调用者暴露 submit 或 launch。应撤销 PUBLIC 执行权，限制 worker_pool 表，并只接受应用生成的标识符和预审批语句。

池身份只按工作进程名匹配，不包含数据库或角色。如果在某工作进程活动期间，不同用户或数据库重用同一名称，作业可能滞留，或者以启动现有工作进程的角色执行。查询会在该角色下作为任意 SQL 执行。应分配集群全局唯一名称，避免混用不同信任级别的角色，并将该扩展隔离到一个受控数据库。

出队查询使用 LIMIT 1 但没有 ORDER BY，所以源码并不像 README 宣称的那样保证提交顺序。后端崩溃可能使作业永久停留在 pending，失败只保存 failed 状态，不保存错误详情。它没有重试、取消、租约、所有权、超时或过期作业恢复 API。

只有在修复参数化和身份隔离后才应使用该项目。应测试提交与回滚时的启动行为、并发提交者、角色变更、工作进程耗尽、重启、崩溃恢复、长事务、DDL 锁和幂等重试。锁定仓库没有提供许可证或生产支持政策。
