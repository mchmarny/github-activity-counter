SELECT repo, actor, count(1) as activities
FROM gheventcounter.events
WHERE event_time >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 28 DAY)
GROUP BY repo, actor
ORDER BY 3 desc