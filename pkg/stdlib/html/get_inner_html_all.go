package html

import (
	"context"

	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"github.com/MontFerret/ferret/pkg/runtime/values/types"
)

// INNER_HTML_ALL returns an array of inner HTML strings of matched elements.
// @param parent (HTMLPage | HTMLDocument | HTMLElement) - Parent document or element.
// @param selector (String) - String of CSS selector.
// @return (String) - An array of inner HTML strings if all matched elements, otherwise empty array.
func GetInnerHTMLAll(ctx context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 2)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[1], types.String)

	if err != nil {
		return values.None, err
	}

	el, err := drivers.ToElement(args[0])

	if err != nil {
		return values.None, err
	}

	selector := args[1].(values.String)

	return el.GetInnerHTMLBySelectorAll(ctx, selector)
}
