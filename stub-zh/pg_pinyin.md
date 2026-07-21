## 用法

来源：

- [pg_pinyin v0.0.5 README](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/readme.md)
- [pg_pinyin v0.0.5 control file](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/pg_pinyin.control)
- [0.0.4到0.0.5升级SQL](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/pg_pinyin--0.0.4--0.0.5.sql)

pg_pinyin 对中文文本进行拉丁化，并提供分词器和查询辅助工具，适用于搜索应用。使用 pg_pinyin 可以创建稳定的拼音搜索键、分汉字文本或扩展拼音输入为 pg_search 正则表达式查询。

版本 0.0.5 主要是打包和工具链更新；其升级脚本未对 SQL 系统目录进行更改，因此用户面向的 API 与 0.0.4 保持兼容。

### 创建扩展

    CREATE EXTENSION pg_pinyin;

该扩展是可重定位的，并不需要 shared_preload_libraries 或服务器重启。

### 拉丁化文本

逐字符拉丁化或使用词感知分割：

    SELECT pinyin_char_romanize('重庆');
    SELECT pinyin_word_romanize('重庆火锅');
    SELECT pinyin_word_romanize('重庆火锅', ' ');

这两个函数接受一个可选后缀，该后缀插入到每个发出的拼音单元之后。字符模式是每字符确定性的；词模式使用捆绑的词典来解决上下文发音。

### 使用 pg_search 分词器输入

当安装了 pg_search 时，词拉丁化也接受 pg_search 分词器的结果：

    SELECT pinyin_word_romanize(
      description::pdb.icu::text[]
    )
    FROM documents;

该重载返回拉丁化文本；它不暴露每分词一行的 API。在不需要 pg_search 分词的情况下使用纯文本重载。

### 构建一个 pg_search 查询

如果在安装 pg_pinyin 之前已安装了 pg_search，pg_pinyin 提供了一个类型化的重载，返回 pdb.query：

    SELECT *
    FROM documents
    WHERE id @@@ pinyin_regex_phrase(
      'chong qing',
      slope => 1,
      max_expansions => 64,
      generated_pinyin => true
    );

如果未安装 pg_search，则相同的入口点将作为错误报告的占位符而不是静默返回不同类型的值。在升级后按预期顺序安装依赖项并测试函数签名。

### 对象索引

- pinyin_char_romanize(text [, suffix]) 返回基于字符的拼音文本。
- pinyin_word_romanize(text [, suffix]) 返回词典分割的拼音文本。
- pinyin_word_romanize(tokenizer_input [, suffix]) 接受一个 pg_search 分词器结果。
- pinyin_regex_phrase(text, slope, max_expansions, generated_pinyin) 当该集成可用时，构建一个 pg_search 拼音短语查询。
- pinyin_regex_phrase_patterns 是一个内部模式构建辅助函数；请优先使用公共查询函数。

### 运行注意事项

扩展在其 pinyin 模式中自带生成的字符和词字典。将这些表视为由扩展管理的数据，而不是应用程序表。拉丁化是规范化，不是翻译，并且可能需要应用端审查以处理模糊或领域特定的读音。
