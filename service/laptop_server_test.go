package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/eswzy/grpc-learn/pb"
	"github.com/eswzy/grpc-learn/sample"
	"github.com/eswzy/grpc-learn/service"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicatedID := sample.NewLaptop()
	StoreDuplicatedID := service.NewInMemoryLaptopStore()
	err := StoreDuplicatedID.Save(laptopDuplicatedID)
	require.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_with_no_id",
			laptop: laptopNoID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failed_invalid_id",
			laptop: laptopInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failed_duplicated_id",
			laptop: laptopDuplicatedID,
			store:  StoreDuplicatedID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		// Save test case to local variable to avoid concurrency issues,
		// because we want to create multiple parallel subtests.
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // make it run in parallel with other tests.

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			// Direct call on SERVER SIDE.
			server := service.NewLaptopServer(tc.store, nil, nil)
			res, err := server.CreateLaptop(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok) // err was produced by gRPC
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
