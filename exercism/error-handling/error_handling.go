package erratum

func Use(opener ResourceOpener, input string) error {
	var resource Resource
	var err error
	for resource, err = opener(); err != nil; resource, err = opener() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer resource.Close()
	defer func() {
		if rec := recover(); rec != nil {
			switch e := rec.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()
	resource.Frob(input)
	return err
}
