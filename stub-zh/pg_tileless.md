## 用法

来源：

- [0.1 版本 SQL 对象](https://github.com/TilelessMap/pg_tileless/blob/e951bafabb65aaacce9fd0bf1a40eece7e9beb68/pg_tileless--0.1.sql)
- [SQLite 写入实现](https://github.com/TilelessMap/pg_tileless/blob/e951bafabb65aaacce9fd0bf1a40eece7e9beb68/sqlite_writer.c)
- [扩展 control 文件](https://github.com/TilelessMap/pg_tileless/blob/e951bafabb65aaacce9fd0bf1a40eece7e9beb68/pg_tileless.control)

`pg_tileless` 是一个已放弃的原型，用于把 PostGIS 几何打包成 TWKB 数据集并写入 SQLite。它会创建 `tileless` 与 `tmp` 模式，并依赖 `plpgsql`、`postgis` 和 `postgis_sfcgal`。

```sql
CREATE EXTENSION pg_tileless CASCADE;
SELECT tileless.pack_twkb_linestring(
  'public.road', 'geom', 'road_id', '',
  '/srv/export/roads.sqlite', 'roads'
);
```

多边形与线串打包器会构造动态 SQL，创建和删除工作表，简化及切分几何，并调用 `tileless.TWKB_Write2SQLite(...)` 写入服务器本地文件。表、列、前缀、查询和输出路径等参数会被拼接进 SQL 或传给原生文件代码；绝不能向不可信调用者开放这些函数。

已复核 README 没有操作文档，项目也已标记为 abandoned。只能由严格受限角色在可丢弃数据上运行。任何使用前都应验证文件系统权限、路径约束、防 SQL 注入、几何有效性、精度损失、失败清理、SQLite 一致性，以及 PostgreSQL/PostGIS/SFCGAL ABI 兼容性。
