package terraform

func Apply(directory string) error {
	return Execute(directory, "apply")
}
