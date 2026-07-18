## 用法

来源：

- [扩展 README](https://github.com/kingluo/pgcat-pgxs/blob/bd8d5975680f6870965240dc9cf1218229d9c3a9/README.md)
- [扩展 SQL](https://github.com/kingluo/pgcat-pgxs/blob/bd8d5975680f6870965240dc9cf1218229d9c3a9/pgcat--1.0.sql)
- [扩展 control 文件](https://github.com/kingluo/pgcat-pgxs/blob/bd8d5975680f6870965240dc9cf1218229d9c3a9/pgcat.control)
- [PGXN 用法文档](https://pgxn.org/dist/pgcat/README.html)

`pgcat` 是外部 pgcat 逻辑复制守护进程的数据库端组件。扩展在固定的 `pgcat` 模式中创建配置和进度表及特权辅助函数，独立守护进程读取这些表并应用变更。

创建扩展前必须已有 `pgcat` 角色。最小单向配置是在发布端创建复制角色与发布，然后在订阅端创建对应角色、扩展、目标表权限和订阅记录：

```sql
-- Publisher
CREATE ROLE pgcat LOGIN REPLICATION PASSWORD 'replace-me';
CREATE SCHEMA pgcat AUTHORIZATION pgcat;
CREATE EXTENSION pgcat;
CREATE PUBLICATION app_pub FOR TABLE public.orders;
GRANT SELECT ON public.orders TO pgcat;

-- Subscriber
CREATE ROLE pgcat LOGIN PASSWORD 'replace-me';
CREATE SCHEMA pgcat AUTHORIZATION pgcat;
CREATE EXTENSION pgcat;
GRANT SELECT, INSERT, UPDATE, DELETE, TRUNCATE ON public.orders TO pgcat;

INSERT INTO pgcat.pgcat_subscription
  (name, hostname, port, username, password, dbname, publications, copy_data, enabled)
VALUES
  ('orders_from_primary', 'primary.example.net', 5432, 'pgcat', 'replace-me',
   'app', ARRAY['app_pub'], true, true);
```

数据库对象准备完毕后，再启动并配置外部 `pgcat` 进程。

### 版本与安全

扩展 control 公开的版本为 `1.0`，PGXN 发行包则标记为 `1.0.0`。应按上游说明在发布端和订阅端都安装扩展。连接密码以明文 `text` 存放在 `pgcat.pgcat_subscription` 中，因此必须严格限制固定模式、`pgcat` 角色及守护进程配置的访问。逻辑复制还要求正确准备 `wal_level`、`pg_hba.conf`、发布、复制标识和目标表。
