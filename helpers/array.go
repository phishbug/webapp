package helpers

// inArray checks if a value exists in a slice of integers
func InArrayStrings(arr []string, value string) bool {
    for _, v := range arr {
        if v == value {
            return true
        }
    }
    return false
}