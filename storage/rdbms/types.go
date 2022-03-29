package rdbms

// ////////////////////////////////////////////////////////////////////////////////// //

type BaculaJob struct {
	Name      string `db:"Name"`
	Level     string `db:"Level"`
	Status    string `db:"JobStatus"`
	SchedTime uint32 `db:"SchedTime"`
	StartTime uint32 `db:"StartTime"`
	EndTime   uint32 `db:"EndTime"`
	JobBytes  uint64 `db:"JobBytes"`
	JobFiles  uint64 `db:"JobFiles"`
}

type BaculaJobSummary struct {
	Name          string `db:"Name"`
	Level         string `db:"Level"`
	TotalJobBytes uint64 `db:"TotalJobBytes"`
	TotalJobFiles uint64 `db:"TotalJobFiles"`
}

type BaculaSummary struct {
	ScheduledJobs uint32 `db:"ScheduledJobs"`
}

// ////////////////////////////////////////////////////////////////////////////////// //
