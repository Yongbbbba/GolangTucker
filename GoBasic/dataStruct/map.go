// hash map 자료구조
// map은 전화번호부와 같이 key(이름)와 value(전화번호)로 이루어져 있음
// Rolling hash 사용
// hash map은 hash를 어떻게 만드느냐가 매우 중요
// map은 list와 함께 가장 많이 사용되는 자료구조
// hash map은 매우 빠르고 효율적, find, add, remove가 모두 O(1)

package dataStruct

func Hash(s string) int {
	h := 0
	A := 256
	B := 3571 // 주로 소수를 사용

	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B // Rolling Hash
	}
	return h // 0 ~ 3570의 범위를 갖는데, 이는 손실압축의 경우임
}

type KeyValue struct {
	key   string
	value string
}

type Map struct {
	keyArray [3571][]KeyValue
}

func CreateMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], KeyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}

	}
	return ""

}
