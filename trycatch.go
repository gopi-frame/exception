package exception

import (
	"reflect"
)

// Trier trier
type Trier struct {
	function       func()
	catches        map[reflect.Type]func(error)
	defaultCatch   func(error)
	finallyHandler func()
}

// Catch set exception catcher
func (t *Trier) Catch(e error, fn func(error)) *Trier {
	typ := reflect.TypeOf(e)
	if _, ok := t.catches[typ]; !ok {
		t.catches[typ] = fn
	}
	return t
}

// CatchAll set default exception catcher
func (t *Trier) CatchAll(fn func(error)) *Trier {
	t.defaultCatch = fn
	return t
}

// Finally finally
func (t *Trier) Finally(fn func()) *Trier {
	t.finallyHandler = fn
	return t
}

// Run run
func (t *Trier) Run() {
	defer func() {
		if t.finallyHandler != nil {
			defer t.finallyHandler()
		}
		if err := recover(); err != nil {
			switch exp := err.(type) {
			case error:
				if catch, ok := t.catches[reflect.TypeOf(err)]; ok {
					catch(exp)
				} else if t.defaultCatch != nil {
					t.defaultCatch(exp)
				} else {
					panic(exp)
				}
			default:
				if t.defaultCatch != nil {
					t.defaultCatch(NewValueException(exp))
				}
			}
		}
	}()
	t.function()
}

// Try try catch
func Try(fn func()) *Trier {
	t := new(Trier)
	t.function = fn
	t.catches = make(map[reflect.Type]func(error))
	return t
}
