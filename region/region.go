package region

import (
	"errors"
	"strings"
)

// Address represents a complete address with province, city, district, and zip code
type Address struct {
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	ZipCode  string `json:"zip_code"`
}

// RegionCode represents an administrative region code
type RegionCode string

// ErrZipCodeNotFound is returned when the zip code is not found
var ErrZipCodeNotFound = errors.New("zip code not found")

// ErrRegionNotFound is returned when the region is not found
var ErrRegionNotFound = errors.New("region not found")

// zipCodeMap maps zip codes to province/city/district
var zipCodeMap = map[string]Address{
	// Example data - in real implementation, this would be a comprehensive map
	"100000": {Province: "北京市", City: "北京市", District: "东城区", ZipCode: "100000"},
	"100010": {Province: "北京市", City: "北京市", District: "朝阳区", ZipCode: "100010"},
	"200000": {Province: "上海市", City: "上海市", District: "黄浦区", ZipCode: "200000"},
	"310000": {Province: "浙江省", City: "杭州市", District: "上城区", ZipCode: "310000"},
	"510000": {Province: "广东省", City: "广州市", District: "越秀区", ZipCode: "510000"},
}

// regionCodeMap maps province/city/district to region code
var regionCodeMap = map[string]RegionCode{
	// Example data - in real implementation, this would be based on GB/T 2260
	"北京市北京市东城区": "110101",
	"北京市北京市朝阳区": "110105",
	"上海市上海市黄浦区": "310101",
	"浙江省杭州市上城区": "330102",
	"广东省广州市越秀区": "440104",
}

// FromZipCode converts a zip code to an Address
func FromZipCode(zipCode string) (Address, error) {
	// Normalize zip code - remove spaces and ensure 6 digits
	zipCode = strings.TrimSpace(zipCode)
	if len(zipCode) != 6 {
		// Try to find by prefix (first 3 digits)
		prefix := zipCode[:3]
		for k, v := range zipCodeMap {
			if strings.HasPrefix(k, prefix) {
				return v, nil
			}
		}
		return Address{}, ErrZipCodeNotFound
	}

	// Exact match
	addr, ok := zipCodeMap[zipCode]
	if ok {
		return addr, nil
	}

	// Try to find by prefix (first 3 digits)
	prefix := zipCode[:3]
	for k, v := range zipCodeMap {
		if strings.HasPrefix(k, prefix) {
			return v, nil
		}
	}

	return Address{}, ErrZipCodeNotFound
}

// ToRegionCode converts a province/city/district to a RegionCode
func ToRegionCode(province, city, district string) (RegionCode, error) {
	// Normalize inputs
	province = strings.TrimSpace(province)
	city = strings.TrimSpace(city)
	district = strings.TrimSpace(district)

	// Create key with all three components
	key := province + city + district

	code, ok := regionCodeMap[key]
	if ok {
		return code, nil
	}

	// Try with province and city only (exact match)
	key = province + city
	code, ok = regionCodeMap[key]
	if ok {
		return code, nil
	}

	// Try to find by province + city prefix (for cases where district might vary)
	for k, v := range regionCodeMap {
		if strings.HasPrefix(k, province+city) {
			return v, nil
		}
	}

	return "", ErrRegionNotFound
}

// AddressToRegionCode converts an Address to a RegionCode
func AddressToRegionCode(addr Address) (RegionCode, error) {
	return ToRegionCode(addr.Province, addr.City, addr.District)
}
