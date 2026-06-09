package mapping

import (
	"encoding/json"
	"os"
	"sync"
)

type CategoryMapping map[string]int

type MappingTable struct {
	Tasks      CategoryMapping `json:"tasks"`
	Extensions CategoryMapping `json:"extensions"`
	Wagons     CategoryMapping `json:"wagons"`
	Engines    CategoryMapping `json:"engines"`
	Versions   CategoryMapping `json:"versions"`
}

type Registry struct {
	mu       sync.Mutex
	filePath string
	table    *MappingTable
}

func NewRegistry(filePath string) (*Registry, error) {
	r := &Registry{
		filePath: filePath,
		table: &MappingTable{
			Tasks:      make(CategoryMapping),
			Extensions: make(CategoryMapping),
			Wagons:     make(CategoryMapping),
			Engines:    make(CategoryMapping),
			Versions:   make(CategoryMapping),
		},
	}

	if err := r.load(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Registry) load() error {
	file, err := os.Open(r.filePath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(r.table)
}

func (r *Registry) save() error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(r.table)
}

// GetID は指定されたカテゴリとキーに対して連番IDを返します。
// キーが未登録の場合は、新しい連番IDを割り当ててマッピングをJSONに保存します。
func (r *Registry) GetID(category string, key string) (int, error) {
	if key == "" {
		return 0, nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	var cat CategoryMapping
	switch category {
	case "tasks":
		cat = r.table.Tasks
	case "extensions":
		cat = r.table.Extensions
	case "wagons":
		cat = r.table.Wagons
	case "engines":
		cat = r.table.Engines
	case "versions":
		cat = r.table.Versions
	default:
		return 0, nil
	}

	if id, exists := cat[key]; exists {
		return id, nil
	}

	newID := len(cat) + 1
	cat[key] = newID

	if err := r.save(); err != nil {
		return 0, err
	}

	return newID, nil
}
