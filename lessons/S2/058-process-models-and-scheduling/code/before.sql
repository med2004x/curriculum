/* FIXTURE: Input data for the process scheduling lab. */
EXPLAIN (ANALYZE, BUFFERS)
SELECT id, title, created_at
FROM app.tasks
WHERE project_id = (SELECT min(id) FROM app.projects)
  AND status = 'open'
ORDER BY created_at DESC
LIMIT 20;
