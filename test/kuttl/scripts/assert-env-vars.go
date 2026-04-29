package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var expects multiFlag
	var absents multiFlag
	flag.Var(&expects, "expect", "container:VAR_NAME=VALUE;container:VAR_NAME=VALUE;...")
	flag.Var(&absents, "absent", "container:VAR_NAME;container:VAR_NAME;...")
	flag.Parse()

	expectedEnvVars := parseEnvVars(expects)
	absentEnvVars := parseEnvVars(absents)

	ss := getSts()
	logBackrestSidecarEnv(os.Stderr, ss)

	if err := checkExpectedEnvVars(os.Stderr, ss, expectedEnvVars); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := checkAbsentEnvVars(os.Stderr, ss, absentEnvVars); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func logBackrestSidecarEnv(w *os.File, ss StatefulSet) {
	for _, c := range ss.Spec.Template.Spec.Containers {
		if c.Name == "backrest-sidecar" {
			fmt.Fprintf(w, "[e2e] assert-env-vars: backrest-sidecar env snapshot: %s\n", c.Env.String())
			return
		}
	}
	fmt.Fprintf(w, "[e2e] assert-env-vars: no backrest-sidecar container in STS template (containers=%v)\n", stsContainerNames(ss))
}

func checkExpectedEnvVars(w *os.File, ss StatefulSet, expected map[string][]EnvVar) error {
	for logicalName, envVars := range expected {
		k8sName := K8sStatefulSetContainerName(logicalName)
		var match *Container
		for i := range ss.Spec.Template.Spec.Containers {
			c := &ss.Spec.Template.Spec.Containers[i]
			if c.Name == k8sName {
				match = c
				break
			}
		}
		if match == nil {
			fmt.Fprintf(w, "[e2e] assert-env-vars: missing container logical=%q k8s=%q; STS containers=%v\n",
				logicalName, k8sName, stsContainerNames(ss))
			fmt.Fprintf(w, "[e2e] hint: kubectl -n %s get pod %s -o jsonpath='{.spec.containers[*].name}'\n",
				kuttlNamespace(), "cassandra-e2e-dc1-rack1-0")
			return fmt.Errorf("container %q (%q) not found in StatefulSet template", logicalName, k8sName)
		}
		envsString := match.Env.String()
		fmt.Fprintf(w, "[e2e] assert-env-vars: expect on %s (logical %s) envs=%s\n", k8sName, logicalName, envsString)
		for _, envVar := range envVars {
			if !strings.Contains(envsString, envVar.String()) {
				return fmt.Errorf("env var missing on container %s (logical %s): want substring %q; full env list: %s",
					k8sName, logicalName, envVar.String(), envsString)
			}
		}
	}
	return nil
}

func checkAbsentEnvVars(w *os.File, ss StatefulSet, absent map[string][]EnvVar) error {
	for logicalName, envVars := range absent {
		k8sName := K8sStatefulSetContainerName(logicalName)
		var match *Container
		for i := range ss.Spec.Template.Spec.Containers {
			c := &ss.Spec.Template.Spec.Containers[i]
			if c.Name == k8sName {
				match = c
				break
			}
		}
		if match == nil {
			fmt.Fprintf(w, "[e2e] assert-env-vars: absent check skipped (no container logical=%q k8s=%q)\n", logicalName, k8sName)
			continue
		}
		envsString := match.Env.String()
		fmt.Fprintf(w, "[e2e] assert-env-vars: absent on %s (logical %s) envs=%s\n", k8sName, logicalName, envsString)
		for _, envVar := range envVars {
			if strings.Contains(envsString, envVar.String()) {
				return fmt.Errorf("env var expected absent on container %s (logical %s) still present: %q; envs: %s",
					k8sName, logicalName, envVar.String(), envsString)
			}
		}
	}
	return nil
}
