// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package schemaexpr

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/sql/catalog"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgcode"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgerror"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/eval"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
)

// RowIndexedVarContainer is used to evaluate expressions over various rows.
type RowIndexedVarContainer struct {
	CurSourceRow tree.Datums

	// Because the rows we have might not be permuted in the same way as the
	// original table, we need to store a mapping between them.

	Cols    []catalog.Column
	Mapping catalog.TableColMap
}

var _ eval.IndexedVarContainer = &RowIndexedVarContainer{}

// IndexedVarEval implements eval.IndexedVarContainer.
func (r *RowIndexedVarContainer) IndexedVarEval(
	ctx context.Context, idx int, e tree.ExprEvaluator,
) (tree.Datum, error) {
	rowIdx, ok := r.Mapping.Get(r.Cols[idx].GetID())
	if !ok {
		return tree.DNull, nil
	}
	return r.CurSourceRow[rowIdx], nil
}

// IndexedVarResolvedType implements tree.IndexedVarContainer.
func (*RowIndexedVarContainer) IndexedVarResolvedType(idx int) *types.T {
	panic("unsupported")
}

// CannotWriteToComputedColError constructs a write error for a computed column.
func CannotWriteToComputedColError(colName string) error {
	return pgerror.Newf(pgcode.ObjectNotInPrerequisiteState,
		"cannot write directly to computed column %q", tree.ErrNameString(colName))
}
