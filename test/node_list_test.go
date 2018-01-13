package gost_test

import (
	"testing"

	"github.com/christat/gost/list"
)

func generateList(size int) *gost.NodeList {
	list := new(gost.NodeList)
	for i := 0; i < size; i++ {
		list.Append(newVector(i))
	}
	return list
}

func TestNodeList_Append(t *testing.T) {
	list := generateList(num)
	if list.Size != num {
		t.Errorf("Append() size update failed; expected: %v, got: %v", num, list.Size)
	}
}

func TestNodeList_Retrieve(t *testing.T) {
	list := generateList(num)
	_, err := list.Retrieve(int(list.Size))
	if err == nil {
		t.Error("Retrieve() did not return error on exceeding size index")
	}
	_, err = list.Retrieve(int(list.Size) - 1)
	if err != nil {
		t.Error("Retrieve() failed retrieving last item")
	}
	value, err := list.Retrieve(int(list.Size) / 2)
	if err != nil {
		t.Error("Retrieve() failed unexpectedly")
	}
	if *(value.(*vector)) != *newVector(int(list.Size) / 2) {
		t.Errorf("Retrieve() error; expected to get: %v, got: %v", newVector(int(list.Size)/2), value)
	}
}

func TestNodeList_Add(t *testing.T) {
	list := generateList(num)
	err := list.Add(int(list.Size)+1, newVector(num))
	if err == nil {
		t.Error("Add() did not return error on exceeding size index")
	}
	list.Add(int(list.Size), newVector(int(list.Size)+5))
	value, err := list.Retrieve(int(list.Size) - 1)
	if *(value.(*vector)) != *newVector(int(list.Size) + 4) {
		t.Errorf("Add() error; expected to retrieve last element: %v, got: %v", newVector(int(list.Size)+4), value)
	}
	err = list.Add(int(list.Size)/2, newVector(int(list.Size)/2+5))
	value, err = list.Retrieve(int(list.Size)/2 - 1)
	if *(value.(*vector)) != *newVector(int(list.Size)/2 + 4) {
		t.Errorf("Add() error; expected to retrieve intermediate element: %v, got: %v", newVector(int(list.Size)/2+4), value)

	}
	list.Add(0, newVector(0))
	value, err = list.Retrieve(0)
	if *(value.(*vector)) != *newVector(0) {
		t.Error("Add() error inserting head item")
	}
}

func TestNodeList_Remove(t *testing.T) {
	list := generateList(num)
	_, err := list.Remove(num)
	if err == nil {
		t.Error("Remove() did not return error on exceeding size index")
	}
	value, err := list.Remove(num - 1)
	if err != nil {
		t.Error("Remove() failed removing last element")
	}
	if *(value.(*vector)) != *newVector(num - 1) {
		t.Errorf("Add() error; expected to retrieve last element: %v, got: %v", newVector(num-1), value)
	}
	value, err = list.Remove(0)
	if err != nil {
		t.Error("Remove() failed removing last element")

	}
	if *(value.(*vector)) != *newVector(0) {
		t.Errorf("Add() error; expected to retrieve first element: %v, got: %v", newVector(0), value)
	}
	value, err = list.Remove(num / 2)
	if err != nil {
		t.Error("Remove() failed removing middle element")

	}
	if *(value.(*vector)) != *newVector(num/2 + 1) {
		t.Errorf("Add() error; expected to retrieve middle element: %v, got: %v", newVector(num/2+1), value)
	}
}
