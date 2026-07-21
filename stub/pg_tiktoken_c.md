## Usage

Sources:

- [pg_tiktoken_c README at the packaged revision](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/README.md)
- [Version 1.1 SQL API](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/sql/pg_tiktoken_c--1.1.sql)
- [Extension control file](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/pg_tiktoken_c.control)
- [Bundled vocabulary data](https://github.com/relytcloud/pg_tiktoken_c/tree/fa2957b6ece483322f4c4ce0c374b3b77e22b892/data)

pg_tiktoken_c implements OpenAI-compatible tiktoken encoding in C inside PostgreSQL. Use it to count or materialize tokens near stored text and to split text into token-bounded chunks before embedding or model requests.

### Create the Extension

    CREATE EXTENSION pg_tiktoken_c;

The implementation depends on PCRE2 10.30 or later at build time. It does not require shared_preload_libraries; vocabulary data is loaded and cached per backend as encodings are used.

### Encode and Count

    SELECT tiktoken_encode('cl100k_base', 'PostgreSQL search');
    SELECT tiktoken_count('cl100k_base', 'PostgreSQL search');

tiktoken_encode returns a bigint array of token identifiers. tiktoken_count returns the token count without requiring the caller to retain the token array.

The bundled selectors include cl100k_base, o200k_base, r50k_base, p50k_base, and p50k_edit, together with aliases documented by the project. Choose the encoding required by the target model rather than assuming all models share a vocabulary.

### Chunk Text

Return chunks as an array:

    SELECT chunk_text(
      'long document text',
      chunk_size => 512,
      chunk_overlap => 64,
      encoding => 'cl100k_base'
    );

Or return one row per chunk:

    SELECT *
    FROM chunk_text_table(
      'long document text',
      chunk_size => 512,
      chunk_overlap => 64,
      encoding => 'cl100k_base'
    );

chunk_text_table returns chunk_index, chunk, and token_count. The chunk index is zero-based. Overlap repeats boundary tokens between neighboring chunks and must be smaller than the chunk size.

### Function Index

- tiktoken_encode(selector, text) returns bigint[] token identifiers.
- tiktoken_count(selector, text) returns bigint token count.
- chunk_text(input_text, chunk_size, chunk_overlap default 0, encoding default cl100k_base) returns text[].
- chunk_text_table(input_text, chunk_size, chunk_overlap default 0, encoding default cl100k_base) returns one row per chunk with its index and token count.

The SQL functions are declared immutable and parallel safe. They can therefore be used in generated expressions or parallel plans only when the selected vocabulary files are deployed consistently across every server.

### Operational Notes

- Tokenization is model-encoding specific. Confirm both the encoding name and the model's current context limits in the application.
- Counting or chunking large text consumes backend CPU and memory; batch large corpora and monitor query latency.
- Backend-local caches avoid repeated parsing but increase memory use in sessions that touch several vocabularies.
- The upstream README's compatibility list can lag packaging. Test the exact pg_tiktoken_c build against the target PostgreSQL major version instead of inferring support from a different binary.
