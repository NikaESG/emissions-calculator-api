package service

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestPatchTrinoCatalogData(t *testing.T) {
	InitK8sClient()

}

func TestExecPod(t *testing.T) {
	InitK8sClient()
	r, e := ExecPod(TrinoConfigMapName, "ls")
	assert.Equal(t, e, nil)
	fmt.Println(r)
}

func TestRestartDeployments(t *testing.T) {
	InitK8sClient()
	e := RestartDeployments([]string{
		TrinoWorker,
		TrinoCoordinator,
	})
	assert.Equal(t, e, nil)
}
