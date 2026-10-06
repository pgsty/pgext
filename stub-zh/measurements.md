## 用法

来源：

- [readme.md](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [packages/measurements/sql/measurements--0.0.1.sql](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/measurements/sql/measurements--0.0.1.sql)
- [packages/measurements/Makefile](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/measurements/Makefile)
- [packages/measurements/measurements.control](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/measurements/measurements.control)

`measurements` 安装预置数据的 `measurements.quantities` 参考表，字段描述物理量名称、标签、单位及说明文字。

### 核心用法

```sql
CREATE EXTENSION "measurements" CASCADE;
SELECT name, unit, description FROM measurements.quantities WHERE name = 'Length';
```

### 运行边界

此历史 SQL 分发提供参考记录，不是单位转换引擎，也不提供量纲类型系统。部分说明条目存在格式错误，用作领域权威数据前应核对预置值。该源码修订中未找到当前 PostgreSQL 主版本矩阵或许可证声明。

须先安装声明的依赖：`plpgsql`, `uuid-ossp`, `launchql-extension-verify`。这是 SQL/PLpgSQL 代码，没有自己的共享库，也不要求预加载。控制文件允许非超级用户安装，但没有标记为 trusted；仍须满足依赖和模式权限。
