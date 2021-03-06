/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package models

import (
	"time"
)

// Repo holds information about repositories.
type Repo struct {
	Repositories []string `json:"repositories"`
}

// RepoItem holds manifest of an image.
type RepoItem struct {
	ID            string    `json:"Id"`
	Parent        string    `json:"Parent"`
	Created       time.Time `json:"Created"`
	CreatedStr    string    `json:"CreatedStr"`
	DurationDays  string    `json:"Duration Days"`
	Author        string    `json:"Author"`
	Architecture  string    `json:"Architecture"`
	DockerVersion string    `json:"Docker Version"`
	Os            string    `json:"OS"`
	//Size           int       `json:"Size"`
}

// Tag holds information about a tag.
type Tag struct {
	Id             int64     `json:"id"`
	Version        string    `json:"version"`
	ImageID        string    `json:"imageId"`
	ProjectID      int64     `orm:"column(project_id)" json:"projectId"`
	ProjectName    string    `json:"projectName"`
	RepositoryName string    `json:"repositoryName"`
	RepositoryID   int64     `orm:"column(repository_id)"json:"respositoryId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type Repository struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	ProjectName    string    `json:"projectName"`
	ProjectID      int64     `orm:"column(project_id)" json:"projectId"`
	UserName       string    `json:"userName"`
	Category       string    `json:"category"`
	IsPublic       int64     `json:"isPublic"`
	LatestTag      string    `json:"latestTag"`
	Description    string    `json:"description"`
	Readme         string    `json:"readme"`
	DockerCompose  string    `json:"dockerCompose"`
	MarathonConfig string    `json:"marathonConfig"`
	Catalog        string    `json:"catalog"`
	Questions      string    `json:"questions"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type RepositoriesResponse struct {
	Code int64        `json:"code"`
	Data []Repository `json:"data"`
}

type RepositoryResponse struct {
	Code    int64       `json:"code"`
	Data    *Repository `json:"data"`
	Message string      `json:"message"`
}

type TagsResponse struct {
	Code int64 `json:"code"`
	Data []Tag `json:"data"`
}

type CategoriesResponse struct {
	Code int64    `json:"code"`
	Data []string `json:"data"`
}
