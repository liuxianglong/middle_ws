// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"middle/internal/dao/internal"
)

// cmsModuleDao is the data access object for the table cms_module.
// You can define custom methods on it to extend its functionality as needed.
type cmsModuleDao struct {
	*internal.CmsModuleDao
}

var (
	// CmsModule is a globally accessible object for table cms_module operations.
	CmsModule = cmsModuleDao{internal.NewCmsModuleDao()}
)

// Add your custom methods and functionality below.
