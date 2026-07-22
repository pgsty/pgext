## 用法

来源：

- [腾讯云 PostgreSQL 支持扩展矩阵](https://cloud.tencent.com/document/product/409/75121)
- [腾讯云 PostgreSQL 产品页](https://cloud.tencent.com/product/postgres)
- [PostgreSQL 11 历史 timetravel 文档](https://www.postgresql.org/docs/11/contrib-spi.html)
- [PostgreSQL 12 删除 timetravel 的发行说明](https://www.postgresql.org/docs/12/release-12.html)

`timetravel` 是腾讯云 PostgreSQL 提供的服务商托管扩展。腾讯云当前矩阵只在 PostgreSQL 10 与 11 上列出 1.0 版本，PostgreSQL 12 及更高版本均未显示支持。依赖该扩展前，应确认实例引擎版本与最新矩阵。

### 服务商流程

腾讯云说明支持矩阵中的扩展可由用户自行创建。在符合条件的测试数据库中，应先创建扩展并立即盘点已安装对象，再设计应用逻辑：

```sql
CREATE EXTENSION timetravel;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'timetravel';

SELECT n.nspname, p.proname, pg_get_function_identity_arguments(p.oid)
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
JOIN pg_depend AS d ON d.objid = p.oid
JOIN pg_extension AS e ON e.oid = d.refobjid
WHERE e.extname = 'timetravel'
ORDER BY 1, 2, 3;
```

应使用获准管理扩展的腾讯云数据库账户，并在与生产相同的内核小版本上测试。可用性、升级行为、备份、只读副本与故障转移都仍属于腾讯云服务边界。

### API 与兼容性边界

PostgreSQL 历史上曾提供同名的 1.0 版本模块，通过触发器与会话本地控制函数实现行历史行为。PostgreSQL 12 因其依赖已淘汰类型与陈旧代码而删除该模块。这些历史文档说明了可能的来源，但腾讯云矩阵没有公开源码、control 文件、对象定义、权限，也没有声明其构建完全相同。

不能仅因名称相同就假设历史函数、触发器参数、数据类型或语义。创建表或触发器前，应把已安装对象清单与服务商文档比对，或提交腾讯云支持工单。由于当前矩阵在 11 之后不再列出 `timetravel`，升级 PostgreSQL 大版本前必须规划替代方案。
