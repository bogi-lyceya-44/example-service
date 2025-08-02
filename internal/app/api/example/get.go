package example

import (
	"context"

	"github.com/bogi-lyceya-44/example-service/internal/app/models"
	desc "github.com/bogi-lyceya-44/example-service/internal/pb/api/example"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bogi-lyceya-44/common/pkg/utils"
)

func (s *Implementation) Get(
	ctx context.Context,
	req *desc.ExampleRequest,
) (*desc.ExampleResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"bad request: %v",
			err,
		)
	}

	ids := utils.Map(
		req.GetIds(),
		models.NewId,
	)

	resp, err := s.exampleService.Get(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something happened: %v", err)
	}

	objs := utils.Map(
		resp,
		func(obj models.Example) *desc.ExampleModel {
			return &desc.ExampleModel{
				Id:       obj.Id.ToInt64(),
				FormedAt: timestamppb.New(obj.FormedAt),
			}
		},
	)

	return &desc.ExampleResponse{
		Objs: objs,
	}, nil
}
