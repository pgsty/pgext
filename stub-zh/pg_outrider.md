## 用法

来源：

- [1.0 版本 SQL 接口](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider--1.0.sql)
- [动态后台工作进程实现](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.c)
- [扩展 control 文件](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.control)

`pg_outrider` 启动动态后台工作进程，监视某个关系的空闲空间映射；当在文件尾部附近看不到指定空闲空间水位时，直接扩展其主分叉。它是基于 PostgreSQL 内部接口的实验性存储管理代码，不是普通的容量预留 API。

### 核心流程

只有数据库管理员在验证关系与大小后才应启动工作进程：

```sql
CREATE EXTENSION pg_outrider;
REVOKE EXECUTE ON FUNCTION pg_outrider_launch(regclass, bigint, bigint) FROM PUBLIC;

CREATE TABLE event_queue (id bigint, payload jsonb);
SELECT pg_outrider_launch('event_queue'::regclass, 16, 32);
```

`pg_outrider_launch(regclass, bigint, bigint)` 把最后两个参数解释为扩展增量与水位，单位为 MiB；默认值是 1 和 2。它返回注册的工作进程 PID。该工作进程为动态注册，所以不需要 `shared_preload_libraries`，但每次启动都会占用一个后台工作进程槽位。

### 安全与运维边界

SQL 函数只检查调用者对关系的 SELECT 权限，但工作进程会修改物理存储，而扩展函数默认可由 PUBLIC 执行。应撤销访问，并只允许可信管理员启动。C 实现会把 `bigint` 大小收窄成 C 整数，却没有充分验证正数范围；调用前必须拒绝零、负数和过大的值。

扩展没有用于停止、查看状态或检测重复工作进程的 SQL 函数。工作进程失败后不会自动重启，并可能意外扩大存储。空闲空间映射只是启发式信息，不是精确的空闲字节计数。2017 年的代码直接使用 heap、page、fork 与空闲空间映射内部接口，因此应在一次性集群中验证关系类型、崩溃恢复、复制、备份、膨胀及确切大版本兼容性。
