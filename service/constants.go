package service

const (
	ListConnectionsPath = "/api/v1/connector/list_connections"
	// GetUserConnectionsPath     = "/api/v1/connector/get_user_connections"
	GetConfigPath              = "/api/v1/connector/get_config"
	AddConnectionPath          = "/api/v1/connector/add_connection"
	UpdateConnectionConfigPath = "/api/v1/connector/update_connection"
	DeleteConnectionConfigPath = "/api/v1/connector/delete_connection"
	TestConnectionConfigPath   = "/api/v1/connector/test_connection"

	GetNoDBDataPath = "/api/v1/nocodedb/get_data"

	K8sNamespace       = "default"
	TrinoConfigMapName = "tcb-trino-catalog"
	TrinoCoordinator   = "tcb-trino-coordinator"
	TrinoWorker        = "tcb-trino-worker"

	TestPort = "0.0.0.0:%s"
)
