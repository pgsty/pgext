

## Usage

> [pg_similarity](https://github.com/eulerto/pg_similarity): Support similarity queries on PostgreSQL.
> Source: [README.md](https://raw.githubusercontent.com/eulerto/pg_similarity/master/README.md)

**pg_similarity** is an extension to support similarity queries on PostgreSQL. The implementation is tightly integrated in the RDBMS in the sense that it defines operators so instead of the traditional operators (`=` and `<>`) you can use `~~~` and `~!~` (any of these operators represents a similarity function).

**pg_similarity** has three main components:

- **Functions**: a set of functions that implements similarity algorithms available in the literature. These functions can be used as UDFs and will be the base for implementing the similarity operators;
- **Operators**: a set of operators defined at the top of similarity functions. They use similarity functions to obtain the similarity threshold and compare its value to a user-defined threshold to decide if it is a match or not;
- **Session Variables**: a set of variables that store similarity function parameters. These variables can be defined at run time.


--------

## Functions and Operators

This extension supports a set of similarity algorithms. The most known algorithms are covered by this extension. You must be aware that each algorithm is suited for a specific domain. The following algorithms are provided:

- L1 Distance (as known as City Block or Manhattan Distance)
- Cosine Distance
- Dice Coefficient
- Euclidean Distance
- Hamming Distance
- Jaccard Coefficient
- Jaro Distance
- Jaro-Winkler Distance
- Levenshtein Distance
- Matching Coefficient
- Monge-Elkan Coefficient
- Needleman-Wunsch Coefficient
- Overlap Coefficient
- Q-Gram Distance
- Smith-Waterman Coefficient
- Smith-Waterman-Gotoh Coefficient
- Soundex Distance

| Algorithm | Function | Operator | Use Index? | Parameters |
|---|---|---|---|---|
| L1 Distance | `block(text, text) returns float8` | `~++` | yes | `pg_similarity.block_tokenizer`, `pg_similarity.block_threshold`, `pg_similarity.block_is_normalized` |
| Cosine Distance | `cosine(text, text) returns float8` | `~##` | yes | `pg_similarity.cosine_tokenizer`, `pg_similarity.cosine_threshold`, `pg_similarity.cosine_is_normalized` |
| Dice Coefficient | `dice(text, text) returns float8` | `~-~` | yes | `pg_similarity.dice_tokenizer`, `pg_similarity.dice_threshold`, `pg_similarity.dice_is_normalized` |
| Euclidean Distance | `euclidean(text, text) returns float8` | `~!!` | yes | `pg_similarity.euclidean_tokenizer`, `pg_similarity.euclidean_threshold`, `pg_similarity.euclidean_is_normalized` |
| Hamming Distance | `hamming(bit varying, bit varying) returns float8` / `hamming_text(text, text) returns float8` | `~@~` | no | `pg_similarity.hamming_threshold`, `pg_similarity.hamming_is_normalized` |
| Jaccard Coefficient | `jaccard(text, text) returns float8` | `~??` | yes | `pg_similarity.jaccard_tokenizer`, `pg_similarity.jaccard_threshold`, `pg_similarity.jaccard_is_normalized` |
| Jaro Distance | `jaro(text, text) returns float8` | `~%%` | no | `pg_similarity.jaro_threshold`, `pg_similarity.jaro_is_normalized` |
| Jaro-Winkler Distance | `jarowinkler(text, text) returns float8` | `~@@` | no | `pg_similarity.jarowinkler_threshold`, `pg_similarity.jarowinkler_is_normalized` |
| Levenshtein Distance | `lev(text, text) returns float8` | `~==` | no | `pg_similarity.levenshtein_threshold`, `pg_similarity.levenshtein_is_normalized` |
| Matching Coefficient | `matchingcoefficient(text, text) returns float8` | `~^^` | yes | `pg_similarity.matching_tokenizer`, `pg_similarity.matching_threshold`, `pg_similarity.matching_is_normalized` |
| Monge-Elkan Coefficient | `mongeelkan(text, text) returns float8` | `~\|\|` | no | `pg_similarity.mongeelkan_tokenizer`, `pg_similarity.mongeelkan_threshold`, `pg_similarity.mongeelkan_is_normalized` |
| Needleman-Wunsch Coefficient | `needlemanwunsch(text, text) returns float8` | `~#~` | no | `pg_similarity.nw_threshold`, `pg_similarity.nw_is_normalized` |
| Overlap Coefficient | `overlapcoefficient(text, text) returns float8` | `~**` | yes | `pg_similarity.overlap_tokenizer`, `pg_similarity.overlap_threshold`, `pg_similarity.overlap_is_normalized` |
| Q-Gram Distance | `qgram(text, text) returns float8` | `~~~` | yes | `pg_similarity.qgram_threshold`, `pg_similarity.qgram_is_normalized` |
| Smith-Waterman Coefficient | `smithwaterman(text, text) returns float8` | `~=~` | no | `pg_similarity.sw_threshold`, `pg_similarity.sw_is_normalized` |
| Smith-Waterman-Gotoh Coefficient | `smithwatermangotoh(text, text) returns float8` | `~!~` | no | `pg_similarity.swg_threshold`, `pg_similarity.swg_is_normalized` |
| Soundex Distance | `soundex(text, text) returns float8` | `~*~` | no | |


--------

## Parameters

The several parameters control the behavior of the pg_similarity functions and operators. They can be classified in three classes:

- **tokenizer**: controls how the strings are tokenized. Valid values are **alnum**, **gram**, **word**, and **camelcase**. All tokens are lowercase. Default is **alnum**.
  - **alnum**: delimiters are any non-alphanumeric characters.
  - **gram**: an n-gram is a subsequence of length n, extracted using sliding-by-one technique.
  - **word**: delimiters are white space characters.
  - **camelcase**: delimiters are capitalized characters but they are also included as first token characters.
- **threshold**: controls how flexible the result set will be. Values range from **0.0** to **1.0**. Default is **0.7**.
- **normalized**: controls whether the similarity coefficient/distance is normalized (between 0.0 and 1.0) or not. Default is **true**.


--------

## Examples

Set parameters at run time:

```sql
SHOW pg_similarity.levenshtein_threshold;
-- 0.7

SET pg_similarity.levenshtein_threshold TO 0.5;

SET pg_similarity.cosine_tokenizer TO camelcase;

SET pg_similarity.euclidean_is_normalized TO false;
```

Simple tables for examples:

```sql
CREATE TABLE foo (a text);
INSERT INTO foo VALUES('Euler'),('Oiler'),('Euler Taveira de Oliveira'),('Maria Taveira dos Santos'),('Carlos Santos Silva');

CREATE TABLE bar (b text);
INSERT INTO bar VALUES('Euler T. de Oliveira'),('Euller'),('Oliveira, Euler Taveira'),('Sr. Oliveira');
```

### Using similarity functions

```sql
SELECT a, b, cosine(a,b), jaro(a, b), euclidean(a, b) FROM foo, bar;
```

### Using the levenshtein operator (~==)

```sql
SHOW pg_similarity.levenshtein_threshold;
-- 0.7

SELECT a, b, lev(a,b) FROM foo, bar WHERE a ~== b;
--              a             |          b           |   lev
-- ---------------------------+----------------------+----------
--  Euler                     | Euller               | 0.833333
--  Euler Taveira de Oliveira | Euler T. de Oliveira |     0.76

SET pg_similarity.levenshtein_threshold TO 0.5;

SELECT a, b, lev(a,b) FROM foo, bar WHERE a ~== b;
--              a             |          b           |   lev
-- ---------------------------+----------------------+----------
--  Euler                     | Euller               | 0.833333
--  Oiler                     | Euller               |      0.5
--  Euler Taveira de Oliveira | Euler T. de Oliveira |     0.76
```

### Using the qgram operator (~~~)

```sql
SET pg_similarity.qgram_threshold TO 0.7;

SELECT a, b, qgram(a, b) FROM foo, bar WHERE a ~~~ b;
--              a             |            b            |  qgram
-- ---------------------------+-------------------------+----------
--  Euler                     | Euller                  |      0.8
--  Euler Taveira de Oliveira | Euler T. de Oliveira    |  0.77551
--  Euler Taveira de Oliveira | Oliveira, Euler Taveira | 0.807692
```

### Comparing different operators

```sql
SELECT * FROM bar WHERE b ~@@ 'euler'; -- jaro-winkler operator
SELECT * FROM bar WHERE b ~~~ 'euler'; -- qgram operator
SELECT * FROM bar WHERE b ~== 'euler'; -- levenshtein operator
SELECT * FROM bar WHERE b ~## 'euler'; -- cosine operator
```
