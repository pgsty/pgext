## 用法

来源：

- [19.0.0 使用说明](https://github.com/pgsty/openhalo/blob/v19.0.0/README.md)
- [2.0 控制文件](https://github.com/pgsty/openhalo/blob/v19.0.0/contrib/aux_mysql/aux_mysql.control)

`aux_mysql` 2.0 随 OpenHalo 19.0.0 提供，支持 MySQL 线协议、语法和函数。运行 `pig install openhalo-19` 安装内核包 `openhalodb-19`。即使当前使用其他 PostgreSQL 大版本，这条命令也会选择 PG19。

### 启用与连接

OpenHalo 19 的新实例默认在 `shared_preload_libraries` 中加载 `aux_mysql`。扩展会创建 ICU 排序规则，数据库编码需支持 ICU，例如 `UTF8`。管理员通过 PostgreSQL 接口连接 `mysql.database` 指定的数据库（默认为 `postgres`），然后创建扩展。

```sql
CREATE EXTENSION aux_mysql;
SELECT extversion FROM pg_extension WHERE extname = 'aux_mysql';
```

在 OpenHalo 19 上查询应返回 `2.0`。启用 `mysql.listener_on` 后，MySQL 客户端可连接 `mysql.port`（默认 3306）。MySQL 数据库映射到该 PostgreSQL 数据库中的模式。扩展在安装库中注册类型转换，需要原生转换行为的工作负载可使用独立数据库。
