package region

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromZipCode(t *testing.T) {
	// Test exact match
	addr, err := FromZipCode("100000")
	assert.NoError(t, err)
	assert.Equal(t, "北京市", addr.Province)
	assert.Equal(t, "北京市", addr.City)
	assert.Equal(t, "东城区", addr.District)
	assert.Equal(t, "100000", addr.ZipCode)

	// Test prefix match
	addr, err = FromZipCode("100010")
	assert.NoError(t, err)
	assert.Equal(t, "北京市", addr.Province)
	assert.Equal(t, "北京市", addr.City)
	assert.Equal(t, "朝阳区", addr.District)
	assert.Equal(t, "100010", addr.ZipCode)

	// Test another province
	addr, err = FromZipCode("510000")
	assert.NoError(t, err)
	assert.Equal(t, "广东省", addr.Province)
	assert.Equal(t, "广州市", addr.City)
	assert.Equal(t, "越秀区", addr.District)
	assert.Equal(t, "510000", addr.ZipCode)

	// Test non-existent zip code
	addr, err = FromZipCode("999999")
	assert.Error(t, err)
	assert.Equal(t, ErrZipCodeNotFound, err)
	assert.Empty(t, addr.Province)
}

func TestToRegionCode(t *testing.T) {
	// Test with full address
	code, err := ToRegionCode("北京市", "北京市", "东城区")
	assert.NoError(t, err)
	assert.Equal(t, RegionCode("110101"), code)

	// Test with another region
	code, err = ToRegionCode("上海市", "上海市", "黄浦区")
	assert.NoError(t, err)
	assert.Equal(t, RegionCode("310101"), code)

	// Test with province and city only
	code, err = ToRegionCode("浙江省", "杭州市", "")
	assert.NoError(t, err)
	assert.Equal(t, RegionCode("330102"), code)

	// Test non-existent region
	code, err = ToRegionCode("不存在的省", "不存在的市", "不存在的区")
	assert.Error(t, err)
	assert.Equal(t, ErrRegionNotFound, err)
	assert.Empty(t, code)
}

func TestAddressToRegionCode(t *testing.T) {
	// Test with full address
	addr := Address{
		Province: "广东省",
		City:     "广州市",
		District: "越秀区",
		ZipCode:  "510000",
	}

	code, err := AddressToRegionCode(addr)
	assert.NoError(t, err)
	assert.Equal(t, RegionCode("440104"), code)
}
