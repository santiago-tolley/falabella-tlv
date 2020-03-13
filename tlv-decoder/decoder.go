package decoder

import (
	"strconv"
)

const HEADER_LENGTH = 5

func Decode(input []byte) (map[string]string, error) {
	// Return an error on empty input
	if len(input) == 0 {
		return map[string]string{}, ErrEmptyInput{}
	}

	result := make(map[string]string)
	fieldNumber := 0
	i := 0
	for i < len(input) {
		// Check minimum length for a valid TLV item
		if len(input) < i+HEADER_LENGTH {
			return map[string]string{}, ErrInvalidFormat{}
		}

		// Extract Type from TLV and convert it to String
		dataType := string(input[i])
		if dataType != "N" && dataType != "A" {
			return map[string]string{}, ErrInvalidType{}
		}

		// Extract Length from TLV and convert it to Int
		length, err := strconv.Atoi(string(input[i+3 : i+5]))
		if err != nil {
			return map[string]string{}, ErrInvalidFormat{}
		}

		// Check if value length does not exceed input length
		if i+length+HEADER_LENGTH > len(input) {
			return map[string]string{}, ErrInvalidLength{}
		}

		// Extract Value from TLV
		value := string(input[i+HEADER_LENGTH : i+length+HEADER_LENGTH])

		// Check if the Value is of the type specified in Type
		if dataType == "N" {
			_, err := strconv.Atoi(value)
			if err != nil {
				return map[string]string{}, ErrInvalidType{}
			}
		}

		// Add the TLV to the result map with the order of the data item
		result[strconv.Itoa(fieldNumber)] = string(input[i : i+length+HEADER_LENGTH])
		fieldNumber++

		// Update i to match the begining of the next item
		i = i + length + HEADER_LENGTH
	}

	return result, nil
}
