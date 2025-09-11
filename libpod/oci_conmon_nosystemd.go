//go:build !remote && (linux || freebsd) && !systemd

package libpod

// addHealthCheckArgs adds healthcheck-related arguments to conmon for non-systemd builds
func (r *ConmonOCIRuntime) addHealthCheckArgs(ctr *Container, args []string) []string {
	// Add healthcheck flag if container has healthcheck config (non-systemd builds only)
	// For systemd builds, healthchecks are managed by systemd, not conmon
	if ctr.HasHealthCheck() {
		args = append(args, "--enable-healthcheck")
	}
	return args
}
