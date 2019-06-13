package store

import (
	"errors"
	"fmt"
	"github.com/gascore/std/store"
)

var listsHandlers = map[string]store.Handler{
	"clearList": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		name, ok := values[0].(string)
		if !ok {
			return nil, errors.New("invalid list name")
		}

		return map[string]interface{}{
			name: []interface{}{},
		}, nil
	},
	"appendToList": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		el, ok := values[1].(string)
		if !ok {
			return nil, errors.New("invalid element value")
		}

		name, ok := values[0].(string)
		if !ok {
			return nil, errors.New("invalid 'name' type")
		}

		listUnTyped, err := s.GetSafely(name)
		if err != nil {
			return nil, err
		}

		list, ok := listUnTyped.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid '%s' type", name)
		}

		list = append(list, el)

		return map[string]interface{}{
			name: list,
		}, nil
	},
	"editActive": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		i := values[0].(int)
		val := values[1].(string)

		active := s.Get("active").([]interface{})

		active[i] = val

		return map[string]interface{}{
			"active": active,
		}, nil
	},
	"deleteFromActive": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		i, ok := values[0].(int)
		if !ok {
			return nil, errors.New("invalid index")
		}

		appendToDeleted, ok := values[1].(bool)
		if !ok {
			return nil, errors.New("invalid appendToDeleted")
		}

		activeUnTyped, err := s.GetSafely("active")
		if err != nil {
			return nil, err
		}

		active, ok := activeUnTyped.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid '%s' type", "activeUnTyped")
		}

		el := active[i]

		copy(active[i:], active[i+1:]) // Shift a[i+1:] left one index
		active[len(active)-1] = ""     // Erase last element (write zero value)
		active = active[:len(active)-1]   // Truncate slice

		out := map[string]interface{}{
			"active": active,
		}

		if appendToDeleted {
			out["deleted"] = append(s.Get("deleted").([]interface{}), el)
		}

		return out, nil
	},
	"completeInActive": func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		i, ok := values[0].(int)
		if !ok {
			return nil, errors.New("invalid index")
		}

		activeUnTyped, err := s.GetSafely("active")
		if err != nil {
			return nil, err
		}

		active, ok := activeUnTyped.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid '%s' type", "activeUnTyped")
		}

		el := active[i]

		copy(active[i:], active[i+1:]) // Shift a[i+1:] left one index
		active[len(active)-1] = ""     // Erase last element (write zero value)
		active = active[:len(active)-1]   // Truncate slice

		out := map[string]interface{}{
			"active": active,
		}

		out["completed"] = append(s.Get("completed").([]interface{}), el)

		return out, nil
	},
}
