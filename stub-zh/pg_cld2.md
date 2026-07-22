## 用法

来源：

- [官方 PGXN 分发包](https://pgxn.org/dist/pg_cld2/1.0.0/)
- [官方 pg_cld2 README](https://github.com/hedges333/pg_cld2/blob/f13f0721009db68fc0333aa7671dcf27ba16c739/README.md)
- [1.0.0 版扩展 SQL](https://github.com/hedges333/pg_cld2/blob/f13f0721009db68fc0333aa7671dcf27ba16c739/sql/pg_cld2--1.0.0.sql)

`pg_cld2` 将 Google Compact Language Detector 2（CLD2）暴露给 PostgreSQL，用于估计文本的自然语言和书写系统，并把候选语言映射到可用的 PostgreSQL 文本搜索配置。结果是统计提示，而不是权威分类。

### 检测语言与文字

扩展在构建和运行时都依赖外部 CLD2 库：

```sql
CREATE EXTENSION pg_cld2;

SELECT is_reliable,
       mll_language_code,
       mll_primary_script_code,
       language_1_language_code,
       language_1_percent,
       language_1_ts_name
FROM pg_cld2_detect_language(
    'This is a sample text to detect the language.'
);
```

`pg_cld2_detect_language` 返回 `pg_cld2_language_detection` 复合类型。除字节计数和 UTF-8 有效性外，它还包含最可能语言字段组（`mll_*`）以及三个按概率排列的候选字段组（`language_1_*` 到 `language_3_*`），其中有语言名称与代码、文字名称与代码、百分比、归一化分数和匹配的 `pg_catalog.pg_ts_config` 名称。

要查看完整调用签名和默认值，请运行：

```sql
SELECT pg_cld2_usage();
```

重要输入包括 `is_plain_text`（HTML 应设为 false）、`content_language_hint`、`tld_hint`、`cld2_language_hint`、`best_effort`、`text_encoding`、`tsconfig_language_hint` 和 `locale_lang_hint`。提示参数会有意影响结果。对于非 UTF-8 输入，可提供正确的 PostgreSQL 编码名进行转换。

### 结果解释与注意事项

根据结果采取动作前应检查 `is_reliable`；混合语言文本需要保留候选百分比。默认的 `best_effort = true` 会对短文本给出猜测而不是 `UNKNOWN`；如果宁可得到不确定结果，也不要强制答案，应将其关闭。`valid_prefix_bytes` 和 `is_valid_utf8` 可帮助识别格式错误的输入。

CLD2 是较旧的统计模型。短字符串、专有名词、音译、混合文字、标记文本和带偏向的提示都可能产生误导分类。不要用其输出做授权、合规、国籍或身份判断。PGXN 元数据只声明了历史 PostgreSQL 最低版本；部署前应针对实际现代服务器版本验证 CLD2 ABI 并重新编译此版本。
