## 用法

来源：

- [官方 README](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/README.md)
- [扩展 control 文件](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/grants_manager.control)
- [权限对齐实现](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/gm_align_permissions.sql)

`grants_manager` 是一个小型 PL/pgSQL 工具集，可把表与序列权限快照成可编辑矩阵、报告偏差，并按需应用声明的授权。它只应在受控管理流程中使用：应用矩阵会执行 `REVOKE ALL` 和新的 `GRANT` 语句。

### 核心流程

```sql
CREATE EXTENSION grants_manager;

CREATE ROLE app_reader;
CREATE TABLE public.orders(id bigint PRIMARY KEY);
GRANT SELECT ON public.orders TO app_reader;

SELECT gm_generate_current();
SELECT * FROM public.grants_manager;

UPDATE public.grants_manager
SET app_reader = ARRAY['r', 'w']
WHERE object_name = 'orders';

SELECT gm_align_permissions(p_execute := false);
SELECT gm_align_permissions(p_execute := true);
```

`gm_generate_current()` 根据当前表与序列 ACL 重建 `public.grants_manager`，并把旧矩阵改名为 `grants_manager_old`。每个角色对应一个数组列。编辑这些数组后，先运行 `gm_align_permissions(false)` 只查看 NOTICE，复核拟议变更，确认无误后才使用 `gm_align_permissions(true)` 修改权限。

### 主要对象

- `gm_generate_current()` 创建可编辑的权限快照。
- `gm_get_status()` 以 `schema_name`、`object_name`、`grantee` 与 `grants` 返回当前授权。
- `gm_align_permissions(boolean)` 报告或应用差异。
- `gm_translate(text)` 在 PostgreSQL 权限名称与 ACL 字母之间转换。

实现中使用的 ACL 字母包括：`r` 表示 SELECT，`w` 表示 UPDATE，`a` 表示 INSERT，`d` 表示 DELETE，`D` 表示 TRUNCATE，`x` 表示 REFERENCES，`t` 表示 TRIGGER，`U` 表示 USAGE。

### 注意事项

版本 `0.0.1` 处理普通表与序列，并未声称覆盖所有可授权对象。上游 README 将组角色对齐列为未完成功能。生成快照会替换工作矩阵，而执行对齐会先撤销全部权限，再按声明集合重新授权。投入生产前应覆盖所有权、带引号标识符、Schema、角色继承与并发变更等代表性情况进行测试。
