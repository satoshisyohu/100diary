package domain

// HealthCheckRepository ヘルスチェックリポジトリのインターフェース
type HealthCheckRepository interface {
	Ping() error
}
