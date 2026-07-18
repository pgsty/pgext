## 用法

来源：

- [上游 README](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/README.md)
- [1.0.0 版安装 SQL](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/tm_postgres_basics--1.0.0.sql)
- [扩展控制文件](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/tm_postgres_basics.control)
- [MIT 许可证](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/LICENSE)

tm_postgres_basics 1.0.0 是一组小型的 2016 年纯 SQL 函数，用于将关系名解析为 OID、递归检查 pg_inherits 祖先关系，以及列出继承表。

### 继承检查

```sql
CREATE EXTENSION tm_postgres_basics;
CREATE TEMP TABLE tm_parent();
CREATE TEMP TABLE tm_child() INHERITS (tm_parent);
SELECT tm_is_inherit_from(
  'tm_child'::regclass,
  'tm_parent'::regclass
);
```

对于这个会话本地示例，直接基于 OID 的祖先函数返回 true。

### 注意事项

- tm_name_to_oid 只按 relname 搜索 pg_class。它忽略模式与关系类型，名称重复时可能返回任意匹配。应优先使用模式限定的 regclass 或 to_regclass 查找。
- tm_find_tables_inherit_from 虽接受用于枚举的模式，却通过不区分模式的辅助函数解析父表与子表，因此同名对象可能产生错误结果。
- 列表函数会扫描请求模式中的每个表，并为每个候选递归查询 pg_inherits。在大型目录上可能产生不必要的高成本。
- 函数与目录引用都没有模式限定，并使用调用者搜索路径。应安装到受控模式并以模式限定方式调用。
- 所有函数都保留默认的 volatile 与并行不安全标记，尽管它们只读取目录，因此限制了规划器优化。
- 上游没有测试、升级脚本或当前兼容声明，最后修改于 2016 年。现代 PostgreSQL 目录函数通常更安全、更快。
