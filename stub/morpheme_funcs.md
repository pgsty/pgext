## Usage

Sources:

- [Upstream README](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/readme.md)
- [Extension control file](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/morpheme_funcs.control)
- [Cargo package metadata](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/Cargo.toml)
- [Rust implementation](https://github.com/cystal-dot/morpheme_funcs/blob/c7607a5affb6a1d3688c92f36aa6456efc648f48/src/lib.rs)

`morpheme_funcs` version `0.0.0` is a Rust/pgrx wrapper around MeCab for Japanese morphological tokenization. `to_morpheme_array(text)` returns a sorted set of unique morphemes, while `calculate_similar_score(text,text)` computes the Jaccard ratio of the two unique-token sets.

### Example

```sql
CREATE EXTENSION morpheme_funcs;
SELECT to_morpheme_array('形態素解析機能');
SELECT calculate_similar_score('大ねじ小ねじ', 'ねじ');
```

Build and runtime require the native MeCab library plus an installed dictionary. Tokenization therefore depends on the exact dictionary and configuration and may vary between servers. The score ignores token order and frequency, so it is not a linguistic or semantic similarity measure. The implementation keeps a mutable process-local tagger and is still version `0.0.0`; pin the full native environment and load-test it before exposing it to arbitrary text.
