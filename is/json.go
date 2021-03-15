package is

import "encoding/json"

func JSON(s string) bool {
    var js json.RawMessage
    return json.Unmarshal([]byte(s), &js) == nil
}
