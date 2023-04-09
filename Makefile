gen:
	rm graph/schema.resolvers.go
	rm graph/model/models_gen.go
	rm graph/generated.go
	go run github.com/99designs/gqlgen generate
