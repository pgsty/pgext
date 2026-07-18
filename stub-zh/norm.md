## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/README.md)
- [0.0.1 版本 SQL 实现](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/norm--0.0.1.sql)
- [PGXS Makefile](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/Makefile)
- [上游 control 文件](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/NORM.control)

`norm` 是为简单 CRUD 函数生成代码的纯 SQL/PLpgSQL 扩展。`norm_insert`、`norm_update`、`norm_delete` 与 `norm_get` 检查表元数据并创建针对具体表的函数。

```sql
CREATE EXTENSION norm;

CREATE TABLE games (
  game_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  game_name text NOT NULL
);

SELECT norm_insert('games', ARRAY['game_name']);
SELECT insert_games('chess');
```

生成的函数使用 `SECURITY DEFINER`，实现会根据传入的表名与列名构造动态 SQL。标识符只能来自可信调用方；授权前必须审查所有权和 `search_path`，并检查生成的函数。

### 早期版本限制

0.0.1 版不支持多表写入、left join、聚合，也不支持连接含多个外键的表。PGXS Makefile 与版本化 SQL 使用小写 `norm`，但仓库中的文件名是 `NORM.control`；在区分大小写的源码构建中，可能需要将它安装为 `norm.control`。
