package zipcode

import (
	"github.com/gopherus/mymodule/checksum"
	"github.com/viant/xdatly/types/core"
	"reflect"
)

var PackageName = "zipcode"

func init() {
	core.RegisterType(PackageName, "Entity", reflect.TypeOf(Entity{}), checksum.GeneratedTime)
}

type Entity struct {
	Entity []*Zipcode `typeName:"Zipcode"`
}

type Zipcode struct {
	Id            int         `sqlx:"name=id,autoincrement,primaryKey,required"`
	CountryCode   *string     `sqlx:"name=country_code,unique,table=CI_ZIPCODE" json:",omitempty"`
	PostalCode    *string     `sqlx:"name=postal_code" json:",omitempty"`
	CityName      *string     `sqlx:"name=city_name" json:",omitempty"`
	StateName     *string     `sqlx:"name=state_name" json:",omitempty"`
	StateCode     *string     `sqlx:"name=state_code" json:",omitempty"`
	CountyName    *string     `sqlx:"name=county_name" json:",omitempty"`
	CountyCode    *string     `sqlx:"name=county_code" json:",omitempty"`
	CommunityName *string     `sqlx:"name=community_name" json:",omitempty"`
	CommunityCode *string     `sqlx:"name=community_code" json:",omitempty"`
	Latitude      *float64    `sqlx:"name=latitude" json:",omitempty"`
	Longitude     *float64    `sqlx:"name=longitude" json:",omitempty"`
	Accuracy      *int        `sqlx:"name=accuracy" json:",omitempty"`
	Has           *ZipcodeHas `presenceIndex:"true" typeName:"ZipcodeHas" json:"-" diff:"presence=true" sqlx:"presence=true" validate:"presence=true"`
}

type ZipcodeHas struct {
	Id            bool
	CountryCode   bool
	PostalCode    bool
	CityName      bool
	StateName     bool
	StateCode     bool
	CountyName    bool
	CountyCode    bool
	CommunityName bool
	CommunityCode bool
	Latitude      bool
	Longitude     bool
	Accuracy      bool
}
