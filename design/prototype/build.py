#!/usr/bin/env python3
"""Assemble the PGEXT.CLOUD prototype from parts.

Reads style.css / app.js / shell.html plus the data snapshot under data/,
and produces two flavors of the same page:

  index.html    : full standalone document — open locally or deploy anywhere
  artifact.html : content-only fragment for Claude Artifact publishing

To refresh the data snapshot from the local catalog database, see README.md.
"""
import json
import datetime
import pathlib

HERE = pathlib.Path(__file__).parent
DATA = HERE / "data" if (HERE / "data").is_dir() else HERE

css = (HERE / "style.css").read_text()
js = (HERE / "app.js").read_text()
shell = (HERE / "shell.html").read_text()
ext = (DATA / "ext.json").read_text()
cats = (DATA / "cat.json").read_text()
docs = (DATA / "docs.json").read_text()
meta = json.dumps({"generated": datetime.date.today().isoformat()})

TITLE = "PGEXT.CLOUD — The PostgreSQL Extension Catalog"

data_block = (
    "<script>window.__EXT__=" + ext.strip()
    + ";window.__CATS__=" + cats.strip()
    + ";window.__DOCS__=" + docs.strip()
    + ";window.__META__=" + meta
    + ";</script>"
)

body = (
    "<title>" + TITLE + "</title>\n"
    + "<style>\n" + css + "\n</style>\n"
    + shell.strip() + "\n"
    + data_block + "\n"
    + "<script>\n" + js + "\n</script>\n"
)

(HERE / "artifact.html").write_text(body)

standalone = (
    "<!doctype html>\n<html lang=\"en\">\n<head>\n"
    "<meta charset=\"utf-8\">\n"
    "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n"
    "<meta name=\"description\" content=\"The catalog of the PostgreSQL extension universe\">\n"
    "<link rel=\"icon\" href=\"data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='0.9em' font-size='90'>🐘</text></svg>\">\n"
    "</head>\n<body>\n"
    + body
    + "</body>\n</html>\n"
)
(HERE / "index.html").write_text(standalone)

print("artifact.html", len(body.encode()) // 1024, "KiB")
print("index.html   ", len(standalone.encode()) // 1024, "KiB")
