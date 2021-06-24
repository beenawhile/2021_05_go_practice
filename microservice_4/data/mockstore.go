package data

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Search(name string) []Kitten {
	// mock 패키지는 메서드의 호출 여부와 사용한 매개 변수를 기록하고 있기 때문에 나중에 이를 바탕으로 메서드에 대한 단정문을 작성할 수 있음
	args := m.Mock.Called(name)

	// mock이 초기 설정에서 제공한 인수의 목록을 리턴
	return args.Get(0).([]Kitten)
}
