## 用法

来源：

- [固定提交的上游 README](https://github.com/labgeo/siguanet_quest/blob/b8cba3984e6d32ad4760efe15dd7b60144a45342/README.md)
- [1.0 版应用 SQL](https://github.com/labgeo/siguanet_quest/blob/b8cba3984e6d32ad4760efe15dd7b60144a45342/siguanet_quest--1.0.sql)
- [固定提交的扩展控制文件](https://github.com/labgeo/siguanet_quest/blob/b8cba3984e6d32ad4760efe15dd7b60144a45342/siguanet_quest.control)
- [SIGUANET 数据库模式项目](https://github.com/labgeo/siguanet-dbsetup)

siguanet_quest 1.0 版是面向 SIGUANET 高校建筑资产地理数据库的专用应用报表层。它在固定的 quest 模式中安装数百个以西班牙语命名的 PL/pgSQL 函数、视图和复合类型，用来按空间层级聚合楼面面积、房间、人员、建筑、楼层、部门和活动组。

### 安装与代表性查询

它要求数据库已经按配套模式项目填充 SIGUANET 表和视图；历史 README 还要求 PostgreSQL 9.2+ 与 PostGIS 2.0+：

```sql
CREATE EXTENSION siguanet_quest;

SELECT quest_edificio_superficie('0037');

SELECT *
FROM quest_edificio_obtenerplantas('0037')
ORDER BY indice;
```

只有建筑 0037 和预期的 SIGUANET 对象确实存在时，这些示例才有意义。管理员还应在 quest_adminroles 中加入获准查看未加密保护数据的角色，例如西班牙 NIF 标识符。

### 应用与隐私边界

它不是通用 PostGIS 分析库。SQL 假设确切的 public 模式表、列约定、活动代码、西班牙机构语义和配套视图。安装会创建大量动态 SQL、视图、类型、注释和权限敏感的报表路径接口。在不同模式或仅部分兼容的数据库上安装，可能中途失败或返回误导结果。

仓库自 2013 年后没有变化，并被归类为已废弃。文档中的 PostgreSQL/PostGIS 最低版本是历史信息，不是对当前主版本的支持证据。使用前应在隔离环境恢复代表性 SIGUANET 数据库，盘点每个创建对象与依赖，测试每个暴露函数，审查动态 SQL 和受保护标识符路径，并在适当位置用当前角色与行安全控制替换旧的应用级管理表。
