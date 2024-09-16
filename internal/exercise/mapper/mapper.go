package mapper

import (
	"encoding/json"
	"github.com/bellyachx/hercules-be/api/exercisepb"
	"github.com/bellyachx/hercules-be/internal/exercise/model"
)

func MapToModel(ex *exercisepb.Exercise) (*model.Exercise, error) {
	reqJson, err := json.Marshal(ex)
	if err != nil {
		return nil, err
	}

	exModel := &model.Exercise{}
	err = json.Unmarshal(reqJson, exModel)
	if err != nil {
		return nil, err
	}
	return exModel, nil
}

func MapFromModel(exModel *model.Exercise) (*exercisepb.Exercise, error) {
	exModelJson, err := json.Marshal(exModel)
	if err != nil {
		return nil, err
	}

	ex := &exercisepb.Exercise{}
	err = json.Unmarshal(exModelJson, ex)
	if err != nil {
		return nil, err
	}
	return ex, nil
}

func MapFromModelSlice(modelExercises []model.Exercise) ([]*exercisepb.Exercise, error) {
	exSlice := make([]*exercisepb.Exercise, len(modelExercises))
	for i, val := range modelExercises {
		var err error
		exSlice[i], err = MapFromModel(&val)
		if err != nil {
			return nil, err
		}
	}

	return exSlice, nil
}
