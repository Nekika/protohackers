package main

const (
	MethodIsPrime = "isPrime"
)

type IsPrimeRequest struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

func (r IsPrimeRequest) Valid() bool {
	return r.Method == MethodIsPrime
}

type IsPrimeResponse struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func NewIsPrimeResponse(n int) IsPrimeResponse {
	return IsPrimeResponse{
		Method: MethodIsPrime,
		Prime:  IsPrime(n),
	}
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if (n%2 == 0) || (n%3 == 0) {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if (n%i == 0) || (n%(i+2) == 0) {
			return false
		}
	}
	return true
}
