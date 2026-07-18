## 用法

来源：

- [上游 README](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/readme.md)
- [扩展 control 文件](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/morpheme_funcs.control)
- [Cargo 软件包元数据](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/Cargo.toml)
- [Rust 实现](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/src/lib.rs)

`morpheme_funcs` `0.0.0` 版是用 Rust/pgrx 编写的 MeCab 封装，用于日语形态素切分。`to_morpheme_array(text)` 返回排序后的唯一形态素集合，`calculate_similar_score(text,text)` 计算两个唯一词元集合的 Jaccard 比率。

### 示例

```sql
CREATE EXTENSION morpheme_funcs;
SELECT to_morpheme_array('形態素解析機能');
SELECT calculate_similar_score('大ねじ小ねじ', 'ねじ');
```

构建和运行需要原生 MeCab 库及已安装字典，因此切分结果取决于具体字典与配置，在不同服务器间可能变化。分数忽略词元顺序和频率，不是语言学或语义相似度。实现保存了进程本地的可变 Tagger，版本仍为 `0.0.0`；向任意文本开放前，应固定完整原生环境并进行负载测试。
