## Usage

Sources:

- [Official database.dev package page](https://database.dev/langchain/embedding_search)

`embedding_search` is the canonical catalog name for the database.dev package `langchain@embedding_search`. It installs the Supabase/PostgreSQL objects expected by the LangChain vector-store example, including document storage and similarity matching backed by `vector`. The registry's latest package release is 1.1.1, while its published extension installation example and default extension version remain 1.1.0.

### Install the Package

Install the namespaced package with `dbdev.install`, create the dependency explicitly, and use the namespaced extension name rather than `embedding_search` in `CREATE EXTENSION`:

```sql
SELECT dbdev.install('langchain@embedding_search');
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION "langchain@embedding_search"
  SCHEMA public
  VERSION '1.1.0';
```

Dependency resolution is not automatic in the documented workflow, so `vector` must already be available in the database.

### Core Workflow

The package creates a `documents` table and a `match_documents` function with the shape used by the linked LangChain/Supabase integration. The client embeds documents, stores content and metadata in Supabase, and invokes the matching function for vector similarity search; the example also demonstrates metadata filtering.

Use the exact table and function signatures installed by the selected package release rather than assuming that examples from another release match. LangChain import paths shown on the package page reflect an older JavaScript client generation, so align client code with the LangChain version in the application.

### Security and Operations

The example uses a Supabase service-role key. Keep that key and embedding-provider credentials in server-side secret storage, never browser code or SQL files. Apply row-level security and least-privilege grants to `documents` and `match_documents`, define who may insert embeddings or read metadata, and test index choice, embedding dimension, filter selectivity, and migration behavior with production-scale data.
