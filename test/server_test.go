package servertest

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	pb "github.com/rprtr258/kvadotest/api"
	"github.com/rprtr258/kvadotest/internal"
	"github.com/rprtr258/kvadotest/internal/repositories"
)

func TestSearchRequests(t *testing.T) {
	type mockRecorder = *repositories.MockBooksRepositoryMockRecorder
	tests := []struct {
		name          string
		expecting     func(mockRecorder)
		searchRequest pb.SearchRequest
	}{
		{
			name:          "by author",
			expecting:     func(r mockRecorder) { r.SearchByAuthor(gomock.Any(), "author") },
			searchRequest: pb.SearchRequest{Request: &pb.SearchRequest_ByAuthor{ByAuthor: "author"}},
		}, {
			name:          "by title",
			expecting:     func(r mockRecorder) { r.SearchByTitle(gomock.Any(), "title") },
			searchRequest: pb.SearchRequest{Request: &pb.SearchRequest_ByTitle{ByTitle: "title"}},
		}, {
			name:          "by content",
			expecting:     func(r mockRecorder) { r.SearchByContent(gomock.Any(), "needle") },
			searchRequest: pb.SearchRequest{Request: &pb.SearchRequest_ByContent{ByContent: "needle"}},
		},
	}
	for i := 0; i < len(tests); i++ {
		test := &tests[i]
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			booksRepo := repositories.NewMockBooksRepository(ctrl)
			srv := internal.NewServer(booksRepo)
			ctx := context.Background()
			test.expecting(booksRepo.EXPECT())
			_, err := srv.Search(ctx, &test.searchRequest)
			if err != nil {
				t.Fatalf("search request #%d failed: %v", i, err)
			}
		})
	}
}
