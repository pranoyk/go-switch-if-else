package main

func main() {

}

func switchFn(state string) []int32 {
	result := []int32{}
	switch state {
	case "solid":
		result = append(result, 0)
	case "liquid":
		result = append(result, 1)
	case "gas":
		result = append(result, 2)
	case "plasma":
		result = append(result, 3)
	}
	return result
}

func ifElseFn(state string) []int32 {
	result := []int32{}
	if state == "solid" {
		result = append(result, 0)
	} else if state == "liquid" {
		result = append(result, 1)
	} else if state == "gas" {
		result = append(result, 2)
	} else if state == "plasma" {
		result = append(result, 3)
	}
	return result
}
