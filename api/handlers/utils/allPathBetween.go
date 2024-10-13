// patched github.com/dominikbraun/graph

package utils

import (
	"errors"
	"fmt"
	"time"

	dominik "github.com/dominikbraun/graph"
	"github.com/rs/zerolog/log"
)

func AllPathsBetween(g dominik.Graph[string, string], start, end string) ([][]string, error) {
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return nil, err
	}

	// The algorithm used relies on stacks instead of recursion. It is described
	// here: https://boycgit.github.io/all-paths-between-two-vertex/
	mainStack := newStack[string]()
	viceStack := newStack[stack[string]]()

	checkEmpty := func() error {
		if mainStack.isEmpty() || viceStack.isEmpty() {
			return errors.New("empty stack")
		}
		return nil
	}

	buildLayer := func(element string) {
		mainStack.push(element)

		newElements := newStack[string]()
		for e := range adjacencyMap[element] {
			var contains bool
			mainStack.forEach(func(k string) {
				if e == k {
					contains = true
				}
			})
			if contains {
				continue
			}
			newElements.push(e)
		}
		viceStack.push(newElements)
	}

	buildStack := func() error {
		if err = checkEmpty(); err != nil {
			return fmt.Errorf("unable to build stack: %w", err)
		}

		elements, _ := viceStack.top()

		for !elements.isEmpty() {
			element, _ := elements.pop()
			buildLayer(element)
			elements, _ = viceStack.top()
		}

		return nil
	}

	removeLayer := func() error {
		if err = checkEmpty(); err != nil {
			return fmt.Errorf("unable to remove layer: %w", err)
		}

		if e, _ := viceStack.top(); !e.isEmpty() {
			return errors.New("the top element of vice-stack is not empty")
		}

		_, _ = mainStack.pop()
		_, _ = viceStack.pop()

		return nil
	}

	buildLayer(start)

	allPaths := make([][]string, 0)

	cnt := 0
	startAt := time.Now()
	for !mainStack.isEmpty() {
		v, _ := mainStack.top()
		adjs, _ := viceStack.top()

		// our patch to limit execution time
		cnt++
		if cnt%100_000 == 0 {
			if time.Since(startAt) > 5*time.Second {
				return nil, errors.New("timeout")
			}
			if cnt > 1_000_000 {
				// log.Warn().Msgf("too many iterations %s", allPaths)
				log.Warn().Msgf("too many iterations %s | main stack %d | adjs %d |", v, mainStack.len(), adjs.len())
			}

		}
		if adjs.isEmpty() {
			if v == end {
				path := make([]string, 0)
				mainStack.forEach(func(k string) {
					path = append(path, k)
				})
				allPaths = append(allPaths, path)
			}

			err = removeLayer()
			if err != nil {
				return nil, err
			}
		} else {
			if err = buildStack(); err != nil {
				return nil, err
			}
		}
	}

	return allPaths, nil
}

type stack[T any] interface {
	push(T)
	pop() (T, error)
	top() (T, error)
	isEmpty() bool
	// forEach iterate the stack from bottom to top
	forEach(func(T))
	len() int
}

func newStack[T any]() stack[T] {
	return &stackImpl[T]{
		elements: make([]T, 0),
	}
}

type stackImpl[T any] struct {
	elements []T
}

func (s *stackImpl[T]) push(t T) {
	s.elements = append(s.elements, t)
}

func (s *stackImpl[T]) pop() (T, error) {
	e, err := s.top()
	if err != nil {
		var defaultValue T
		return defaultValue, err
	}

	s.elements = s.elements[:len(s.elements)-1]
	return e, nil
}

func (s *stackImpl[T]) top() (T, error) {
	if s.isEmpty() {
		var defaultValue T
		return defaultValue, errors.New("no element in stack")
	}

	return s.elements[len(s.elements)-1], nil
}

func (s *stackImpl[T]) isEmpty() bool {
	return len(s.elements) == 0
}
func (s *stackImpl[T]) len() int {
	return len(s.elements)
}

func (s *stackImpl[T]) forEach(f func(T)) {
	for _, e := range s.elements {
		f(e)
	}
}
