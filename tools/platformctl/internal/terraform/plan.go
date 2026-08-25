package terraform

func Plan(directory string) error {
	return Execute(directory, "plan")
}
