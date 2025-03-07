// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package sql

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/eval"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
)

// filterNode implements a filtering stage. It is intended to be used
// during plan optimizations in order to avoid instantiating a fully
// blown selectTopNode/renderNode pair.
type filterNode struct {
	source      planDataSource
	filter      tree.TypedExpr
	ivarHelper  tree.IndexedVarHelper
	reqOrdering ReqOrdering
}

// filterNode implements eval.IndexedVarContainer
var _ eval.IndexedVarContainer = &filterNode{}

// IndexedVarEval implements the eval.IndexedVarContainer interface.
func (f *filterNode) IndexedVarEval(
	ctx context.Context, idx int, e tree.ExprEvaluator,
) (tree.Datum, error) {
	return f.source.plan.Values()[idx].Eval(ctx, e)
}

// IndexedVarResolvedType implements the tree.IndexedVarContainer interface.
func (f *filterNode) IndexedVarResolvedType(idx int) *types.T {
	return f.source.columns[idx].Typ
}

func (f *filterNode) startExec(runParams) error {
	return nil
}

// Next implements the planNode interface.
func (f *filterNode) Next(params runParams) (bool, error) {
	panic("filterNode cannot be run in local mode")
}

func (f *filterNode) Values() tree.Datums {
	panic("filterNode cannot be run in local mode")
}

func (f *filterNode) Close(ctx context.Context) { f.source.plan.Close(ctx) }
