package main

import (
  "testing"
)

func TestPosStringOK(t *testing.T) {
    var pos int
    okResponses := []string{"y", "Y", "yes", "Yes", "YES"}
    // test regular 1
    tval1 := "y"
    pos = posString(okResponses, tval1)
    if pos == -1 {
        t.Errorf("When value is %s, not expected -1.", tval1)
    }

    // test regular 2
    tval2 := "Yes"
    pos = posString(okResponses, tval2)
    if pos == -1 {
        t.Errorf("When value is %s, not expected -1.", tval2)
    }    

    // test irregular
    tval3 := "yeah"
    pos = posString(okResponses, tval3)
    if pos != -1 {
        t.Errorf("When value is %s, expected -1. pos:%d", tval3, pos)
    }

    // test empty
    tval4 := ""
    pos = posString(okResponses, tval4)
    if pos != -1 {
        t.Errorf("When value is empty, expected -1. pos:%d", pos)
    }
}

func TestPosStringNG(t *testing.T) {
    var pos int
    ngResponses := []string{"n", "N", "no", "No", "NO"}
    // test regular 1
    tval1 := "N"
    pos = posString(ngResponses, tval1)
    if pos == -1 {
        t.Errorf("When value is %s, not expected -1.", tval1)
    }

    // test regular 2
    tval2 := "no"
    pos = posString(ngResponses, tval2)
    if pos == -1 {
        t.Errorf("When value is %s, not expected -1.", tval2)
    }    

    // test irregular
    tval3 := "noph"
    pos = posString(ngResponses, tval3)
    if pos != -1 {
        t.Errorf("When value is %s, expected -1. pos:%d", tval3, pos)
    }

    // test empty
    tval4 := ""
    pos = posString(ngResponses, tval4)
    if pos != -1 {
        t.Errorf("When value is empty, expected -1. pos:%d", pos)
    }
}

func TestContainsStringOK(t *testing.T) {
    var result bool
    okResponses := []string{"y", "Y", "yes", "Yes", "YES"}
    // test regular 1
    tval1 := "y"
    result = containsString(okResponses, tval1)
    if result == false {
        t.Errorf("When value is %s, expected true.", tval1)
    }

    // test regular 2
    tval2 := "Yes"
    result = containsString(okResponses, tval2)
    if result == false {
        t.Errorf("When value is %s, expected true.", tval2)
    }    

    // test irregular
    tval3 := "yeah"
    result = containsString(okResponses, tval3)
    if result == true {
        t.Errorf("When value is %s, expected false.", tval3)
    }

    // test empty
    tval4 := ""
    result = containsString(okResponses, tval4)
    if result == true {
        t.Errorf("When value is empty, expected false.")
    }
}

func TestContainsStringNG(t *testing.T) {
    var result bool
    ngResponses := []string{"n", "N", "no", "No", "NO"}
    // test regular 1
    tval1 := "N"
    result = containsString(ngResponses, tval1)
    if result == false {
        t.Errorf("When value is %s, expected true.", tval1)
    }

    // test regular 2
    tval2 := "no"
    result = containsString(ngResponses, tval2)
    if result == false {
        t.Errorf("When value is %s, not expected -1.", tval2)
    }    

    // test irregular
    tval3 := "noph"
    result = containsString(ngResponses, tval3)
    if result == true {
        t.Errorf("When value is %s, expected false.", tval3)
    }

    // test empty
    tval4 := ""
    result = containsString(ngResponses, tval4)
    if result == true {
        t.Errorf("When value is empty, expected false.")
    }
}