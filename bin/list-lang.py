#!/usr/bin/env python3
"""Generate language index pages for Hugo content."""

from __future__ import annotations

import argparse
import json
import os
from collections import defaultdict
from dataclasses import dataclass, field
from pathlib import Path
from typing import Any, Dict, Iterable, List, Optional, Tuple

import psycopg2


# ---------------------------------------------------------------------------
# Data structures & helpers
# ---------------------------------------------------------------------------

LANGUAGE_DESCRIPTIONS_EN: Dict[str, str] = {
    "C": "The traditional PostgreSQL extension language",
    "C++": "Extensions leveraging C++ features and libraries",
    "Rust": "Extensions written in Rust with the pgrx framework",
    "Python": "Extensions written in Python",
    "SQL": "Pure SQL extensions and functions",
    "Java": "Extensions running on JVM",
    "Data": "Data-only extensions",
}

LANGUAGE_DESCRIPTIONS_ZH: Dict[str, str] = {
    "C": "传统的 PostgreSQL 扩展开发语言",
    "C++": "使用 C++ 特性和库的扩展",
    "Rust": "使用 pgrx 框架用 Rust 编写的扩展",
    "Python": "使用 Python 编写的扩展",
    "SQL": "纯 SQL 扩展和函数",
    "Java": "在 JVM 上运行的扩展",
    "Data": "仅包含数据的扩展",
}

NAV_LANGUAGE_ORDER: Tuple[str, ...] = ("C", "C++", "Rust", "Java", "Python", "SQL", "Data")


@dataclass
class ExtensionRow:
    """Subset of extension metadata required for language index pages."""

    ext_id: int
    name: str
    pkg: str
    lead_ext: str
    lang: str
    en_desc: Optional[str]
    zh_desc: Optional[str]
    extra: Dict[str, Any] = field(default_factory=dict)


# ---------------------------------------------------------------------------
# Shortcode helpers
# ---------------------------------------------------------------------------
def parse_extra(raw: Any) -> Dict[str, Any]:
    """Parse extra column into a dictionary."""
    if not raw:
        return {}
    if isinstance(raw, dict):
        return raw
    if isinstance(raw, (bytes, bytearray, memoryview)):
        try:
            return json.loads(bytes(raw).decode("utf-8"))
        except Exception:
            return {}
    if isinstance(raw, str):
        try:
            return json.loads(raw)
        except Exception:
            return {}
    return {}


def ext_shortcode(ext_name: str, display: Optional[str] = None) -> str:
    """Render the ext shortcode with optional custom display text."""
    safe_ext = ext_name.replace('"', '\\"')
    if display and display != ext_name:
        safe_display = display.replace('"', '\\"')
        return f'{{{{< ext "{safe_ext}" "{safe_display}" >}}}}'
    return f'{{{{< ext "{safe_ext}" >}}}}'


def language_shortcode(lang: str) -> str:
    """Render the language shortcode."""
    safe_lang = lang.replace('"', '\\"')
    return f'{{{{< language "{safe_lang}" >}}}}'


def badge_shortcode(text: str, color: str = "gray", icon: Optional[str] = None) -> str:
    """Render a badge shortcode."""
    safe_text = text.replace('"', '\\"')
    safe_color = color.replace('"', '\\"')
    parts = [f'content="{safe_text}"', f'color="{safe_color}"']
    if icon:
        parts.append(f'icon="{icon}"')
    args = " ".join(parts)
    return f"{{{{< badge {args} >}}}}"


def sanitize(text: Optional[str], fallback: str) -> str:
    """Prepare text for table cells."""
    value = (text or fallback).replace("|", r"\|").replace("\n", " ")
    return " ".join(value.split())


# ---------------------------------------------------------------------------
# Data loading
# ---------------------------------------------------------------------------

def fetch_extensions(pgurl: str) -> List[ExtensionRow]:
    """Fetch extension rows required for language pages."""
    query = """
        SELECT
            id,
            name,
            pkg,
            COALESCE(lead_ext, name) AS lead_ext,
            lang,
            en_desc,
            zh_desc,
            extra
        FROM pgext.extension
        WHERE lang IS NOT NULL AND lang <> ''
        ORDER BY lang, id
    """

    with psycopg2.connect(pgurl) as conn, conn.cursor() as cur:
        cur.execute(query)
        rows = cur.fetchall()

    return [
        ExtensionRow(
            ext_id=row[0],
            name=row[1],
            pkg=row[2],
            lead_ext=row[3],
            lang=row[4],
            en_desc=row[5],
            zh_desc=row[6],
            extra=parse_extra(row[7]),
        )
        for row in rows
    ]


# ---------------------------------------------------------------------------
# Rendering
# ---------------------------------------------------------------------------

def render_summary_table(
    lang_order: Iterable[Tuple[str, List[ExtensionRow]]],
    locale: str,
) -> str:
    """Render the summary table for a locale."""
    en_headers = (
        "| Language | Count | Description |",
        "|:-------:|:-----:|:------------|",
    )
    zh_headers = (
        "| 语言 | 数量 | 描述 |",
        "|:-------:|:-----:|:------------|",
    )

    if locale == "zh":
        lines = [zh_headers[0], zh_headers[1]]
    else:
        lines = [en_headers[0], en_headers[1]]

    for lang, extensions in lang_order:
        badge = language_shortcode(lang)
        if locale == "zh":
            desc = sanitize(
                LANGUAGE_DESCRIPTIONS_ZH.get(lang, f"使用 {lang} 编写的扩展"),
                f"使用 {lang} 编写的扩展",
            )
        else:
            desc = sanitize(
                LANGUAGE_DESCRIPTIONS_EN.get(lang, f"Extensions written in {lang}"),
                f"Extensions written in {lang}",
            )
        lines.append(f"| {badge} | {len(extensions)} | {desc} |")

    return "\n".join(lines)


def render_language_table(extensions: List[ExtensionRow], locale: str, lang: str) -> str:
    """Render the per-language extensions table for a locale."""
    is_rust = lang.lower() == "rust"

    if locale == "zh":
        if is_rust:
            header = (
                "| ID | 扩展 | 包 | PGRX Ver | 描述 |",
                "|:---:|:---|:---|:---|:---|",
            )
        else:
            header = (
                "| ID | 扩展 | 包 | 描述 |",
                "|:---:|:---|:---|:---|",
            )
    else:
        if is_rust:
            header = (
                "| ID | Extension | Package | PGRX Ver | Description |",
                "|:---:|:---|:---|:---|:---|",
            )
        else:
            header = (
                "| ID | Extension | Package | Description |",
                "|:---:|:---|:---|:---|",
            )

    rows = [header[0], header[1]]
    for ext in extensions:
        pkg_display = ext.pkg or ext.name
        if locale == "zh":
            description = sanitize(ext.zh_desc or ext.en_desc, "暂无描述")
        else:
            description = sanitize(ext.en_desc, "No description")
        if is_rust:
            pgrx_ver = "-"
            if isinstance(ext.extra, dict):
                pgrx_value = ext.extra.get("pgrx")
                if pgrx_value:
                    pgrx_ver = sanitize(str(pgrx_value), "-")
            rows.append(
                f"| {ext.ext_id} | {ext_shortcode(ext.name)} | "
                f"{ext_shortcode(ext.lead_ext, pkg_display)} | {pgrx_ver} | {description} |"
            )
        else:
            rows.append(
                f"| {ext.ext_id} | {ext_shortcode(ext.name)} | "
                f"{ext_shortcode(ext.lead_ext, pkg_display)} | {description} |"
            )

    return "\n".join(rows)


def render_language_sections(
    lang_order: Iterable[Tuple[str, List[ExtensionRow]]],
    locale: str,
) -> str:
    """Render all per-language sections for a locale."""
    sections: List[str] = []
    for lang, extensions in lang_order:
        if locale == "zh":
            count_badge = badge_shortcode(f"{len(extensions)} 个扩展", icon="box")
            desc = LANGUAGE_DESCRIPTIONS_ZH.get(lang, f"使用 {lang} 编写的扩展")
        else:
            count_badge = badge_shortcode(f"{len(extensions)} Extensions", icon="box")
            desc = LANGUAGE_DESCRIPTIONS_EN.get(lang, f"Extensions written in {lang}")
        section = (
            f"\n## {lang}\n\n"
            f"{language_shortcode(lang)} {count_badge}\n\n"
            f"{desc}\n\n"
            f"{render_language_table(extensions, locale, lang)}\n"
        )
        sections.append(section)
    return "\n".join(sections)


def build_nav(lang_order: Iterable[Tuple[str, List[ExtensionRow]]]) -> str:
    """Render navigation badges."""
    ordered = []
    seen = set()

    for lang in NAV_LANGUAGE_ORDER:
        seen.add(lang)
        ordered.append(lang)

    for lang, _ in lang_order:
        if lang not in seen:
            ordered.append(lang)
            seen.add(lang)

    badges = [language_shortcode(lang) for lang in ordered if any(lang == item[0] for item in lang_order)]
    return " ".join(badges)


def generate_language_pages(pgurl: str, output_dir: Path) -> None:
    """Generate both English and Chinese language index pages."""
    extensions = fetch_extensions(pgurl)
    if not extensions:
        print("No extensions with language metadata found; skipping language index generation.")
        return

    grouped: Dict[str, List[ExtensionRow]] = defaultdict(list)
    for ext in extensions:
        grouped[ext.lang].append(ext)

    # Sort languages by count desc, then name
    lang_order = sorted(grouped.items(), key=lambda item: (-len(item[1]), item[0].lower()))
    for _, items in lang_order:
        items.sort(key=lambda ext: (ext.ext_id, ext.name.lower()))

    output_dir.mkdir(parents=True, exist_ok=True)

    # English page
    english_content = (
        "---\n"
        'title: "By Language"\n'
        'description: "PostgreSQL extensions organized by implementation language"\n'
        "---\n\n"
        f"{build_nav(lang_order)}\n\n"
        "## Summary\n\n"
        f"{render_summary_table(lang_order, 'en')}\n\n"
        f"{render_language_sections(lang_order, 'en')}\n"
    )
    (output_dir / "lang.md").write_text(english_content, encoding="utf-8")
    print(f"Generated: {(output_dir / 'lang.md')}")

    # Chinese page
    chinese_content = (
        "---\n"
        'title: "按语言"\n'
        'description: "按实现语言组织的 PostgreSQL 扩展"\n'
        "---\n\n"
        f"{build_nav(lang_order)}\n\n"
        "## 概览\n\n"
        f"{render_summary_table(lang_order, 'zh')}\n\n"
        f"{render_language_sections(lang_order, 'zh')}\n"
    )
    (output_dir / "lang.zh.md").write_text(chinese_content, encoding="utf-8")
    print(f"Generated: {(output_dir / 'lang.zh.md')}")


# ---------------------------------------------------------------------------
# CLI
# ---------------------------------------------------------------------------

def resolve_output_dir(value: Optional[str]) -> Path:
    """Resolve the output directory argument."""
    if value:
        return Path(value).expanduser().resolve()
    return Path(__file__).resolve().parent.parent / "content" / "list"


def main() -> None:
    parser = argparse.ArgumentParser(description="Generate /list/lang Hugo index pages.")
    parser.add_argument("pgurl", nargs="?", default=None, help="PostgreSQL connection URL (defaults to $PGURL or postgres:///)")
    parser.add_argument("output_dir", nargs="?", default=None, help="Output directory (defaults to content/list)")
    args = parser.parse_args()

    pgurl = args.pgurl or os.getenv("PGURL") or "postgres:///"
    output_dir = resolve_output_dir(args.output_dir)

    generate_language_pages(pgurl, output_dir)


if __name__ == "__main__":
    main()
