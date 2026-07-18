## 用法

来源：

- [上游 README](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/README.md)
- [扩展 control 文件](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/cassandra2_fdw.control)
- [SQL 安装脚本](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/cassandra2_fdw--2.0.sql)
- [C FDW 实现](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/cassandra2_fdw.c)

`cassandra2_fdw` `2.0` 版是使用 DataStax C/C++ 驱动的只读 Cassandra 2.x–4.x 外部数据包装器。服务器对象指定 Cassandra 主机、端口和协议，用户映射提供凭据，外部表把 PostgreSQL 行结构映射到 Cassandra keyspace 和表。

### 示例

```sql
CREATE EXTENSION cassandra2_fdw;
CREATE SERVER cass_serv FOREIGN DATA WRAPPER cassandra2_fdw
  OPTIONS (host '127.0.0.1,127.0.0.2', port '9042', protocol '4');
CREATE USER MAPPING FOR CURRENT_USER SERVER cass_serv
  OPTIONS (username 'reader', password 'secret');
CREATE FOREIGN TABLE cass_orders (id integer, description text)
  SERVER cass_serv OPTIONS (schema_name 'example', table_name 'orders');
SELECT * FROM cass_orders LIMIT 5;
```

该包装器不能写入 Cassandra，只支持有限的标量类型映射；集合、map 和用户定义类型尚未实现。上游报告的测试范围是 PostgreSQL 11–15 和 Cassandra 3/4，不包括当前 PostgreSQL 版本。凭据保存在 PostgreSQL 目录中。应使用最小权限 Cassandra 账户，保护目录可见性和网络流量，并在部署前验证协议/驱动兼容性及查询行为。
