

## Usage

> [pg_summarize](https://github.com/HexaCluster/pg_summarize): Text Summarization using LLMs, built using pgrx.
> Source: [README.md](https://raw.githubusercontent.com/HexaCluster/pg_summarize/main/README.md)

`pg_summarize` is a PostgreSQL extension written in Rust (using `pgrx`) that integrates with the OpenAI API. It includes a basic "Hello, pg_summarize!" function and a `summarize` function that summarizes text using OpenAI's models.


--------

### Getting Started

```sql
CREATE EXTENSION pg_summarize;

-- Test the hello function
SELECT hello_pg_summarize();
--  hello_pg_summarize
-- ----------------------
--  Hello, pg_summarize
```


--------

### Configuration

The extension retrieves configuration from PostgreSQL settings. Set the following before using the `summarize` function:

```sql
-- Set the OpenAI API key (required)
ALTER SYSTEM SET pg_summarizer.api_key = 'your_openai_api_key';

-- Optionally set the model (default: gpt-3.5-turbo)
ALTER SYSTEM SET pg_summarizer.model = 'gpt-3.5-turbo';

-- Or set the prompt at session level
SET pg_summarizer.prompt = 'Your custom prompt here';

-- Reload the configuration if set at SYSTEM level
SELECT pg_reload_conf();
```


--------

### Summarize Function

The `summarize` function takes text input, sends it to the OpenAI API, and returns a summary:

```sql
-- Summarize a text input
SELECT summarize('<This is the text to be summarized.>');

-- Create a summary table from existing data
CREATE TABLE blogs_summary AS
  SELECT blog_url, summarize(blogs_text)
  FROM hexacluster_blogs;

-- Use a different model
SET pg_summarizer.model = 'gpt-4o';
CREATE TABLE blogs_summary_4o AS
  SELECT blog_url, summarize(blogs_text)
  FROM hexacluster_blogs;
```


--------

### How It Works

- **Configuration Retrieval**: The `summarize` function retrieves settings (API key, model, prompt) from PostgreSQL using `current_setting()`. Defaults are used if settings are not found.
- **Default Prompt**: A built-in prompt instructs the AI to summarize text from `<text>` tags, focusing on capturing the most important information concisely.
- **API Call**: The function sends a POST request to the OpenAI chat completions endpoint with the configured model and prompt, returning the summary content.
