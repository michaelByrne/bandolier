package projections

import "reflect"

type Projection interface {
	CanHandle(reflect.Type) bool
	Handle(reflect.Type, interface{}) error
	GetHandledTypes() []reflect.Type
}

type ProjectionBase struct {
	Projection

	handlers     []EventHandler
	handledTypes map[reflect.Type]bool
	types        []reflect.Type
}

func NewProjection() ProjectionBase {
	return ProjectionBase{
		handledTypes: make(map[reflect.Type]bool, 0),
	}
}

func (p *ProjectionBase) When(event interface{}, handler func(interface{}) error) {
	t := reflect.ValueOf(event).Type()
	p.handlers = append(p.handlers, NewEventHandler(t, handler))
	p.handledTypes[t] = true
}

func (p *ProjectionBase) CanHandle(t reflect.Type) bool {
	_, exists := p.handledTypes[t]
	return exists
}

func (p *ProjectionBase) Handle(t reflect.Type, event interface{}) error {
	for _, h := range p.handlers {
		if h.Type == t {
			err := h.Handler(event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type EventHandler struct {
	Type    reflect.Type
	Handler func(interface{}) error
}

func NewEventHandler(t reflect.Type, handler func(interface{}) error) EventHandler {
	return EventHandler{
		Type:    t,
		Handler: handler,
	}
}
