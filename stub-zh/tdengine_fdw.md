## 用法

来源：

- [Official README](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/README.md)
- [Extension control file](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/tdengine_fdw.control)
- [Installation SQL](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/tdengine_fdw--1.0.sql)
- [Build configuration](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/Makefile)

`tdengine_fdw` 通过 WebSocket 连接访问 TDengine 时序数据。已核验的实现支持 PostgreSQL 15/16，上游要求 TDengine 3.4+ 及兼容的 C 客户端。

### 基本用法

应使用逐用户映射保存凭据，并使外部表列与远端表一致。

```sql
CREATE EXTENSION tdengine_fdw;
CREATE SERVER tdengine_svr FOREIGN DATA WRAPPER tdengine_fdw
  OPTIONS (host 'localhost', port '6041', dbname 'mydb');
CREATE USER MAPPING FOR CURRENT_USER SERVER tdengine_svr
  OPTIONS (user 'app_user', password 'replace-with-password');
CREATE FOREIGN TABLE sensor_data (ts timestamptz, temperature float4)
  SERVER tdengine_svr OPTIONS (table 'sensor_data');
SELECT * FROM sensor_data WHERE temperature > 25;
```

### 能力与限制

支持读取、谓词下推、插入和 `IMPORT FOREIGN SCHEMA`。表选项包括 `table`、`tags` 和 `schemaless`，后者通过 JSONB 处理标签与字段数据。客户端依赖为 `libtaos >= 3.3.6.0`，文档中的连接端口为 `6041`。DELETE、聚合下推及其他 PostgreSQL 大版本支持仍列为未完成功能，不能据此推断支持。外部写入需要独立规划持久性与回滚。
