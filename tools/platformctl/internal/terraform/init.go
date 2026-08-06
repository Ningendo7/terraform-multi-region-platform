package terraform

func Init(directory string) error {

	return Execute(directory, "init")
}