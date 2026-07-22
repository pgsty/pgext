## 用法

来源：

- [官方仓库](https://github.com/torrespri/pg_popyramids_datamarts)
- [扩展控制文件](https://github.com/torrespri/pg_popyramids_datamarts/blob/master/pg_popyramids_datamarts.control)
- [1.0.0 版扩展 SQL](https://github.com/torrespri/pg_popyramids_datamarts/blob/master/sql/pg_popyramids_datamarts--1.0.0.sql)
- [上游项目文档](https://github.com/torrespri/pg_popyramids_datamarts/blob/master/doc/pg_popyramids_datamarts.md)

pg_popyramids_datamarts 安装 popyramids 应用的数据集市层，包括人口金字塔类型与转换、PostGIS 几何辅助函数、生成 GeoJSON 的函数和已填充物化视图。它属于应用模式，而不是可独立使用的通用扩展。

### 必需的应用环境

控制文件要求 PostGIS，并把扩展标记为不可重定位，但 1.0.0 版 SQL 还有未编码在控制文件中的前提。安装前，数据库必须已有目标模式、包含金字塔类型与源表的操作数据存储模式，以及授权语句引用的硬编码角色。

```sql
CREATE EXTENSION postgis;
CREATE SCHEMA dms;

-- The popyramids ODS deployment must already provide ods.pyrint and ods.main.
-- Provision the roles referenced by the upstream grants before continuing.
CREATE EXTENSION pg_popyramids_datamarts;
```

不要在通用数据库中执行最后一条语句：创建扩展时会根据现有操作数据存储填充物化视图；任何应用专用类型、关系、列、角色或模式缺失都会导致失败。

### 主要对象族

- 人口统计类型表示年龄分组、男女计数或百分比、变量和分类后的金字塔形状。
- 数组与金字塔辅助函数用于汇总年龄段、计算总数与百分比、转换操作数据存储值，并生成 PostGIS 几何。
- 查询辅助函数用于选择人口金字塔标识符，并向地图客户端返回类似 GeoJSON 的要素集合。
- 物化视图提供原始、五年、十年和大分组版本、百分比版本、统计信息及用户界面目录。

### 刷新与所有权边界

这些物化视图创建时会填充数据，并依赖既有源表及彼此。扩展没有提供自动刷新计划；源数据变化后应按依赖顺序刷新，并在生产规模数据上评估锁、耗时和存储。

1.0.0 版嵌入了应用专用的授权与对象名，其中包括 PostgreSQL 角色。安装前应对照目标 popyramids 部署审查完整 SQL，不能把控制文件的依赖列表当作充分条件。扩展拥有大量带模式限定的类型、类型转换、函数和已填充视图，因此应在恢复副本上测试安装与卸载。

扩展没有声明预加载或重启要求。数据库级安装仍需要创建全部应用对象并向所引用角色授权的权限。
