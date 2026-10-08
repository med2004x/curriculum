BEGIN;
UPDATE app.tasks
SET status = 'done'
WHERE id = 42;
INSERT INTO app.task_events (task_id, event_type)
VALUES (42, 'completed');
COMMIT;
