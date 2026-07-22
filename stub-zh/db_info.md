## 用法

来源：

- [Pinned official README](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/README.md)
- [Pinned extension SQL](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/db_info--0.1.0.sql)

`db_info` 提供纯 SQL 的清点辅助工具，用于汇总当前数据库的对象计数、数据库大小与所有者、角色、已安装扩展与语言、已用数据类型及表空间。它适合快速检查，不应视为稳定的监控 API。

### 核心流程

扩展会创建固定的 `db_info` 模式。查询其中的视图获得紧凑概览，调用函数获得文本报告：

```sql
CREATE EXTENSION db_info;

SELECT * FROM db_info.db_details;
SELECT * FROM db_info.db_obj_count;
SELECT * FROM db_info.get_extension_installed();
SELECT * FROM db_info.get_datatype_used();
```

### 重要对象

- `db_info.db_details`：当前数据库名称、所有者、编码、排序规则/字符类型、表空间和大小。
- `db_info.db_obj_count`：模式、表、分区、序列、函数/过程、触发器、规则、约束、索引和视图的数量。
- `db_info.get_db_roles()`：角色与对象所有权的文本报告。
- `db_info.get_extension_installed()`：以文本表示的已安装扩展名称与版本。
- `db_info.get_language_installed()`：已安装的过程语言。
- `db_info.get_datatype_used()`：用户列引用的数据类型。
- `db_info.get_tb_names()`：表、索引和物化视图使用的表空间。

### 注意事项

这些函数返回格式化的 `text` 或 `SETOF text`，而不是规范化的行，因此输出主要面向人工阅读。视图直接依赖 PostgreSQL 系统目录和 `information_schema`；对象类型语义与目录列可能随服务端版本改变。

上游 README 称其在 PostgreSQL 11 及更高版本上测试过，并明确警告可能仍有缺陷。0.2.0 版只修订 `db_obj_count` 视图，其余基础对象定义仍来自 0.1.0。将结果用于审计、容量决策或自动告警前应先复核。无需预加载或重启。
