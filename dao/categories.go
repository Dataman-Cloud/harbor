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

package dao

import (
	"github.com/vmware/harbor/models"

	//"errors"
	//"fmt"
	//"log"
	//"strings"

	"github.com/astaxie/beego/orm"
)

func AddCategories(categories *models.Categories) (*models.Categories, error) {
	o := orm.NewOrm()

	p, err := o.Raw("insert into categories(name) values(?)").Prepare()
	if err != nil {
		return nil, err
	}
	c, err := p.Exec(categories.Name)
	categoriesID, err := c.LastInsertId()
	if err != nil {
		return nil, err
	}
	categories.Id = categoriesID
	return categories, nil
}

func GetAllCategories() ([]models.Categories, error) {
	o := orm.NewOrm()
	var c []models.Categories
	_, err := o.Raw("select * form categories").QueryRows(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
