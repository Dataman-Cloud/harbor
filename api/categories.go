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

package api

import (
	"net/http"

	"github.com/vmware/harbor/dao"

	"github.com/astaxie/beego"
	"github.com/vmware/harbor/models"
)

type CategoriesAPI struct {
	BaseAPI
}

// Get ...
func (c *CategoriesAPI) Get() {
	categories, err := dao.GetAllCategories()
	if err != nil {
		beego.Error("Failed to get categories from mysql error: ", err)
		c.CustomAbort(http.StatusInternalServerError, "Failed to get categories")
	}
	c.Data["json"] = categories
	c.ServeJSON()
}

func (c *CategoriesAPI) Post() {
	name := c.GetString("name")

	categories := &models.Categories{
		Name: name,
	}
	categories, err := dao.AddCategories(categories)
	if err != nil {
		beego.Error("Failed to insert into categories error: ", err)
		c.CustomAbort(http.StatusInternalServerError, "Failed to insert into categories error")
	}

	c.Data["json"] = categories
	c.ServeJSON()
}
