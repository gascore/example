package store

import "github.com/gascore/std/store"

// S global store instance
var S *store.Store

// InitStore init store
func InitStore() error {
	handlers := store.NewHandlers()
	handlers.AddMany(listsHandlers)

	s, err := store.New(&store.Store{
		Data: map[string]interface{}{
			"count":     0,
			"active":    []interface{}{},
			"completed": []interface{}{},
			"deleted":   []interface{}{},
		},
		Handlers:  handlers,
		OnCreate:  []store.OnCreateHook{store.LSSyncOnCreate},
		AfterEmit: []store.AfterEmitHook{store.LSSyncAfterEmit},
	})
	if err != nil {
		return err
	}

	S = s

	return nil
}
