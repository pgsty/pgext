## 用法

来源：

- [0.1.3 版本指南](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/README.md)
- [发布版本与包内安装 SQL](https://github.com/mapseekai/geometry_pyramid/releases/tag/v0.1.3)
- [控制文件](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/geometry_pyramid.control)
- [目录定义](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/sql/bootstrap.sql)
- [生命周期函数](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/pyramid.rs)
- [瓦片访问](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/tile.rs)
- [存储检查](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/storage.rs)
- [工作进程执行上下文](https://github.com/mapseekai/geometry_pyramid/blob/v0.1.3/src/worker.rs)

`geometry_pyramid` 在 PostGIS 几何表上构建物化矢量瓦片，并通过 SQL 返回 MVT 字节。0.1.3 版本支持 PostgreSQL 15–18，要求 PostGIS 3.1 或更新版本具备 MVT/protobuf 支持。扩展安装到 `public`，不能迁移到其他模式。它使用原生库，无需共享预加载或外部守护进程。

### 构建与读取瓦片

```sql
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION geometry_pyramid;

CREATE TABLE gp_points (
    id integer PRIMARY KEY,
    name text,
    geom geometry(Point, 4326)
);
INSERT INTO gp_points VALUES
    (1, 'Central', ST_SetSRID(ST_MakePoint(116.4074, 39.9042), 4326));
CREATE INDEX ON gp_points USING gist (geom);

SELECT ST_BuildPyramid(
    'gp_points', 'geom', 'id',
    '{"name":"gp_points_pyr","minZoom":0,"maxZoom":5,"attributes":["name"],"parallel":1}'
);
SELECT octet_length(ST_Tile('gp_points_pyr', 0, 0, 0));
SELECT ST_Pyramid('gp_points_pyr');
SELECT * FROM ST_Pyramids();
```

表名、列名和金字塔名称必须是不超过 45 个字符的 ASCII 裸标识符，不接受带模式限定的名称。使用稳定的要素 ID，并为几何列建立 GiST 索引。输出瓦片采用 Web Mercator；源坐标系不是默认的 4326 时，应设置 `sourceSRS`。

`ST_Tile` 返回 `bytea`；瓦片为空、缺失或坐标越界时返回 `NULL`。`splitSize` 大于零时，稀疏瓦片按需生成。`dynamicBeyondMaxZoom` 控制超过物化最大缩放级别后的动态切片。`attributes` 和 `buildRules` 控制瓦片属性与各缩放级别的过滤。`parallel` 指定构建并发度；工作进程需要可用的 PostgreSQL 工作进程槽位，并要求源数据已经提交。`tileSizeBudgetBytes` 不能与并行构建同时使用。应先限定缩放范围：默认最大缩放级别 18 可能带来大量计算。

### 刷新与删除

```sql
SELECT ST_UpdatePyramid('gp_points_pyr', 'id = 1');
SELECT ST_UpdatePyramidByExtent(
    'gp_points_pyr', 116.35, 39.85, 116.45, 39.95, 4326
);
SELECT ST_DropPyramid('gp_points_pyr');
```

刷新需要显式调用，不由触发器自动执行。要素移动或删除后，应刷新覆盖新旧位置的范围；条件谓词无法重新找到已删除的要素。重建已有金字塔会保留已发布的存储关系。删除源表前，应先删除金字塔。

### 权限与升级

函数以调用者权限运行。构建需要读取源表、写入输出模式，并具有访问 `geometry_pyramid_catalog`、`geometry_pyramid_stats` 和瓦片表所需的权限。安装 SQL 不会将这些表权限授予应用角色。缓存瓦片是已存储的数据：瓦片表的读取权限允许访问物化内容，不会随源行策略的后续变化自动调整。应据此授权并刷新过期数据。当有效的 RLS 策略过滤瓦片存储中的行时，替换操作会被拒绝。

`buildRules[].filter` 以及传给 `ST_UpdatePyramid` 的谓词都是以调用者身份求值的可信 SQL 表达式，不能直接传入不可信输入。工作进程保留有效角色、搜索路径和行安全上下文。

升级前应停止扩展操作，安装匹配的库并重新连接会话。Windows 上替换已加载的 DLL 前，必须先停止 PostgreSQL。使用此版本前，需要将 PostgreSQL 13–14 升级到受支持版本。

```sql
ALTER EXTENSION geometry_pyramid UPDATE TO '0.1.3';
SELECT geometry_pyramid_version();
```

从 0.1.1 或 0.1.2 升级无需重建瓦片。从 0.1.0 升级时，应按官方迁移示例，在同一事务内执行更新并显式填写 `source_rel`/`storage_rel` 绑定；提交前确认所有已注册金字塔都已绑定关系。这些绑定使用 `regclass`，防止后续搜索路径替换对象。
