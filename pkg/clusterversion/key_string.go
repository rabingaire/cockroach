// Code generated by "stringer -type=Key"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Start21_1-0]
	_ = x[replacedTruncatedAndRangeAppliedStateMigration-1]
	_ = x[replacedPostTruncatedAndRangeAppliedStateMigration-2]
	_ = x[TruncatedAndRangeAppliedStateMigration-3]
	_ = x[PostTruncatedAndRangeAppliedStateMigration-4]
	_ = x[V21_1-5]
	_ = x[Start21_1PLUS-6]
	_ = x[Start21_2-7]
	_ = x[JoinTokensTable-8]
	_ = x[AcquisitionTypeInLeaseHistory-9]
	_ = x[SerializeViewUDTs-10]
	_ = x[ExpressionIndexes-11]
	_ = x[DeleteDeprecatedNamespaceTableDescriptorMigration-12]
	_ = x[FixDescriptors-13]
	_ = x[SQLStatsTable-14]
	_ = x[DatabaseRoleSettings-15]
	_ = x[TenantUsageTable-16]
	_ = x[SQLInstancesTable-17]
	_ = x[NewRetryableRangefeedErrors-18]
	_ = x[AlterSystemWebSessionsCreateIndexes-19]
	_ = x[SeparatedIntentsMigration-20]
	_ = x[PostSeparatedIntentsMigration-21]
	_ = x[RetryJobsWithExponentialBackoff-22]
	_ = x[RecordsBasedRegistry-23]
	_ = x[AutoSpanConfigReconciliationJob-24]
	_ = x[PreventNewInterleavedTables-25]
	_ = x[EnsureNoInterleavedTables-26]
	_ = x[DefaultPrivileges-27]
	_ = x[ZonesTableForSecondaryTenants-28]
	_ = x[UseKeyEncodeForHashShardedIndexes-29]
	_ = x[DatabasePlacementPolicy-30]
	_ = x[GeneratedAsIdentity-31]
	_ = x[OnUpdateExpressions-32]
	_ = x[SpanConfigurationsTable-33]
	_ = x[BoundedStaleness-34]
	_ = x[SQLStatsCompactionScheduledJob-35]
	_ = x[DateAndIntervalStyle-36]
	_ = x[PebbleFormatVersioned-37]
	_ = x[MarkerDataKeysRegistry-38]
	_ = x[PebbleSetWithDelete-39]
}

const _Key_name = "Start21_1replacedTruncatedAndRangeAppliedStateMigrationreplacedPostTruncatedAndRangeAppliedStateMigrationTruncatedAndRangeAppliedStateMigrationPostTruncatedAndRangeAppliedStateMigrationV21_1Start21_1PLUSStart21_2JoinTokensTableAcquisitionTypeInLeaseHistorySerializeViewUDTsExpressionIndexesDeleteDeprecatedNamespaceTableDescriptorMigrationFixDescriptorsSQLStatsTableDatabaseRoleSettingsTenantUsageTableSQLInstancesTableNewRetryableRangefeedErrorsAlterSystemWebSessionsCreateIndexesSeparatedIntentsMigrationPostSeparatedIntentsMigrationRetryJobsWithExponentialBackoffRecordsBasedRegistryAutoSpanConfigReconciliationJobPreventNewInterleavedTablesEnsureNoInterleavedTablesDefaultPrivilegesZonesTableForSecondaryTenantsUseKeyEncodeForHashShardedIndexesDatabasePlacementPolicyGeneratedAsIdentityOnUpdateExpressionsSpanConfigurationsTableBoundedStalenessSQLStatsCompactionScheduledJobDateAndIntervalStylePebbleFormatVersionedMarkerDataKeysRegistryPebbleSetWithDelete"

var _Key_index = [...]uint16{0, 9, 55, 105, 143, 185, 190, 203, 212, 227, 256, 273, 290, 339, 353, 366, 386, 402, 419, 446, 481, 506, 535, 566, 586, 617, 644, 669, 686, 715, 748, 771, 790, 809, 832, 848, 878, 898, 919, 941, 960}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
