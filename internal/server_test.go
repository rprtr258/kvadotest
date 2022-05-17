// Unit tests for protobuf server
package internal

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rprtr258/kvadotest/internal/repositories"
	pb "github.com/rprtr258/kvadotest/pkg/api"
)

// TestSearchRequests checks that requests trigger according repository methods with according parameters
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
			// Init mock book repository
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			booksRepo := repositories.NewMockBooksRepository(ctrl)

			// Init app server
			srv := NewServer(booksRepo)

			// Check request handling
			ctx := context.Background()
			test.expecting(booksRepo.EXPECT())
			_, err := srv.Search(ctx, &test.searchRequest)
			if err != nil {
				t.Fatalf("search request #%d failed: %v", i, err)
			}
		})
	}
}

// TestSearchResponse checks that book from repository gets to response
func TestSearchResponse(t *testing.T) {
	type mockRecorder = *repositories.MockBooksRepositoryMockRecorder
	// Init mock book repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	booksRepo := repositories.NewMockBooksRepository(ctrl)

	// Init app server
	srv := NewServer(booksRepo)

	// Check request handling
	ctx := context.Background()
	book := repositories.Book{
		Title:   "title",
		Content: "content",
		Authors: []string{"author"},
	}
	booksRepo.EXPECT().SearchByContent(gomock.Any(), "needle").Return([]repositories.Book{book}, nil)
	response, err := srv.Search(ctx, &pb.SearchRequest{Request: &pb.SearchRequest_ByContent{ByContent: "needle"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(response.Books) != 1 {
		t.Fatal("response len must be 1")
	}
	responseBook := response.Books[0]
	if responseBook.Title != book.Title ||
		responseBook.Content != book.Content ||
		len(responseBook.Authors) != 1 ||
		responseBook.Authors[0] != book.Authors[0] {
		t.Fatalf("expected %v, found %v", book, responseBook)
	}
}

// TestRepositoryErrorHandling checks that error from repository gets to response
func TestRepositoryErrorHandling(t *testing.T) {
	type mockRecorder = *repositories.MockBooksRepositoryMockRecorder
	// Init mock book repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	booksRepo := repositories.NewMockBooksRepository(ctrl)

	// Init app server
	srv := NewServer(booksRepo)

	// Check request handling
	ctx := context.Background()
	expectedError := errors.New("error")
	booksRepo.EXPECT().SearchByContent(gomock.Any(), "needle").Return(nil, expectedError)
	_, err := srv.Search(ctx, &pb.SearchRequest{Request: &pb.SearchRequest_ByContent{ByContent: "needle"}})
	if err != expectedError {
		t.Fatalf("expected %v, found %v", expectedError, err)
	}
}
