package mapper

import (
	"github.com/bellyachx/hercules-be/api/exercisepb"
	"github.com/bellyachx/hercules-be/internal/exercise/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestMapToModel(t *testing.T) {
	type args struct {
		ex *exercisepb.Exercise
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Exercise
		wantErr bool
	}{
		{
			name: "Should map successfully",
			args: args{
				ex: &exercisepb.Exercise{
					UserId:      "123321",
					Name:        "test_name",
					Description: "test_desc",
					MuscleGroup: "all",
					Difficulty:  "hard",
					Type:        "cardio",
					SetsCount:   3,
					RepsCount:   15,
					Duration:    1500,
				},
			},
			want: &model.Exercise{
				UserID:      "123321",
				Name:        "test_name",
				Description: "test_desc",
				MuscleGroup: "all",
				Difficulty:  "hard",
				Type:        "cardio",
				SetsCount:   3,
				RepsCount:   15,
				Duration:    1500,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapToModel(tt.args.ex)
			if tt.wantErr {
				assert.Error(t, err)
				require.Nil(t, got)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMapFromModel(t *testing.T) {
	type args struct {
		exModel *model.Exercise
	}
	tests := []struct {
		name    string
		args    args
		want    *exercisepb.Exercise
		wantErr bool
	}{
		{
			name: "Should map successfully",
			args: args{
				exModel: &model.Exercise{
					UserID:      "123321",
					Name:        "test_name",
					Description: "test_desc",
					MuscleGroup: "all",
					Difficulty:  "hard",
					Type:        "cardio",
					SetsCount:   3,
					RepsCount:   15,
					Duration:    1500,
				},
			},
			want: &exercisepb.Exercise{
				UserId:      "123321",
				Name:        "test_name",
				Description: "test_desc",
				MuscleGroup: "all",
				Difficulty:  "hard",
				Type:        "cardio",
				SetsCount:   3,
				RepsCount:   15,
				Duration:    1500,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapFromModel(tt.args.exModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapFromModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFromModel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapFromModelSlice(t *testing.T) {
	type args struct {
		modelExercises []model.Exercise
	}
	tests := []struct {
		name    string
		args    args
		want    []*exercisepb.Exercise
		wantErr bool
	}{
		{
			name: "Should map successfully",
			args: args{
				modelExercises: []model.Exercise{
					{
						UserID:      "123321",
						Name:        "test_name",
						Description: "test_desc",
						MuscleGroup: "all",
						Difficulty:  "hard",
						Type:        "cardio",
						SetsCount:   3,
						RepsCount:   15,
						Duration:    1500,
					},
					{
						UserID:      "111111",
						Name:        "tesssst",
						Description: "description1",
						MuscleGroup: "legs",
						Difficulty:  "middle",
						SetsCount:   3,
						RepsCount:   7,
						Duration:    3000,
					},
				},
			},
			want: []*exercisepb.Exercise{
				{
					UserId:      "123321",
					Name:        "test_name",
					Description: "test_desc",
					MuscleGroup: "all",
					Difficulty:  "hard",
					Type:        "cardio",
					SetsCount:   3,
					RepsCount:   15,
					Duration:    1500,
				},
				{
					UserId:      "111111",
					Name:        "tesssst",
					Description: "description1",
					MuscleGroup: "legs",
					Difficulty:  "middle",
					SetsCount:   3,
					RepsCount:   7,
					Duration:    3000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapFromModelSlice(tt.args.modelExercises)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapFromModelSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFromModelSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
