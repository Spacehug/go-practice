package react

type reactor struct{}

// New returns a newly created Reactor.
func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(value int) InputCell {
	return &inputCell{*newCell(value)}
}

func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	compute := newComputeCell(func() int { return f(c.Value()) })
	if cell, ok := c.(cellWithDependents); ok {
		cell.addDependent(compute)
	}
	return compute
}

func (r *reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	compute := newComputeCell(func() int { return f(c1.Value(), c2.Value()) })
	if cell1, ok := c1.(cellWithDependents); ok {
		cell1.addDependent(compute)
	}
	if cell2, ok := c2.(cellWithDependents); ok {
		cell2.addDependent(compute)
	}
	return compute
}

type cellWithDependents interface {
	addDependent(compute *computeCell)
}

type cell struct {
	value      int
	dependents map[*computeCell]bool
}

func newCell(value int) *cell {
	return &cell{value: value, dependents: map[*computeCell]bool{}}
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) addDependent(compute *computeCell) {
	c.dependents[compute] = true
}

func (c *cell) populateDependentValues(dependents map[*computeCell]int) {
	for d := range c.dependents {
		dependents[d] = d.Value()
		d.populateDependentValues(dependents)
	}
}

func (c *cell) dependentValues() map[*computeCell]int {
	depValues := map[*computeCell]int{}
	c.populateDependentValues(depValues)
	return depValues
}

type inputCell struct {
	cell
}

func (c *inputCell) SetValue(v int) {
	c.value = v

	oldDepValues := c.dependentValues()
	for d := range oldDepValues {
		d.compute()
	}

	for d, v := range c.dependentValues() {
		if oldDepValues[d] != v {
			d.invokeCallbacks()
		}
	}
}

type computeCell struct {
	cell
	computeFunc    func() int
	callbacks      map[int]func(int)
	nextCallbackID int
}

func newComputeCell(computeFunc func() int) *computeCell {
	return &computeCell{
		cell:        *newCell(computeFunc()),
		computeFunc: computeFunc,
		callbacks:   map[int]func(int){},
	}
}

func (c *computeCell) AddCallback(cb func(int)) Canceler {
	cancel := canceler{computeCell: c, callbackID: c.nextCallbackID}
	c.callbacks[c.nextCallbackID] = cb
	c.nextCallbackID++
	return cancel
}

func (c *computeCell) invokeCallbacks() {
	for _, cb := range c.callbacks {
		cb(c.value)
	}
}

func (c *computeCell) compute() {
	c.value = c.computeFunc()
	for d := range c.dependents {
		d.compute()
	}
}

type canceler struct {
	computeCell *computeCell
	callbackID  int
}

func (c canceler) Cancel() {
	delete(c.computeCell.callbacks, c.callbackID)
}
