## 用法

来源：

- [官方文档](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/doc/user_info.md)
- [官方扩展 SQL](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/sql/user_info--0.0.1.sql)
- [官方扩展控制文件](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/user_info.control)

`user_info` 版本 `0.0.1` 提供目录报告函数，用于查看角色成员关系、对象所有权与直接授予的权限。项目已经归档，其 SQL 面向旧版 PostgreSQL 目录，因此应把它视为历史诊断代码，而非当前授权判定依据。

### 核心流程

函数接受角色 OID、角色名称；部分函数也可以不传参数，以当前用户为目标：

```sql
CREATE EXTENSION user_info;

SELECT * FROM user_objects('app_user');
SELECT * FROM granted_roles('app_user');
SELECT * FROM granted_roles_pretty('app_user');
SELECT * FROM accessible_objects('app_user');
```

`user_objects(...)` 会合并多个系统目录与信息模式报告，因此同一张表可能分别以关系、表、行类型和数组类型出现。跨对象类别的重复属于预期行为，不会被自动去重。

### 报告限制

`granted_roles(...)` 递归返回成员关系路径与层级。`granted_roles_pretty(...)` 以便于阅读的方式格式化同一层次。`accessible_objects(...)` 报告直接显式授予所选角色的权限。

权限报告不包含默认权限，也不包含通过其他角色继承的权限，因此不能计算有效访问权。它还依赖扩展开发年代的目录布局与对象类别；在现代 PostgreSQL 上，安装或查询可能失败，而且不会覆盖新增对象类型。用于审查前必须用当前 PostgreSQL 目录与权限函数逐项验证，不能单独依赖它进行安全或合规决策。
