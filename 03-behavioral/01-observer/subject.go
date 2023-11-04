package main

// Define the interface for the observable type
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// The DataSubject will have a list of listeners
// and a field that gets changed, triggering them
type DataSubject struct {
	observers []DataListener
	field     string
}

// ChangeItem function will cause the Listeners to be called
func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

// registerObserver adds an observer to the list
func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

// unregisterObserver removes an observer from the list
func (ds *DataSubject) unregisterObserver(o DataListener) {
	var newObs []DataListener
	for _, obs := range ds.observers {
		if obs.Name != o.Name {
			newObs = append(newObs, obs)
		}
	}
	ds.observers = newObs
}

// notifyAll calls all the listeners with the changed data
func (ds *DataSubject) notifyAll() {
	for _, obs := range ds.observers {
		obs.onUpdate(ds.field)
	}
}
