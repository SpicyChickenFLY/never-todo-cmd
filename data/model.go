package data

import "time"

type Task struct {
	ID        int    `json:"id" mapstructure:"id"`
	Content   string `json:"content" mapstructure:"content"`
	Deleted   bool   `json:"deleted" mapstructure:"deleted"`
	Completed bool   `json:"completed" mapstructure:"completed"`
	Important bool   `json:"important" mapstructure:"important"`
}

type Tag struct {
	ID      int    `json:"id" mapstructure:"id"`
	Content string `json:"content" mapstructure:"content"`
	Deleted bool   `json:"deleted" mapstructure:"deleted"`
	Color   string `json:"color" mapstructure:"color"`
}

type TaskTag struct {
	TaskID int `json:"task_id" mapstructure:"task_id"`
	TagID  int `json:"tag_id" mapstructure:"tag_id"`
}

type TimeGroup struct {
	level int
	start time.Time
	end   time.Time
}

type Loop struct {
	year   int
	month  int
	week   int
	hour   int
	minute int
}

type Log struct {
	Target string                 `json:"target" mapstructure:"target"`
	Type   string                 `json:"type" mapstructure:"type"`
	Data   map[string]interface{} `json:"data" mapstructure:"data"`
}

type Model struct {
	Data struct {
		Tasks          []Task    `json:"tasks" mapstructure:"tasks"`
		Tags           []Tag     `json:"tags" mapstructure:"tags"`
		TaskTags       []TaskTag `json:"task_tags" mapstructure:"task_tags"`
		TaskAutoIncVal int       `json:"taskAutoIncVal" mapstructure:"taskAutoIncVal"`
		TagAutoIncVal  int       `json:"tagAutoIncVal" mapstructure:"tagAutoIncVal"`
	}
	Log []Log `json:"log" mapstructure:"log"`
}
