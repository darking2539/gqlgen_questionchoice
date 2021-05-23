package graph

import "github.com/darking2539/gqlgen/graph/model"

type Resolver struct{
	questions []*model.Question
	choices   []*model.Choice
}

