package terraform

func Destroy(directory string) error {
	return Execute(directory, "destroy")
}