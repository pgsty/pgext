## Usage

Sources:

- [Official control file](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag/rag.control)
- [Official pgrag README](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [Official package manifest](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag/Cargo.toml)

`rag` is the base extension from the deprecated pgrag project. It provides local document-to-text conversion and character-based chunking, plus HTTPS clients for several remote embedding, reranking, and chat APIs. The project is no longer developed or maintained, so use it only with an explicit migration and security plan.

### Core Workflow

Install its declared `vector` dependency, then use the local conversion helpers without a model extension:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rag;

SELECT rag.markdown_from_html(
    '<html><body><h1>Title</h1><p>A short paragraph.</p></body></html>'
);

SELECT rag.chunks_by_character_count(
    'The quick brown fox jumps over the lazy dog', 20, 4
);
```

### Main API Groups

- Local conversion: `rag.markdown_from_html(text)`, `rag.text_from_pdf(bytea)`, and `rag.text_from_docx(bytea)`. PDF extraction has no OCR or complex-layout support.
- Local chunking: `rag.chunks_by_character_count(text, integer, integer)` returns overlapping text chunks.
- OpenAI: `rag.openai_set_api_key(text)`, `rag.openai_get_api_key()`, embedding helpers, and `rag.openai_chat_completion(json)`.
- Anthropic and Fireworks: matching key setters/getters plus `rag.anthropic_messages(text, json)`, Fireworks embedding helpers, and `rag.fireworks_chat_completion(json)`.
- Voyage AI: key setters/getters, embedding helpers, and scalar/array reranking score and distance functions.

### Operational Notes

The reviewed version is 0.0.0, non-relocatable, and requires `vector`; the control file also sets `superuser = true`. It declares no preload requirement. Remote helpers perform outbound HTTPS requests inside database sessions, so latency, provider failures, quotas, and response handling become query concerns. Key getter functions return stored credentials; restrict function privileges and never expose them to untrusted roles. Pin dependencies because the upstream repository is archived and deprecated.
