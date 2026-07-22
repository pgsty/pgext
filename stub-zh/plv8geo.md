## 用法

来源：

- [官方 README](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/README.md)
- [扩展 control 文件](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/plv8geo.control)
- [0.0.2 版扩展 SQL](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/plv8geo--0.0.2.sql)
- [D3 hexbin 包装器](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/src/d3_hexbin.sql)

`plv8geo` 把选定的 JavaScript 几何库打包进 PostgreSQL，并公开用于 TopoJSON、等值线、六边形分箱、三角剖分、多边形三角剖分与 Voronoi 操作的 PL/v8 包装器。它同时依赖 `plv8` 与 `postgis`；`0.0.2` 版软件栈较旧，上游测试环境是 PostgreSQL 9.4 与 PL/v8 2.0.3。

### 加载模块并调用函数

创建扩展会把内嵌 JavaScript 保存到 `public.plv8_modules`，并在 `plv8` 模式下创建函数。调用函数前，应初始化加载器并加载该函数所需的各模块。

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION postgis;
CREATE EXTENSION plv8geo;

SELECT plv8.plv8_startup();
DO LANGUAGE plv8 'load_module("d3")';
DO LANGUAGE plv8 'load_module("d3-hexbin")';

SELECT plv8.d3_hexbin(
  '[[1,2],[0.5,0.5],[2,2]]'::json,
  '["foo","bar","baz"]'::json,
  1
);
```

README 建议用 `plv8.startproc = 'plv8.plv8_startup'` 配置会话启动。PL/v8 启动设置在不同版本间有变化，因此应确认已安装 PL/v8 版本实际支持的 GUC；上面的显式加载方式是可审计的后备方案。

### 重要对象

- `public.plv8_modules`：内嵌 JavaScript 源码与加载标志。
- `plv8.plv8_startup()`：定义当前会话的 `load_module` 辅助函数。
- `plv8.d3_totopojson`、`plv8.d3_simplifytopology`、`plv8.d3_mergetopology`、`plv8.d3_topologytofeatures`：TopoJSON 转换与操作。
- `plv8.d3_contour`、`plv8.d3_slopecontours`、`plv8.d3_hexbin`：栅格等值线、坡度等值线与点六边形分箱。
- `plv8.delaunator`、`plv8.earcut`、`plv8.jsts_voronoi`：三角剖分与 Voronoi 包装器。

### 兼容性与安全

软件包内嵌的库版本已经较旧，包括 D3 4.x、TopoJSON 3.0.0、GeoTIFF 0.4.1、Delaunator 1.0.2、Earcut 2.1.1 和 JSTS 1.4.0。不要假定它具备当前 JavaScript API、错误修复或安全属性。应使用实际 PostgreSQL、PostGIS、PL/v8 与 V8 构建验证兼容性，并将结果和 SRID 与原生 PostGIS 等价方案对比。

JavaScript 在数据库后端内部求值，可能消耗大量内存和 CPU。应限制 CREATE 与 EXECUTE 权限、审计内嵌源码，并限制不可信几何与 JSON 的大小。即使该软件栈依赖会话全局加载模块，多个包装器仍声明为 IMMUTABLE；用于表达式索引或生成列前必须验证确定性。control 文件指定模式 `plv8geo`，但安装 SQL 显式创建并使用 `plv8` 以及 public 模块表，因此不能依赖模式局部性假设。
