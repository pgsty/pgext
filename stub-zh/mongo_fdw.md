## 用法

来源：

- [Official source](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_6_0/README.md)
- [Official source](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_6_0/mongo_fdw.control)
- [Official source](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_6_0/mongo_fdw--1.1.sql)
- [Official source](https://github.com/EnterpriseDB/mongo_fdw/releases/tag/REL-5_6_0)

`mongo_fdw` 5.6.0 在 PostgreSQL 14–19 中读写 MongoDB 集合，SQL/control 版本仍为 1.1。安装匹配的 MongoDB 客户端库后，以超级用户创建扩展，无需共享预加载。

### 映射集合

对于已有集合，配置拥有预期远端权限的账户：

```sql
CREATE EXTENSION mongo_fdw;
CREATE SERVER mongo_server FOREIGN DATA WRAPPER mongo_fdw
  OPTIONS (address '127.0.0.1', port '27017');
CREATE USER MAPPING FOR CURRENT_USER SERVER mongo_server
  OPTIONS (username 'mongo_user', password 'replace_with_remote_password');
CREATE FOREIGN TABLE warehouse (
  _id name, warehouse_id integer, warehouse_name text
) SERVER mongo_server OPTIONS (database 'mydb', collection 'warehouse');
SELECT warehouse_id, warehouse_name FROM warehouse;
```

外部表第一列必须是 `name` 类型的 `_id`。让 MongoDB 生成有效标识，不要编造过短的 ObjectId 字符串。列类型必须与 BSON 表示匹配。使用 `INSERT`、`UPDATE`、`DELETE` 前确认远端权限和标识映射；远端写入不能替代分布式事务协议。

### 连接与计划

服务器选项包括 `authentication_database`、`replica_set`、`read_preference` 和 TLS 证书设置。限制用户映射凭据的访问，并验证远端证书。外部表通过 `database`、`collection` 选择数据。该 FDW 不实现 `IMPORT FOREIGN SCHEMA`。

连接、聚合、排序和行数限制下推由服务器、表选项及会话 GUC 控制。`mongo_fdw.log_remote_query` 有助于诊断下推查询。BSON 文本要求 UTF8；嵌套字段或含大写字母的字段需要引号。不支持以美元符号开头或在点号后包含美元符号的字段名。

### 变化与升级检查

5.6.0 修复 URI 选项注入及过短标识的越界读取。游标失败现在会报错，不再静默返回空结果。BSON 浮点数转换为整数时四舍五入到最近整数，非有限值和越界输入会报错。应测试曾依赖截断行为或隐藏错误的应用。

聚合下推可能绕过上述整数输入检查，使远端聚合保留非有限值。需要本地验证时，将 `mongo_fdw.enable_aggregate_pushdown` 设为 `false`。安装新二进制并重连会话，SQL 版本保持 1.1。升级后用有代表性的远端数据核对行数、标识与执行计划。
