package main

import "github.com/zapj/zap/cmd/deamon/cmd"

var (
	BuildDate string
	// GitCommit, GitBranch string
	BuildVersion string
	Version      string
)

func main() {
	cmd.Execute()
	// fmt.Printf("Build %s-%s-%s-%s ,Version %s \n", runtime.GOOS, runtime.GOARCH, BuildVersion, BuildDate, Version)
	// syscall.Setegid(65534)
	// syscall.Seteuid(65534)
	// ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// defer stop()

	// Listen for the interrupt signal.
	// <-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	// stop()
	// log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the assist it has 5 seconds to finish
	// the request it is currently handling
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// log.Fatal("Server forced to shutdown: ", err)
	// }

	// log.Println("Server exiting")
}
