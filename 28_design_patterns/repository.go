package main

import (
	"log"

	"github.com/gmdmgithub/go_first/28_design_patterns/inmem"
	"github.com/gmdmgithub/go_first/28_design_patterns/project"
)

//  Abstract persistence/retrieval from business logic
func repositoryPattern() {
	log.Println("## Repository pattern")
	defer log.Println("-- Bye from Repository")

	projectsRepo := inmem.ProjectRepository{}
	newProject, _ := projectsRepo.Store(
		project.Project{
			Name: "First project",
		},
	)
	newProject.Name = "First project - updated!"
	p, err := projectsRepo.Store(newProject)
	if err != nil {
		log.Printf("We got the problem %v", err)
		return
	}
	log.Printf("Success! project stored! %+v", p)
}
