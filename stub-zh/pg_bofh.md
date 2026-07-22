## 用法

来源：

- [官方 README](https://github.com/rdunklau/pg_bofh/blob/master/README.md)
- [扩展 control 文件](https://github.com/rdunklau/pg_bofh/blob/master/pg_bofh.control)
- [规划器钩子实现](https://github.com/rdunklau/pg_bofh/blob/master/pg_bofh.c)

`pg_bofh` 是一个无 SQL 对象的规划器钩子示例，尝试拒绝缺少有效条件的语句。它不安装 SQL 对象，也没有版本化的扩展脚本。只能在一次性测试服务器中加载该共享库；它不是生产级查询安全或授权系统。

### 核心流程

构建并安装该库，把它加入服务器配置，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_bofh'
```

重启后，应先测试有代表性的读、写、维护和应用语句，再考虑在任何其他环境启用。它没有 SQL 函数、GUC 绕过开关、角色允许列表、数据库允许列表或状态视图。

### 实际强制边界

README 把规则概括为禁止没有 WHERE 子句的查询，但已复核的实现更狭窄，也更不可靠：其树遍历器只有在规划器条件中发现操作符表达式后才会判定成功。`WHERE TRUE`、裸布尔列或 `IS NULL` 等条件仍可能被拒绝，而出现无关的操作符表达式也不能证明语句安全。

该钩子先运行标准规划器，之后才报错。它会影响表扫描之外的规划语句，包括普通 SELECT 与某些 INSERT 形式，而 utility 命令走另一条路径。旧式钩子签名和对规划器树的直接假设都依赖 PostgreSQL 版本。由于没有内置逃生机制，意外拒绝可能阻塞正常操作；必须保留带外删除该库并重启服务器的方法。
