package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"log"
	"net/http"
	"net/url"
)

const ENDPOINT = "http://factordb.com/api"

type FactorDBResponse struct {
	Status  string
	Id      string
	Factors []Factor
}

type Factor struct {
	Number string
	Power  int
}

func ConvertToFactorDB(b []byte) (FactorDBResponse, error) {
	var base interface{}
	err := json.Unmarshal(b, &base)
	if err != nil {
		return FactorDBResponse{}, errors.New("Cannot parse the input")
	}
	s := base.(map[string]interface{})

	var factor FactorDBResponse
	factor.Status = s["status"].(string)
	factor.Id = s["id"].(string)

	factors := s["factors"].([]interface{})

	for _, f := range factors {
		tmp := f.([]interface{})
		number, _ := (tmp[0].(string))
		power := int(tmp[1].(float64))
		factor.Factors = append(factor.Factors, Factor{number, power})
	}

	return factor, nil
}

type FactorDB struct {
	Number string
	Result FactorDBResponse
}

func (f *FactorDB) Empty() bool {
	if f.Result.Status == "" {
		return true
	}
	return false
}

func (f *FactorDB) Connect() error {
	values := url.Values{}
	values.Add("query", fmt.Sprintf("%s", f.Number))
	resp, err := http.Get(fmt.Sprintf("%s?%s", ENDPOINT, values.Encode()))
	if err != nil {
		return errors.New("cannot connect" + ENDPOINT)
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Empty Body")
	}

	response, err := ConvertToFactorDB(b)
	if err != nil {
		return errors.New("Cannot converting data")
	}

	f.Result = response
	return nil
}

func (f *FactorDB) GetId() (string, error) {
	if f.Empty() {
		return "", errors.New("Empty Result")
	}
	return f.Result.Id, nil
}

func (f *FactorDB) GetStatus() (string, error) {
	if f.Empty() {
		return "", errors.New("Empty Result")
	}
	return f.Result.Status, nil
}

func (f *FactorDB) GetFactorList() ([]Factor, error) {
	if f.Empty() {
		return []Factor{}, errors.New("Empty Result")
	}
	return f.Result.Factors, nil
}

func GetFactors(c string) ([]Factor, error) {
	n := c

	f := FactorDB{Number: n}
	if err := f.Connect(); err != nil {
		log.Fatal("Connection Error")
	}

	factors, _ := f.GetFactorList()
	return factors, nil
}
