package mock_repository

//go:generate mockgen -destination=recipe_repository_mock.go github.com/dbiagi/gororoba/src/repository  RecipeRepositoryInterface
