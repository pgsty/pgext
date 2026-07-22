## 用法

来源：

- [官方扩展 SQL](https://github.com/pindlebot/pg_slugify/blob/fab77da2df8306eac3f8010ff2c7cad8c3dfb75e/pg_slugify--0.0.1.sql)
- [官方控制文件](https://github.com/pindlebot/pg_slugify/blob/fab77da2df8306eac3f8010ff2c7cad8c3dfb75e/pg_slugify.control)

`pg_slugify` 版本 `0.0.1` 提供一个纯 SQL 函数，用于把文本转换为小写、面向 ASCII 的 URL slug。它依赖 `unaccent`，并始终在 `public` 模式中创建函数。

### 核心流程

通过扩展依赖安装所需组件，并在写入应用记录时生成 slug：

```sql
CREATE EXTENSION pg_slugify CASCADE;

SELECT public.slugify('Crème brûlée: Summer Menu 2026');
```

在标准上游 unaccent 规则下，结果是 `creme-brulee-summer-menu-2026`。

### 转换规则

`public.slugify(text)` 依次执行 `unaccent`、转小写、把 `[a-z0-9_-]` 之外的每段字符替换为连字符，并移除首尾连字符。原有下划线和连字符会被保留。函数为 strict，因此 NULL 输入返回 NULL。

### 运维注意事项

SQL 将 `public.slugify(text)` 声明为 `IMMUTABLE`，但输出取决于已安装的 `unaccent` 规则。如果规则变化，已存 slug 与表达式索引条目不会自动更新；应按需重新生成数据或重建索引。该函数面向 ASCII，可能把不同源字符串折叠为同一个 slug，也不强制唯一性或最大长度。如果 slug 用作标识符，应增加符合应用需求的唯一约束与冲突策略。由于模式被硬编码，安装前还要检查 `public` 中的名称冲突与权限。
