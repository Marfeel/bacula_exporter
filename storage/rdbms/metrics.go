package rdbms

// ////////////////////////////////////////////////////////////////////////////////// //

// GetLatestJobs return total list of latest jobs
func (db *DB) GetLatestJobs() ([]*BaculaJob, error) {
	baculaJobs := make([]*BaculaJob, 0)

	sqlState := `
          SELECT
		t.Name,
                t.Level,
                t.JobStatus,
                unix_timestamp(t.SchedTime) as SchedTime,
                unix_timestamp(t.StartTime) as StartTime,
                unix_timestamp(t.EndTime) as EndTime,
                cast(t.JobBytes AS UNSIGNED INTEGER) as JobBytes,
                cast(t.JobFiles AS UNSIGNED INTEGER) as JobFiles
          FROM
                Job t
          INNER JOIN (
                SELECT
                      Name,
                      Level,
                      MAX(StartTime) as MaxStartTime
                FROM
                      Job
                GROUP BY
                      Name,
                      Level
                ) tm
          ON
                t.Name = tm.Name
                AND
                t.Level = tm.Level
                AND
                t.StartTime = tm.MaxStartTime
          WHERE
                t.Type = 'B'`

	err := db.Select(&baculaJobs, sqlState)

	return baculaJobs, err
}

// GetJobsSummary return summary of all jobs
func (db *DB) GetJobsSummary() ([]*BaculaJobSummary, error) {
	jobsSummary := make([]*BaculaJobSummary, 0)

	sqlState := `
          SELECT
                Name,
                Level,
                cast(SUM(JobBytes) AS UNSIGNED INTEGER) as TotalJobBytes,
                cast(SUM(JobFiles) AS UNSIGNED INTEGER) as TotalJobFiles
          FROM
                Job
          WHERE
                Name IN (
                      SELECT DISTINCT
                            Name
                      FROM
                            Job
                      WHERE
                            SchedTime = DATE(NOW())
                )
          GROUP BY
                Name,
                Level`

	err := db.Select(&jobsSummary, sqlState)

	return jobsSummary, err
}
