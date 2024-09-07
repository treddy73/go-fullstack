package db

import "strings"

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Collection struct {
	todos []*Todo
}

func NewCollection() *Collection {
	return &Collection{}
}

func (c *Collection) Add(name string) *Todo {
	t := &Todo{Id: len(c.todos), Name: name, Completed: false}
	c.todos = append(c.todos, t)
	return t
}

func (c *Collection) Filter(keyword string) []*Todo {
	if keyword == "" {
		return c.todos
	}

	keyword = strings.ToLower(keyword)
	var result []*Todo
	for _, t := range c.todos {
		if strings.Contains(strings.ToLower(t.Name), keyword) {
			result = append(result, t)
		}
	}

	return result
}
