package magic

import (
	"encoding/json"
	"strings"

	"github.com/x-research-team/utils/is"
)

func Jsonify(s string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		if strings.Contains(err.Error(), "array") {
			array := make([]map[string]interface{}, 0)
			if err := json.Unmarshal([]byte(s), &array); err != nil {
				return nil, err
			}
			for i := range array {
				for k, v := range array[i] {
					switch t := v.(type) {
					case string:
						if !is.JSON(t) {
							continue
						}
						array[i][k], err = Jsonify(t)
						return array[i], err
					default:
						continue
					}
				}
			}
		} else {
			return nil, err
		}
	}
	for k, v := range result {
		switch t := v.(type) {
		case string:
			if !is.JSON(t) {
				continue
			}
			var err error
			result[k], err = Jsonify(t)
			return result, err
		default:
			continue
		}
	}
	return result, nil
}