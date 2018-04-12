package funk

import (
	"encoding/json"
	"os"
)

type Config struct {
	GithubURL        string            `json:"github_url"`
	MainSource       string            `json:"main_file"`
	InputMethod      string            `json:"input_method"`
	InputParameters  map[string]string `json:"input_parameters"`
	OutputMethod     string            `json:"output_method"`
	OutputParameters map[string]string `json:"output_parameters"`
}

func ReadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)

	if err != nil {
		return Config{}, err
	}

	var config Config

	err = json.NewDecoder(file).Decode(&config)

	return config, err
}
