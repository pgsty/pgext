## 用法

来源：

- [官方 Rust 实现](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/lib.rs)
- [官方图查询生成器](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/queries/mod.rs)
- [官方 Rust 软件包清单](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/Cargo.toml)

`pgautograph` 0.0.0 是未完成的 pgrx 实验，从 `public` 模式读取外键并构造拟议的属性图 DDL 字符串。其公开函数返回外键元数据，并将生成字符串写入 PostgreSQL 日志；它不会创建图。

### 检查外键

在一次性数据库中创建扩展并调用主要函数：

```sql
CREATE EXTENSION pgautograph;

SELECT * FROM list_foreign_keys();
SELECT hello_pgautograph();
```

`list_foreign_keys()` 返回 `source_table`、`source_pk`、`source_columns`、`target_table` 和 `target_column`。随后它会以 LOG 级别记录一个以 `CREATE PROPERTY GRAPH graph` 开头的字符串。评估原型时，需要手工捕获并审查该日志文本。

### 硬编码模型假设

目录查询只考虑源表位于 `public` 的外键。它使用不带模式限定的关系名，假定每个源表都有主键，每个外键行只处理一列，并应用简单的英语首字母大写与单数化。生成的顶点定义没有属性，边生成器则根据匹配引用的数量推断有向或无向形式。

### 开发中边界

函数既不返回也不执行图 DDL。它直接解包 NULL 目录值和 SPI 结果，所以没有主键的外键表可能引发 Rust panic。Hash set 迭代还会使生成子句顺序不确定。随附的问候测试所期待的文本与函数实际包含感叹号的结果不同，说明 0.0.0 源码树甚至没有通过自身测试进行发布门禁。

只能将它作为源码材料或一次性原型。实际使用前，必须加入带模式限定的目录处理、复合键支持、确定性输出、标识符引用、NULL 与错误处理、生成 DDL 返回值，以及针对准确 PostgreSQL 属性图语法的测试。不要仅为获取一个可以安全地用 SQL 编写的目录查询，就向它授予生产超级用户安装权限。
