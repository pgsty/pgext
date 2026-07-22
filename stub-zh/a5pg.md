## 用法

来源：

- [官方 README](https://github.com/decision-labs/a5pg/blob/2555ce912ae6d135b831b21385b6fc2d4bcff989/README.md)
- [官方控制文件](https://github.com/decision-labs/a5pg/blob/2555ce912ae6d135b831b21385b6fc2d4bcff989/a5pg.control)
- [官方构建清单](https://github.com/decision-labs/a5pg/blob/2555ce912ae6d135b831b21385b6fc2d4bcff989/Cargo.toml)

`a5pg` 在 PostgreSQL 中提供 A5 等面积离散全球网格函数。它可将经纬度坐标转换成分层的 64 位单元标识，用于聚合与制图；它不是 PostGIS 索引访问方法，而且可选的几何包装函数只有在兼容的 `geometry` 类型存在时才可用。

### 核心流程

```sql
CREATE EXTENSION a5pg;

SELECT a5_lonlat_to_cell(-73.9857, 40.7580, 10) AS cell_id;

SELECT a5_cell_to_lonlat(2742822465196523520) AS center,
       a5_get_resolution(2742822465196523520) AS resolution,
       a5_cell_to_parent(2742822465196523520, 8) AS parent;
```

按分析尺度选择分辨率，把返回的 `bigint` 单元 ID 与源数据一同持久化，再按该值分组或连接。分辨率转换必须明确验证：父单元更粗，而 `a5_cell_to_children` 会展开成更细的单元。

### 重要对象

- `a5_lonlat_to_cell` 把经度、纬度与分辨率转换为单元 ID。
- `a5_cell_to_lonlat` 与 `a5_cell_to_boundary` 返回中心点及多边形坐标数组。
- `a5_get_resolution`、`a5_cell_to_parent` 与 `a5_cell_to_children` 用于浏览层级。
- 已定义 `geometry` 时，`a5_point_to_cell` 可作为 PostGIS 便捷包装函数。
- `a5pg_version` 与 `a5pg_info` 返回扩展信息。

### 限制

控制文件将版本 `0.6.1` 标记为不可迁移、不可信且只能由超级用户安装。构建清单声明 PostgreSQL 13 至 18 的特性，而本次核验的 README 只报告 15 至 17 的 CI；请测试实际打包的服务器目标。

经纬度顺序、坐标参考系、单元边界处理与所选分辨率都属于应用语义。输入应规范化为 WGS84 角度并拒绝无效坐标，同时回归测试单元边缘附近的连接。A5 单元 ID 是实现标识，并不表示距离，也不是全局可排序的地理坐标。
