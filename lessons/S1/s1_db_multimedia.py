# S1 COURSE EXERCISE - DB-MULTIMEDIA
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Initiation aux bases de données + Technologies multimédias
#
# **Semester:** S1  
# **Attached engineering lesson:** 048 — PostgreSQL, psql, and Relational Thinking
#
# This pack contains practice material for the database and multimedia subjects attributed to S1 in the supplied checklist. Confirm the official status and exact scope of each course with the current institution syllabus.
#
# ## Part A — Introduction aux bases de données
#
# ### Scope
#
# Data versus information; database and DBMS; relation/table/row/column; primary and foreign keys; entity identification; simple ER modelling; attributes; one-to-one/one-to-many relationships; integrity constraints; relational schema; basic SQL `SELECT`, `INSERT`, `UPDATE`, `DELETE`, filtering, ordering, simple joins, grouping.
#
# ### Teaching notes
#
# A table models a relation, not a spreadsheet-shaped storage pile. Keys identify records and foreign keys express relationships. A good schema makes invalid states difficult to represent.
#
# Before writing SQL, state the intended result in one sentence: what relation should the query produce? Then choose the tables, join condition, filter, grouping, and ordering.
#
# ### Worked example
#
# Model customers, orders, and order lines. Identify entities, primary keys, foreign keys, cardinalities, then implement the schema and answer “Which customer spent the most?” and “Which products have never been ordered?”
#
# ### Exercises
#
# 1. Convert a flat customer/order table into entities.
# 2. Choose primary keys for five schemas.
# 3. Draw relationships and cardinalities for a library and a clinic.
# 4. Write basic `SELECT` filters.
# 5. Write joins across two and three tables.
# 6. Write grouped queries with `COUNT` and `SUM`.
# 7. Add constraints that block invalid references.
# 8. Predict query results by hand before running them.
# 9. Repair a deliberately denormalised schema.
# 10. Build `schema.sql`, `seed.sql`, and `queries.sql` for a small business case.
#
# ### Coding workspace
#
# Use s1_database_exercise.sql, plus the lesson's PostgreSQL setup. No answer SQL is stored in the student file.
#
# ## Part B — Technologies multimédias
#
# ### Scope
#
# Raster versus vector; pixels; dimensions; aspect ratio; colour channels; RGB; bit depth; image compression; JPEG/PNG concepts; audio sample rate/bit depth/bitrate; video frame rate/resolution/codec/container; file-size estimation; quality versus size trade-offs; basic web-media constraints and accessibility.
#
# ### Worked calculations
#
# 1. Estimate raw RGB image storage from width × height × bytes/pixel.
# 2. Estimate uncompressed PCM audio size from duration × sample rate × channels × bytes/sample.
# 3. Compare two video settings and explain the storage/bandwidth trade-off.
#
# ### Exercises
#
# 1. Convert dimensions to pixel counts and aspect ratios.
# 2. Calculate raw image size for several bit depths.
# 3. Calculate PCM audio size for different sample rates/channels.
# 4. Interpret bitrate and duration to estimate file size.
# 5. Distinguish codec from container.
# 6. Choose an image format for a logo, photograph, and transparent asset.
# 7. Explain why resizing after compression can degrade quality.
# 8. Design a media budget for a small website.
# 9. Identify accessibility issues in a media-heavy page.
# 10. Compare two delivery strategies under a fixed bandwidth limit.
#
# ### Coding workspace
#
# Use s1_db_multimedia.py for arithmetic verification only.
#
# ## Mastery gate
#
# Complete one database modelling task plus SQL queries and one multimedia calculation set with formulas shown. Explain every schema choice and every unit.

# --- Executable exercise/checking helpers ---
"""Student arithmetic verifier for multimedia calculations.

Write units and formulas in the academic workbook before using this file.
"""
from __future__ import annotations

from math import isfinite


def raw_rgb_bytes(width: int, height: int, bytes_per_pixel: int = 3) -> int:
    if width <= 0 or height <= 0 or bytes_per_pixel <= 0:
        raise ValueError("dimensions and bytes_per_pixel must be positive")
    return width * height * bytes_per_pixel


def pcm_bytes(seconds: float, sample_rate: int, channels: int, bytes_per_sample: int) -> int:
    if not isfinite(seconds) or seconds < 0:
        raise ValueError("seconds must be finite and non-negative")
    if sample_rate <= 0 or channels <= 0 or bytes_per_sample <= 0:
        raise ValueError("sample rate, channels, and bytes per sample must be positive")
    return int(seconds * sample_rate * channels * bytes_per_sample)


def main() -> None:
    print("S1 multimedia calculation verifier")
    print("Enter your own problem values after solving them by hand.")


if __name__ == "__main__":
    main()
