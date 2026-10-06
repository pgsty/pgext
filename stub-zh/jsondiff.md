## 用法

来源：

- [DBMS_Course_Project/JsonDiffAuditLog/README.md](https://github.com/NMGorovenko/ThirdSemester/blob/a43a67849c923d8dca6bc4f2dd6b2093b91ae93c/DBMS_Course_Project/JsonDiffAuditLog/README.md)
- [DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff--1.0.sql](https://github.com/NMGorovenko/ThirdSemester/blob/a43a67849c923d8dca6bc4f2dd6b2093b91ae93c/DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff--1.0.sql)
- [DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff.control](https://github.com/NMGorovenko/ThirdSemester/blob/a43a67849c923d8dca6bc4f2dd6b2093b91ae93c/DBMS_Course_Project/JsonDiffAuditLog/ext/jsondiff/jsondiff.control)

`jsondiff` 是课程审计应用中的小型 C 组件，只提供一个浅层 JSONB 差异函数；审计表、触发器和快照重建属于独立的应用脚本。

### 核心用法

```sql
CREATE EXTENSION jsondiff;
SELECT public.compute_json_diff('{"a":1,"b":2}'::jsonb, '{"a":3,"c":4}'::jsonb);
```

### 运行边界

`public.compute_json_diff(before_json, after_json)` 返回一个对象，`set` 包含新增或变更的键，`unset` 包含被删除的键。嵌套值整体替换，不递归计算差异。应使用对象形态输入，并在将差异应用到业务记录前验证边界情况。

版本化 SQL 将函数固定在 public，虽然控制文件标记为 relocatable，也不能依赖模式迁移来移动该函数。安装要求超级用户及编译后的 C 库，不要求预加载。上游构建示例使用 PostgreSQL 16，未声明更广的支持矩阵或许可证。
