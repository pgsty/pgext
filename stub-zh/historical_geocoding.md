## 用法

来源：

- [已复核 commit 的上游工作流文档](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/doc/overall.md)
- [上游 README](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/README.md)
- [扩展 control 文件](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/historical_geocoding.control)

`historical_geocoding` 是面向研究的 SQL 模式，用于把文本地址和日期与不完整历史地址数据匹配，并按置信度排序可能地点。其 control 文件要求 `geohistorical_objects`、`postal` 和 `unaccent`。

```sql
CREATE EXTENSION historical_geocoding CASCADE;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname IN (
  'historical_geocoding', 'geohistorical_objects', 'postal', 'unaccent'
);
```

安装只提供表和函数模板，并不包含可直接使用的历史地名库。文档工作流要求规范化地址/日期输入，导入并验证从地图提取的道路和门牌数据，按文本、时间和空间匹配，再排序候选。结果具有概率性，应保留来源和置信度，不能覆盖原始事实。

已复核项目聚焦法国及巴黎历史数据，并描述了部分人工研究流程。必须审计所有函数和依赖扩展，尤其是动态 SQL、继承、模糊时间模型、PostGIS 操作与 libpostal 行为。任何操作使用前，应通过黄金样本评估偏差、缺失数据、别名、时间不确定性、可复现性和人工修正。
