package config

import (
	"fmt"
	stdlog "log"
	"regexp"
	"strings"

	"github.com/go-kit/kit/log/level"
	flag "github.com/spf13/pflag"
)

var (
	validRule    = regexp.MustCompile(`^[_A-Za-z][\w]*$`)
	validPackage = regexp.MustCompile(`^[_A-Za-z][\w]*(\.[_A-Za-z][\w]*)*$`)
)

type Config struct {
	KubeconfigPath string
	DebugToken     string
	Name           string
	Mappings       map[string]string

	LogFormat string
	LogLevel  level.Option

	Opa    OPAConfig
	Server ServerConfig
}

type OPAConfig struct {
	Pkg     string
	Rule    string
	Matcher string
}

type ServerConfig struct {
	Listen         string
	ListenInternal string
	HealthcheckURL string
}

func ParseFlags() (*Config, error) {
	cfg := &Config{}
	// Logger flags
	flag.StringVar(&cfg.Name, "debug.name", "opa-openshift", "A name to add as a prefix to log lines.")
	logLevelRaw := flag.String("log.level", "info", "The log filtering level. Options: 'error', 'warn', 'info', 'debug'.")
	flag.StringVar(&cfg.LogFormat, "log.format", "logfmt", "The log format to use. Options: 'logfmt', 'json'.")

	// Server flags
	flag.StringVar(&cfg.Server.Listen, "web.listen", ":8080", "The address on which the public server listens.")
	flag.StringVar(&cfg.Server.ListenInternal, "web.internal.listen", ":8081", "The address on which the internal server listens.")
	flag.StringVar(&cfg.Server.HealthcheckURL, "web.healthchecks.url", "http://localhost:8080", "The URL against which to run healthchecks.")

	// OpenShift API flags
	flag.StringVar(&cfg.KubeconfigPath, "openshift.kubeconfig", "", "A path to the kubeconfig against to use for authorizing client requests.")
	mappingsRaw := flag.StringSlice("openshift.mappings", nil, "A map of tenantIDs to resource api groups to check to apply a given role to a user, e.g. tenant-a=observatorium.openshift.io") //nolint:lll

	// OPA flags
	flag.StringVar(&cfg.Opa.Pkg, "opa.package", "", "The name of the OPA package that opa-openshift should implement, see https://www.openpolicyagent.org/docs/latest/policy-language/#packages.")          //nolint:lll
	flag.StringVar(&cfg.Opa.Rule, "opa.rule", "allow", "The name of the OPA rule for which opa-openshift should provide a result, see https://www.openpolicyagent.org/docs/latest/policy-language/#rules.") //nolint:lll
	flag.StringVar(&cfg.Opa.Matcher, "opa.matcher", "", "The label key of the OPA label matcher returned to the requesting client.")

	// Integration testing flags
	flag.StringVar(&cfg.DebugToken, "debug.token", "", "Debug bearer token used for integration tests.")

	err := flag.CommandLine.MarkHidden("debug.token")
	if err != nil {
		stdlog.Fatal("failed to mark flag hidden")
	}

	flag.Parse()

	ll, err := parseLogLevel(logLevelRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	cfg.LogLevel = ll

	if len(cfg.Opa.Pkg) > 0 && !validPackage.Match([]byte(cfg.Opa.Pkg)) {
		return nil, fmt.Errorf("invalid OPA package name: %s", cfg.Opa.Pkg) //nolint:goerr113
	}

	if len(cfg.Opa.Rule) > 0 && !validRule.Match([]byte(cfg.Opa.Rule)) {
		return nil, fmt.Errorf("invalid OPA rule name: %s", cfg.Opa.Rule) //nolint:goerr113
	}

	if *mappingsRaw == nil {
		stdlog.Fatal("missing tenant mappings")
	}

	cfg.Mappings = make(map[string]string)

	for _, m := range *mappingsRaw {
		parts := strings.Split(m, "=")
		if len(parts) != 2 { //nolint:gomnd
			return nil, fmt.Errorf("invalid mapping: %q", m) //nolint:goerr113
		}

		cfg.Mappings[parts[0]] = parts[1]
	}

	return cfg, nil
}

func parseLogLevel(logLevelRaw *string) (level.Option, error) {
	switch *logLevelRaw {
	case "error":
		return level.AllowError(), nil
	case "warn":
		return level.AllowWarn(), nil
	case "info":
		return level.AllowInfo(), nil
	case "debug":
		return level.AllowDebug(), nil
	default:
		return nil, fmt.Errorf("unexpected log level: %s", *logLevelRaw) //nolint:goerr113
	}
}
