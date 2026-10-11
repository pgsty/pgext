## 用法

来源：

- [README.md](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/README.md)
- [src/pg_settimeout--1.1.sql](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/src/pg_settimeout--1.1.sql)
- [src/pg_settimeout.c](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/src/pg_settimeout.c)
- [src/pg_settimeout.control](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/src/pg_settimeout.control)

`pg_settimeout` 启动动态后台工作进程，在延迟后执行 SQL。上游报告在 PostgreSQL 9.6、10 和 11 上测试，更新主版本的兼容性尚未确认。

### 核心用法

```sql
CREATE EXTENSION pg_settimeout;
SELECT pg_settimeout('SELECT 1', 1000);
```

### 运行边界

`pg_settimeout(query, timeout)` 执行一次，`pg_setinterval(query, timeout)` 无限重复。延迟单位为毫秒，返回值为工作进程 PID。工作进程以调用者身份连接当前数据库，在独立事务中执行；调度事务的回滚不会撤销已经独立完成的任务。可用工作进程槽位限制调度数量。模块没有启动初始化逻辑，无需预加载。安装采用控制文件的默认超级用户要求。
