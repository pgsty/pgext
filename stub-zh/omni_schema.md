


## 用法

> [omni_schema: 高级模式管理工具](https://docs.omnigres.org/omni_schema/reference/)

`omni_schema` 扩展通过迁移和从文件系统重新加载对象来管理应用模式。

### 运行迁移

```sql
SELECT * FROM omni_schema.migrate_from_fs(omni_vfs.local_fs('/path/to/migrations'));
```

递归查找 `.sql` 文件，按路径名顺序应用，跳过已应用的迁移。执行记录保存在 `omni_schema.migrations` 表中。

### 重新加载模式对象

```sql
SELECT * FROM omni_schema.load_from_fs(omni_vfs.local_fs('/path/to/project'));
```

从文件系统重新加载函数、策略和视图。要忽略某个文件，在其中放置 `omni_schema[[ignore]]`（通常在注释中）。

### 迁移表

| 列 | 类型 | 描述 |
|:---|:-----|:-----|
| `id` | int | 唯一标识符 |
| `name` | text | 迁移文件名 |
| `migration` | text | 源代码 |
| `applied_at` | timestamp | 应用时间 |

### 多语言支持

支持 SQL、PL/pgSQL（`.sql`）、Python（`.py`）、Perl（`.pl`）、Tcl（`.tcl`）和 Rust（`.rs`）。使用 `SQL[[...]]` 指令指定函数签名：

```python
# times_ten.py
# SQL[[create function times_ten(a integer) returns integer]]
return a * 10
```
