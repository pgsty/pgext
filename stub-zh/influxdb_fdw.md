## 用法

来源：

- [项目 README 与兼容性说明](https://github.com/pgspider/influxdb_fdw/blob/9cb845ecee3f6a66c2abd2bcdf6a7a5266e756f7/README.md)
- [扩展 control 文件](https://github.com/pgspider/influxdb_fdw/blob/9cb845ecee3f6a66c2abd2bcdf6a7a5266e756f7/influxdb_fdw.control)
- [1.3 版 SQL API](https://github.com/pgspider/influxdb_fdw/blob/9cb845ecee3f6a66c2abd2bcdf6a7a5266e756f7/influxdb_fdw--1.3.sql)

`influxdb_fdw` 1.3 将 InfluxDB measurement 映射为 PostgreSQL 外部表。审查到的上游兼容矩阵覆盖 PostgreSQL 13 至 17、InfluxDB 1.8，以及通过 v1 兼容 API 访问的 InfluxDB 2.7。它支持读取、插入、删除、模式导入和若干查询下推；不支持 `TRUNCATE`。

### 定义外部服务器和表

下面的服务器连接 InfluxDB 1.x 数据库。远端凭据应放在独立管理且权限最小化的用户映射中，不要写入应用 SQL 或日志。

```sql
CREATE EXTENSION influxdb_fdw;

CREATE SERVER influxdb_svr
  FOREIGN DATA WRAPPER influxdb_fdw
  OPTIONS (dbname 'metrics', host 'http://127.0.0.1', port '8086');

CREATE FOREIGN TABLE sensor_reading (
  time timestamptz,
  device_id text,
  temperature double precision
)
SERVER influxdb_svr
OPTIONS (table 'sensor_reading', tags 'device_id');

SELECT time, device_id, temperature
FROM sensor_reading
WHERE time >= now() - interval '15 minutes';
```

InfluxDB 2.x 需要 C++ 客户端路径、`version '2'` 以及用户映射中的 `auth_token` 选项。应保护包含用户映射的 PostgreSQL 目录和备份。无模式模式下，外部表使用固定的 `time`、`tags` 和 `fields` 列，后两列用 `jsonb` 表示。

### 语义与运行边界

当远端目标列表指定具体字段时，InfluxDB 可能省略字段为 null 的行，而远端 `SELECT *` 可以返回这些行。因此，仅改变 PostgreSQL 投影列也可能改变可见行。依赖下推行为之前，应检查 `EXPLAIN VERBOSE` 并与原生 InfluxDB 查询结果比较。

写入跨越外部系统边界，不具备 PostgreSQL 本地事务的全部保证。重试、部分失败、时间精度、保留策略和远端认证都应由应用明确处理。`time` 列使用微秒精度；需要纳秒文本表示时可使用 `time_text`。审查到的兼容矩阵不能证明 PostgreSQL 17 以后的支持，因此应验证具体的 PostgreSQL、客户端库和 InfluxDB 组合。
