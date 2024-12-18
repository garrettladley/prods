package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/model/product"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/garrettladley/prods/internal/xhttp"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	go_json "github.com/goccy/go-json"
)

type submitRequestBody struct {
	URL string `json:"url"`
}

func (s *Service) Submit(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return xerr.BadRequest(fmt.Errorf("failed to parse token. got: %s", rawToken))
	}

	var r submitRequestBody
	if err := c.BodyParser(&r); err != nil {
		slog.Error("invalid JSON request data", "error", err)
		return xerr.InvalidJSON()
	}

	ctx, cancel := context.WithTimeout(c.Context(), 30*time.Second)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	var expected algo.Solution
	eg.Go(func() error {
		solution, err := s.storage.Solution(ctx, token)
		if err != nil {
			return err
		}
		expected = solution
		return nil
	})

	var actual [][]product.Product
	eg.Go(func() error {
		solution, err := test(ctx, r.URL)
		if err != nil {
			return err
		}
		actual = solution
		return nil
	})

	if err := eg.Wait(); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return xerr.Timeout("took too long to test your solution")
		} else {
			return err
		}
	}

	score := s.algo.Score(ctx, expected, actual)
	if err := s.storage.Submit(ctx, token, score); err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(-1)
}

func test(ctx context.Context, url string) ([][]product.Product, error) {
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(8)

	solutions := make([][]product.Product, len(algo.EncodedFilters))

	for idx, filter := range algo.EncodedFilters {
		func(idx int, filter string) {
			eg.Go(func() error {
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, url+filter, nil)
				if err != nil {
					return fmt.Errorf("failed to create request: %w", err)
				}
				req.Header.Set(xhttp.HeaderContentType, xhttp.HeaderApplicationJSON)

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return fmt.Errorf("failed to make request: %w", err)
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
				}

				var products []product.Product
				if err := go_json.NewDecoder(resp.Body).Decode(&products); err != nil {
					return fmt.Errorf("failed to decode response: %w", err)
				}

				solutions[idx] = products

				return nil
			})
		}(idx, filter)
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return solutions, nil
}
