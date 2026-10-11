## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/f716914e8a8fb22faf4e932944b6411753b294fa/CHANGELOG.md)
- [extensions/pg_solid/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/f716914e8a8fb22faf4e932944b6411753b294fa/extensions/pg_solid/Cargo.toml)
- [extensions/pg_solid/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/f716914e8a8fb22faf4e932944b6411753b294fa/extensions/pg_solid/src/lib.rs)

- [extensions/pg_solid/pg_solid.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/pg_solid.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_solid/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/Cargo.toml)
- [extensions/pg_solid/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_solid/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/pgbrew.toml)
- [extensions/pg_solid/build.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/build.rs)

`pg_solid` 0.4.0 在 `pgsolid` 中提供实体几何，包括构造、布尔运算、度量、变换、空间索引与 CAD／BIM 导入导出。

### 核心用法

```sql
CREATE EXTENSION pg_solid;
SELECT pgsolid.solid_volume(pgsolid.solid_box(100, 200, 300));
SELECT pgsolid.solid_is_valid(pgsolid.solid_sphere(10));
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。没有预加载声明，需要 OpenCASCADE 7.6+ 原生库。应限制文件／模型导入并控制资源消耗。构造与度量函数使用调用方提供的坐标和单位，导入或计算形体时需注意有效性与容差。这是独立的实体几何实现，不能据此推断依赖 PostGIS 或使用可互换的几何编码。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 0.4.0 的 IFC 朝向

仓库发布 0.6.0 中的扩展 0.4.0 新增 `pgsolid.ifc_true_north(path)`，读取模型真北方向和局部轴旋转。文件没有地图转换信息时，`pgsolid.solid_georeference` 在转换到 EPSG 4326/4978 时应用该方向，修复此前把局部 Y 轴当成北方的 IFC2x3 模型。受影响模型的地理配准结果会改变，应核验派生数据。0.3.0 及之后的版本具有 SQL 升级路径。
