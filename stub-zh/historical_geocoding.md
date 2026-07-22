## 用法

来源：

- [已复核 commit 的上游工作流文档](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/doc/overall.md)
- [声明与实际依赖脚本](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/creating_dependency_extensions.sql)
- [地理编码 API 源码](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/src/050_geocoding_api.sql)
- [扩展 control 文件](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/historical_geocoding.control)

`historical_geocoding` 是面向研究的 SQL 模式，把地址和模糊日期与不完整历史地址数据匹配。它按语义、时间、门牌号、尺度、空间距离和来源精度对候选地点排序；安装后并不包含可直接查询的地名库。

### 安装与工作流

control 文件声明 `geohistorical_objects`、`postal` 和 `unaccent`。上游依赖脚本还安装 `postgis`、`pgsfti` 和 `pg_trgm`，生成的 SQL 会使用这些扩展，却没有在 `requires` 中声明。创建本扩展前必须准备并验证全部六项依赖。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pgsfti;
CREATE EXTENSION pg_trgm;
CREATE EXTENSION geohistorical_objects;
CREATE EXTENSION postal;
CREATE EXTENSION unaccent;
CREATE EXTENSION historical_geocoding;

SELECT *
FROM historical_geocoding.geocode_name_optimised(
  '10 rue du temple, Paris',
  sfti_makesfti('1872-11-15'::date)
)
LIMIT 10;
```

主要 API 包括 `geocode_name_base`、`geocode_name_optimised_inter` 和 `geocode_name_optimised`。真实工作流必须规范化地址与日期，导入并验证道路和门牌来源，填充精确与粗略定位表，保留来源血缘，并用人工确认样本评估排序候选。

### 安全与安装风险

`ordering_priority_function` 经过一个删除部分字符和关键字的黑名单函数后，被直接插入动态 SQL。这不是安全的 SQL 注入防线。绝不能传入用户控制文本；应通过包装函数只暴露固定且已审计的评分表达式，并撤销不可信角色直接执行 API 的权限。

PGXS 构建会拼接源码文件，而其中包含未加保护的示例 `SELECT` 和测试 `CREATE TABLE` 语句。因此必须审查生成的 `historical_geocoding--1.0.sql`，并只在可丢弃数据库中进行首次安装；不能假定创建扩展没有额外副作用。

已复核项目聚焦法国及巴黎历史数据，且包含部分人工研究流程。结果具有概率性：应使用黄金样本评估偏差、缺失数据、别名、时间不确定性、SRID 2154 假设、libpostal 行为、可复现性和人工修正，并保留来源与置信度，不能覆盖原始事实。
