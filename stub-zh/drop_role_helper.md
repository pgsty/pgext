## 用法

来源：

- [上游 README](https://github.com/cybertec-postgresql/drop_role_helper/blob/b56500c0b4167ed8cb744c390c0b7cf24c0fd9d9/README.md)
- [扩展 SQL](https://github.com/cybertec-postgresql/drop_role_helper/blob/b56500c0b4167ed8cb744c390c0b7cf24c0fd9d9/drop_role_helper--1.0.sql)
- [扩展 control 文件](https://github.com/cybertec-postgresql/drop_role_helper/blob/b56500c0b4167ed8cb744c390c0b7cf24c0fd9d9/drop_role_helper.control)

`drop_role_helper` 提供表函数 `drop_role_helper(regrole)`，用于生成撤销某个角色在当前数据库中权限的 SQL 语句。在 `psql` 中，处理完该角色拥有的对象后，可将返回的语句交给 `\gexec` 执行：

```sql
REASSIGN OWNED BY doomed_role TO replacement_role;
DROP OWNED BY doomed_role;

CREATE EXTENSION drop_role_helper SCHEMA public;
SELECT * FROM public.drop_role_helper('doomed_role') \gexec
```

需要在集群的每个数据库中重复这套流程，确认不再存在依赖后再执行 `DROP ROLE doomed_role`。

### 权限与安全

调用 `drop_role_helper(regrole)` 本身不需要特殊权限。执行其生成的 SQL 则需要超级用户权限，或拥有所有受影响对象；处理默认权限时还可能需要成为其定义角色的成员。使用 `\gexec` 前应先审阅生成的各行，尤其是在共享集群或生产集群中。
