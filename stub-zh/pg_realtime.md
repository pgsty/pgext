## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/pg_realtime.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/pg_realtime--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_realtime/pg_realtime.c)

`pg_realtime` `1.0` 提供基于触发器的变更事件与通知。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_realtime;
CREATE TABLE realtime_demo(id int PRIMARY KEY);
SELECT pg_realtime_subscribe('realtime_demo');
INSERT INTO realtime_demo VALUES (1);
SELECT * FROM pg_realtime_poll(0);
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

订阅会安装行触发器，pg_realtime_unsubscribe 移除订阅，pg_realtime_cleanup 清理事件。轮询读取事件表，与消费 LISTEN/NOTIFY 不同。序列化器将非空列值表示为 JSON 字符串。应检查事件表授权和负载敏感性并安排清理，通知机制不提供持久的恰好一次投递保证。

此模块对应上游旧名 `alohadb_realtime`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
