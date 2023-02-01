/* Auto generated code. Do not modify it */
package models

type Users []User

func (items *Users) IndexOf(item User) int {
	for idx, p := range *items {
		if p == item {
			return idx
		}
	}
	return -1
}

func (items *Users) Includes(item User) bool {
	return items.IndexOf(item) != -1
}

func (items *Users) Any(predicate func(User) bool) bool {
	for _, item := range *items {
		if criteria(item) {
			return true
		}
	}
	return false
}

func (items Users) Filter(predicate func(item User) bool) Users {
	result := Users{}
	for _, p := range items {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}
