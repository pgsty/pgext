## 用法

来源：

- [docs/merge.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/merge.rst)
- [merge/merge--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/merge/merge--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [merge/merge.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/merge/merge.control)

`merge` 是 Oliphant 的表间传输工具，按列名匹配字段，并根据键对齐源关系和目标关系。

### 核心用法

```sql
CREATE EXTENSION "merge";
CREATE TABLE merge_source (id integer PRIMARY KEY, label text);
CREATE TABLE merge_target (id integer PRIMARY KEY, label text);
INSERT INTO merge_source VALUES (1, 'example');
SELECT auto_insert('merge_source', 'merge_target');
SELECT * FROM merge_target;
```

### 运行边界

这些是最初面向 PostgreSQL 9.x 编写文档的纯 SQL/PLpgSQL 工具，尚未验证当前主版本兼容性。安装沿用 control 默认权限，需要管理员；不含共享库，也不需要预加载。 `auto_insert` 复制共有列，目标独有列使用默认值。`auto_merge` 通过可写 CTE 更新匹配的键并插入缺少的行，不能作为原生冲突处理的并发安全替代品。`auto_delete` 会删除源中不存在的目标行，应明确限定范围。调用者需要读取和写入权限；批量执行前应核查重复项、键、约束和源数据范围。
