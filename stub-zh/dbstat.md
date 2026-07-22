## 用法

来源：

- [官方 README](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/README.md)
- [官方扩展 SQL](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/sql/dbstat--0.1.sql)
- [官方采集器源码](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/dbstat.c)

`dbstat` 是一个实验性的表行数计数器，由数据库触发器、`NOTIFY` 和独立的 libpq 采集进程组成。它属于学习项目，而不是可靠的统计服务：版本 `0.1` 只观察安装时 `public` 模式中已经存在的表，任何漏掉的通知都会让计数失真。

### 核心流程

创建扩展会安装固定的 `dbstat` 模式，扫描 `public` 中的普通表，保存初始精确计数，并为每张表增加一个 `AFTER INSERT OR DELETE OR UPDATE` 行级触发器：

```sql
CREATE EXTENSION dbstat;

SELECT schema_name, table_name, row_count
FROM dbstat.tb_catalog_table
ORDER BY schema_name, table_name;
```

需要为同一个数据库持续运行另外构建的采集器。它为每个登记表监听一个通道，并在触发器发送操作后更新保存的计数：

```console
dbstat -h localhost -p 5432 -U postgres -d appdb
```

未设置 `dbstat.enable_logging` 时默认记录。会话可以用 `SET dbstat.enable_logging = false` 抑制通知，但相应变更将永久缺失于派生计数中。

### 对象索引

- `dbstat.tb_catalog_table` 记录表 OID、名称和估算的当前行数。
- `dbstat.tb_catalog_table_modification` 是由采集器写入的非日志事件表。
- `dbstat.fn_install_triggers()` 登记表并计算其初始精确计数。
- `dbstat.fn_issue_notify()` 是行级触发器函数。
- `dbstat.fn_get_table_oid()` 和 `dbstat.fn_log_action()` 支持登记与事件记录。

### 运维边界

安装只能由超级用户执行，而且会修改 `public` 中每一张既有普通表；使用前必须审查这些触发器变更。之后创建的表不会自动登记。触发器只发送操作类型，因此采集器停机、事务通知合并、关闭记录、崩溃和非日志表重置都可能导致历史永久丢失，且没有恢复机制。登记函数会用未引用的模式名和表名拼接 SQL，代码也面向较旧的 PostgreSQL API。评估时应使用 `COUNT(*)` 对账、最小权限采集凭据和进程守护；不能把结果当作审计证据或权威基数。
