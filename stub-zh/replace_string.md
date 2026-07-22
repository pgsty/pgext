## 用法

来源：

- [官方 README](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/README.md)
- [函数实现](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/src/lib.rs)
- [构建特性矩阵](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/Cargo.toml)
- [扩展控制文件](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/replace_string.control)

`replace_string` 提供 `repl_str(text, text, text) → text`，通过 Rust 正则库把一个字符串的所有出现替换为另一个字符串。搜索字符串会被转义，因而按字面匹配；但替换字符串仍使用正则替换展开规则。

### 核心流程

```sql
CREATE EXTENSION replace_string;

SELECT repl_str('h@llo', '@', 'e');
SELECT repl_str('a.c a.c', '.', '-');
```

两个示例都执行全局字面搜索。PostgreSQL 已提供 `replace(text, from, to)` 进行完全字面的替换，因此除非现有应用必须使用该扩展，否则应优先选择内置函数。

### 替换语义

实现会对搜索参数应用 `regex::escape`，然后调用 `Regex::replace_all`。因此，搜索文本中的正则元字符没有特殊含义。相反，替换文本中的 `$name`、`$1` 和 `$$` 遵循 Rust 正则捕获展开语法。由于转义后的搜索模式没有捕获组，捕获引用可能展开为空文本，而不是保留字面内容。替换文本可能包含美元符号时，应使用 PostgreSQL 内置的 `replace`。

空搜索字符串会生成空正则表达式，并在每个匹配边界插入替换文本；这可能出乎意料，也可能大幅增大结果。SQL 空输入遵循已构建软件包中 pgrx 生成函数包装器的严格性行为；如果空值处理很重要，应实际核验。

### 兼容性

扩展版本为 `0.0.0`，安装需要超级用户，且不可重定位。当前 Cargo 特性面向 PostgreSQL 13、14 与 15，默认 14，并使用 `pgrx` 0.16.1 和 regex 1.12.2。它不需要预加载或运行时配置。

正则替换行为属于结果语义的一部分，因此应固定实际软件包版本。在批量迁移数据前，应测试美元符号、空搜索字符串、Unicode、大值及目标 PostgreSQL 兼容性。
