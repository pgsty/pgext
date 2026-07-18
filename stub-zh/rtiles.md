## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/gportenko/pg_rtiles/blob/b92acdbf435f8cec12df9d2ab6223a31662e417c/README.md)
- [Rust 实现](https://github.com/gportenko/pg_rtiles/tree/b92acdbf435f8cec12df9d2ab6223a31662e417c/pg_rtiles/src)
- [Cargo 清单](https://github.com/gportenko/pg_rtiles/blob/b92acdbf435f8cec12df9d2ab6223a31662e417c/pg_rtiles/Cargo.toml)

`rtiles` 是基于 pgrx/PostGIS 的扩展，为自托管交互地图生成 Mapbox Vector Tile。图层定义位于 `rtiles.layers`，指定源表、几何列、选择字段、可选聚类，以及允许的数据库角色。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION rtiles;
INSERT INTO rtiles.layers
  (name, source, geom_field, fields, clustered, roles)
VALUES
  ('roads', 'public.road', 'geom', ARRAY['class', 'name'], false,
   ARRAY['map_reader']);
```

配套 `tileserver` 通过 `DB_HOST`、`DB_PORT`、`DB_DBNAME`、`DB_USER` 和 `DB_PASSWORD` 连接。凭据应保存在秘密存储中，使用最小权限登录，强制 TLS 与网络控制，并验证数据库角色检查不会被服务器连接角色绕过。

图层元数据驱动动态表列访问。应限制 `rtiles.layers` 写权限，引用并验证所有源标识符，只暴露精心选择的字段，避免泄漏私有列或执行昂贵表达式。应基准 tile 大小、空间过滤、聚类、无效几何、大缩放范围、取消、并发请求和语句超时。已复核 README 只记录 PostgreSQL 16 测试；升级前必须验证每组 pgrx 与 PostGIS ABI。
