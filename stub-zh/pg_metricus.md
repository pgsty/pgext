## 用法

来源：

- [官方 README](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/README.md)
- [1.0 扩展 SQL](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/pg_metricus--1.0.sql)
- [UDP 发送实现](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/pg_metricus.c)

`pg_metricus` 从 SQL 或 PL/pgSQL 向 Brubeck、Graphite 等 IPv4 UDP 端点发送由应用格式化的指标。它只提供传输能力，调用者必须自行构造接收器所需的精确行协议。

### 核心流程

配置目标地址与端口，重新加载 PostgreSQL 配置，再将扩展安装到选定模式。

```conf
pg_metricus.host = '10.9.5.164'
pg_metricus.port = 8124
```

```sql
CREATE SCHEMA metricus;
CREATE EXTENSION pg_metricus SCHEMA metricus;

SELECT metricus.send_metric(
  format(E'%s.%s:%s|%s\n', 'db.sql', 'latency', 12, 'ms')
);
```

目标端点设置可通过重新加载配置来修改；该扩展不需要预加载。

### SQL 对象与交付语义

- `send_metric(text) RETURNS void`：把提供的字节作为一个 UDP 数据报发送。

1.0 版本的 C 实现对 `AF_INET` 使用 `inet_pton`，因此只接受数值形式的 IPv4 地址。UDP 没有确认与重试，函数也不会让交付具备事务性：即使外围数据库事务之后回滚，已经发送的数据报也不会撤回。不要在指标文本中放入秘密；函数成功返回只表示本地发送尝试，不代表收集器已经保存指标。
