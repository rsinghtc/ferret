package expressions

import (
	"context"

	"github.com/MontFerret/ferret/pkg/runtime/collections"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

type ForWhileIterableExpression struct {
	src         core.SourceMap
	valVariable string
	condition   core.Expression
}

func NewForWhileIterableExpression(
	src core.SourceMap,
	valVariable string,
	condition core.Expression,
) (collections.Iterable, error) {
	if condition == nil {
		return nil, core.Error(core.ErrMissedArgument, "condition")
	}

	return &ForWhileIterableExpression{
		src:         src,
		valVariable: valVariable,
		condition:   condition,
	}, nil
}

func (iterable *ForWhileIterableExpression) Iterate(_ context.Context, _ *core.Scope) (collections.Iterator, error) {
	return collections.NewWhileIterator(iterable.valVariable, func(ctx context.Context, scope *core.Scope) (bool, error) {
		res, err := iterable.condition.Exec(ctx, scope)

		if err != nil {
			return false, err
		}

		return res == values.True, nil
	})
}
