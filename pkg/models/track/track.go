package track

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/tunes-anywhere/anywhere-api/pkg/db"
	"github.com/tunes-anywhere/anywhere-api/pkg/util"
)

var (
	_ json.Unmarshaler = (*TrackAttributes)(nil)
)

type TrackAttributes struct {
	Source   string `json:"source"`
	Duration uint   `json:"duration"`
	Title    string `json:"title,omitempty"`
	Album    string `json:"album,omitempty"`
	Artist   string `json:"artist,omitempty"`
}

func (ta TrackAttributes) UnmarshalJSON(raw []byte) error {
	if err := json.Unmarshal(raw, &ta); err != nil {
		return err
	}

	if ta.Source == "" || ta.Title == "" || ta.Duration < 1 {
		return fmt.Errorf("track missing required fields: %s", string(raw))
	}

	return nil
}

func (ta TrackAttributes) Save(ctx context.Context, path string) (*Track, error) {
	var (
		err    error
		track  = New(ta, path)
		values map[string]types.AttributeValue
		client *dynamodb.Client
	)

	if values, err = attributevalue.MarshalMap(track); err != nil {
		return nil, err
	}

	if client, err = db.GetClient(ctx); err != nil {
		return nil, err
	}

	client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &util.Config.TableName,
		Item:      values,
	})

	return nil, nil
}

type Track struct {
	TrackAttributes
	ID          uuid.UUID `json:"id"`
	Path        string    `json:"path"`
	AccessCount int       `json:"access_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func New(ta TrackAttributes, path string) Track {
	return Track{
		TrackAttributes: ta,
		ID:              uuid.New(),
		Path:            path,
		AccessCount:     0,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}
}

func NewFromString(raw string) (*Track, error) {
	var track Track
	if err := json.Unmarshal([]byte(raw), &track); err != nil {
		return nil, err
	}

	return &track, nil
}
