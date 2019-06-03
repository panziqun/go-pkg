package rbac

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/laughmaker/go-pkg/db"
	"github.com/laughmaker/go-pkg/model"
	"github.com/laughmaker/go-pkg/resp"

	"github.com/gin-gonic/gin"
)

type postPermissionForm struct {
	PermissionID int    `form:"permission_id"`
	Name         string `form:"name" binding:"required"`
	Method       string `form:"method"`
	Route        string `form:"route" binding:"required"`
}

// @Summary 权限-保存
// @Tags RBAC
// @Produce  json
// @Param permission_id query int false "权限ID"
// @Param name query string true "权限名"
// @Param method query string false "方法:GET、POST、PUT、DELETE、空代表所有"
// @Param route query string true "路由: /after-sale/order /after-sale/*"
// @Success 200 {object} resp.Data
// @Router /rbac/permission [post]
func PostPermission(c *gin.Context) {
	r := resp.Resp{C: c}
	var form postPermissionForm
	if err := c.ShouldBind(&form); err != nil {
		r.Error(http.StatusBadRequest)
		return
	}
	var p Permission
	if form.PermissionID != 0 {
		err := db.DB.Where("id = ?", form.PermissionID).First(&p).Error
		if err != nil {
			r.Failure(resp.DataNotExist)
		}
		err = db.DB.Model(&p).Updates(Permission{
			Name:   form.Name,
			Method: form.Method,
			Route:  form.Route,
		}).Error
		if err != nil {
			panic(err)
		}
		r.Success(nil)
		return
	}
	p = Permission{
		Name:   form.Name,
		Method: form.Method,
		Route:  form.Route,
	}
	err := db.DB.Create(&p).Error
	if err != nil {
		panic(err)
	}
	r.Success(nil)
	return
}

type deletePermissionForm struct {
	PermissionID int `form:"permission_id" binding:"required"`
}

// @Summary 权限-删除
// @Tags RBAC
// @Produce  json
// @Param permission_id query int true "permission_id"
// @Success 200 {object} resp.Data
// @Router /rbac/permission [delete]
func DeletePermission(c *gin.Context) {
	r := resp.Resp{C: c}
	var form deletePermissionForm
	if err := c.ShouldBind(&form); err != nil {
		r.Error(http.StatusBadRequest)
		return
	}
	var p Permission
	if err := db.DB.Where("id = ?", form.PermissionID).First(&p).Error; err != nil {
		r.Failure(resp.DataNotExist)
		return
	}
	if err := db.DB.Delete(&p).Error; err != nil {
		panic(err)
	}
	r.Success(nil)
	return
}

// @Summary 权限-列表
// @Tags RBAC
// @Produce  json
// @Param page_index query int false "分页索引 默认1"
// @Param page_size query int false "分页大小 默认20"
// @Success 200 {object} resp.Data
// @Router /rbac/permission/list [get]
func GetPermissionList(c *gin.Context) {
	r := resp.Resp{C: c}
	var p []Permission
	query := db.DB.Model(&Permission{})
	page := model.DbPage(c, query)
	if err := query.Order("name desc, id").Find(&p).Error; err != nil {
		panic(err)
	}
	r.List(p, page)
	return
}

type postRoleForm struct {
	RoleID      int    `form:"role_id"`
	Name        string `form:"name" binding:"required"`
	Permissions string `form:"permissions" binding:"required"`
}

// @Summary 角色-保存
// @Tags RBAC
// @Produce  json
// @Param role_id query int false "角色ID"
// @Param name query string true "角色名"
// @Param permissions query string true "权限"
// @Success 200 {object} resp.Data
// @Router /rbac/role [post]
func PostRole(c *gin.Context) {
	r := resp.Resp{C: c}
	var form postRoleForm
	if err := c.ShouldBind(&form); err != nil {
		r.Error(http.StatusBadRequest)
		return
	}
	var role Role
	tx := db.DB.Begin()
	if form.RoleID != 0 {
		err := tx.Where("id = ?", form.RoleID).First(&role).Error
		if err != nil {
			r.Failure(resp.DataNotExist)
			return
		}
		err = tx.Model(&role).Updates(Role{
			Name: form.Name,
		}).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	} else {
		role = Role{
			Name: form.Name,
		}
		err := tx.Create(&role).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}
	var pemIds []int
	if err := json.Unmarshal([]byte(form.Permissions), &pemIds); err != nil {
		tx.Rollback()
		panic(err)
	}
	if err := tx.Where("role_id = ?", role.ID).Delete(RolePermission{}).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	for _, v := range pemIds {
		rp := RolePermission{
			RoleID:       int(role.ID),
			PermissionID: v,
		}
		err := tx.Create(&rp).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}
	tx.Commit()
	r.Success(nil)
	return
}

// @Summary 角色-列表
// @Tags RBAC
// @Produce  json
// @Param page_index query int false "分页索引 默认1"
// @Param page_size query int false "分页大小 默认20"
// @Success 200 {object} resp.Data
// @Router /rbac/role/list [get]
func GetRoleList(c *gin.Context) {
	r := resp.Resp{C: c}
	var roles []Role
	query := db.DB.Model(&Role{})
	page := model.DbPage(c, query)
	if err := query.Order("id desc").Find(&roles).Error; err != nil {
		panic(err)
	}
	for k, v := range roles {
		db.DB.Model(&v).Related(&roles[k].RolePermissions)
		for i, j := range roles[k].RolePermissions {
			db.DB.Model(&j).Related(&roles[k].RolePermissions[i].Permission)
		}
	}
	r.List(roles, page)
	return
}

type deleteRoleForm struct {
	RoleID int `form:"role_id" binding:"required"`
}

// @Summary 角色-删除
// @Tags RBAC
// @Produce  json
// @Param role_id query int true "role_id"
// @Success 200 {object} resp.Data
// @Router /rbac/role [delete]
func DeleteRole(c *gin.Context) {
	r := resp.Resp{C: c}
	var form deleteRoleForm
	if err := c.ShouldBind(&form); err != nil {
		r.Error(http.StatusBadRequest)
		return
	}
	var role Role
	if err := db.DB.Where("id = ?", form.RoleID).First(&role).Error; err != nil {
		r.Failure(resp.DataNotExist)
		return
	}
	if err := db.DB.Delete(&role).Error; err != nil {
		panic(err)
	}
	r.Success(nil)
	return
}

func SyncSwagList() {
	var list []map[string]string
	data, err := ioutil.ReadFile("./docs/swagger.json")
	if err != nil {
		panic(err)
	}
	var re map[string]interface{}
	err = json.Unmarshal(data, &re)
	if err != nil {
		panic(err)
	}

	var paths map[string]interface{} = re["paths"].(map[string]interface{})
	for k, v := range paths {
		for i, j := range v.(map[string]interface{}) {
			var path map[string]string = make(map[string]string)
			tags := j.(map[string]interface{})["tags"]
			path["name"] = tags.([]interface{})[0].(string) + "-" + j.(map[string]interface{})["summary"].(string)
			path["method"] = strings.ToUpper(i)
			path["route"] = k
			list = append(list, path)
			var p Permission
			err := db.DB.Where("method = ? and route = ?", path["method"], path["route"]).First(&p).Error
			if err != nil {
				p = Permission{
					Name:   path["name"],
					Method: path["method"],
					Route:  path["route"],
				}
				err := db.DB.Create(&p).Error
				if err != nil {
					panic(err)
				}
			} else {
				err := db.DB.Model(&p).Updates(Permission{Name: path["name"]}).Error
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
