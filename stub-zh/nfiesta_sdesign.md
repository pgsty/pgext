## 用法

来源：

- [官方仓库 README](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/README.md)
- [官方扩展控制文件](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/nfiesta_sdesign.control)
- [官方扩展 SQL](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/nfiesta_sdesign--1.0.0.sql)
- [官方导入函数](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/functions/fn_import_data.sql)

`nfiesta_sdesign` 安装 NFIesta 森林清查系统使用的抽样设计数据模型。它在固定的 `sdesign` 模式中描述国家、分层、面板、样地簇、样地、清查活动、参考年集合、抽样权重和 PostGIS 几何；它属于领域应用模式，而不是通用统计函数库。

### 安装与检查

版本 `1.1.17` 依赖 `plpgsql` 和 `postgis`，且不可重定位：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION nfiesta_sdesign;

SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'sdesign'
ORDER BY table_name;
```

主要表包括 `sdesign.c_country`、`sdesign.t_strata_set`、`sdesign.t_stratum`、`sdesign.t_panel`、`sdesign.t_cluster_configuration`、`sdesign.t_cluster`、`sdesign.f_p_plot`、`sdesign.t_inventory_campaign` 和 `sdesign.t_reference_year_set`，另有映射表连接面板、样地簇、清查活动和参考年。

### 导入与校验

`sdesign.fn_import_data(text)` 从调用方指定的暂存模式导入完整数据集。该模式必须包含函数所期待的全部暂存关系，包括国家、样地簇配置、分层集合、分层、面板、样地簇、样地、清查活动、测量日期和参考年映射。

校验触发器会强制检查清查活动起止配对、面板/参考年映射、样地测量日期、样地簇配置组合和抽样权重总和等关系。`sdesign.fn_create_buffered_geometry()` 为支持的样带布局派生 MultiPolygon 抽样框几何。

### 运维边界

应把安装、导入、升级和删除都当成应用模式迁移：它们会创建并修改大量持久表、序列、约束、触发器和 PostGIS 值。导入函数使用调用方提供的模式字符串拼接动态 SQL，但没有引用标识符，因此只允许预先验证的可信模式名，并以最小权限所有者运行。导入不是幂等行更新，且可能到延迟校验阶段才失败；应在事务中执行，提前验证暂存计数和 SRID，并在副本上测试回滚。版本 `1.1.17` 为面板/参考年映射增加唯一约束，升级前必须清理重复项。
