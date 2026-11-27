"""Shared Hugo root option for catalog generation and release-note ingestion."""

import os
from pathlib import Path


def add_site_dir(parser):
    parser.add_argument(
        "-D", "--dir", type=Path,
        default=Path(os.environ.get("PGEXT_SITE_DIR") or "~/pgsty/pgext.cloud"),
        help="Hugo catalog root (PGEXT_SITE_DIR or ~/pgsty/pgext.cloud)",
    )
