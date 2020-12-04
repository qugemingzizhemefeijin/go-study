package store

import (
	"testing"
)

//TestStore 测试序列化
func TestStore(t *testing.T) {
	var monster = &Monster{
		Name:  "狐狸精",
		Age:   1000,
		Skill: "蛊惑人心",
	}

	b := monster.Store("E:/111.json")
	if b {
		t.Logf("Store() Testing Success")
	} else {
		t.Fatalf("Store() Testing Fail, monster=%v\n", monster)
	}
}

//TestRestore 测试反序列化
func TestRestore(t *testing.T) {
	var monster *Monster = &Monster{}

	b := monster.ReStore("E:/111.json")
	if b {
		if monster.Name == "狐狸精" {
			t.Logf("ReStore() Testing Success, monster=%v\n", monster)
		} else {
			t.Fatalf("ReStore() Testing Fail, Name Not Match, monster=%v\n", monster)
		}
	} else {
		t.Fatalf("ReStore() Testing Fail\n")
	}
}
