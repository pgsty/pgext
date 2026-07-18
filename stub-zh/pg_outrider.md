## 用法

来源：

- [1.0 版本 SQL 接口](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider--1.0.sql)
- [动态后台工作进程实现](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.c)
- [扩展 control 文件](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.control)

`pg_outrider` 会启动一个动态后台工作进程，监视某个堆关系的空闲空间水位。空闲空间低于指定水位时，工作进程按指定兆字节数扩展关系的 main fork，并更新其空闲空间映射。

```sql
CREATE EXTENSION pg_outrider;

CREATE TABLE event_queue (id bigint, payload jsonb);
SELECT pg_outrider_launch('event_queue'::regclass, 16, 32);
```

第二、第三个参数分别是扩展增量与水位，单位为 MiB，默认值是 1 和 2。调用成功时返回新注册工作进程的进程 ID。该进程为动态注册，因此上游不要求 `shared_preload_libraries`。

### 注意事项

每次调用都会占用一个后台工作进程槽位，并为该关系启动工作进程；应避免重复启动并监控存储增长。工作进程失败后会停止，不会自动重启。已复核的实现来自 2017 年，使用 PostgreSQL 内部存储 API，且没有发布兼容性矩阵或 README，因此部署前必须针对准确的服务端大版本验证。
