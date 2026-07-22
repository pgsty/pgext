## 用法

来源：

- [官方 README](https://github.com/obartunov/dict_regex/blob/master/README.dict_regex)
- [官方扩展 SQL](https://github.com/obartunov/dict_regex/blob/master/dict_regex--1.0.sql)
- [官方示例规则](https://github.com/obartunov/dict_regex/blob/master/data/dict_regex.rules)
- [官方 control 文件](https://github.com/obartunov/dict_regex/blob/master/dict_regex.control)

`dict_regex` 是 PostgreSQL 全文搜索词典，通过 PCRE 兼容正则规则改写一个或多个输入 token。它在一个词典中组合了同义词扩展、类似词库的多词归并、捕获组替换与停用词删除。

### 核心流程

把经过审查的规则文件放到 PostgreSQL 服务器进程可读的位置，创建扩展，并让已安装的 `regex` 词典指向该文件：

```text
# synonym expansion
catalogue catalogue catalog

# collapse a phrase
regular\sexpressions? regex

# reorder two captured groups
(\d\d\d\d)(\d\d\d\d) $2$1
```

```sql
CREATE EXTENSION dict_regex;
ALTER TEXT SEARCH DICTIONARY regex
  (RULES = '/srv/postgresql/dict_regex.rules');

SELECT ts_lexize('regex', 'catalogue');
```

直接用 `ts_lexize` 测试后，可通过 `ALTER TEXT SEARCH CONFIGURATION ... ALTER MAPPING ...` 把 `regex` 挂到全文搜索配置的指定 token 类型上。

### 规则语义

每个非注释行包含以空白分隔的模式与输出配方。配方为空时，匹配项成为停用词。词典会自动锚定模式，选择匹配 token 序列最长的规则；长度相同时按文件顺序取最早者。多个输出词会以同一位置的同义词返回；`$1` 到 `$9` 替换捕获组，输出会转为小写。

部分匹配引擎限制某些单字符/元序列重复。例如应写成 `(\d)+` 而非 `\d+`，写成 `(a){2,4}` 而非 `a{2,4}`。

### 注意事项

扩展依赖 PCRE，并读取服务器端规则文件。上游明确指出版本 `1.0` 不安全，因为它没有充分限制文件路径或内容：能够修改词典选项的用户可以指向 PostgreSQL 服务器可读的任意文件，并可能重建其内容。若不对源码加固，不应向不可信角色开放词典修改，也不能用于安全关键环境。

规则顺序会影响结果，多 token 匹配还会改变词位。应使用代表性文本同时测试索引端与查询端配置；修改文件后按实际情况重新加载或重建会话。
