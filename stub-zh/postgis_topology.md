

## 用法

> [PostGIS Topology：PostGIS 的拓扑数据模型支持](https://github.com/postgis/postgis)

PostGIS Topology 为 PostgreSQL 实现了 SQL/MM 拓扑模型。它将空间数据表示为节点、边和面的图结构，确保共享边界只存储一次，并强制保证几何一致性。

- [拓扑参考手册](https://postgis.net/docs/Topology.html)

### 安装

```sql
CREATE EXTENSION postgis_topology;
```

这将创建包含管理表和函数的 `topology` 模式。

--------

## 创建拓扑

拓扑是一个具有指定 SRID 和捕捉容差的命名模式：

```sql
-- 创建名为 'city_topo' 的拓扑，SRID 4326，容差 0.00001 度
SELECT topology.CreateTopology('city_topo', 4326, 0.00001);
```

列出所有拓扑：

```sql
SELECT * FROM topology.topology;
```

--------

## 构建拓扑基元

### 添加边

边是基本构建单元。节点会在边的端点处自动创建。

```sql
-- 添加孤立节点
SELECT topology.ST_AddIsoNode('city_topo', NULL, ST_Point(-73.98, 40.75));

-- 添加两点之间的边（如需要会自动创建节点）
SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.99, 40.74), ST_Point(-73.98, 40.74)));

SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.98, 40.74), ST_Point(-73.98, 40.75)));

SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.98, 40.75), ST_Point(-73.99, 40.75)));

SELECT topology.ST_AddEdgeModFace('city_topo',
    ST_MakeLine(ST_Point(-73.99, 40.75), ST_Point(-73.99, 40.74)));
```

### 查看基元

```sql
-- 列出节点
SELECT node_id, ST_AsText(geom) FROM city_topo.node;

-- 列出边
SELECT edge_id, ST_AsText(geom) FROM city_topo.edge_data;

-- 列出面（face_id 0 是全局面）
SELECT face_id, ST_AsText(mbr) FROM city_topo.face;

-- 获取面的几何形状
SELECT topology.ST_GetFaceGeometry('city_topo', 1);
```

--------

## 拓扑几何（TopoGeometry）

拓扑几何是通过引用拓扑基元而非坐标来定义的空间对象。这确保了共享边界始终保持一致。

### 创建拓扑几何图层

```sql
-- 创建包含拓扑几何列的表
CREATE TABLE city_topo.blocks (
    id serial PRIMARY KEY,
    name text
);

-- 注册拓扑几何列（返回 layer_id）
SELECT topology.AddTopoGeometryColumn('city_topo', 'city_topo', 'blocks', 'topo', 'POLYGON');
```

### 赋值拓扑几何

```sql
-- 从常规几何创建拓扑几何（捕捉到已有拓扑）
INSERT INTO city_topo.blocks (name, topo) VALUES
    ('Block A', topology.toTopoGeom(
        ST_GeomFromText('POLYGON((-73.99 40.74,-73.98 40.74,-73.98 40.75,-73.99 40.75,-73.99 40.74))', 4326),
        'city_topo', 1));

-- 将拓扑几何转换回常规几何
SELECT name, topo::geometry FROM city_topo.blocks;
```

### 验证拓扑

```sql
-- 验证整个拓扑
SELECT * FROM topology.ValidateTopology('city_topo');

-- 检查拓扑错误
SELECT error FROM topology.ValidateTopology('city_topo')
WHERE error IS NOT NULL;
```

--------

## 清理

```sql
-- 删除拓扑及其所有对象
SELECT topology.DropTopology('city_topo');
```
