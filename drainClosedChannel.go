//prerequisite: channel is closed
func drain(c chan struct{}) {
	for value := range c {
		//do something here
	}
}