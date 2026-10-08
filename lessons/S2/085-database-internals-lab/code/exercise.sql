-- STUDENT TASK: complete the constraint and descending query described in WORK.md.
CREATE TEMP TABLE buffer_events (
    page_id integer NOT NULL,
    hit_count integer NOT NULL
);

INSERT INTO buffer_events (page_id, hit_count) VALUES
    (4, 1), (7, 3), (2, 0);

-- TODO: reject negative hit_count values with a CHECK constraint.
-- TODO: select rows by hit_count in descending order.
