package json


import (
	"encoding/json"
	"os"
	"path/filepath"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)


type fileData struct {
	Version int             `json:"version"`
	Tasks   []*domain.Task  `json:"tasks"`
}

const fileVersion = 1


//load

func (r *TaskRepo) load() (*fileData, error) {
	if _, err := os.Stat(r.path); os.IsNotExist(err) {
		return &fileData{
			Version: fileVersion,
			Tasks:   []*domain.Task{},
		}, nil
	}

	b, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	var data fileData
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	if data.Version != fileVersion {
		return nil, ErrInvalidData
	}

	return &data, nil
}



//save

func (r *TaskRepo) save(data *fileData) error {
	dir := filepath.Dir(r.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	tmp := r.path + ".tmp"

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, r.path)
}
