## 用法

来源：

- [1.2.2 版手册](https://pgxn.org/dist/temporal_tables/1.2.2/)
- [1.2.2 版变更日志](https://api.pgxn.org/src/temporal_tables/temporal_tables-1.2.2/CHANGELOG)
- [控制文件](https://api.pgxn.org/src/temporal_tables/temporal_tables-1.2.2/temporal_tables.control)
- [函数定义与权限](https://api.pgxn.org/src/temporal_tables/temporal_tables-1.2.2/temporal_tables--1.2.2.sql)

时态表（temporal table）是一种记录行有效时间段的表。系统时段（system period）是由系统自动维护的列（或列对），包含该行在数据库层面有效的时间范围。当你向这样的表中插入行时，系统会自动生成时段的起止值。当你更新或删除系统时段时态表中的行时，旧版本会自动归档到另一张表——即历史表（history table）。


### 创建系统时段时态表

该扩展使用一个通用触发器函数来维护系统时段时态表的行为：

```
versioning(<system_period_column_name>, <history_table_name>, <adjust>)
```

首先，创建一张表并添加系统时段列：

```sql
CREATE EXTENSION temporal_tables;

CREATE TABLE employees (
  name text NOT NULL PRIMARY KEY,
  department text,
  salary numeric(20, 2)
);

ALTER TABLE employees ADD COLUMN sys_period tstzrange NOT NULL;
```

然后创建历史表：

```sql
CREATE TABLE employees_history (LIKE employees);
```

历史表必须包含与原表同名、同类型的系统时段列。如果两张表都包含某列，则数据类型必须相同。

最后，创建触发器将其与历史表关联：

```sql
CREATE TRIGGER versioning_trigger
BEFORE INSERT OR UPDATE OR DELETE ON employees
FOR EACH ROW EXECUTE PROCEDURE versioning('sys_period',
                                          'employees_history',
                                          true);
```


### 插入数据

向系统时段时态表插入数据与普通表类似：

```sql
INSERT INTO employees (name, department, salary)
VALUES ('Bernard Marx', 'Hatchery and Conditioning Centre', 10000);

INSERT INTO employees (name, department, salary)
VALUES ('Lenina Crowne', 'Hatchery and Conditioning Centre', 7000);
```

`sys_period` 列的起始值表示该行何时变为当前版本，由 `CURRENT_TIMESTAMP` 自动生成。


### 更新数据

当用户更新行时，触发器会将旧行的副本插入历史表。如果单个事务内对同一行进行了多次更新，只会生成一条历史记录：

```sql
UPDATE employees SET salary = 11200 WHERE name = 'Bernard Marx';
```

此时历史表包含之前的版本：

| name         | department                       | salary | sys_period              |
|--------------|----------------------------------|--------|-------------------------|
| Bernard Marx | Hatchery and Conditioning Centre | 10000  | [2006-08-08, 2007-02-27)|

#### 更新冲突与时间调整

当多个事务更新同一行时，可能会发生更新冲突。当 `adjust` 参数设为 `true` 时，`sys_period` 的起始值会通过加上一个微小的时间增量（通常为 1 微秒）来避免冲突，否则会抛出 SQLSTATE 22000 错误。


### 删除数据

当用户删除数据时，触发器同样会将行添加到历史表：

```sql
DELETE FROM employees WHERE name = 'Helmholtz Watson';
```


### 高级用法

你可以为版本控制触发器设置自定义系统时间，这在从已记录时间戳的系统构建数据仓库时非常有用：

```sql
SELECT set_system_time('1985-08-08 06:42:00+08');
```

恢复默认行为：

```sql
SELECT set_system_time(NULL);
```

如果在随后被中止的事务中执行，所有更改都会撤销。如果已提交，更改将持续到会话结束。


### 示例与技巧

#### 使用继承创建历史表

```sql
CREATE TABLE employees_history (
  name text NOT NULL,
  department text,
  salary numeric(20, 2),
  sys_period tstzrange NOT NULL
);

CREATE TABLE employees (PRIMARY KEY(name)) INHERITS (employees_history);
```

#### 历史表清理

历史表会持续增长。以下是几种清理策略：

  1. 定期删除历史表中的旧数据。
  2. 使用分区，并将旧分区从历史表中分离。
  3. 仅保留每行的最近 N 个版本。
  4. 当时态表中的对应行被删除时，同步清理历史记录。
  5. 按业务规则清理符合条件的行。

你也可以将历史表设置到另一个表空间，将其迁移到更廉价的存储上。

#### 数据审计

你可以添加触发器来记录修改或删除当前行的用户：

```sql
ALTER TABLE employees ADD COLUMN user_modified text;
ALTER TABLE employees_history ADD COLUMN user_modified text;
ALTER TABLE employees_history ADD COLUMN user_deleted text;

CREATE FUNCTION employees_modify()
RETURNS TRIGGER AS $$
BEGIN
  NEW.user_modified = SESSION_USER;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER employees_modify
BEFORE INSERT OR UPDATE ON employees
FOR EACH ROW EXECUTE PROCEDURE employees_modify();

CREATE FUNCTION employees_delete()
RETURNS TRIGGER AS $$
BEGIN
  NEW.user_deleted = SESSION_USER;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER employees_delete
BEFORE INSERT ON employees_history
FOR EACH ROW EXECUTE PROCEDURE employees_delete();
```

### 版本与权限边界

`temporal_tables` 维护系统时段历史，不实现应用时段版本控制。`1.2.2` 版本新增 PostgreSQL 16 支持。上游手册要求 PostgreSQL 9.2 或更高版本，但这一历史发布记录不能证明更新主版本的软件包支持矩阵。

由超级用户安装扩展。SQL 脚本撤销了 `PUBLIC` 对 `versioning()` 的访问，因此应明确授予需要创建版本控制触发器的角色执行权限。时段列须为 `tstzrange`，历史表与当前表共有列的类型必须一致。历史数据增长与保留策略仍由管理员管理。
