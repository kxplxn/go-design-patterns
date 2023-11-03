package main

func main() {
	// Construct two DataListener observers and
	// give each one a name
	_ = DataListener{
		Name: "Listener 1",
	}
	_ = DataListener{
		Name: "Listener 2",
	}

	// Create the DataSubject that the listeners will observe
	_ = &DataSubject{}
	// TODO: Register each listener with the DataSubject

	// TODO: Change the data in the DataSubject - this will cause the
	// onUpdate method of each listener to be called

	// TODO: Try to unregister one of the observers

}
