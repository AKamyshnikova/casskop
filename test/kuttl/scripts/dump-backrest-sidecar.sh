#!/usr/bin/env bash
# Non-fatal diagnostics for kuttl e2e when debugging backrest-sidecar (backup/restore, env vars, etc.).
# Intended to run on every TestAssert retry so CI logs capture sidecar state while resources converge.
set +e

NS="${NAMESPACE:-default}"
POD="${E2E_CASSANDRA_POD:-cassandra-e2e-dc1-rack1-0}"
CONTAINER="${E2E_BACKREST_CONTAINER:-backrest-sidecar}"

echo "=== [e2e] dump-backrest-sidecar ns=${NS} pod=${POD} container=${CONTAINER} ==="
kubectl get pod -n "${NS}" "${POD}" -o wide 2>&1
echo "=== [e2e] pod events (recent) ==="
kubectl get events -n "${NS}" --field-selector "involvedObject.name=${POD}" --sort-by='.lastTimestamp' 2>&1 | tail -n 25
echo "=== [e2e] containerStatuses ==="
kubectl get pod -n "${NS}" "${POD}" -o jsonpath='{range .status.containerStatuses[*]}{.name}{"\t"}{"ready="}{.ready}{"\t"}{"restarts="}{.restartCount}{"\t"}{.state}{"\n"}{end}' 2>&1
echo ""
echo "=== [e2e] describe pod (backrest / tail) ==="
kubectl describe pod -n "${NS}" "${POD}" 2>&1 | sed -n '/backrest-sidecar/,/^$/p' | head -n 80
echo "=== [e2e] logs ${CONTAINER} --previous (tail 80) ==="
kubectl logs -n "${NS}" "${POD}" -c "${CONTAINER}" --tail=80 --previous 2>&1
echo "=== [e2e] logs ${CONTAINER} (tail 120) ==="
kubectl logs -n "${NS}" "${POD}" -c "${CONTAINER}" --tail=120 2>&1

exit 0
