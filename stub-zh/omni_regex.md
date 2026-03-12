


## 用法

> [omni_regex: PCRE 兼容正则表达式](https://docs.omnigres.org/omni_regex/regex/)

`omni_regex` 扩展提供基于 PCRE2 的正则表达式支持，包含命名捕获组功能。

### 运算符

| 运算符 | 描述 |
|:-------|:-----|
| `~` | 字符串匹配正则表达式 |
| `!~` | 不匹配运算符 |
| `=~` | 替代匹配方式（同 `~`） |

```sql
SELECT 'foo' ~ regex 'fo+';   -- true
SELECT 'bar' !~ regex 'foo';  -- true
```

### 函数

**`regex_match(text, regex)`** -- 返回第一个匹配的捕获组：

```sql
SELECT regex_match('ABC123', '([A-Z]*)(\d+)');  -- {ABC,123}
```

**`regex_matches(text, regex)`** -- 返回所有匹配结果的文本数组集合：

```sql
SELECT regex_matches('foo1bar', '(fo+|bar)(\d?)');
-- {foo,1}
-- {bar,""}
```

**`regex_named_groups(regex)`** -- 提取命名捕获组：

```sql
SELECT index FROM regex_named_groups('(fo+|bar)(?<num>\d?)')
WHERE name = 'num';  -- 2
```

命名捕获组语法：`(?<name>REGEX)`
