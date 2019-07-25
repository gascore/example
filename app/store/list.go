package store

import (
	"github.com/gascore/std/store"
)

var listsHandlers = map[string]store.Handler{
	"clearList": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		return map[string]interface{}{
			values[0].(string): []interface{}{},
		}, nil
	},
	"appendToList": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		var (
			name = values[0].(string)
			el   = values[1].(string)
			list = s.Get(name).([]interface{})
		)

		list = append(list, el)
		
		return map[string]interface{}{
			name: list,
		}, nil
	},
	"editActive": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		active := s.Get("active").([]interface{})

		active[values[0].(int)] = values[1].(interface{})

		return map[string]interface{}{
			"active": active,
		}, nil
	},
	"deleteFromActive": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		var (
			i	   = values[0].(int)
			active = s.Get("active").([]interface{})
			el	   = active[i]
		)

		copy(active[i:], active[i+1:])  // Shift a[i+1:] left one index
		active[len(active)-1] = ""      // Erase last element (write zero value)
		active = active[:len(active)-1] // Truncate slice

		return map[string]interface{}{
			"active":  active,
			"deleted": append(s.Get("deleted").([]interface{}), el),
		}, nil
	},
	"completeInActive": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		var (
			i      = values[0].(int)
			active = s.Get("active").([]interface{})
			el     = active[i]
		)

		copy(active[i:], active[i+1:])  // Shift a[i+1:] left one index
		active[len(active)-1] = ""      // Erase last element (write zero value)
		active = active[:len(active)-1] // Truncate slice

		return map[string]interface{}{
			"active":    active,
			"completed": append(s.Get("completed").([]interface{}), el),
		}, nil
	},
}
