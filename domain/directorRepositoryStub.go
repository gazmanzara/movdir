package domain

type DirectorRepositoryStub struct {
	directors []Director
}

func (d DirectorRepositoryStub) FindAll() ([]Director, error) {
	return d.directors, nil
}

func NewDirectorRepositoryStub() DirectorRepositoryStub {
	directors := []Director{
		{4762, "James Cameron", 2},
		{4763, "Gore Verbinski", 2},
		{4764, "Sam Mendes", 2},
		{4765, "Christopher Nolan", 2},
		{4766, "Andrew Stanton", 2},
	}

	return DirectorRepositoryStub{
		directors: directors,
	}
}
