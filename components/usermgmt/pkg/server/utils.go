package server

import (
	v3 "github.com/RafaySystems/rcloud-base/components/common/proto/types/commonpb/v3"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func getStatus(err error) *v3.Status {
	if err != nil {
		return &v3.Status{
			ConditionStatus: v3.ConditionStatus_StatusFailed,
			LastUpdated:     timestamppb.Now(),
			Reason:          err.Error(),
		}
	}
	return &v3.Status{
		ConditionStatus: v3.ConditionStatus_StatusOK,
	}
}
