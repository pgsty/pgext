## Usage

Sources:

- [Amazon Aurora machine learning documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-ml.html)

`aws_ml` is an Amazon Aurora PostgreSQL extension for invoking Amazon Comprehend, SageMaker AI, and Bedrock from SQL. It is an AWS-managed feature, not a portable extension for self-managed PostgreSQL, and all participating services must be available in the same AWS Region.

### Core Workflow

First attach the service-specific IAM role and policy to the Aurora cluster, using the corresponding Aurora ML feature. Then connect to the writer instance and install the extension; `CASCADE` also installs `aws_commons`.

```sql
CREATE EXTENSION IF NOT EXISTS aws_ml CASCADE;

SELECT aws_bedrock.invoke_model_get_embeddings(
  model_id := 'amazon.titan-embed-text-v1',
  content_type := 'application/json',
  json_key := 'embedding',
  model_input := '{"inputText":"PostgreSQL"}'
);
```

The Bedrock example requires extension version `2.0`, model access, and a model whose request and response match the supplied JSON. Version `1.0` exposes `aws_comprehend.detect_sentiment` and `aws_sagemaker.invoke_endpoint`; version `2.0` adds `aws_bedrock.invoke_model` and `aws_bedrock.invoke_model_get_embeddings`.

### Security and Operations

Installation creates the `aws_ml` administrative role and the `aws_comprehend`, `aws_sagemaker`, and `aws_bedrock` schemas. Public execution is revoked by default. An administrator must grant schema usage and execution on only the required functions.

Availability depends on the Aurora PostgreSQL version and Region. IAM, service quotas, network calls, model costs, payload limits, and provider upgrades remain operational dependencies; test errors and latency on the exact Aurora cluster before using calls in transactions or latency-sensitive queries.
