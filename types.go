package main

type Social struct {
	Platform string `json:"Platform"`
	URL      string `json:"Url"`
}
type Skill struct {
	Name  string `json:"Name"`
	Balls int    `json:"Balls"`
}
type SkillWithUsage struct {
	Skill
	Learning bool `json:"Learning"`
	Using    bool `json:"Using"`
	Show     bool `json:"Show"`
}
type Education struct {
	Start       string  `json:"Start"`
	End         *string `json:"End"`
	Institution string  `json:"Institution"`
	Place       string  `json:"Place"`
	Degree      string  `json:"Degree"`
	Niveau      string  `json:"Niveau"`
	Finished    bool    `json:"Finished"`
}
type Job struct {
	Start   string `json:"Start"`
	End     string `json:"End"`
	Company string `json:"Company"`
	Place   string `json:"Place"`
	Role    string `json:"Role"`
}
type Internship = Job
type Address struct {
	Street     string `json:"Street"`
	PostalCode string `json:"PostalCode"`
	City       string `json:"City"`
	Country    string `json:"Country"`
}
type Contact struct {
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}
type CV struct {
	Name              string           `json:"Name"`
	Picture           string           `json:"Picture"`
	Born              string           `json:"Born"`
	Goals             []string         `json:"Goals"`
	Address           Address          `json:"Address"`
	Contact           Contact          `json:"Contact"`
	Socials           []Social         `json:"Socials"`
	Interests         []string         `json:"Interests"`
	ProgrammingSkills []SkillWithUsage `json:"ProgrammingSkills"`
	Languages         []Skill          `json:"Languages"`
	Educations        []Education      `json:"Educations"`
	Jobs              []Job            `json:"Jobs"`
	Internships       []Internship     `json:"Internships"`
}
