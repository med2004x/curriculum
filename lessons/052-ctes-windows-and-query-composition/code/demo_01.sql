WITH ranked_events AS (
	SELECT task_id, status, changed_at,
		   ROW_NUMBER() OVER (
			   PARTITION BY task_id
			   ORDER BY changed_at DESC, id DESC
		   ) AS row_num
	FROM app.task_events
)
SELECT task_id, status, changed_at
FROM ranked_events
WHERE row_num = 1
ORDER BY task_id;
