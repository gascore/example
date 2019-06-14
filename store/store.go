package store

import "github.com/gascore/std/store"

// S store instance
var S *store.Store

// InitStore init store
func InitStore() error {
	var err error

	handlers := store.NewHandlers()
	handlers.Add("updateCount", func(s *store.Store, values ...interface{}) (map[string]interface{}, error) {
		return map[string]interface{}{
			"count": s.Get("count").(int) + 1,
		}, nil
	})
	handlers.AddMany(listsHandlers)

	S, err = store.New(&store.Store{
		Data: map[string]interface{}{
			"count":     0,
			"active":    []interface{}{},
			"completed": []interface{}{},
			"deleted":   []interface{}{},
		},
		Handlers: handlers,
		OnCreate: []store.OnCreateHook{
			store.LSSyncOnCreate,
		},
		AfterEmit: []store.AfterEmitHook{
			store.LSSyncAfterEmit,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
