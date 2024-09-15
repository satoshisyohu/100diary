package repository

// HealthCheckRepository ヘルスチェックリポジトリのインターフェース
type HealthCheckRepository interface {
	Ping() error
}
