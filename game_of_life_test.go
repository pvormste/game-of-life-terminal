package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestWorld_CheckNeighborIsAlive(t *testing.T) {
	type testFacts struct {
		inhabitants             [][]bool
		inputX                  int
		inputY                  int
		expectedIsNeighborAlive types.GomegaMatcher
	}

	t.Run("should return false for out of bounds on:", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		tests := []testFacts{
			{
				inhabitants: [][]bool{
					{true, true, true},
					{true, true, true},
					{true, true, true},
				},
				inputX:                  -1,
				inputY:                  0,
				expectedIsNeighborAlive: BeFalse(),
			},
			{
				inhabitants: [][]bool{
					{true, true, true},
					{true, true, true},
					{true, true, true},
				},
				inputX:                  0,
				inputY:                  -1,
				expectedIsNeighborAlive: BeFalse(),
			},
			{
				inhabitants: [][]bool{
					{true, true, true},
					{true, true, true},
					{true, true, true},
				},
				inputX:                  3,
				inputY:                  0,
				expectedIsNeighborAlive: BeFalse(),
			},
			{
				inhabitants: [][]bool{
					{true, true, true},
					{true, true, true},
					{true, true, true},
				},
				inputX:                  0,
				inputY:                  3,
				expectedIsNeighborAlive: BeFalse(),
			},
		}

		for _, test := range tests {
			test := test

			t.Run(fmt.Sprintf("(%d,%d)", test.inputX, test.inputY), func(t *testing.T) {
				world := World{
					inhabitants: test.inhabitants,
					size:        3,
				}

				actualNeighborIsAlive := world.checkNeighborIsAlive(test.inputX, test.inputY)
				tt.Expect(actualNeighborIsAlive).To(test.expectedIsNeighborAlive)
			})
		}
	})

	t.Run("should return true for a living neighbor:", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		tests := []testFacts{
			{
				inhabitants: [][]bool{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
				inputX:                  0,
				inputY:                  0,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
				inputX:                  1,
				inputY:                  0,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
				inputX:                  2,
				inputY:                  0,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, false, false},
					{true, false, false},
					{false, false, false},
				},
				inputX:                  0,
				inputY:                  1,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
				inputX:                  1,
				inputY:                  1,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, false, false},
					{false, false, true},
					{false, false, false},
				},
				inputX:                  2,
				inputY:                  1,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, false, false},
					{false, false, false},
					{true, false, false},
				},
				inputX:                  0,
				inputY:                  2,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, false, false},
					{false, false, false},
					{false, true, false},
				},
				inputX:                  1,
				inputY:                  2,
				expectedIsNeighborAlive: BeTrue(),
			},
			{
				inhabitants: [][]bool{
					{false, false, false},
					{false, false, false},
					{false, false, true},
				},
				inputX:                  2,
				inputY:                  2,
				expectedIsNeighborAlive: BeTrue(),
			},
		}

		for _, test := range tests {
			test := test

			t.Run(fmt.Sprintf("(%d,%d)", test.inputX, test.inputY), func(t *testing.T) {
				world := World{
					inhabitants: test.inhabitants,
					size:        3,
				}

				actualNeighborIsAlive := world.checkNeighborIsAlive(test.inputX, test.inputY)
				tt.Expect(actualNeighborIsAlive).To(test.expectedIsNeighborAlive)
			})
		}
	})
}
