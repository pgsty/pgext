## 用法

来源：

- [pg_readme 0.7.1 README](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/README.md)
- [pg_readme 0.7.1 控制文件](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme.control)
- [pg_readme 0.7.1 升级 SQL](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/sql/pg_readme--0.7.0--0.7.1.sql)
- [Pigsty 软件包矩阵](https://pgext.cloud/ext/pg_readme)

`pg_readme` 根据 `COMMENT` 对象和实时目录元数据，为 PostgreSQL 扩展或模式生成 Markdown 文档。使用它可以让扩展的 README 与其 SQL 定义保持接近，并在源代码管理中审查生成结果。

### 安装并生成 Markdown

```sql
CREATE EXTENSION pg_readme CASCADE;

SELECT pg_extension_readme('my_extension'::name);
SELECT pg_schema_readme('my_schema'::regnamespace);
```

控制文件要求 `hstore`，扩展可重定位；只要调用者能够安装依赖并创建相应对象，就允许非超级用户安装。

### 添加处理指令

将 Markdown 和处理指令放入扩展或模式的注释中：

```sql
COMMENT ON EXTENSION my_extension IS $markdown$
### `my_extension`

What the extension does.

### Reference

<?pg-readme-reference?>

### Colophon

<?pg-readme-colophon?>
$markdown$;
```

`<?pg-readme-reference?>` 会展开为根据目录生成的对象参考。`<?pg-readme-colophon?>` 会添加生成元数据。将生成的章节嵌入其他内容时，可通过可选的指令属性调整标题深度。

### 设置

- `pg_readme.include_view_definitions`：包含视图定义；默认为 `true`。
- `pg_readme.include_routine_definitions_like`：需要包含定义的例程名称模式数组；默认为 `'{test__%}'`。
- `pg_readme.include_this_routine_definition`：是否包含当前定义的例程局部覆盖项。
- `pg_readme.readme_url`：生成内容使用的上游 README 链接。

项目需要可复现的生成设置时，请在包装函数或事务中使用 `SET` 选项。

### 版本 0.7.1 与注意事项

- 版本 0.7.1 修复了 PostgreSQL 18 参考文档生成问题，该问题可能重复列出数组/复合表类型和 `NOT NULL` 标记。
- 上游和当前 Pigsty DEB 软件包为 0.7.1，而当前 Pigsty RPM 软件包仍为 0.7.0。在依赖 PostgreSQL 18 修复前，请检查 `pg_available_extension_versions`。
- 生成结果反映当前数据库目录、已安装扩展版本、注释以及生成时间。应审查差异，不要假定两个环境会生成完全相同的文本。
- 目录自省不能替代人工编写的运维指导。请在维护的正文中保留前置条件、预加载/重启行为、升级说明和不安全操作。
- 旧 README 的包装函数示例中出现了单数设置 `pg_readme.include_routine_definition_like`，但当前文档中的 GUC 是复数形式 `pg_readme.include_routine_definitions_like`。
