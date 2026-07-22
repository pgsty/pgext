## 用法

来源：

- [官方 README](https://github.com/nyurik/postile/blob/a11e6479aa435fc17379a7c23ad1f9501c43aca8/README.md)
- [SQL 函数实现](https://github.com/nyurik/postile/blob/a11e6479aa435fc17379a7c23ad1f9501c43aca8/src/pg_funcs.rs)
- [扩展控制文件](https://github.com/nyurik/postile/blob/a11e6479aa435fc17379a7c23ad1f9501c43aca8/postile.control)

`postile` 是一个 pgrx 扩展，用于在 PostgreSQL 内生成地图图层瓦片字节并压缩二进制载荷。0.0.1 版提供精简且仍在演进的 SQL 接口；应把当前源代码，而不是宽泛的项目描述，作为 API 契约。

### 核心流程

在每个需要生成瓦片的数据库中安装扩展。对于坐标已经处于瓦片空间的表，可将其行编码为一个 MapLibre Tile 图层：

```sql
CREATE EXTENSION postile;

CREATE TABLE tile_features (
  id bigint PRIMARY KEY,
  geom point NOT NULL,
  name text,
  rank integer
);

INSERT INTO tile_features VALUES
  (1, point(10, 20), 'one', 7),
  (2, point(30, 40), 'two', 9);

SELECT PT_AsMLT(
  'tile_features'::regclass,
  'places',
  4096,
  'geom',
  'id'
);
```

`PT_AsMLT` 会读取指定关系中几何值非空的所有行。其参数依次选择关系、图层名称、范围、几何列以及可选的整数要素 ID 列。标量布尔、整数、浮点、text、varchar 与 character 列会成为要素属性；不支持的属性类型会被跳过。

### 压缩辅助函数

扩展还提供不可变、可并行安全执行并返回 `bytea` 的压缩辅助函数：

```sql
SELECT pt_gzip('Hello, world!'::bytea);
SELECT pt_gzip('Hello, world!'::bytea, 9);
SELECT pt_brotli('Hello, world!'::bytea);
SELECT pt_version();
```

- `pt_gzip` 接受 0 到 9 的可选压缩级别；省略时使用库默认值。
- `pt_brotli` 使用实现中固定的 Brotli 参数。
- `pt_version` 返回编译时的扩展版本。

### 几何与运维边界

原生 PostgreSQL `point` 值可直接使用。其他几何列通过 `ST_AsEWKB` 读取，因此环境必须提供该函数，通常由 PostGIS 提供。支持 point、linestring、polygon、multipoint、multilinestring 与 multipolygon EWKB；geometry collection 会被拒绝。坐标会舍入为有符号 32 位瓦片坐标，解析后忽略 Z/M 纵坐标。

`PT_AsMLT` 针对所给关系构造查询，并将结果物化为一个编码图层。调用前应把源行筛选、转换到有界关系或视图，不要直接指向无界基础表。当前实现不会代为完成 Web Mercator 投影、瓦片裁剪或瓦片包络过滤。

控制文件将扩展标记为不可重定位、仅超级用户可安装且不受信任，并且没有预加载设置。即使非原生几何依赖 `ST_AsEWKB`，它也没有声明 PostGIS 扩展依赖。应使用目标瓦片消费者测试生成字节，并固定上游修订，因为 0.0.1 API 仍在变化。
