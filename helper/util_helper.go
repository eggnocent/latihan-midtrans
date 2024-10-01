package helper

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func MessageForTag(tag string) string {
	switch tag {
	case "required":
		return "this field is required"
	}
	return ""
}
