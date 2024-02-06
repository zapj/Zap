package task

type TaskCommand struct {
	Title     string   `json:"title"`
	Cmd       string   `json:"cmd"`
	TargetDir string   `json:"target_dir"`
	LogFile   string   `json:"log_file"`
	Args      []string `json:"args"`
	Env       []string `json:"env"`
}
