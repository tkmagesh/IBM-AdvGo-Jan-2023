func logWrapper(f interface{}) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		// Get the function type of the input function
		fnType := reflect.TypeOf(f)
		// Check if the input function has correct number of arguments
		if len(args) != fnType.NumIn() {
			return nil
		}
		// Create a slice of reflect.Value to hold the input arguments
		inputArgs := make([]reflect.Value, fnType.NumIn())
		// Convert each input argument to a reflect.Value
		for i, arg := range args {
			inputArgs[i] = reflect.ValueOf(arg)
		}
		// Call the input function using reflect.Value.Call()
		result := reflect.ValueOf(f).Call(inputArgs)
		// Log the input arguments and the result
		fmt.Println("Input args:", args)
		fmt.Println("Result:", result[0].Interface())
		// Return the result as interface{}
		return result[0].Interface()
	}
}