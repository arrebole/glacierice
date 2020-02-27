package service

import "github.com/arrebole/glacierice/model"

var conn = connect()

// Resource3RD ...
func Resource3RD() []model.Resource3RD {
	var result []model.Resource3RD
	conn.Find(&result)
	return result
}

// ResourceGithub ...
func ResourceGithub() []model.ResourceGithub {
	var result []model.ResourceGithub
	conn.Find(&result)
	return result
}

// Contributor ...
func Contributor() []model.Contributor {
	var result []model.Contributor
	conn.Find(&result)
	return result
}
