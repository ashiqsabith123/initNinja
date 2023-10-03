package architecture

func DDD(project_name, language string) {
	folderStructure := []string{
		"cmd/",
		"internal/app/domain/entity",
		"internal/app/domain/valueobject",
		"internal/app/domain/repository",
		"internal/app/domain/service",
		"internal/app/domain/factory",
		"internal/app/domain/aggregate",
		"internal/app/domain/event",

		"internal/app/application/usecase",
		"internal/app/application/service",

		"internal/app/infrastructure/persistence/repository_impl",
		"internal/app/messaging",

		"internal/interfaces/http",
		"internal/interfaces/grpc",

		"tests/unit",
		"tests/integration",
		"tests/e2e",

		"scripts"



	}
}
