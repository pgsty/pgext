## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/README.md)
- [0.0.1 版本 SQL 实现](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/norm--0.0.1.sql)
- [扩展 control 文件](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/NORM.control)

`norm` 是用于生成简单 CRUD 函数的纯 SQL/PLpgSQL 代码生成器。生成器检查目录元数据，并创建表专属的 `SECURITY DEFINER` 函数。只能在可信的模式迁移流程中使用它，并复核每个标识符与每个生成的定义。

### 核心流程

```sql
CREATE EXTENSION norm;

CREATE TABLE games (
  game_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  game_name text NOT NULL
);

SELECT norm_insert('games', ARRAY['game_name']);
SELECT norm_get('games', ARRAY['game_id', 'game_name']);

SELECT insert_games('chess');
SELECT * FROM get_games();
```

公开生成器包括 `norm_insert(table, columns, function_name)`、`norm_update(table, columns, filters, function_name)`、`norm_delete(table, filters, function_name)` 和 `norm_get(tables, columns, filters, function_name)`。最后的名称参数可省略；默认值由表名推导。

生成的更新函数使用 `coalesce(parameter, column)`，所以不能把目标列设为 SQL NULL。多表读取从外键推导内连接，在受支持路径中拒绝带有多个外键的表，也不提供外连接、聚合或任意子查询。

### 安全边界

已复核的 SQL 在插入调用者提供的表名、列名和自定义函数名时，没有始终使用标识符引用。生成的 `SECURITY DEFINER` 函数也没有建立加固的 `search_path`。因此，不可信的生成器输入可能导致 SQL 注入、对象替换或权限提升。

不要向应用角色授予生成器执行权。只能把它作为受控部署步骤运行，避免跨模式存在有歧义的同名表，检查生成函数的定义与所有者，设置安全的执行路径，并在暴露任何生成的辅助函数前撤销默认权限。
