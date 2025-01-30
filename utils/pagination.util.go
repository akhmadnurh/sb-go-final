package utils

func GetOffset(page, limit int) (offset int) {
	offset = (page - 1) * limit

	return
}
