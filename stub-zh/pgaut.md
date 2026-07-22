## 用法

来源：

- [已复核 commit 的归档官方 README](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/README.md)
- [1.0.0 版本 SQL 实现](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut--1.0.0.sql)
- [扩展 control 文件](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut.control)

`pgaut` 使用 `auto_update_timestamp` 域、数据库级事件触发器，以及为每张表生成的行触发器来模拟自动更新时间戳。该域基于不带时区的 `timestamp`。这个已归档的早期实现需要超级用户级事件触发器安装；在任何共享数据库中使用前都必须谨慎复核。

### 核心流程

```sql
CREATE EXTENSION pgaut;

CREATE TABLE app_user (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  updated_at auto_update_timestamp DEFAULT current_timestamp
);

INSERT INTO app_user (name) VALUES ('alice');
UPDATE app_user SET name = 'Alice' WHERE name = 'alice';
```

事件触发器响应表的创建、修改和删除。它们为每张表创建 BEFORE UPDATE 触发器函数，把所有被识别为 `auto_update_timestamp` 域的列赋值为 `clock_timestamp()`。

### 实际语义与注意事项

与 README 的描述不同，生成的行触发器不会比较 OLD 与 NEW 值。每次 UPDATE 都会重置所有已识别的时间戳域列，即使其他业务列没有改变，或更新时间戳本身就是唯一目标。

生成器甚至会为不含该域列的表创建函数和触发器。它构造模式、表、列、触发器与函数标识符时没有可靠引用；特殊或恶意名称可能破坏生成过程，或改变 SQL 的解释方式。域查询也不能可靠区分域所在模式，生成名称还可能与现有对象冲突。

只能安装在受控的模式命名空间中，检查每个生成对象，并测试表的创建、修改与删除迁移。该仓库已经归档，源自 PostgreSQL 9 时代；尚未证明与当前确切大版本兼容。
