## 用法

来源：

- [项目 README 与示例](https://github.com/avito-tech/pgmock/blob/f786e6f835771c65b2721fd7f275a673d627193d/README.md)
- [扩展 control 文件](https://github.com/avito-tech/pgmock/blob/f786e6f835771c65b2721fd7f275a673d627193d/pgmock.control)
- [0.2 版 SQL 实现](https://github.com/avito-tech/pgmock/blob/f786e6f835771c65b2721fd7f275a673d627193d/pgmock--0.2.sql)
- [回归测试](https://github.com/avito-tech/pgmock/tree/f786e6f835771c65b2721fd7f275a673d627193d/sql)

`pgmock` 0.2 是最初面向 PostgreSQL 9.4 及以上版本编写文档的纯 SQL 单元测试扩展。其 `mock()` 函数在 `pg_temp` 中重新创建所选函数及表依赖，并提供 `pg_temp.setup()` 和 `pg_temp.teardown()` 来重置测试状态。

### 隔离函数测试

将它安装到专用模式以避免名称冲突，并用事务包住每个测试套件：

```sql
CREATE SCHEMA pgmock;
CREATE EXTENSION pgmock WITH SCHEMA pgmock;

BEGIN;

SELECT pgmock.mock($$"${'public.universal_answer'::regproc}"$$);
SELECT pg_temp.setup();
SELECT pg_temp.universal_answer();
SELECT pg_temp.teardown();

ROLLBACK;
```

JSON 请求可以标识 `regproc` 和 `regclass` 依赖，并选择性重建约束、默认值、常量函数和触发器。只有明确请求的行为会被重建；临时表不会自动成为源表语义的完整副本。

### 仅限测试的安全与兼容性边界

`mock()` 读取系统目录，并生成 SQL 以在临时搜索路径下重建对象。只接受来自受信测试固件的定义，使用隔离的非生产数据库和角色，并保留外层回滚。绝不能在生产会话中运行未经审查的 mock 请求。

0.2 版实现引用了包括 `pg_attrdef.adsrc` 在内的旧目录字段，该字段已从现代 PostgreSQL 删除。仓库没有当前兼容矩阵或 release 活动，因此原始 PostgreSQL 9.4+ 声明不能证明它支持当前大版本。应在目标大版本上测试安装和每一种所需对象形态，尤其关注重载函数、identity 与生成列、分区表、策略、security-definer 函数、触发器和搜索路径解析。
