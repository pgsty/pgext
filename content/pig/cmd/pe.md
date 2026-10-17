---
title: "pig pe"
description: "Read PostgreSQL metrics and reload pg_exporter configuration"
weight: 675
icon: fas fa-chart-line
---

Use `pig pe` to read metrics from a running [pg_exporter](https://github.com/Vonng/pg_exporter)
and reload its configuration. The full command name is `pig pg_exporter`.
Aliases are `pe`, `pgexp`, and `pgexporter`.

```bash
pig pe get                     # Show PostgreSQL metrics
pig pe list                    # List PostgreSQL metric families
pig pe stat                    # Show exporter statistics
pig pe reload                  # Reload exporter configuration
```

## Select an exporter

The default endpoint is `http://127.0.0.1:9630`. Use these options to select another instance:

| Option | Default | Description |
|:---|:---|:---|
| `--host` | `127.0.0.1` | Host name or IPv4 address |
| `-p, --port` | `9630` | HTTP port |
{.full-width}

```bash
pig pe --host pg-meta -p 9630 stat
pig pe --host 10.10.10.10 get
```

PIG uses HTTP and the standard `/metrics`, `/stat`, and `/reload` paths.
For HTTPS, authentication, client certificates, or a custom path, use an HTTP client such as `curl`.

## pe get

Read `/metrics`. Print samples whose names start with `pg_`, plus their Prometheus `HELP` and `TYPE` lines.

```bash
pig pe get
pig pe --host pg-meta get
```

To include process and runtime metrics that do not start with `pg_`, read the complete endpoint:

```bash
curl -fsS http://127.0.0.1:9630/metrics
```

## pe list

Read `/metrics`. Print one `HELP` line for each `pg_` metric family in the response.
The list reflects the exporter's current configuration.

```bash
pig pe list
```

## pe stat

Read `/stat` and print its response. The response format depends on the installed pg_exporter version.

```bash
pig pe stat
```

## pe reload

After editing the exporter's configuration, request a reload through `/reload`:

```bash
pig pe reload
```

## Output and errors

Use `-o json` or `-o yaml` for a PIG result that contains the command's text output.
Prometheus metrics remain text inside that result.

Network errors, response-read errors, and HTTP 4xx/5xx responses return a nonzero exit code.
HTTP error details contain the status code and standard description, without the response body.
Successful responses retain the filtering or raw text behavior described above.

To install pg_exporter, use `pig ext add pg-exporter`. Configure and start the exporter before
using `pig pe`; these commands access its HTTP API.
