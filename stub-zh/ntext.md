## 用法

来源：

- [固定提交的日文归一化规则](https://github.com/nuko-yokohama/ntext/blob/4ec273e7ac4a6f45a25d8eb3b5c12602a7dd919d/readme_ja.txt)
- [固定提交的 SQL 类型与操作符定义](https://github.com/nuko-yokohama/ntext/blob/4ec273e7ac4a6f45a25d8eb3b5c12602a7dd919d/ntext--1.0.sql)
- [固定提交的扩展 control 文件](https://github.com/nuko-yokohama/ntext/blob/4ec273e7ac4a6f45a25d8eb3b5c12602a7dd919d/ntext.control)

`ntext` 定义了一个用于日文归一化比较的 UTF-8 文本类型。它把平假名、片假名和半角片假名统一为片假名，把全角与半角拉丁字母折叠为半角小写形式，统一符号宽度，并在比较前应用额外的日文替换规则。

```sql
CREATE EXTENSION ntext;

SELECT 'かな'::ntext = 'カナ'::ntext AS same_after_normalization;
SELECT pg_ntext_normalize('ＡＢＣ'::ntext);
SELECT pg_similar_rate('東京都'::ntext, '東京'::ntext);
```

该扩展提供比较操作符、默认 B-tree operator class、`pg_ntext_normalize(ntext)`、`pg_similar_rate(ntext,ntext)` 和近似匹配操作符 `/=`。因此索引与等值语义基于归一化表示，而不是源文本的逐字节形式。

部分规则会有意丢失信息：上游文档把 `ヴァ` 行映射到 `バ` 行，转换 `ゐ`/`ヰ`、`ゑ`/`ヱ` 等历史假名，重写 `を`/`ヲ`，并删除末尾长音符。应根据应用搜索策略验证这些选择，并另外保存原始文本。已复核版本为 `1.0`；上游没有发布许可证说明或 PostgreSQL 兼容性矩阵。
